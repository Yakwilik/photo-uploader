package version

import "fmt"

// Version information
var (
	Version = "dev"
	Commit  = "unknown"
	Date    = "unknown"
)

// Info returns version information
func Info() string {
	return fmt.Sprintf("File Uploader %s (commit: %s, built: %s)", Version, Commit, Date)
}

// Print prints version information
func Print() {
	fmt.Println(Info())
}
