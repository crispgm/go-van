package deploy

import (
	"os/exec"
)

// RSync ...
type RSync struct{}

// Run do deployment with rsync
func (r RSync) Run(src, dst string, extraArgs []string) ([]byte, error) {
	args := []string{"-avl", src, dst}
	if extraArgs != nil && len(extraArgs) > 0 {
		args = append(args, extraArgs...)
	}
	cmd := exec.Command("rsync", args...)
	output, err := cmd.Output()
	return output, err
}
