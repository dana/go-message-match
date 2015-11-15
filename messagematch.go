package messagematch

import (
	"fmt"
	"github.com/kr/pretty"
	"reflect"
)

func Match(message map[string]interface{}, match map[string]interface{}) (bool, error) {
	return matchMapMap(message, match)
}

func matchStringString(message string, match string) (bool, error) {
	fmt.Println("In matchStringString")
	doesMatch := message == match
	return doesMatch, nil
}

func matchIntInt(message int, match int) (bool, error) {
	fmt.Println("In matchIntInt")
	doesMatch := message == match
	return doesMatch, nil
}

func matchMapMap(message map[string]interface{}, match map[string]interface{}) (bool, error) {
	for key, value := range match {
		fmt.Println(key, " -> ", reflect.TypeOf(value))
		//		switch matchValueType := value.(type) {
		switch value.(type) {
		case string:
			//			switch messageValueType := message[key].(type) {
			switch message[key].(type) {
			case string:
				if _, ok := message[key]; ok {
					doMatch, _ := matchStringString(value.(string), message[key].(string))
					return doMatch, nil
				}
				return false, nil
			default:
				return false, nil
			}
		case int:
			//			fmt.Println(key, "is int", matchValueType)
			fmt.Println(key, "is int", message[key])
			/*			if _, ok := message[key]; ok {
							doMatch, _ := matchIntInt(value.(int), message[key].(int))
							return doMatch, nil
						}
						return false, nil */
		case float64:
			//			fmt.Println(key, "is float64", matchValueType)
			fmt.Println(key, "is float64", message[key])
		case []interface{}:
			fmt.Println(key, "is an array:")
		case map[string]interface{}:
			pretty.Println(value, "is a map:")
			if _, ok := message[key]; ok {
				doMatch, _ := matchMapMap(value.(map[string]interface{}), message[key].(map[string]interface{}))
				return doMatch, nil
			}
			return false, nil
		default:
			//			fmt.Println(key, "is of a type I don't know how to handle", matchValueType)
			return false, nil
		}
	}
	return false, nil
}

/*
func _match(message map[string]interface{}, match map[string]interface{}) (boolean, error) {
	error := ""
	ret := true
	return ret, error
}
*/
