package simplentp

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

// Time returns the current time using information from a remote NTP server.
// On error, it returns the local system time.
func Time(host string) {
	ntpTime, err := ntp.Time(host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		fmt.Printf("Local time: %v\n", ntpTime)
		os.Exit(1)
	}

	fmt.Printf("Current time: %v\n", ntpTime)
	os.Exit(0)
}
