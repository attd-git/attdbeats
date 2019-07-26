package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/attd-git/attdbeats/beater"
)

func main() {
	err := beat.Run("attdbeats", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
