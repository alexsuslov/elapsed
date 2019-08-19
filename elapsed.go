package elapsed

import (
	"fmt"
	"log"
	"time"
)

// Log log elapsed time
func Log(start time.Time, params ...interface{}) {
	params = append(params, time.Since(start).String())
	log.Print("Elapsed", params)
}

// Fmt Fmt.Print elapsed time
func Fmt(start time.Time, params ...interface{}) {
	params = append(params, time.Since(start).String())
	fmt.Print("Elapsed", params)
}
