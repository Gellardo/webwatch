package clock

import (
	"errors"
	"fmt"
	"time"

	"github.com/Gellardo/webwatch/rand"
)

type Clock struct {
	Cid    string
	Create time.Time
}

var clocklist []Clock
var ErrClockNotFound = errors.New("Clock: not found")

// Create a new clock.  If cid is left empty, a random identifier is chosen.
// Returns the cid of the created clock.
func Create(cid string) string {
	if cid == "" {
		cid = rand.RandomString(5)
	}
	clocklist = append(clocklist, Clock{cid, time.Now()})
	fmt.Println("added clock; list: ", clocklist)

	return cid
}

// Get a clock with the corresponding clock identifier.
// The returned pointer will be nil if the clock does not exist
func Get(cid string) (*Clock, error) {
	for _, clock := range clocklist {
		if clock.Cid == cid {
			return &clock, nil
		}
	}
	return nil, ErrClockNotFound
}
