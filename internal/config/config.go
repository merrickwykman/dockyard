package config

import (
	"errors"
	"io"
	"os"
	"path/filepath"
)

func Path() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "starship.toml"), nil
}

func Read() (string, error) {
	p, err := Path()
	if err != nil {
		return "", err
	}
	data, err := os.ReadFile(p)
	if errors.Is(err, os.ErrNotExist) {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func Write(content string) error {
	p, err := Path()
	if err != nil {
		return err
	}
	if err := backup(p); err != nil {
		return err
	}
	return os.WriteFile(p, []byte(content), 0644)
}

func Revert() error {
	p, err := Path()
	if err != nil {
		return err
	}
	bak := p + ".bak"
	if _, err := os.Stat(bak); errors.Is(err, os.ErrNotExist) {
		return errors.New("no backup found — nothing to revert")
	}
	return copyFile(bak, p)
}

func backup(p string) error {
	if _, err := os.Stat(p); errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return copyFile(p, p+".bak")
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}
