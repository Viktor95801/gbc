//go:build !windows

package backend

import (
	"bytes"
	"fmt"
	"strings"

	"modernc.org/libqbe"
)

// expects validated target
// eg. `linux-arm64`, `windows-amd64`, etc
func CompileQBE(qbeIR, target string) (string, error) {
	fmt.Println("Calling libqbe on our QBE IR...")

	targetParts := strings.Split(target, "-")
	target = libqbe.DefaultTarget(targetParts[0], targetParts[1])

	var asmBuf bytes.Buffer
	err := libqbe.Main(target, "input.ssa", strings.NewReader(qbeIR), &asmBuf, nil)
	if err != nil {
		return "", fmt.Errorf("libqbe: %w", err)
	}

	return asmBuf.String(), nil
}
