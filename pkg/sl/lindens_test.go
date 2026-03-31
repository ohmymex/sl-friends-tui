package sl_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ohmymex/sl-friends-tui/pkg/sl"
)

func TestFetchLindens_Success(t *testing.T) {
	html, err := os.ReadFile("testdata/lindens.html")
	if err != nil {
		t.Fatal(err)
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/my/widget-linden-dollar.php" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Write(html)
	}))
	defer srv.Close()

	client := sl.NewClient("test-token", sl.WithBaseURL(srv.URL))
	lindens, err := client.FetchLindens(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if lindens.Balance != "4,250" {
		t.Errorf("expected '4,250', got %q", lindens.Balance)
	}
}

func TestFetchLindens_Empty(t *testing.T) {
	html, err := os.ReadFile("testdata/lindens_empty.html")
	if err != nil {
		t.Fatal(err)
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(html)
	}))
	defer srv.Close()

	client := sl.NewClient("test-token", sl.WithBaseURL(srv.URL))
	lindens, err := client.FetchLindens(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if lindens.Balance != "" {
		t.Errorf("expected empty balance, got %q", lindens.Balance)
	}
}

func TestFetchLindens_ServerError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	client := sl.NewClient("bad-token", sl.WithBaseURL(srv.URL))
	_, err := client.FetchLindens(context.Background())
	if err == nil {
		t.Fatal("expected error for 500 response")
	}
}
