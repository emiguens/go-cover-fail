package pkg1_test

import (
	"net/http"
	"testing"

	"github.com/emiguens/go-cover-fail/pkg1"
)

func TestHasHeader(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("x-ttl", "value")

	if hh := pkg1.HasHeader(req, "x-ttl"); hh == true {
		t.Fatal("expected header not to be found")
	}
}
