package router

import (
	"github.com/gin-gonic/gin"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *service) GetPodByID(c *gin.Context, in *SharedIn2) (*coreV1.Pod, error) {
	pod, err := s.kubeClient.CoreV1().Pods(in.Namespace).Get(c.Request.Context(), in.ID, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return pod, nil
}

func (s *service) ListPods(c *gin.Context, in *SharedIn) (*coreV1.PodList, error) {
	pods, err := s.kubeClient.CoreV1().Pods(in.Namespace).List(c.Request.Context(), metaV1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return pods, nil
}
