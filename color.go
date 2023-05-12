/*
 * Author: DG9Jww
 * Date: 2023/4/25
 */

package color

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

// prefix
const (
	INFO = "[INFO]"
	ERRO = "[ERRO]"
	WARN = "[WARN]"
)

// color
var p = map[string]string{
	INFO: color.CyanString(INFO),
	ERRO: color.RedString(ERRO),
	WARN: color.YellowString(WARN),
}

/*
******
println
******
*/
func PrintlnInfo(a ...any) {
	fmt.Print(p[INFO])
	fmt.Println(a...)
}

func PrintlnError(a ...any) {
	fmt.Print(p[ERRO])
	fmt.Println(a...)
}

func PrintlnWarn(a ...any) {
	fmt.Print(p[WARN])
	fmt.Println(a...)
}

/*
******
printf
******
*/
func PrintfInfo(format string, a ...any) {
	fmt.Print(p[INFO])
	fmt.Printf(format, a...)
}

func PrintfError(format string, a ...any) {
	fmt.Print(p[ERRO])
	fmt.Printf(format, a...)
}

func PrintfWarn(format string, a ...any) {
	fmt.Print(p[WARN])
	fmt.Printf(format, a...)
}

/*
******
logPrintln
******
*/
func LogPrintlnInfo(a ...any) {
	print(p[INFO])
	log.Println(a...)
}

func LogPrintlnError(a ...any) {
	print(p[ERRO])
	log.Println(a...)
}

func LogPrintlnWarn(a ...any) {
	print(p[WARN])
	log.Println(a...)
}

/*
******
logPrintf
******
*/
func LogPrintfInfo(format string, a ...any) {
	print(p[INFO])
	log.Printf(format, a...)
}

func LogPrintfError(format string, a ...any) {
	print(p[ERRO])
	log.Printf(format, a...)
}

func LogPrintfWarn(format string, a ...any) {
	print(p[WARN])
	log.Printf(format, a...)
}

/*
******
test
******
*/
func LogP() {
	warn := log.New(os.Stdout, color.YellowString("[WARN]"), 0)
	warn.Println("test")
}
