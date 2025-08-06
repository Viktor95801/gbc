package main

import (
	"strings"

	"slices"
)

var targetArch = []string{
	// 64-bit
	"amd64",
	"arm64",
	"riscv64",
	"s390x",
	"ppc64le",
	"mips64le",
	// 32-bit
	"386",
	"arm",
	"mipsle",
}
var targetOs = []string{
	"linux",
	"windows",
}

// true if target is valid
func validateTarget(target string) bool {
	targetParts := strings.Split(target, "-")
	if len(targetParts) != 2 {
		return false
	}

	return slices.Contains(targetOs, targetParts[0]) && slices.Contains(targetArch, targetParts[1])
}

var backends = []string{
	"qbe",
	"c-lang",
}

func validateBackend(backend string) bool {
	return slices.Contains(backends, backend)
}
