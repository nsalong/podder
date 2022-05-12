package handler

import (
	"fmt"
	"podder/k8s"
)

func HandlePods(kubeContext string, overridePath string) (string, error) {
	kubeConfig, err := k8s.InitConfig(kubeContext, overridePath)
	if err != nil {
		return "", err
	}

	pods, err := k8s.GetPods(kubeConfig)
	if err != nil {
		return "", err
	}

	return formatPodResponse(pods), nil

}

func formatPodResponse(pods []*k8s.PodDetails) string {
	responseStr := ""

	for _, pod := range pods {
		responseStr = responseStr + fmt.Sprintf("%s | %s | %s | %s | %s | %s \n",
			pod.PodNamespace,
			pod.PodName,
			*pod.PodStatus,
			pod.LatestImage,
			pod.MemoryLimit,
			pod.CPULimit)
	}

	return responseStr
}
