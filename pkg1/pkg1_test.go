package pkg1_test

import (
	"net/http"
	"testing"

	"github.com/emiguens/go-cover-fail/pkg1"
)

func TestHasHeader(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("x-ttl", "value")

	if hh := pkg1.HasHeader(req, "x-ttl"); hh == false {
		t.Fatal("expected header to be found")
	}
}
