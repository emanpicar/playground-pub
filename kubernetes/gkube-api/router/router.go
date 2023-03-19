package router

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
)

type (
	Router interface {
		Init()

		Plain(c *gin.Context, in *PlainIn) (*PlainOut, error)
		// GetSAToken(c *gin.Context, in *PlainIn) (*PlainOut, error)
		// GetConfigMapsSecrets(c *gin.Context, in *PlainIn) (*PlainOut, error)
		// GetPodByID(c *gin.Context, in *PlainIn) (*PlainOut, error)
		// ListPods(c *gin.Context, in *PlainIn) (*PlainOut, error)
		// GetStockCRD(c *gin.Context, in *PlainIn) (*PlainOut, error)
		// ListStockCRD(c *gin.Context, in *PlainIn) (*PlainOut, error)
		// CreateStockCRD(c *gin.Context, in *PlainIn) (*PlainOut, error)
	}

	service struct{}

	PlainIn  struct{}
	PlainOut struct {
		Message string
		Now     string
		Random  int
	}
)

func New() Router {
	return &service{}
}

func (s *service) Init() {

}

func (s *service) Plain(c *gin.Context, in *PlainIn) (*PlainOut, error) {
	randResult := rand.Intn(505)
	switch randResult {
	case http.StatusBadRequest:
		return nil, errors.NewBadRequest(nil, "so lucky - bad request!!")
	case http.StatusForbidden:
		return nil, errors.NewForbidden(nil, "so lucky - forbidden!!")
	}

	return &PlainOut{
		Message: "Hello from gkube-api",
		Now:     time.Now().Format(time.Layout),
		Random:  randResult,
	}, nil
}
