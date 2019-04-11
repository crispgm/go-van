package deploy

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRsyncDeploy(t *testing.T) {
	var rsync RSync
	os.Remove("../fixtures/testoutput/file.txt")
	os.Remove("../fixtures/testoutput/file_excluded.txt")
	rsync.Run("../fixtures/testsrc/file.txt", "../fixtures/testoutput", nil)
	assert.FileExists(t, "../fixtures/testoutput/file.txt")
	os.Remove("../fixtures/testoutput/file.txt")
	os.Remove("../fixtures/testoutput/file_excluded.txt")
}

func TestRsyncDeployWithArgs(t *testing.T) {
	var rsync RSync
	os.Remove("../fixtures/testoutput/file.txt")
	os.Remove("../fixtures/testoutput/file_excluded.txt")
	rsync.Run("../fixtures/testsrc/file.txt", "../fixtures/testoutput", []string{"--exclude=file_excluded.txt"})
	assert.FileExists(t, "../fixtures/testoutput/file.txt")
	_, err := os.Stat("../fixtures/testoutput/file_excluded.txt")
	assert.True(t, os.IsNotExist(err))
	os.Remove("../fixtures/testoutput/file.txt")
}
