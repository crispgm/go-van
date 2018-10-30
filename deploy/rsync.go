package deploy

import (
	"fmt"
	"os/exec"
)

// RSync ...
type RSync struct{}

// Run do deployment with rsync
func (r RSync) Run(src, dst Path) error {
	cmd := exec.Command("rsync", "-avl", string(src), string(dst))
	output, err := cmd.Output()
	fmt.Println(output)
	return err
}
