package main

func main() {
    println("In main before calling greeting")
    greeting()
    println("In main after calling greeting")
}

func greeting() {
    println("In greeting: Hi!!!!!")
}

/*

In main before calling greeting
In greeting: Hi!!!!!
In main after calling greeting

*/