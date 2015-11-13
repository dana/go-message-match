package messagematch

import (
	"fmt"
	"reflect"
	"github.com/kr/pretty"
)

func Match(message map[string]interface{}, match map[string]interface{}) (bool, error) {
/*	pretty.Println("message: ", message)
	pretty.Println("match: ", match)
	pretty.Println(reflect.TypeOf(message))
	pretty.Println(reflect.TypeOf(match))
	*/
	for k, v := range match {
		fmt.Println(k, " -> ", reflect.TypeOf(v))
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
			fmt.Println(k, "is string", message[k])
			//need also to check datatype of message[k]
			if _, ok := message[k]; ok {
				doMatch := vv == message[k]
				return doMatch, nil
			}
			return false, nil
		default:
			fmt.Println(k, "is of a type I don't know how to handle", vv)
		}
		return true, nil
	}
	error := error(nil)
	ret := true
	pretty.Println("")
	return ret, error
}

/*
func _match(message map[string]interface{}, match map[string]interface{}) (boolean, error) {
	error := ""
	ret := true
	return ret, error
}
*/
