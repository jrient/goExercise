package main

import "fmt"

func main() {
    str := "Go is a beautiful language!"
    fmt.Printf("The length of str is: %d\n", len(str))
    for pos, char := range str {
        fmt.Printf("Character on position %d is: %c \n", pos, char)
    }
    fmt.Println()
    str2 := "Japanese: 日本語"
    fmt.Printf("The length of str2 is: %d\n", len(str2))
    for pos, char := range str2 {
        fmt.Printf("character %c starts at byte position %d\n", char, pos)
    }
    fmt.Println()
    fmt.Println("index int(rune) rune    char bytes")
    for index, rune := range str2 {
        fmt.Printf("%-2d      %d      %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
    }
}


/*
The length of str is: 27
Character on position 0 is: G 
Character on position 1 is: o 
Character on position 2 is:   
Character on position 3 is: i 
Character on position 4 is: s 
Character on position 5 is:   
Character on position 6 is: a 
Character on position 7 is:   
Character on position 8 is: b 
Character on position 9 is: e 
Character on position 10 is: a 
Character on position 11 is: u 
Character on position 12 is: t 
Character on position 13 is: i 
Character on position 14 is: f 
Character on position 15 is: u 
Character on position 16 is: l 
Character on position 17 is:   
Character on position 18 is: l 
Character on position 19 is: a 
Character on position 20 is: n 
Character on position 21 is: g 
Character on position 22 is: u 
Character on position 23 is: a 
Character on position 24 is: g 
Character on position 25 is: e 
Character on position 26 is: ! 

The length of str2 is: 19
character J starts at byte position 0
character a starts at byte position 1
character p starts at byte position 2
character a starts at byte position 3
character n starts at byte position 4
character e starts at byte position 5
character s starts at byte position 6
character e starts at byte position 7
character : starts at byte position 8
character   starts at byte position 9
character 日 starts at byte position 10
character 本 starts at byte position 13
character 語 starts at byte position 16

index int(rune) rune    char bytes
0       74      U+004A 'J' 4A
1       97      U+0061 'a' 61
2       112      U+0070 'p' 70
3       97      U+0061 'a' 61
4       110      U+006E 'n' 6E
5       101      U+0065 'e' 65
6       115      U+0073 's' 73
7       101      U+0065 'e' 65
8       58      U+003A ':' 3A
9       32      U+0020 ' ' 20
10      26085      U+65E5 '日' E6 97 A5
13      26412      U+672C '本' E6 9C AC
16      35486      U+8A9E '語' E8 AA 9E

*/