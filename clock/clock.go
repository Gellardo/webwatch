package clock

import (
	"fmt"
	"time"

	"github.com/gellardo/webwatch/rand"
)

type Clock struct {
	Cid    string
	Create time.Time
}

var clocklist []Clock

func Create(cid string, t time.Time) string {
	if cid == "" {
		cid = rand.RandomString(5)
	}
	clocklist = append(clocklist, Clock{cid, time.Now()})
	fmt.Println("added clock; list: ", clocklist)

	return cid
}

func Get(cid string) *Clock {
	for _, clock := range clocklist {
		if clock.Cid == cid {
			return &clock
		}
	}
	return nil
}
