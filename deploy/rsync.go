package deploy

import (
	"os/exec"
)

// RSync ...
type RSync struct{}

// Run do deployment with rsync
func (r RSync) Run(src, dst string) ([]byte, error) {
	cmd := exec.Command("rsync", "-avl", src, dst)
	output, err := cmd.Output()
	return output, err
}
