package test

import (
	"log"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	startTime := time.Now()
	rc := m.Run()
	log.Printf("Time elapsed = %v", time.Since(startTime).Seconds())

	// rc 0 means we've passed,
	// and CoverMode will be non empty if run with -cover
	if rc == 0 && testing.CoverMode() != "" {
		c := testing.Coverage()
		if c < 0.8 {
			log.Println("Tests passed but coverage failed at", c)
			rc = -1
		} else {
			log.Printf("Code coverage = %v", c)
		}
	}
	os.Exit(rc)
}
