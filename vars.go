package log

import "regexp"

var (
	level        Level
	path         = "./logs"
	
	logToFile    = true
	logToMsg     = true

	ansiRegex, _ = regexp.Compile(`(\x9B|\x1B\[)[0-?]*[ -\/]*[@-~]`)
)
