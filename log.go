package log

import (
	"runtime"
	"strconv"
	"strings"
	"time"
	"fmt"
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

func (l Log) Msg() Log {
	if logToMsg && l.lvl >= level {
		out := os.Stdout
		if l.lvl == LevelError || l.lvl == LevelFatal {
			out = os.Stderr
		}

		if _, err := out.WriteString(fmt.Sprintf("\u001b[1m\u001b[90m%s %s\u001b[90m ~\u001b[0m %s\n", time.Now().Format("15:04:05"), l.lvl.string(true), l.text)); err != nil {
			panic(err)
		}
	}

	if l.lvl == LevelFatal {
		os.Exit(1)
	}

	return l
}

func (l Log) File() Log {
	if logToFile {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			panic(err)
		}
		
		f, err := os.OpenFile(fmt.Sprintf("%s/%s", path, strings.ToLower(l.lvl.string(false)) + ".log"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
	
		if _, err := f.WriteString(fmt.Sprintf("%s %s", time.Now().Format("2006-01-02 15:04:05"), stripAnsiCodes(l.text)) + "\n"); err != nil {
			panic(err)
		}
	}
	
	if l.lvl == LevelFatal {
		os.Exit(1)
	}

	return l
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
