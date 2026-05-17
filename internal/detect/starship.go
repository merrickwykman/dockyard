package detect

import (
	"os/exec"
	"runtime"
	"strings"
)

type Result struct {
	Found    bool
	Version  string
	Platform string // "mac", "windows", "linux"
}

func Starship() Result {
	platform := resolvePlatform()
	path, err := exec.LookPath("starship")
	if err != nil || path == "" {
		return Result{Found: false, Platform: platform}
	}
	out, err := exec.Command("starship", "--version").Output()
	version := ""
	if err == nil {
		version = strings.TrimSpace(string(out))
	}
	return Result{Found: true, Version: version, Platform: platform}
}

func resolvePlatform() string {
	switch runtime.GOOS {
	case "windows":
		return "windows"
	case "darwin":
		return "mac"
	default:
		return "linux"
	}
}
