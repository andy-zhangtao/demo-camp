package main

import (
	"github.com/sirupsen/logrus"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "config")
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	info, err := clientset.ServerVersion()
	if err != nil {
		panic(err)
	}

	logrus.Infof("Kubernetes Version: %s", info.String())
}

func Deployments(c *kubernetes.Clientset, namespace string) {
	deploy, err := c.AppsV1beta2().Deployments(namespace).List(v1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, d := range deploy.Items {
		logrus.Infof("name: %s", d.GetName())
	}
}
