/*
 * Author: DG9Jww
 * Date: 2023/4/25
 */

package color

import (
	"testing"
)

func TestNewPrefix(t *testing.T) {
	out := NewLogger("xx", GREEN)
	ErrorLogger.Println("It's a error message")
	ErrorLogger.Printf("It's a %s message", "succ")

	out.Println("nihao")
}
