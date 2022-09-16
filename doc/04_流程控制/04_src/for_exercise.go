package main

import "fmt"

func main() {
	fmt.Println("练习1：")
	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("练习2：")
	for i := 0; ; i++ {
		fmt.Println("Value of i is now:", i)
		if (i > 10) {
			break
		}
	}
	fmt.Println()

	fmt.Println("练习3：")
	// j用来限制无限循环
	j := 0
	for i := 0; i < 3; {
		fmt.Println("Value of i:", i)
		if (j > 10) {
			break
		}
		j++
	}
	fmt.Println()

	fmt.Println("练习4：")
	s := ""
	for ; s != "aaaaa"; {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}
	fmt.Println()

	
	fmt.Println("练习5：")
	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s + "a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
	fmt.Println()
}


/*

练习1：
0 0 0 0 0 

练习2：
Value of i is now: 0
Value of i is now: 1
Value of i is now: 2
Value of i is now: 3
Value of i is now: 4
Value of i is now: 5
Value of i is now: 6
Value of i is now: 7
Value of i is now: 8
Value of i is now: 9
Value of i is now: 10
Value of i is now: 11

练习3：
Value of i: 0
Value of i: 0
Value of i: 0
Value of i: 0
Value of i: 0
Value of i: 0
Value of i: 0
Value of i: 0
Value of i: 0
Value of i: 0
Value of i: 0
Value of i: 0

练习4：
Value of s: 
Value of s: a
Value of s: aa
Value of s: aaa
Value of s: aaaa

练习5：
Value of i, j, s: 0 5 a
Value of i, j, s: 1 6 aa
Value of i, j, s: 2 7 aaa

*/