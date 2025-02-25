package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/retry"
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

	// update a pod
	fmt.Println("Updating pod...")
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {

		// retrive the latest pod
		currentPod, updateErr := podsClient.Get(context.TODO(), "demo-k8s-7p7w9", metav1.GetOptions{})
		if updateErr != nil {
			panic(updateErr.Error())
		}

		// change container image
		currentPod.Spec.Containers[0].Image = "nginx:1.25.4"

		// update pod
		updatedPod, updateErr := podsClient.Update(context.TODO(), currentPod, metav1.UpdateOptions{})
		fmt.Printf("Updated pod: %s", updatedPod.Name)
		return updateErr
	})
	if retryErr != nil {
		panic(retryErr.Error())
	}

}
