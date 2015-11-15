package messagematch

import (
	"fmt"
	"github.com/kr/pretty"
	//	"reflect"
)

func Match(message map[string]interface{}, match map[string]interface{}) (bool, error) {
	pretty.Println(message, match)
	return matchMapMap(message, match)
}

func matchArrayArray(message []interface{}, match []interface{}) (bool, error) {
	fmt.Println("In matchArrayArray")
	//	doesMatch := message == match
	//	return doesMatch, nil
	for index, value := range match {
		switch value.(type) {
		case string:
			switch message[index].(type) {
			case string:
				doMatch, _ := matchStringString(value.(string), message[index].(string))
				return doMatch, nil
			default:
				return false, nil
			}
		case int:
			switch message[index].(type) {
			case int:
				doMatch, _ := matchIntInt(value.(int), message[index].(int))
				return doMatch, nil
			default:
				return false, nil
			}
		}
	}
	return false, nil
}
func matchStringString(message string, match string) (bool, error) {
	fmt.Println("In matchStringString")
	doesMatch := message == match
	return doesMatch, nil
}

func matchFloat64Float64(message float64, match float64) (bool, error) {
	fmt.Println("In matchFloat64Float64")
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
		//fmt.Println(key, " -> ", reflect.TypeOf(value))
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
			switch message[key].(type) {
			case int:
				if _, ok := message[key]; ok {
					doMatch, _ := matchIntInt(value.(int), message[key].(int))
					return doMatch, nil
				}
				return false, nil
			default:
				return false, nil
			}
		case float64:
			switch message[key].(type) {
			case float64:
				if _, ok := message[key]; ok {
					doMatch, _ := matchFloat64Float64(value.(float64), message[key].(float64))
					return doMatch, nil
				}
				return false, nil
			default:
				return false, nil
			}
		case []interface{}:
			switch message[key].(type) {
			case []interface{}:
				if _, ok := message[key]; ok {
					doMatch, _ := matchArrayArray(value.([]interface{}), message[key].([]interface{}))
					return doMatch, nil
				}
			default:
				return false, nil
			}
			fmt.Println(key, "is an array:")
		case map[string]interface{}:
			switch message[key].(type) {
			case map[string]interface{}:
				if _, ok := message[key]; ok {
					doMatch, _ := matchMapMap(value.(map[string]interface{}), message[key].(map[string]interface{}))
					return doMatch, nil
				}
			default:
				return false, nil
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
