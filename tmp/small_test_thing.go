package main

import (
	"github.com/dana/go-message-match"
	"github.com/kr/pretty"
)

//func Match(message map[string]interface{}, match map[string]interface{}) (bool, error) {
func main() {
	x := []interface{}{"foo"}
	message := map[string]interface{}{
		"Name": "Wednesday",
		"Age":  6,
		"Parents": map[string]interface{}{
			"bee": "boo",
			"foo": map[string]interface{}{
				"hi": []interface{}{"a", "b", x},
			},
		},
	}
	pretty.Println(message)
	match := map[string]interface{}{
		"Name": "Wednesday",
		"Age":  6,
	}

	doesMatch, matchErr := messagematch.Match(message, match)
	pretty.Println(matchErr)
	pretty.Println(doesMatch)
}
