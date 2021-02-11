// Package buildinfo provides high-level build information injected during
// build.
package buildinfo

var (
	// buildID is the unique build identifier.
	buildID string = "unknown"

	// buildTime which is the time build created
	buildTime string = "unknown"

	// RBACServer provides the build information about the rbac server.
	RBACServer buildinfo
)

// info provides the build information about the key server.
type buildinfo struct{}

// ID returns the build ID.
func (buildinfo) ID() string {
	return buildID
}

// Commit returns the build time
func (buildinfo) Time() string {
	return buildTime
}
