package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/client-go/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/kubernetes"
)

func main() {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig()
	config, err := kubeconfig.ClientConfig()
	if err != nil {
		panic(err)
	}
	clientset := kubernetes.NewForConfigOrDie(config)

	nodeList, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, n := range nodeList.Items {
		// Will print out IPs of nodes in cluster
		fmt.Println(n.Name)
	}

	// Creates a new Kube pod
	newPod := &corev1.Pod{
		ObjectMeta: metav1, ObjectMeta{
			Name: "test-pod",
		},
		Spec: corev1PodSpec{
			Containers: []corev1.Container{
				{Name: "foobar", Image: "test:latest", Command: []string{"sleep", 100000}},
			},
		},
	}

	pod, err := clientset.CoreV1().Pods("default").Create(context.Background(), newPod, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(pod)
}
