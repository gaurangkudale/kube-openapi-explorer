// Package discovery provides utilities to connect to a Kubernetes cluster
// and retrieve API group and resource information via the discovery client.
package discovery

import (
	"fmt"

	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
)

// ListAPIGroups connects to the cluster specified by the kubeconfig resolved
// via standard rules (KUBECONFIG env var, then ~/.kube/config) and returns a
// list of all available API group/version paths, e.g. "api/v1", "apis/apps/v1".
func ListAPIGroups() ([]string, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, &clientcmd.ConfigOverrides{})

	config, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("building kubeconfig: %w", err)
	}

	client, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("creating discovery client: %w", err)
	}

	apiGroupList, err := client.ServerGroups()
	if err != nil {
		return nil, fmt.Errorf("listing server groups: %w", err)
	}

	var paths []string
	for _, group := range apiGroupList.Groups {
		for _, version := range group.Versions {
			if group.Name == "" {
				paths = append(paths, "api/"+version.Version)
			} else {
				paths = append(paths, "apis/"+version.GroupVersion)
			}
		}
	}

	return paths, nil
}
