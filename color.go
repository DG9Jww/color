/*
 * Author: DG9Jww
 * Date: 2023/4/25
 */

package color

import (

	"github.com/fatih/color"
	"log"
	"os"
)

// builtin prefix
const (
	INFO = "[INFO] "
	ERRO = "[ERRO] "
	WARN = "[WARN] "
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
)

// builtin logger
var (
	WarnLogger  = NewPrefix(WARN, YELLOW)
	InfoLogger  = NewPrefix(INFO, CYAN)
	ErrorLogger = NewPrefix(ERRO, RED)
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
	return log.New(os.Stdout, colors[colour](prefix), 0)
}
