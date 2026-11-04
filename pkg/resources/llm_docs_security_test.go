package resources_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/firebolt-db/mcp-server/pkg/resources"
	"github.com/firebolt-db/mcp-server/pkg/tools/tool_docs_llm"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type docsSecurityTransport func(*http.Request) (*http.Response, error)

func (f docsSecurityTransport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

// These tests are deliberately serial: they replace the default transport used
// by the public constructor, and restore it before parallel tests resume.
func docsSecurityUseTransport(t *testing.T, f docsSecurityTransport) {
	t.Helper()
	previous := http.DefaultTransport
	http.DefaultTransport = f
	t.Cleanup(func() { http.DefaultTransport = previous })
}

func docsSecurityResponse(r *http.Request, location string) *http.Response {
	response := &http.Response{
		StatusCode: http.StatusOK, Header: make(http.Header),
		Body: io.NopCloser(strings.NewReader("fixture docs")), Request: r,
	}
	if location != "" {
		response.StatusCode = http.StatusFound
		response.Header.Set("Location", location)
	}
	return response
}

func TestLLMDocsSecurityRejectsUntrustedURLs(t *testing.T) {
	untrusted := []string{
		"https://example.com/private", "http://docs.firebolt.io/article.md",
		"http://127.0.0.1/private", "https://127.0.0.1/private", "https://[::1]/private",
		"http://169.254.169.254/latest/meta-data/",
		"https://docs.firebolt.io:8443/article.md", "https://docs.firebolt.io:/article.md",
		"https://docs.firebolt.io:0443/article.md", "https://user@docs.firebolt.io/article.md",
		"https://docs.firebolt.io@evil.example/article.md", "https://evil.example@docs.firebolt.io/article.md",
		"https://docs.firebolt.io.evil.example/article.md", "https://docs.firebolt.io./article.md",
		"https://docs.firebolt.io%2eevil.example/article.md", "https:docs.firebolt.io/article.md",
		"//docs.firebolt.io/article.md", "https://docs.firebolt.io\\@evil.example/article.md",
	}
	for _, articleURL := range untrusted {
		t.Run(articleURL, func(t *testing.T) {
			requests := 0
			docsSecurityUseTransport(t, func(r *http.Request) (*http.Response, error) {
				requests++
				return docsSecurityResponse(r, ""), nil
			})
			docs := resources.NewLLMDocs()
			_, err := docs.Handler(t.Context(), &mcp.ReadResourceRequest{Params: &mcp.ReadResourceParams{
				Meta: map[string]any{"articleURL": articleURL},
			}})
			if err == nil {
				t.Error("untrusted resource URL succeeded")
			}
			if requests != 0 {
				t.Errorf("untrusted URL reached transport %d times", requests)
			}
		})
	}
}

func TestLLMDocsSecurityRejectsUntrustedRedirects(t *testing.T) {
	for _, location := range []string{
		"https://evil.example/private", "//evil.example/private", "http://docs.firebolt.io/private",
		"http://127.0.0.1/private", "https://[::1]/private", "http://169.254.169.254/latest/meta-data/",
		"https://docs.firebolt.io:8443/private", "https://user@docs.firebolt.io/private",
		"https://docs.firebolt.io:/private", "https://docs.firebolt.io:0443/private",
		"https://docs.firebolt.io.evil.example/private", "https://docs.firebolt.io./private",
		"https:docs.firebolt.io/private",
	} {
		t.Run(location, func(t *testing.T) {
			requests := 0
			docsSecurityUseTransport(t, func(r *http.Request) (*http.Response, error) {
				requests++
				if requests == 1 {
					return docsSecurityResponse(r, location), nil
				}
				return docsSecurityResponse(r, ""), nil
			})
			articleURL := "https://docs.firebolt.io/start.md"
			_, err := resources.NewLLMDocs().FetchLLMDocsResources(t.Context(), &articleURL)
			if err == nil {
				t.Error("untrusted redirect succeeded")
			}
			if requests != 1 {
				t.Errorf("redirect destination reached transport: %d requests", requests)
			}
		})
	}
}

func TestLLMDocsSecurityAllowsDocsAndSameOriginRedirects(t *testing.T) {
	for _, articleURL := range []string{
		"https://docs.firebolt.io/article.md", "https://DOCS.FIREBOLT.IO/article.md",
		"https://docs.firebolt.io:443/article.md",
	} {
		t.Run(articleURL, func(t *testing.T) {
			requests := 0
			docsSecurityUseTransport(t, func(r *http.Request) (*http.Response, error) {
				requests++
				switch requests {
				case 1:
					return docsSecurityResponse(r, "/relative.md"), nil
				case 2:
					return docsSecurityResponse(r, "https://DOCS.FIREBOLT.IO:443/final.md"), nil
				default:
					return docsSecurityResponse(r, ""), nil
				}
			})
			result, err := resources.NewLLMDocs().FetchLLMDocsResources(t.Context(), &articleURL)
			if err != nil {
				t.Fatal(err)
			}
			if requests != 3 || len(result.Contents) != 1 || result.Contents[0].Text != "fixture docs" {
				t.Fatalf("unexpected result: requests=%d result=%+v", requests, result)
			}
		})
	}
}

func TestLLMDocsSecurityRedirectLimit(t *testing.T) {
	requests := 0
	docsSecurityUseTransport(t, func(r *http.Request) (*http.Response, error) {
		requests++
		return docsSecurityResponse(r, "/loop.md"), nil
	})
	articleURL := "https://docs.firebolt.io/loop.md"
	_, err := resources.NewLLMDocs().FetchLLMDocsResources(t.Context(), &articleURL)
	if err == nil || !strings.Contains(err.Error(), "stopped after 10 redirects") {
		t.Fatalf("unexpected error: %v", err)
	}
	if requests != 10 {
		t.Errorf("got %d requests, want default limit of 10", requests)
	}
}

func TestLLMDocsSecurityToolHandlerLocalFixture(t *testing.T) {
	requests := 0
	fixture := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		if r.URL.Path == "/start.md" {
			http.Redirect(w, r, "/final.md", http.StatusFound)
			return
		}
		_, _ = io.WriteString(w, "local fixture docs")
	}))
	defer fixture.Close()
	fixtureURL, err := url.Parse(fixture.URL)
	if err != nil {
		t.Fatal(err)
	}
	transport := &http.Transport{}
	defer transport.CloseIdleConnections()
	docsSecurityUseTransport(t, func(r *http.Request) (*http.Response, error) {
		// Only the test transport routes approved-origin requests to the fixture.
		clone := r.Clone(r.Context())
		clone.URL.Scheme, clone.URL.Host = fixtureURL.Scheme, fixtureURL.Host
		return transport.RoundTrip(clone)
	})
	handler := tool_docs_llm.NewDocsLLM(resources.NewLLMDocs(), true).Handler()
	for _, articleURL := range []string{"", "https://docs.firebolt.io/start.md"} {
		_, result, err := handler(t.Context(), &mcp.CallToolRequest{}, tool_docs_llm.Input{ArticleURL: articleURL})
		if err != nil {
			t.Fatal(err)
		}
		if len(result.Articles) != 1 {
			t.Fatalf("unexpected articles: %+v", result.Articles)
		}
		content, ok := result.Articles[0].(*mcp.TextContent)
		if !ok || content.Text != "local fixture docs" {
			t.Fatalf("unexpected content: %+v", result.Articles[0])
		}
	}
	before := requests
	_, _, err = handler(t.Context(), &mcp.CallToolRequest{}, tool_docs_llm.Input{ArticleURL: fixture.URL + "/secret"})
	if err == nil {
		t.Error("tool handler accepted loopback URL")
	}
	if requests != before {
		t.Error("tool handler's blocked request reached fixture")
	}
	if before != 3 {
		t.Errorf("expected index and two article requests, got %d", before)
	}
}
