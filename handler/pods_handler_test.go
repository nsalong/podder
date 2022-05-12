package handler

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/rest"
	"podder/k8s"
	"testing"
)

func TestHandlePods_OK(t *testing.T) {
	k8s.GetPods = func(config *rest.Config) (*k8s.PodDetailsList, error) {
		return &k8s.PodDetailsList{}, nil
	}

	res, err := HandlePods("", "")

	assert.Equal(t, "", res)
	assert.Equal(t, nil, err)
}

func TestHandlePods_Failed(t *testing.T) {
	testError := errors.New("test error")

	k8s.GetPods = func(config *rest.Config) (*k8s.PodDetailsList, error) {
		return nil, testError
	}

	res, err := HandlePods("", "")

	assert.Equal(t, "", res)
	assert.Equal(t, testError, err)
}