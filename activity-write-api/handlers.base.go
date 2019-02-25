package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RespondOK method is in charge of sending the output payload
// based on the provided "Accept" request header.
func respondOK(c *gin.Context, output gin.H) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, output["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, output["payload"])
	default:
		// Respond with JSON
		c.JSON(http.StatusOK, output["payload"])
	}
}
