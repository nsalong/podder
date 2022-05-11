package handler

import (
	"fmt"
	"podder/k8s"
)

func HandleVerify() string {
	_, err := k8s.InitConfig("", "")
	if err != nil {
		return fmt.Sprintf("error getting Kubernetes config: %v\n", err)
	} else {
		return fmt.Sprintf("Kubernetes config exists! Yay")
	}
}
