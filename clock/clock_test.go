package clock

import (
	"testing"
	"time"
)

// Keep at the beginning of the file, since "0" is supposed to be non-existant for sure
func TestGetNonExistant(t *testing.T) {
	c, err := Get("0")
	if err == nil || c != nil {
		t.Error("Get error not set even though element should not exist", err, c)
	}
}

func TestCreateGetClock(t *testing.T) {
	before := time.Now()
	cid := Create("1")
	after := time.Now()
	if cid != "1" {
		t.Error("returned Cid does not equal the correct value 1 : ", cid)
	} else if c, err := Get(cid); err != nil || c == nil {
		t.Error("Get does not return a clock", err, c)
	} else if c.Cid != "1" {
		t.Error("Cid does not equal the required value 1 : ", c.Cid)
	} else if c.Create.Before(before) || c.Create.After(after) {
		t.Error("creation time is not within the correct time bounds ", before, " < ", c.Create, " < ", after)
	}
}

func TestRandomClockid(t *testing.T) {
	cid1 := Create("")
	cid2 := Create("")
	if cid1 == cid2 {
		t.Error("Create does not assign random identifiers", cid1, cid2)
	} else if c2, err := Get(cid2); err != nil || c2 == nil || c2.Cid != cid2 {
		t.Error("Get did not return second clock")
	} else if c1, err := Get(cid1); err != nil || c1 == nil || c1.Cid != cid1 {
		t.Error("Get did not return first clock")
	}
}
