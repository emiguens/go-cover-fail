package pkg2

import "net/http"

// HasHeader ...
func HasHeader(req *http.Request, header string) bool {
	return req.Header.Get(header) != ""
}
