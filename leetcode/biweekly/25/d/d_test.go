// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		{
			`[[3,4],[4,5],[5]]`, 
			`1`,
		},
		{
			`[[3,5,1],[3,5]]`, 
			`4`,
		},
		{
			`[[1,2,3,4],[1,2,3,4],[1,2,3,4],[1,2,3,4]]`, 
			`24`,
		},
		{
			`[[1,2,3],[2,3,5,6],[1,3,7,9],[1,8,9],[2,5,7]]`, 
			`111`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, numberWays, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-25/problems/number-of-ways-to-wear-different-hats-to-each-other/
