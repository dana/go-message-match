package messagematch

import (
	"fmt"
	//	"github.com/kr/pretty"
	//	"reflect"
)

func Match(message map[string]interface{}, match map[string]interface{}) (bool, error) {
	//pretty.Println(message, match)
	return matchMapMap(message, match)
}

func matchArrayArray(message []interface{}, match []interface{}) (bool, error) {
	// This needs to validate that the length of the array on both sides is the same
	for index, value := range match {
		switch value.(type) {
		case string:
			switch message[index].(type) {
			case string:
				doMatch, _ := matchStringString(message[index].(string), value.(string))
				return doMatch, nil
			default:
				return false, nil
			}
		case int:
			switch message[index].(type) {
			case int:
				doMatch, _ := matchIntInt(message[index].(int), value.(int))
				return doMatch, nil
			default:
				return false, nil
			}
		}
	}
	return false, nil
}
func matchIntString(message int, match string) (bool, error) {
	panic("matchIntString: unimplemented")
}
func matchStringInt(message string, match int) (bool, error) {
	panic("matchStringInt: unimplemented")
}

func matchStringString(message string, match string) (bool, error) {
	doesMatch := message == match
	return doesMatch, nil
}

func matchFloat64Float64(message float64, match float64) (bool, error) {
	doesMatch := message == match
	return doesMatch, nil
}
func matchArrayInt(message []interface{}, match int) (bool, error) {
	for _, value := range message {
		switch value.(type) {
		case int:
			doMatch, _ := matchIntInt(value.(int), match)
			if doMatch {
				return true, nil
			}
		case string:
			doMatch, _ := matchStringInt(value.(string), match)
			if doMatch {
				return true, nil
			}
		default:
			return false, nil
		}
	}
	return false, nil
}
func matchIntInt(message int, match int) (bool, error) {
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
					doMatch, _ := matchStringString(message[key].(string), value.(string))
					return doMatch, nil
				}
				return false, nil
			default:
				return false, nil
			}
		case int:
			switch message[key].(type) {
			case []interface{}:
				if _, ok := message[key]; ok {
					doMatch, _ := matchArrayInt(message[key].([]interface{}), value.(int))
					return doMatch, nil
				}
				return false, nil
			case int:
				if _, ok := message[key]; ok {
					doMatch, _ := matchIntInt(message[key].(int), value.(int))
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
					doMatch, _ := matchFloat64Float64(message[key].(float64), value.(float64))
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
					doMatch, _ := matchArrayArray(message[key].([]interface{}), value.([]interface{}))
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
					doMatch, _ := matchMapMap(message[key].(map[string]interface{}), value.(map[string]interface{}))
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
