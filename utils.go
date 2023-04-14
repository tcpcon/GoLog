package log

import "fmt"

func stripAnsiCodes(str string) string {
	return ansiRegex.ReplaceAllLiteralString(str, "")
}

func formatParams(format string, args []any) string {
	var paramsString string

	if len(args) > 0 {
		if params, ok := args[len(args) - 1].(Params); ok {
			for k, v := range params {
				paramsString += fmt.Sprintf("\u001b[1m\u001b[90m%s=\u001b[0m%v ", k, v)
			}

			args = args[:len(args) - 1]
		}
	}

	if len(paramsString) > 0 {
		paramsString = paramsString[:len(paramsString) - 1]
	}

	return fmt.Sprintf("%s %s", fmt.Sprintf(format, args...), paramsString)
}

func (l Log) checkFatal() {
	if l.lvl == LevelFatal {
		os.Exit(1)
	}
}

func (l Log) msg() {
	if logToMsg && l.lvl >= level {
		out := os.Stdout
		if l.lvl == LevelError || l.lvl == LevelFatal {
			out = os.Stderr
		}

		if _, err := out.WriteString(fmt.Sprintf("\u001b[1m\u001b[90m%s %s\u001b[90m ~\u001b[0m %s\n", time.Now().Format("15:04:05"), l.lvl.string(true), l.text)); err != nil {
			panic(err)
		}
	}
}

func (l Log) file() {
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
}
