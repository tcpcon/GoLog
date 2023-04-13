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
