package handler

import (
	"podder/k8s"
)

func HandleVerify() (string, error) {
	_, err := k8s.InitConfig("", "")
	if err != nil {
		return "", err
	} else {
		return "Kubernetes config exists! Yay\n", nil
	}
}
