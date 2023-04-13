package log

import (
	"time"
	"strconv"
)

type (
	Level  uint32

	Log    struct {
		text, ts string
		lvl      Level
	}

	Params map[string]any
)

func new(lvl Level, text string) Log {
	return Log{
		text: text,
		ts: strconv.Itoa(int(time.Now().Unix())),
		lvl: lvl,
	}
}

func (lvl Level) string(coloured bool) string {
	var str string

	switch lvl {
	case LevelDebug:
		str = "\u001b[32mDBG"
	case LevelInfo:
		str = "\u001b[32mINF"
	case LevelWarn:
		str = "\u001b[33mWRN"
	case LevelError:
		str = "\u001b[31mERR"
	case LevelFatal:
		str = "\u001b[31mFAT"
	}

	if coloured {
		return str
	}

	return stripAnsiCodes(str)
}
