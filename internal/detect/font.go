package detect

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const nerdFontZipURL = "https://github.com/ryanoasis/nerd-fonts/releases/latest/download/JetBrainsMono.zip"
const nerdFontTTF = "JetBrainsMonoNerdFont-Regular.ttf"

type FontResult struct {
	Supported    bool
	Cached       bool
	Platform     string
	InstallError string // non-empty if a previous install attempt failed
}

func Font() FontResult {
	platform := resolvePlatform()
	if cacheExists() {
		return FontResult{Supported: true, Cached: true, Platform: platform}
	}
	found := checkFonts(platform)
	if found {
		_ = writeCache()
	}
	return FontResult{Supported: found, Platform: platform}
}

func InstallFont(platform string) error {
	switch platform {
	case "mac":
		return exec.Command("brew", "install", "--cask", "font-jetbrains-mono-nerd-font").Run()
	case "linux":
		return installFontLinux()
	default:
		return errors.New("manual install required on Windows")
	}
}

func checkFonts(platform string) bool {
	if platform == "windows" {
		return checkFontsWindows()
	}
	out, err := exec.Command("fc-list").Output()
	if err != nil {
		return false
	}
	return strings.Contains(strings.ToLower(string(out)), "nerd")
}

func checkFontsWindows() bool {
	dirs := []string{
		filepath.Join(os.Getenv("WINDIR"), "Fonts"),
		filepath.Join(os.Getenv("LOCALAPPDATA"), "Microsoft", "Windows", "Fonts"),
	}
	for _, dir := range dirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, e := range entries {
			name := strings.ToLower(e.Name())
			if strings.Contains(name, "nerd") || strings.Contains(name, "nf-") {
				return true
			}
		}
	}
	return false
}

func installFontLinux() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	fontDir := filepath.Join(home, ".local", "share", "fonts")
	if err := os.MkdirAll(fontDir, 0755); err != nil {
		return err
	}
	data, err := downloadBytes(nerdFontZipURL)
	if err != nil {
		return err
	}
	if err := extractFromZip(data, nerdFontTTF, filepath.Join(fontDir, nerdFontTTF)); err != nil {
		return err
	}
	return exec.Command("fc-cache", "-fv").Run()
}

func downloadBytes(url string) ([]byte, error) {
	resp, err := http.Get(url) //nolint:gosec
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func extractFromZip(data []byte, name, dest string) error {
	r, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return err
	}
	for _, f := range r.File {
		if filepath.Base(f.Name) != name {
			continue
		}
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		out, err := os.Create(dest)
		if err != nil {
			return err
		}
		defer out.Close()
		_, err = io.Copy(out, rc)
		return err
	}
	return errors.New("font file not found in zip")
}

func cacheExists() bool {
	dir, err := cacheDir()
	if err != nil {
		return false
	}
	_, err = os.Stat(filepath.Join(dir, "font-ok"))
	return err == nil
}

func writeCache() error {
	dir, err := cacheDir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, "font-ok"), []byte("1"), 0644)
}

func cacheDir() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, "dockyard"), nil
}
