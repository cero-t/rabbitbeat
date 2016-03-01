package main

import (
	"os"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/cero-t/rabbitbeat/beater"
)

func main() {
	err := beat.Run("rabbitbeat", "", beater.New())
	if err != nil {
		os.Exit(1)
	}
}
