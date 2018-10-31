package deploy

import (
	"log"
	"os/exec"
)

// RSync ...
type RSync struct{}

// Run do deployment with rsync
func (r RSync) Run(src, dst string) error {
	cmd := exec.Command("rsync", "-avl", src, dst)
	output, err := cmd.Output()
	log.Println(string(output))
	return err
}
