package main

import (
	"github.com/alecaivazis/survey"
	"github.com/alecaivazis/survey/tests/util"
)

var value = ""

var table = []TestUtil.TestTableEntry{
	{
		"standard", &survey.Password{"Please type your password:"}, &value,
	},
	{
		"please make sure paste works", &survey.Password{"Please paste your password:"}, &value,
	},
}

func main() {
	TestUtil.RunTable(table)
}
