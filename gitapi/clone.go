package gitapi

import (
	"os"
	"os/exec"
)

var EnableLogging = true

func Clone(url string, where string) error {
	err := os.RemoveAll(where)
	if err != nil {
		return err
	}
	
	cloneCmd := exec.Command("git", "clone", "--depth=1", url, where)
	if EnableLogging {
		cloneCmd.Stdout = os.Stdout
		cloneCmd.Stderr = os.Stderr
	}
	
	cloneCmd.Stdin = os.Stdin
	return cloneCmd.Run()
}
