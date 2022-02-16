package version

var version string

// Version returns the global version.
func Version() string {
	if version == "" {
		return "indev"
	}

	return version
}
