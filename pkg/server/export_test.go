package server

import "time"

// SetShutdownTimeout overrides how long Serve waits for in-flight HTTP requests after shutdown starts.
func SetShutdownTimeout(s Server, d time.Duration) {
	s.(*serverImpl).shutdownTimeout = d
}
