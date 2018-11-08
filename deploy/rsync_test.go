package deploy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRsyncDeploy(t *testing.T) {
	var rsync RSync
	rsync.Run("../fixtures/testsrc/file.txt", "../fixtures/testoutput")
	assert.FileExists(t, "../fixtures/testoutput/file.txt")
}
