package sl_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ohmymex/sl-friends-tui/pkg/sl"
)

func TestFetchAccount_Success(t *testing.T) {
	html, err := os.ReadFile("testdata/account.html")
	if err != nil {
		t.Fatal(err)
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/my/account/index.php" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Write(html)
	}))
	defer srv.Close()

	client := sl.NewClient("test-token", sl.WithBaseURL(srv.URL))
	account, err := client.FetchAccount(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if account.Username != "testuser" {
		t.Errorf("expected username 'testuser', got %q", account.Username)
	}
	if account.Plan != "Premium Membership" {
		t.Errorf("expected plan 'Premium Membership', got %q", account.Plan)
	}
	if account.Status != "Active" {
		t.Errorf("expected status 'Active', got %q", account.Status)
	}
	if account.Country != "indonesia" {
		t.Errorf("expected country 'indonesia', got %q", account.Country)
	}
	if account.LBalance != "L$ 4,250" {
		t.Errorf("expected L$ balance 'L$ 4,250', got %q", account.LBalance)
	}
	if account.USDBalance != "US $12.50" {
		t.Errorf("expected USD balance 'US $12.50', got %q", account.USDBalance)
	}
	if account.LandCurrent != "512 square meters" {
		t.Errorf("expected land '512 square meters', got %q", account.LandCurrent)
	}
}

func TestFetchAccount_ServerError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	}))
	defer srv.Close()

	client := sl.NewClient("bad-token", sl.WithBaseURL(srv.URL))
	_, err := client.FetchAccount(context.Background())
	if err == nil {
		t.Fatal("expected error for 403 response")
	}
}
