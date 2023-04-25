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
