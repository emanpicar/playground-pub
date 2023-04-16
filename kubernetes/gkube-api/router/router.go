package router

import (
	"gkube-api/kubernetes"

	"github.com/gin-gonic/gin"
	coreV1 "k8s.io/api/core/v1"
)

type (
	Router interface {
		Plain(c *gin.Context, in *PlainIn) (*PlainOut, error)
		// GetSAToken(c *gin.Context, in *SharedIn) (*coreV1.PodList, error)
		// GetConfigMapsSecrets(c *gin.Context, in *PlainIn) (*PlainOut, error)
		GetPodByID(c *gin.Context, in *SharedIn2) (*coreV1.Pod, error)
		ListPods(c *gin.Context, in *SharedIn) (*coreV1.PodList, error)

		CRGraphicsCard
		// CreateStockCRD(c *gin.Context, in *PlainIn) (*PlainOut, error)
	}

	service struct {
		kubeClient kubernetes.K8s
	}

	SharedIn struct {
		Namespace string `path:"namespace"`
	}

	SharedIn2 struct {
		Namespace string `path:"namespace"`
		ID        string `path:"id"`
	}

	SAOut struct {
		Message string
		Now     string
		Random  int
	}
)

func New(k8s kubernetes.K8s) Router {
	return &service{kubeClient: k8s}
}
