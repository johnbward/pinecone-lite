package pinecone

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

type roundTripFunc func(req *http.Request) *http.Response

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func TestQueryByVectors(t *testing.T) {
	t.Run("valid_query_response", func(t *testing.T) {
		mockResponse := QueryByVectorResponse{
			Matches: []MatchResult{
				{
					ID:    "rec1",
					Score: 0.92,
					Metadata: map[string]any{
						"category":   "test",
						"chunk_text": "example text",
					},
				},
			},
			Namespace: "test-namespace",
		}

		data, _ := json.Marshal(mockResponse)

		client := &Client{
			IndexURL: "https://example-index.svc.us-east1-gcp.io",
			APIKey:   "test-key",
			HTTPClient: &http.Client{
				Transport: roundTripFunc(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: 200,
						Body:       io.NopCloser(bytes.NewReader(data)),
						Header:     make(http.Header),
					}
				}),
			},
		}

		filter := map[string]any{
			"genre": map[string]any{
				"$eq": "documentary",
			},
			"year": 2019,
		}

		req := &QueryByVectorRequest{
			Vector:          []float64{0.1, 0.2, 0.3},
			TopK:            1,
			Namespace:       "test-namespace",
			Filter:          filter,
			IncludeMetadata: true,
			IncludeValues:   false,
		}

		resp, err := client.QueryByVector(context.Background(), req)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(resp.Matches) != 1 {
			t.Fatalf("expected 1 match, got %d", len(resp.Matches))
		}

		if resp.Matches[0].ID != "rec1" {
			t.Errorf("expected ID 'rec1', got %s", resp.Matches[0].ID)
		}
	})
}
