// Package main is the entry point for the kube-openapi-explorer CLI.
// Run this binary to list all Kubernetes API groups available in the
// cluster pointed to by your current kubeconfig context.
package main

import (
	"fmt"
	"os"

	"github.com/gaurangkudale/kube-openapi-explorer/pkg/discovery"
)

func main() {
	groups, err := discovery.ListAPIGroups()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	for _, group := range groups {
		fmt.Println(group)
	}
}
