package k8s

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Assumes there is a kubernetes config file in default location
func TestInitConfig(t *testing.T) {
	res, err := InitConfig("", "")
	assert.Equal(t, nil, err)
	assert.NotNil(t, res)
}

// Pass context as a random UUID, should fail to init config
func TestInitConfig_FailsContext(t *testing.T) {
	res, err := InitConfig(uuid.New().String(), "")
	assert.NotNil(t, err)
	assert.Nil(t, res)
}

// Pass path as a random UUID, should fail to init config
func TestInitConfig_FailsPath(t *testing.T) {
	res, err := InitConfig("", uuid.New().String())
	assert.NotNil(t, err)
	assert.Nil(t, res)
}
