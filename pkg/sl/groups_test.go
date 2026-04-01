package sl_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ohmymex/sl-friends-tui/pkg/sl"
)

func TestFetchGroups_Success(t *testing.T) {
	html, err := os.ReadFile("testdata/groups.html")
	if err != nil {
		t.Fatal(err)
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/my/widget-groups.php" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Write(html)
	}))
	defer srv.Close()

	client := sl.NewClient("test-token", sl.WithBaseURL(srv.URL))
	groups, err := client.FetchGroups(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(groups) != 3 {
		t.Fatalf("expected 3 groups, got %d", len(groups))
	}
	if groups[0].Name != "Builders United" {
		t.Errorf("expected 'Builders United', got %q", groups[0].Name)
	}
	if groups[0].MemberCount != "42" {
		t.Errorf("expected member count '42', got %q", groups[0].MemberCount)
	}
	if groups[2].Name != "Fashion Circle" {
		t.Errorf("expected 'Fashion Circle', got %q", groups[2].Name)
	}
	if groups[2].MemberCount != "15" {
		t.Errorf("expected member count '15', got %q", groups[2].MemberCount)
	}
}

func TestFetchGroups_Empty(t *testing.T) {
	html, err := os.ReadFile("testdata/groups_empty.html")
	if err != nil {
		t.Fatal(err)
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(html)
	}))
	defer srv.Close()

	client := sl.NewClient("test-token", sl.WithBaseURL(srv.URL))
	groups, err := client.FetchGroups(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(groups) != 0 {
		t.Fatalf("expected 0 groups, got %d", len(groups))
	}
}

func TestFetchGroups_ServerError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	}))
	defer srv.Close()

	client := sl.NewClient("bad-token", sl.WithBaseURL(srv.URL))
	_, err := client.FetchGroups(context.Background())
	if err == nil {
		t.Fatal("expected error for 403 response")
	}
}
