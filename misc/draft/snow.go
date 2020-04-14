package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://vjudge.net/contest/367107#problem/G
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 6

	s := map[[mx]int]bool{}
	var n int
	var a, b [mx]int
	for Fscan(in, &n); n > 0; n-- {
		for i := range a {
			Fscan(in, &a[i])
		}
		todo := [][mx]int{}
		for j := 0; j < mx; j++ {
			for k, v := range a {
				b[(k+j)%mx] = v
			}
			if s[b] {
				Fprintln(out, "Twin snowflakes found.")
				return
			}
			todo = append(todo, b)
			for k := range a {
				b[(k+j)%mx] = a[mx-1-k]
			}
			if s[b] {
				Fprintln(out, "Twin snowflakes found.")
				return
			}
			todo = append(todo, b)
		}
		for _, v := range todo {
			s[v] = true
		}
	}
	Fprintln(out, "No two snowflakes are alike.")
}

//func main() { run(os.Stdin, os.Stdout) }
