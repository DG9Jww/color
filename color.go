/*
 * Author: DG9Jww
 * Date: 2023/4/25
 */

package color

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
)

// builtin prefix
const (
	INFO = "INFO"
	ERRO = "ERRO"
	WARN = "WARN"
)

// color
type Color int

const (
	BLUE Color = iota
	RED
	YELLOW
	BLACK
	CYAN
	GREEN
	WHITE

	FgHiRed = iota + 84
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// builtin logger
var (
	WarnLogger  = NewLogger(WARN, YELLOW)
	InfoLogger  = NewLogger(INFO, CYAN)
	ErrorLogger = NewLogger(ERRO, RED)
)

type Logger struct {
	prefix string
	time   string
	logger *log.Logger
}

// 自定义prefix
// 将传入的颜色和对应的染色函数对应上，我第一想法是使用map
// 这里问的gpt，他竟然用数组完成，这是我没想到的，数组索引正好和颜色对应
func NewLogger(prefix string, colour Color) *Logger {
	colors := []func(format string, a ...interface{}) string{
		color.BlueString,
		color.RedString,
		color.YellowString,
		color.BlackString,
		color.CyanString,
		color.GreenString,
		color.WhiteString,
	}

	logger := &Logger{logger: log.New(os.Stdout, "", 0), prefix: colors[colour](prefix)}

	return logger
}

func (l *Logger) setTime() {
	t := time.Now()
	l.time = color.HiGreenString(fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second()))
}

func (l *Logger) Println(v ...any) {
	t := time.Now()
	l.time = color.HiGreenString(fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second()))

	l.logger.SetPrefix(fmt.Sprintf("[%s] [%s] ", l.time, l.prefix))
	l.logger.Println(v...)
}

func (l *Logger) Printf(format string, v ...any) {
	t := time.Now()
	l.time = color.HiGreenString(fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second()))

	l.logger.SetPrefix(fmt.Sprintf("[%s] [%s] ", l.time, l.prefix))
	l.logger.Printf(format, v...)
}
