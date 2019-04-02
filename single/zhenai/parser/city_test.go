package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	contents, err := ioutil.ReadFile("userlist.html")

	if err != nil {
		panic(err)
	}

	result := ParseCity(contents)

	fmt.Printf("%v", result)

}
