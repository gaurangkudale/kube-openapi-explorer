# kube-openapi-explorer

A lightweight Go tool to explore and inspect the Kubernetes OpenAPI schema directly from a running cluster.

This project connects to a Kubernetes cluster using your kubeconfig and dynamically retrieves the API schema exposed by the Kubernetes API server (`/openapi/v3`).

It helps developers, platform engineers, and operator authors understand Kubernetes resources programmatically — without relying on external documentation.

---

## Features

- Discover Kubernetes APIs dynamically
- Fetch OpenAPI v3 schema from live clusters
- Inspect built-in and Custom Resource Definitions (CRDs)
- Human-readable schema exploration
- Works with any Kubernetes distribution (KIND, EKS, GKE, AKS, etc.)
- Uses official `client-go` discovery APIs

---

## 🚀 Why this exists

Kubernetes exposes its entire API surface via OpenAPI, but:

- There is no native `/docs` endpoint
- Swagger UI is not bundled with Kubernetes
- Debugging CRDs and schemas can be difficult

`kube-openapi-explorer` provides a developer-friendly way to inspect the Kubernetes API directly from code.

---

## How it works

```
Kubernetes API Server
        ↓
/openapi/v3 (live schema)
        ↓
Go program (client-go discovery)
        ↓
OpenAPI JSON aggregation
        ↓
Swagger UI server
        ↓
http://localhost:8080/docs
```

The tool authenticates using your local kubeconfig and queries the API server securely.

---

## Installation

```bash
git clone https://github.com/gaurangkudale/kube-openapi-explorer.git
cd kube-openapi-explorer

go mod tidy
go run main.go
```

---

## ⚙️ Requirements

- Go 1.22+
- Kubernetes cluster access
- Valid kubeconfig (`~/.kube/config`)

---

## 🔧 Usage

### List available API groups

```bash
go run main.go
```

Example output:

```
api/v1
apis/apps/v1
apis/batch/v1
apis/networking.k8s.io/v1
```

---

### Explore schemas (example)

Future versions will support:

- Resource filtering
- CRD schema inspection
- Markdown documentation generation
- Swagger UI export

---

## Use Cases

- Kubernetes Operator development
- CRD validation debugging
- API discovery tooling
- Platform engineering workflows
- Learning Kubernetes internals

---

## Related Kubernetes Concepts

- Kubernetes API Discovery
- OpenAPI v3 Specification
- `kubectl explain`
- client-go discovery client

---

## 🤝 Contributing

Contributions are welcome!

Ideas for improvement:

- Interactive CLI
- Swagger UI generation
- API diff between clusters
- CRD schema validation tools

---

## 📜 License

MIT License
