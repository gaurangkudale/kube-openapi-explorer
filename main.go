package main

import (
	"fmt"
	"os"

	"github.com/gaurangkudale/kube-openapi-explorer/pkg/discovery"
)

func main() {
	groups, err := discovery.ListAPIGroups()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listing API groups: %v\n", err)
		os.Exit(1)
	}

	for _, group := range groups {
		fmt.Println(group)
	}
}
