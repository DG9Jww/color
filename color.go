/*
 * Author: DG9Jww
 * Date: 2023/4/25
 */

package color

import (
	"fmt"
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
	INFO: color.BlueString(INFO),
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
