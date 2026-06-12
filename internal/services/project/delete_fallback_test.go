package project

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CloudCIX/gocloudcix"
	"github.com/CloudCIX/gocloudcix/option"
)

func TestProjectIsDeletedOrClosed(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		statusCode int
		body       string
		want       bool
	}{
		{
			name:       "project no longer exists",
			statusCode: http.StatusNotFound,
			body:       `{"error":"not found"}`,
			want:       true,
		},
		{
			name:       "project is closed",
			statusCode: http.StatusOK,
			body:       `{"content":{"closed":true}}`,
			want:       true,
		},
		{
			name:       "project is still open",
			statusCode: http.StatusOK,
			body:       `{"content":{"closed":false}}`,
			want:       false,
		},
		{
			name:       "project response cannot be decoded",
			statusCode: http.StatusOK,
			body:       `{`,
			want:       false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet {
					t.Fatalf("expected GET request, got %s", r.Method)
				}
				if r.URL.Path != "/project/123/" {
					t.Fatalf("expected project path, got %s", r.URL.Path)
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.statusCode)
				_, _ = w.Write([]byte(tt.body))
			}))
			defer server.Close()

			client := gocloudcix.NewClient(
				option.WithBaseURL(server.URL),
				option.WithAPIKey("test-api-key"),
				option.WithMaxRetries(0),
			)
			resource := &ProjectResource{client: &client}

			got := resource.projectIsDeletedOrClosed(context.Background(), 123)
			if got != tt.want {
				t.Fatalf("projectIsDeletedOrClosed() = %v, want %v", got, tt.want)
			}
		})
	}
}
