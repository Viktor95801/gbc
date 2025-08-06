//go:build windows

package backend

import (
	"fmt"
)

func CompileQBE(qbeIR, target string) (string, error) {
	return "", fmt.Errorf("libQBE is not supported on Windows")
}
