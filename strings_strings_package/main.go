package main

import (
	"fmt"
	"strings"
)

func main() {
	p := fmt.Println
	result := strings.Contains("I love Go programming", "lovex")
	p(result)

	result = strings.ContainsAny("success", "xys")
	p(result)

	result = strings.ContainsRune("golang", 'g')
	p(result)

	n := strings.Count("cheese", "e")
	p(n)

	n = strings.Count("Five", "")
	p(n)

	p(strings.ToLower("GO PyTHON jaVa"))
	p(strings.ToUpper("GO PyTHON jaVa"))

	p("go" == "go")
	p("GO" == "go")
	p(strings.ToLower("GO") == strings.ToLower("go")) // not performant
	p(strings.EqualFold("GO", "go"))

	myStr := strings.Repeat("ab", 10)
	p(myStr)

	myStr = strings.Replace("192.168.0.1", ".", ":", 2) // -1 replaces all
	p(myStr)
	myStr = strings.Replace("192.168.0.1", ".", ":", -1) // -1 replaces all
	p(myStr)

	myStr = strings.ReplaceAll("192.168.0.1", ".", ":")
	p(myStr)

	s := strings.Split("a, b, c", ",")
	fmt.Printf("%T\n", s)
	fmt.Printf("%#v\n", s)

	s = strings.Split("Go for Go!", "")
	fmt.Printf("%#v\n", s)

	s = []string{"I", "learn", "Golang"}
	myStr = strings.Join(s, " ")
	p(myStr)

	myStr = "Orange Green \n Blue Yellow"
	fields := strings.Fields(myStr)
	fmt.Printf("%T\n", fields)
	fmt.Printf("%#v\n", fields)

	s1 := strings.TrimSpace("\t Goodbye Windows, Welcome Linux! \n")
	fmt.Printf("%q\n", s1)

	s2 := strings.Trim("...Hello Gophers!!!!?", ".!?")
	p(s2)
}
