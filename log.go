package log

import (
	"runtime"
	"strconv"
	"os"

	"golang.org/x/sys/windows"
)

const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

func init() {
	if runtime.GOOS == "windows" {
		var (
			stdout = windows.Handle(os.Stdout.Fd())
			mode   uint32
		)

		windows.GetConsoleMode(stdout, &mode)
		windows.SetConsoleMode(stdout, mode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
	}
}

func EnableLogToFile() {
	logToFile = true
}

func DisableLogToFile() {
	logToFile = false
}

func EnableLogToMsg() {
	logToMsg = true
}

func DisableLogToMsg() {
	logToMsg = false
}

func SetLevel(lvl Level) {
	if !(lvl <= LevelFatal) {
		panic("level " + strconv.Itoa(int(lvl)) + " is out of bounds")
	}

	level = lvl
}

func SetPath(p string) {
	path = p
}

func (l Log) Msg() {
	l.msg()
	l.checkFatal()
}

func (l Log) File() {
	l.file()
	l.checkFatal()
}

func (l Log) Full() {
	l.msg()
	l.file()
	l.checkFatal()
}

func Debug(format string, args ...any) Log {
	return new(LevelDebug, formatParams(format, args))
}

func Info(format string, args ...any) Log {
	return new(LevelInfo, formatParams(format, args))
}

func Warn(format string, args ...any) Log {
	return new(LevelWarn, formatParams(format, args))
}

func Error(format string, args ...any) Log {
	return new(LevelError, formatParams(format, args))
}

func Fatal(format string, args ...any) Log {
	return new(LevelFatal, formatParams(format, args))
}
