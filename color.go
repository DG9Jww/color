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
	WarnLogger  = NewPrefixWithTime(WARN, YELLOW)
	InfoLogger  = NewPrefixWithTime(INFO, CYAN)
	ErrorLogger = NewPrefixWithTime(ERRO, RED)
)

// 自定义prefix
// 将传入的颜色和对应的染色函数对应上，我第一想法是使用map
// 这里问的gpt，他竟然用数组完成，这是我没想到的，数组索引正好和颜色对应
func NewPrefix(prefix string, colour Color) *log.Logger {
	colors := []func(format string, a ...interface{}) string{
		color.BlueString,
		color.RedString,
		color.YellowString,
		color.BlackString,
		color.CyanString,
		color.GreenString,
		color.WhiteString,
	}
	return log.New(os.Stdout, fmt.Sprintf("[%s]", colors[colour](prefix)), 0)
}

func NewPrefixWithTime(prefix string, colour Color) *log.Logger {
	t := time.Now()
	timeFormat := color.HiGreenString(fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second()))

	colors := []func(format string, a ...interface{}) string{
		color.BlueString,
		color.RedString,
		color.YellowString,
		color.BlackString,
		color.CyanString,
		color.GreenString,
		color.WhiteString,
	}

	return log.New(os.Stdout, fmt.Sprintf("[%s] [%s] ", timeFormat, colors[colour](prefix)), 0)
}
