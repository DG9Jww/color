/*
 * Author: DG9Jww
 * Date: 2023/4/25
 */

package color

import "testing"

func TestPrintln(t *testing.T) {
	PrintlnInfo("This Info", 1, 1)
	PrintlnWarn("This Warn", 2, 2)
	PrintlnError("This Error", 3, 3)
}

func TestPrintf(t *testing.T) {
	PrintfInfo("I'm %s,how are you?", "888")
	PrintfError("I'm %s,how are you?", "888")
	PrintfWarn("I'm %s,how are you?", "888")
}

func TestLogPrintln(t *testing.T) {
	LogPrintlnInfo("This Info", 1, 1)
	LogPrintlnError("This error", 1, 1)
	LogPrintlnWarn("This warning", 1, 1)
}

func TestLogPrintf(t *testing.T) {
	LogPrintfInfo("I'm %s,how are you?", "888")
	LogPrintfError("I'm %s,how are you?", "888")
	LogPrintfWarn("I'm %s,how are you?", "888")
}
