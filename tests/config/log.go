package config

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "", log.Flags())
