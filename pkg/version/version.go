package version

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

// Version is set during build time to the current version of the application.
var Version = "0.0.0-dev"

// GetFullVersion returns the version of the application along with the Go version
func GetFullVersion() string {

	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		return fmt.Sprintf(
			"%s (%s, %s/%s)",
			Version,
			buildInfo.GoVersion,
			runtime.GOOS,
			runtime.GOARCH,
		)
	}

	return Version
}
