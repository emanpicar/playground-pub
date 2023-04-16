package router

import (
	gpuV1 "github.com/emanpicar/playground-pub/kubernetes/gkube-crd/api/v1"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type (
	CRGraphicsCard interface {
		ListGpu(c *gin.Context, in *SharedIn) (*gpuV1.GraphicsCardList, error)
		GetGpu(c *gin.Context, in *SharedIn2) (*gpuV1.GraphicsCard, error)
		CreateGpu(c *gin.Context, in *gpuV1.GraphicsCard) (*gpuV1.GraphicsCard, error)
		DeleteGpu(c *gin.Context, in *SharedIn2) error
	}

	GpuCreateIn struct {
		Namespace string
		Name      string
		Model     string
	}
)

func (s *service) ListGpu(c *gin.Context, in *SharedIn) (*gpuV1.GraphicsCardList, error) {
	gpus := &gpuV1.GraphicsCardList{}
	if err := s.kubeClient.Client().List(c.Request.Context(), gpus, client.InNamespace(in.Namespace)); err != nil {
		return nil, err
	}

	return gpus, nil
}

func (s *service) GetGpu(c *gin.Context, in *SharedIn2) (*gpuV1.GraphicsCard, error) {
	gpu := &gpuV1.GraphicsCard{}
	key := types.NamespacedName{Namespace: in.Namespace, Name: in.ID}
	if err := s.kubeClient.Client().Get(c.Request.Context(), key, gpu); err != nil {
		return nil, err
	}

	return gpu, nil
}

func (s *service) CreateGpu(c *gin.Context, in *gpuV1.GraphicsCard) (*gpuV1.GraphicsCard, error) {
	if err := s.kubeClient.Client().Create(c.Request.Context(), in); err != nil {
		return nil, err
	}

	gpu, err := s.GetGpu(c, &SharedIn2{Namespace: in.Namespace, ID: in.Name})
	if err != nil {
		return nil, err
	}

	return gpu, nil
}

func (s *service) DeleteGpu(c *gin.Context, in *SharedIn2) error {
	err := s.kubeClient.Client().Delete(c.Request.Context(), &gpuV1.GraphicsCard{
		ObjectMeta: v1.ObjectMeta{
			Name:      in.ID,
			Namespace: in.Namespace,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
