//go:build unix && !darwin

package findfont

import (
	"os"
	"path/filepath"
	"runtime"
)

func getFontDirectories() []string {
	switch runtime.GOOS {
	case "android":
		return []string{"/system/fonts"}
	default:
		directories := getUserFontDirs()
		directories = append(directories, getSystemFontDirs()...)
		return directories
	}
}

func getUserFontDirs() []string {
	if dataPath := os.Getenv("XDG_DATA_HOME"); dataPath != "" {
		return []string{expandUser("~/.fonts/"), filepath.Join(expandUser(dataPath), "fonts")}
	}
	return []string{expandUser("~/.fonts/"), expandUser("~/.local/share/fonts/")}
}

func getSystemFontDirs() []string {
	if dataPaths := os.Getenv("XDG_DATA_DIRS"); dataPaths != "" {
		var paths []string
		for _, dataPath := range filepath.SplitList(dataPaths) {
			paths = append(paths, filepath.Join(expandUser(dataPath), "fonts"))
		}
		return paths
	}
	return []string{"/usr/local/share/fonts/", "/usr/share/fonts/"}
}
