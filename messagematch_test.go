package messagematch

import (
	"fmt"
	"testing"
)

//func Match(message map[string]interface{}, match map[string]interface{}) (bool, error) {

func TestBasic(t *testing.T) {
	fmt.Println("in TestBasic")
	message := map[string]interface{}{
		"Name": "Wednesday",
		"Age":  6,
		"Parents": map[string]interface{}{
			"bee": "boo",
			"foo": map[string]interface{}{
				"hi": []string{"a", "b"},
			},
		},
	}
	match := map[string]interface{}{
		"Name": "Wednesday",
	}

	doesMatch, sendErr := Match(message, match)
	if sendErr != nil {
		t.Error(sendErr)
		return
	}
	if !doesMatch {
		t.Error("match succeeded when it should have failed")
		return
	}
}
