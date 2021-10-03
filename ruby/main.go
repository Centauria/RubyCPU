package main

import (
	"flag"
	"fmt"
	"github.com/Centauria/RubyCPU/engine"
	"log"
	"os"
	"runtime"
)

const (
	name   = "RubyCPU"
	author = "Centauria CHEN"
)

var (
	versionName = "dev"
	buildDate   = "(null)"
	gitRevision = "(null)"
)

func main() {
	flag.Parse()

	var logger = log.New(os.Stderr, "", log.LstdFlags)

	logger.Println(name,
		"VersionName", versionName,
		"BuildDate", buildDate,
		"GitRevision", gitRevision,
		"RuntimeVersion", runtime.Version())

	var ruby = engine.NewEngine(1)
	var cmd string

	ruby.Start()

	for {
		_, err := fmt.Scan(&cmd)
		if err != nil {
			return
		}
		ruby.Input(cmd)
	}
}
