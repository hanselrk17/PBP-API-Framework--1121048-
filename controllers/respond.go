package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendRespond(c *gin.Context, message string, req interface{}) {
	var response Response
	response.Status = 200
	response.Message = message
	response.Data = req
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func SendRespondDoang(c *gin.Context, message string) {
	var response ResponseDoang
	response.Status = 200
	response.Message = message
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func SendErrorResponse(c *gin.Context, message string) {
	var response ErrorResponse
	response.Status = 400
	response.Message = message
	c.JSON(http.StatusInternalServerError, gin.H{"data": response})
}
