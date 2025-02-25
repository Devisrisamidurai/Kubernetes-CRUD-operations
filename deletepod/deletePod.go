package main

import (
	"context"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	// get kubeconfig
	home, _ := os.UserHomeDir()
	kubeConfigPath := filepath.Join(home, ".kube/config")

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		panic(err.Error())
	}

	// create a new client
	client := kubernetes.NewForConfigOrDie(config)

	// define the namespace
	namespace := "default"

	// define the pods client (easy for later use)
	podsClient := client.CoreV1().Pods(namespace)

	// delete a pod
	deleteErr := podsClient.Delete(context.TODO(), "demo-k8s-7p7w9", metav1.DeleteOptions{})
	if deleteErr != nil {
		panic(deleteErr.Error())
	}

}
