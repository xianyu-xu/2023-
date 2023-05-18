package main

import (
	"context"
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", "path/to/etcdconfig")
	if err != nil {
		panic(err.Error())
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	createPod(clientset)
	deletePod(clientset)
}

func createPod(clientset *kubernetes.Clientset) {
	// Define the pod
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-pod",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "test-container",
					Image: "指定的镜像",
				},
			},
		},
	}

	// Create the pod
	fmt.Println("Creating pod...")
	result, err := clientset.CoreV1().Pods("default").Create(context.Background(), pod, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created pod %q.\n", result.GetObjectMeta().GetName())

	waitForPodRunning(clientset)
}

func waitForPodRunning(clientset *kubernetes.Clientset) {
	// Wait for the pod to be running
	for {
		pod, err := clientset.CoreV1().Pods("default").Get(context.Background(), "test-pod", metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				fmt.Println("Pod not found...")
			} else {
				panic(err)
			}
		} else {
			if pod.Status.Phase == corev1.PodRunning {
				fmt.Println("Pod is running...")
				break
			} else {
				fmt.Println("Pod is not running yet...")
			}
		}
		time.Sleep(5 * time.Second)
	}
}

func deletePod(clientset *kubernetes.Clientset) {
	// Delete the pod
	fmt.Println("Deleting pod...")
	err := clientset.CoreV1().Pods("default").Delete(context.Background(), "test-pod", metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Pod deleted.")
}
