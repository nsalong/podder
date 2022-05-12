package handler

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/rest"
	"podder/k8s"
	"testing"
)

func getHandler() *CmdHandler {
	return &CmdHandler{}
}

func TestCmdHandler_HandleVerify_OK(t *testing.T) {
	handler := getHandler()

	k8s.InitConfig = func(context, path string) (*rest.Config, error) {
		return &rest.Config{}, nil
	}

	handler.HandleVerify()

	assert.Equal(t, nil, handler.Error)
	assert.Equal(t, "Kubernetes config exists! Yay\n", handler.Response)
	assert.Equal(t, true, handler.Finished)
}

func TestCmdHandler_HandleVerify_Failed(t *testing.T) {
	handler := getHandler()
	testError := errors.New("test error")

	k8s.InitConfig = func(context, path string) (*rest.Config, error) {
		return nil, testError
	}

	handler.HandleVerify()

	assert.Equal(t, testError, handler.Error)
	assert.Equal(t, "", handler.Response)
	assert.Equal(t, true, handler.Finished)
}

func TestCmdHandler_HandlePods_OK(t *testing.T) {
	handler := getHandler()

	k8s.GetPods = func(config *rest.Config) (*k8s.PodDetailsList, error) {
		return &k8s.PodDetailsList{}, nil
	}
	k8s.InitConfig = func(context, path string) (*rest.Config, error) {
		return &rest.Config{}, nil
	}

	handler.HandlePods("", "")

	assert.Equal(t, "", handler.Response)
	assert.Equal(t, nil, handler.Error)
	assert.Equal(t, true, handler.Finished)
}

func TestCmdHandler_HandlePods_Failed(t *testing.T) {
	handler := getHandler()
	testError := errors.New("test error")

	k8s.GetPods = func(config *rest.Config) (*k8s.PodDetailsList, error) {
		return nil, testError
	}
	k8s.InitConfig = func(context, path string) (*rest.Config, error) {
		return &rest.Config{}, nil
	}

	handler.HandlePods("", "")

	assert.Equal(t, "", handler.Response)
	assert.Equal(t, testError, handler.Error)
	assert.Equal(t, true, handler.Finished)
}
