package k8s

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Dont expect any pods, but successful run of the pipeline
func TestGetPods(t *testing.T) {
	// we can ignore no pods found error, since there could not be any pods up in testing env
	config, err := initConfig("", "")
	pods, err := GetPods(config)
	if err == nil {
		assert.NotEqual(t, 0, len(pods))
	} else {
		assert.Equal(t, errors.New("no pods found"), err)
	}

}
