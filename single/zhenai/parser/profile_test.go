package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("user.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents)

	fmt.Printf("%v", result)

}
