# GoLog
Simplistic golang log package.

Functions

```go
func SetLevel(lvl Level)
func SetPath(p string) // set path to log dir for logging to file

func EnableLogToFile()
func DisableLogToFile()
func EnableLogToMsg()
func DisableLogToMsg()

func Debug(format string, args ...any) Log
func Info(format string, args ...any) Log
func Warn(format string, args ...any) Log
func Error(format string, args ...any) Log
func Fatal(format string, args ...any) Log

func (l Log) Msg() Log // log to stdout/stderr
func (l Log) File() Log // log to file
```

## Info

- `Msg()` and `File()` funcs also return the supplied Log struct allowing for clean statements like `log.Info("Info Message").Msg().File()`
- Log functions will format string with args such as `fmt.Printf` in every case except below
- If the last argument to a log function is of type `log.Params (map[string]any)` the output will format `log.Params` and append it to the log message, see example below

```go
log.Debug("Hello, %s", "World", log.Params{"time": "3pm", "afternoon": true}).Msg()
```
![example_1](./assets/example_1.png)
