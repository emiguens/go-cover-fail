package pkg1_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emiguens/go-cover-fail/pkg1"
	"github.com/gin-gonic/gin"
)

func TestHasHeader(t *testing.T) {
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	c.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)
	c.Header("x-ttl", "value")

	if hh := pkg1.HasHeader(c, "x-ttl"); hh == true {
		t.Fatal("expected header not to be found")
	}
}
