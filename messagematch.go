package messagematch

import (
	"fmt"
)

func Match(message map[string]interface{}, match map[string]interface{}) (bool, error) {
	fmt.Println(message)
	fmt.Println(match)
	error := error(nil)
	ret := true
	return ret, error
}

/*
func _match(message map[string]interface{}, match map[string]interface{}) (boolean, error) {
	error := ""
	ret := true
	return ret, error
}
*/
