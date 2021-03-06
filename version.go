package version

import (
	"fmt"
	"os"
)

var (
	BuildRevision string
	BuildTime     string
	BuildHost     string
	BuildGopath   string
)

func init() {

	for i := 0; i < len(os.Args); i++ {
		if os.Args[i] == "--build-version" {
			Print()
			break
		}
	}
}

func Print() {
	fmt.Printf("Git Commit Hash: %s\nUTC Build Time : %s\nBuild Host     : %s\n", BuildRevision, BuildTime, BuildHost)
}
