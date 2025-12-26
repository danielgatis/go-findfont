//go:build darwin

package findfont

func getFontDirectories() []string {
	return []string{
		expandUser("~/Library/Fonts/"),
		"/Library/Fonts/",
		"/System/Library/Fonts/",
	}
}
