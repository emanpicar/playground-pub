package kubernetes

import (
	"flag"
	"path/filepath"

	gpuV1 "github.com/emanpicar/playground-pub/kubernetes/gkube-crd/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	coreV1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type (
	K8s interface {
		InClusterConfig() error
		OutOfClusterConfig() error
		GetRawClientSet() *kubernetes.Clientset
		CoreV1() coreV1.CoreV1Interface
		Client() client.Client
	}

	service struct {
		clientset *kubernetes.Clientset
		client    client.Client
	}
)

func New() K8s {
	return &service{}
}

func (s *service) InClusterConfig() error {
	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	client, err := s.NewCtrlRuntimeClient(config)
	if err != nil {
		return err
	}

	s.clientset = clientset
	s.client = client

	return nil
}

func (s *service) OutOfClusterConfig() error {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	client, err := s.NewCtrlRuntimeClient(config)
	if err != nil {
		return err
	}

	s.clientset = clientset
	s.client = client

	return nil
}

func (s *service) NewCtrlRuntimeClient(config *rest.Config) (client.Client, error) {
	scheme := runtime.NewScheme()
	if err := gpuV1.AddToScheme(scheme); err != nil {
		return nil, err
	}

	client, err := client.New(config, client.Options{Scheme: scheme})
	if err != nil {
		return nil, err
	}

	return client, nil
}

// hard to mock/test if used directly
func (s *service) GetRawClientSet() *kubernetes.Clientset {
	return s.clientset
}

func (s *service) CoreV1() coreV1.CoreV1Interface {
	return s.clientset.CoreV1()
}

func (s *service) Client() client.Client {
	return s.client
}
