package gitapi

import (
	"os"
	"os/exec"
)

func Clone(url string, where string) error {
	err := os.MkdirAll(where, 0777)
	if err != nil {
		return err
	}

	cloneCmd := exec.Command("git", "clone", "--depth=1", url, where)
	cloneCmd.Stdout = os.Stdout
	cloneCmd.Stderr = os.Stderr
	cloneCmd.Stdin = os.Stdin
	return cloneCmd.Run()
}
