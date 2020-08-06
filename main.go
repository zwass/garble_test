package main

import (
	"github.com/zwass/garble_test/imported"
)

func main() {
	println("var:", imported.ExportedVar)
	println("private var:", imported.Private())
}
