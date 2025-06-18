# pinecone-lite

**pinecone-lite** is a minimal and idiomatic Go client for interacting with [Pinecone](https://www.pinecone.io/), a vector database for semantic search and similarity matching. This SDK focuses on the core functionality required for simple, reliable use of Pinecone's HTTP API with no unnecessary complexity.

---

## 🚀 Features

- Upsert vectors to an index
- Query vectors by similarity
- Delete vectors by ID or entire namespace
- Handles API error responses cleanly
- Zero external dependencies
- Supports float64 vectors (auto-aligned with Pinecone's float32 backend)

---

## 🔧 Installation

```bash
go get github.com/yourusername/pinecone-lite
```

---

## ✍️ Usage

### Initialize Client

```go
client := pinecone.NewClient("https://your-index.svc.your-region.pinecone.io", "your-api-key")
```

### Upsert Vectors

```go
vectors := []*pinecone.Vector{
  {
    ID: "vec1",
    Values: []float64{0.1, 0.2, 0.3},
    Metadata: map[string]any{"label": "example"},
  },
}

count, err := client.UpsertVectors(ctx, vectors, "my-namespace")
```

### Query Vectors

```go
resp, err := client.QueryByVectors(ctx, &pinecone.QueryByVectorRequest{
  Vector: []float64{0.1, 0.2, 0.3},
  TopK: 3,
  Namespace: "my-namespace",
  IncludeMetadata: true,
})
```

### Delete Vectors

```go
err := client.DeleteVectorsByID(ctx, []string{"vec1"}, "my-namespace")
```

---

## 📘 API Reference

See [GoDoc](https://pkg.go.dev/github.com/qhenkart/pinecone-lite) for full reference.

---

## 🧪 Testing

The SDK is tested using a custom HTTP transport layer for mocking API responses. See `*_test.go` files for examples.

---

## 🧪 License

MIT License

---

## 🙏 Acknowledgements

- [Pinecone](https://www.pinecone.io/)
- Inspired by the official Pinecone HTTP API spec

---

## 📬 Contributions

Contributions welcome. Please open issues or submit pull requests with clear commit messages and test coverage.


