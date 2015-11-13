package main

import (
	"github.com/kr/pretty"
	"github.com/dana/go-message-match"
)

//func Match(message map[string]interface{}, match map[string]interface{}) (bool, error) {
func main() {
	message := map[string]interface{}{
		"Name": "Wednesday",
		"Age":  6,
		"Parents": map[string]interface{}{
			"bee": "boo",
			"foo": map[string]interface{}{
				"hi": []interface{}{"a", "b"},
			},
		},
	}
	match := map[string]interface{}{
		"Name": "Wednesday",
		"Age":  6,
	}

	doesMatch, matchErr := messagematch.Match(message, match)
	pretty.Print(matchErr)
	pretty.Print(doesMatch)
}

