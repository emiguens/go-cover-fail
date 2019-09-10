package pkg2

import "github.com/gin-gonic/gin"

// HasHeader ...
func HasHeader(c *gin.Context, header string) bool {
	return c.GetHeader(header) != ""
}
