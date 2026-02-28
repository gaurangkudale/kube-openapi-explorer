package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"

	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var mergedSpec map[string]interface{}

func main() {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")

	cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	dc, err := discovery.NewDiscoveryClientForConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Fetching Kubernetes OpenAPI v3 schemas...")

	mergedSpec = buildMergedOpenAPI(dc)

	log.Println("✅ OpenAPI 3.0 aggregation complete")

	http.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mergedSpec)
	})

	log.Println("Swagger JSON: http://localhost:8080/swagger.json")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func buildMergedOpenAPI(dc *discovery.DiscoveryClient) map[string]interface{} {

	openapiClient := dc.OpenAPIV3()

	pathsClient, err := openapiClient.Paths()
	if err != nil {
		log.Fatalf("failed to list OpenAPI paths: %v", err)
	}

	final := map[string]interface{}{
		"openapi": "3.0.0",
		"info": map[string]interface{}{
			"title":   "Kubernetes API",
			"version": "v1",
		},
		"paths": map[string]interface{}{},
		"components": map[string]interface{}{
			"schemas": map[string]interface{}{},
		},
	}

	finalPaths := final["paths"].(map[string]interface{})
	finalSchemas := final["components"].(map[string]interface{})["schemas"].(map[string]interface{})

	for name, gvClient := range pathsClient {
		log.Printf("Merging %s", name)

		// ✅ NEW API — returns JSON already
		raw, err := gvClient.Schema("application/json")
		if err != nil {
			log.Printf("skip %s: %v", name, err)
			continue
		}

		var spec map[string]interface{}
		if err := json.Unmarshal(raw, &spec); err != nil {
			log.Printf("parse failed %s: %v", name, err)
			continue
		}

		// merge paths
		if p, ok := spec["paths"].(map[string]interface{}); ok {
			for k, v := range p {
				finalPaths[k] = v
			}
		}

		// merge schemas
		if comps, ok := spec["components"].(map[string]interface{}); ok {
			if schemas, ok := comps["schemas"].(map[string]interface{}); ok {
				for k, v := range schemas {
					finalSchemas[k] = v
				}
			}
		}
	}

	return final
}