package main

import (
	"gkube-api/kubernetes"
	"gkube-api/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/loopfz/gadgeto/tonic/utils/jujerr"
)

func main() {
	k8s := kubernetes.New()
	if err := k8s.OutOfClusterConfig(); err != nil {
		panic(err)
	}

	route := router.New(k8s)

	tonic.SetErrorHook(jujerr.ErrHook)
	root := gin.Default()

	root.GET("/", tonic.Handler(route.Plain, http.StatusOK))
	root.GET("/pods/:namespace", tonic.Handler(route.ListPods, http.StatusOK))
	root.GET("/pod/:namespace/:id", tonic.Handler(route.GetPodByID, http.StatusOK))
	root.GET("/graphics/:namespace", tonic.Handler(route.ListGpu, http.StatusOK))
	root.GET("/graphics/:namespace/:id", tonic.Handler(route.GetGpu, http.StatusOK))
	root.POST("/graphics", tonic.Handler(route.CreateGpu, http.StatusOK))
	root.DELETE("/graphics/:namespace/:id", tonic.Handler(route.DeleteGpu, http.StatusOK))

	root.Run(":8080")
}
