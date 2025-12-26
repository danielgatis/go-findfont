//go:build windows

package findfont

import (
	"os"
	"path/filepath"
)

func getFontDirectories() []string {
	return []string{
		filepath.Join(os.Getenv("windir"), "Fonts"),
		filepath.Join(os.Getenv("localappdata"), "Microsoft", "Windows", "Fonts"),
	}
}
