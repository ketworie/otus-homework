package main

import "golang.org/x/example/stringutil"

func main() {
	reversed := stringutil.Reverse("Hello, OTUS!")
	println(reversed)
}
