package git

import (
	"bytes"
	"os/exec"
)

func GetDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--staged")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return out.String(), nil
}
