package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// sendResponse method is in charge of sending back response
// based on provided HTTP Status Code.
func sendResponse(c *gin.Context, statusCode int, output gin.H) {
	acceptHeader := c.Request.Header.Get("Accept")
	payload := output["payload"]

	switch(acceptHeader) {
		case "application/json":
			c.JSON(statusCode, payload)
		case "application/xml":
			c.XML(statusCode, payload)
		default:
			c.JSON(statusCode, payload)
	}
}

// RespondOK method is in charge of sending the output payload
// based on the provided "Accept" request header.
func RespondOK(c *gin.Context, output gin.H) {
	sendResponse(c, http.StatusOK, output)
}

// SendUnprocessableEntity method is in charge of
// sending 422 Code that indicates invalid parameter contents.
func SendUnprocessableEntity(c *gin.Context, output gin.H) {
	sendResponse(c, http.StatusUnprocessableEntity, output)
}