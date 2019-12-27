package version

import (
	"fmt"
)

var (
	// BuildRevision describes the revision of the current build.
	BuildRevision = "UNKNOWN"

	// BuildDate describes the date of the current build.
	BuildDate = "UNKNOWN"

	// BuildTime describes the time of the current build.
	BuildTime = "UNKNOWN"

	// BuildUser describes the user that initiliazed the current build.
	BuildUser = "UNKNOWN"

	// BuildMachine describes the name of the machine on which the build was performed.
	BuildMachine = "UNKNOWN"

	// BuildInfo describes user friendly build information.
	BuildInfo = fmt.Sprintf("[%s] ( Date: %s %s ) %s@%s", BuildRevision, BuildDate, BuildTime, BuildUser, BuildMachine)
)
