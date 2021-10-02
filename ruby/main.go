package main

import (
	"flag"
	"github.com/Centauria/RubyCPU/engine"
	"github.com/Centauria/RubyCPU/uci"
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

	var ruby = engine.NewEngine()

	var protocol = uci.New(name, author, versionName, ruby,
		[]uci.Option{
			&uci.IntOption{Name: "Hash", Min: 4, Max: 1 << 16, Value: &ruby.Hash},
			&uci.IntOption{Name: "Threads", Min: 1, Max: runtime.NumCPU(), Value: &ruby.Threads},
		},
	)

	uci.RunCli(logger, protocol)
}
