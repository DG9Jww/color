/*
 * Author: DG9Jww
 * Date: 2023/4/25
 */

package color

import (
	"testing"
)

func TestNewPrefix(t *testing.T) {
	out := NewPrefixWithTime("xx", GREEN)
	ErrorLogger.Println("It's a error message")
	out.Println("nihao")
}
