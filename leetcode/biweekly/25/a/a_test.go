// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	examples := [][]string{
		{
			`[2,3,5,1,3]`, `3`, 
			`[true,true,true,false,true]`,
		},
		{
			`[4,2,1,1,2]`, `1`, 
			`[true,false,false,false,false]`,
		},
		{
			`[12,1,12]`, `10`, 
			`[true,false,true]`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, kidsWithCandies, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-25/problems/kids-with-the-greatest-number-of-candies/
