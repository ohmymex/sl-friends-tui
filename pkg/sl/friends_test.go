package sl_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ohmymex/sl-friends-tui/pkg/sl"
)

func TestFetchFriends_Online(t *testing.T) {
	html, err := os.ReadFile("testdata/friends_online.html")
	if err != nil {
		t.Fatal(err)
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/my/widget-friends.php" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		cookie, err := r.Cookie("session-token")
		if err != nil || cookie.Value != "test-token" {
			t.Errorf("expected session-token cookie, got err=%v", err)
		}
		w.Write(html)
	}))
	defer srv.Close()

	client := sl.NewClient("test-token", sl.WithBaseURL(srv.URL))
	friends, err := client.FetchFriends(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(friends) != 3 {
		t.Fatalf("expected 3 friends, got %d", len(friends))
	}
	if friends[0].DisplayName != "Alice Wonder" {
		t.Errorf("expected 'Alice Wonder', got %q", friends[0].DisplayName)
	}
	if friends[0].InternalName != "alice.doe" {
		t.Errorf("expected 'alice.doe', got %q", friends[0].InternalName)
	}
	if !friends[0].Online {
		t.Error("expected Alice to be online")
	}
	if friends[2].DisplayName != "Charlie Rose" {
		t.Errorf("expected 'Charlie Rose', got %q", friends[2].DisplayName)
	}
	if friends[2].Online {
		t.Error("expected Charlie to be offline")
	}
}

func TestFetchFriends_Empty(t *testing.T) {
	html, err := os.ReadFile("testdata/friends_empty.html")
	if err != nil {
		t.Fatal(err)
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(html)
	}))
	defer srv.Close()

	client := sl.NewClient("test-token", sl.WithBaseURL(srv.URL))
	friends, err := client.FetchFriends(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(friends) != 0 {
		t.Fatalf("expected 0 friends, got %d", len(friends))
	}
}

func TestFetchFriends_ServerError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	}))
	defer srv.Close()

	client := sl.NewClient("bad-token", sl.WithBaseURL(srv.URL))
	_, err := client.FetchFriends(context.Background())
	if err == nil {
		t.Fatal("expected error for 403 response")
	}
}
