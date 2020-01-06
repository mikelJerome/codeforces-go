package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol101628K(t *testing.T) {
	// just copy from website
	rawText := `
3
aaa
ab
acc
10
2 1 3 abb
2 3 3 abb
2 1 1 aaa
3 1 3 a
3 1 3 ac
3 3 3 b
3 3 3 ac
3 1 3 e
1 1 eae
3 1 3 e
outputCopy
Y
N
Y
Y
Y
N
Y
N
Y`
	testutil.AssertEqualCase(t, rawText, 0, Sol101628K)
}
