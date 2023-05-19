package toolkit

import (
	"fmt"
	"testing"
)

func TestStringPurgerForMatch(t *testing.T) {
	t.Helper()
	input := "123一二三， ；)(哈罗abcDEF"
	output := StringPurgerForMatch(input)
	fmt.Printf("【input】: %s\n【output】: %s\n", input, output)
}
