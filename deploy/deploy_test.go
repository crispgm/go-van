package deploy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeployer(t *testing.T) {
	rsync := NewDeployer("rsync")
	shell := NewDeployer("shell")
	assert.NotNil(t, rsync)
	assert.Nil(t, shell)
}
