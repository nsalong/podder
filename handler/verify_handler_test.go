package handler

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/rest"
	"podder/k8s"
	"testing"
)

func TestHandleVerify_OK(t *testing.T) {
	k8s.InitConfig = func(context, path string) (*rest.Config, error) {
		return &rest.Config{}, nil
	}

	res, err := HandleVerify()
	assert.Equal(t, nil, err)
	assert.Equal(t, "Kubernetes config exists! Yay\n", res)
}

func TestHandleVerify_Failed(t *testing.T) {
	testError := errors.New("test error")

	k8s.InitConfig = func(context, path string) (*rest.Config, error) {
		return nil, testError
	}

	res, err := HandleVerify()
	assert.Equal(t, testError, err)
	assert.Equal(t, "", res)
}
