package main

import (
	"flag"
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

	logger.Println(name)
	logger.Println("Author:", author)
	logger.Println("VersionName:", versionName)
	logger.Println("BuildDate:", buildDate)
	logger.Println("GitRevision:", gitRevision)
	logger.Println("RuntimeVersion:", runtime.Version())

	var ruby = uci.ProtocolUCI{}

	//ruby.Start()

	uci.RunCli(logger, ruby)
}
