// +build linux,!windows,!darwin,!js

package dlgs

import (
	"errors"
	"os/exec"
)

// cmdPath looks for supported programs in PATH
func cmdPath() (string, error) {
	cmd, err := exec.LookPath("qarma")
	if err != nil {
		e := err
		cmd, err = exec.LookPath("zenity")
		if err != nil {
			err2 := err
			cmd, err = exec.LookPath("kdialog")
			if err2 != nil {
				return "", errors.New("dlgs: " + err.Error() + "; " + e.Error() + "; " + err2.Error())
			}
		}
	}

	return cmd, err
}
