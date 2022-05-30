package handlers

import "github.com/gin-gonic/gin"

type Handler[T any] interface {
	Run(context *gin.Context)
}
