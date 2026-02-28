# Examples

This directory contains example configurations and scripts demonstrating how to use `kube-openapi-explorer`.

## Running against a local KIND cluster

```bash
# Start a local cluster
kind create cluster

# Run the explorer
go run main.go
```

## Running against a remote cluster

```bash
# Point to a specific kubeconfig
KUBECONFIG=/path/to/kubeconfig go run main.go
```
