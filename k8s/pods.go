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

type PodDetailsList struct {
	ListOfPodDetails []*PodDetails
}

func getPods(config *rest.Config) (*PodDetailsList, error) {
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

	var listOfPods *PodDetailsList

	for _, pod := range pods.Items {
		podDetails := &PodDetails{
			PodName:      pod.Name,
			PodNamespace: pod.Namespace,
			PodStatus:    &pod.Status.Phase,
			LatestImage:  findLatestImage(pod.Status.ContainerStatuses),
			MemoryLimit:  pod.Spec.Overhead.Memory(),
			CPULimit:     pod.Spec.Overhead.Cpu(),
		}

		listOfPods.ListOfPodDetails = append(listOfPods.ListOfPodDetails, podDetails)
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
