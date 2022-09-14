package main
import (
    "fmt"
    "time"
)

var week time.Duration
func main() {
    t := time.Now()
    fmt.Println(t) // e.g. Wed Dec 21 09:52:14 +0100 RST 2011
    fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
    // 21.12.2011
    t = time.Now().UTC()
    fmt.Println(t) // Wed Dec 21 08:52:14 +0000 UTC 2011
    fmt.Println(time.Now()) // Wed Dec 21 09:52:14 +0100 RST 2011
    // calculating times:
    week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
    week_from_now := t.Add(week)
    fmt.Println(week_from_now) // Wed Dec 28 08:52:14 +0000 UTC 2011
    // formatting times:
    fmt.Println(t.Format(time.RFC822)) // 21 Dec 11 0852 UTC
    fmt.Println(t.Format(time.ANSIC)) // Wed Dec 21 08:56:34 2011
    fmt.Println(t.Format("21 Dec 2011 08:52")) // 21 Dec 2011 08:52
    s := t.Format("20060102")
    fmt.Println(t, "=>", s)
    // Wed Dec 21 08:52:14 +0000 UTC 2011 => 20111221
}

/*

2022-09-14 11:28:23.887048768 +0800 CST m=+0.000022394
14.09.2022
2022-09-14 03:28:23.887167421 +0000 UTC
2022-09-14 11:28:23.887172994 +0800 CST m=+0.000146610
2022-09-21 03:28:23.887167421 +0000 UTC
14 Sep 22 03:28 UTC
Wed Sep 14 03:28:23 2022
149 Dec 14099 08:2314
2022-09-14 03:28:23.887167421 +0000 UTC => 20220914

*/