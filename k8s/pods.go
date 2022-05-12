package k8s

import (
	"context"
	"errors"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var GetPods = getPods

type PodDetails struct {
	PodName      string
	PodNamespace string
	PodStatus    *v12.PodPhase
	LatestImage  string
	MemoryLimit  *resource.Quantity
	CPULimit     *resource.Quantity
}

func getPods(config *rest.Config) ([]*PodDetails, error) {
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	pods, err := clientSet.CoreV1().Pods("").List(context.Background(), v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	if len(pods.Items) == 0 {
		return nil, errors.New("no pods found")
	}

	var listOfPods []*PodDetails

	for _, pod := range pods.Items {
		podDetails := &PodDetails{
			PodName:      pod.Name,
			PodNamespace: pod.Namespace,
			PodStatus:    &pod.Status.Phase,
			LatestImage:  findLatestImage(pod.Status.ContainerStatuses),
			MemoryLimit:  pod.Spec.Overhead.Memory(),
			CPULimit:     pod.Spec.Overhead.Cpu(),
		}

		listOfPods = append(listOfPods, podDetails)
	}

	return listOfPods, nil
}

func findLatestImage(containerStatuses []v12.ContainerStatus) string {
	for _, status := range containerStatuses {
		if status.Ready {
			return status.Image
		}
	}

	return "Not Found"
}
