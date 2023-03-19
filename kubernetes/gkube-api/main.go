package main

import (
	"gkube-api/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/loopfz/gadgeto/tonic/utils/jujerr"
)

func main() {
	route := router.New()

	tonic.SetErrorHook(jujerr.ErrHook)
	root := gin.Default()

	root.GET("/", tonic.Handler(route.Plain, http.StatusOK))

	root.Run(":8080")
}
