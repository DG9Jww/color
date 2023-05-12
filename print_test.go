/*
 * Author: DG9Jww
 * Date: 2023/4/25
 */

package color

import "testing"

func TestNewPrefix(t *testing.T) {
	WarnLogger.Println("你好")
	MyLogger := NewPrefix("[Vuln]", RED)
	MyLogger.Println("Vulnerabilities..........")
}

