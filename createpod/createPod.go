package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func createPod() {

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

	// create a pod defintion
	podDefintion := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "demo-k8s-",
			Namespace:    "default",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "nginx-container",
					Image: "nginx:latest",
				},
			},
		},
	}

	// create a new pod
	newPod, err := podsClient.Create(context.TODO(), podDefintion, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Pod '%s' is created!", newPod.Name)
}
