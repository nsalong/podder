package handler

import (
	"context"
	"fmt"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"podder/k8s"
)

func HandlePods(kubeContext string, overridePath string) string {
	kubeConfig, err := k8s.InitConfig(kubeContext, overridePath)
	if err != nil {
		return fmt.Sprintf("Failed to find k8s config, %v \n", err)
	}

	clientSet, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return fmt.Sprintf("error getting kubernetes clientSet: %v \n", err)
	}

	pods, err := clientSet.CoreV1().Pods("").List(context.Background(), v1.ListOptions{})
	if err != nil {
		return fmt.Sprintf("error getting pods: %v \n", err)
	}

	if len(pods.Items) == 0 {
		return fmt.Sprintf("Found no pods, context: %s", kubeContext)
	}

	for _, pod := range pods.Items {
		fmt.Printf("%s | %s | %s | %s | %s | %s \n",
			pod.Namespace,
			pod.Name,
			pod.Status.Phase,
			findLatestImage(pod.Status.ContainerStatuses),
			pod.Spec.Overhead.Memory(),
			pod.Spec.Overhead.Cpu())
	}

	return ""

}

func findLatestImage(containerStatuses []v12.ContainerStatus) string {
	for _, status := range containerStatuses {
		if status.Ready {
			return status.Image
		}
	}

	return "Not Found"
}
