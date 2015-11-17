package messagematch

import (
	//	"fmt"
	//"github.com/kr/pretty"
	//	"reflect"
	sameish "github.com/dana/go-value-sameish"
)

func Match(message map[string]interface{}, match map[string]interface{}) (bool, error) {
	return matchMapMap(message, match)
}

func matchArrayArray(message []interface{}, match []interface{}) (bool, error) {
	if len(match) > len(message) {
		return false, nil
	}
	for index, value := range match {
		switch value.(type) {
		case string:
			switch message[index].(type) {
			case string:
				didMatch := sameish.Sameish(message[index].(string), value.(string))
				if !didMatch {
					return didMatch, nil
				}
			default:
				return false, nil
			}
		case int:
			switch message[index].(type) {
			case int:
				didMatch := sameish.Sameish(message[index].(int), value.(int))
				if !didMatch {
					return didMatch, nil
				}
			default:
				return false, nil
			}
		}
	}
	return true, nil
}

func matchArrayInt(message []interface{}, match int) (bool, error) {
	for _, value := range message {
		switch value.(type) {
		case int:
			didMatch := sameish.Sameish(value.(int), match)
			if didMatch {
				return true, nil
			}
		case string:
			didMatch := sameish.Sameish(value.(string), match)
			if didMatch {
				return true, nil
			}
		default:
			return false, nil
		}
	}
	return false, nil
}

func matchMapMap(message map[string]interface{}, match map[string]interface{}) (bool, error) {
	for key, value := range match {
		switch value.(type) {
		case string:
			switch message[key].(type) {
			case string:
				if _, ok := message[key]; ok {
					didMatch := sameish.Sameish(message[key].(string), value.(string))
					if !didMatch {
						return didMatch, nil
					}
				}
			default:
				return false, nil
			}
		case int:
			switch message[key].(type) {
			case []interface{}:
				if _, ok := message[key]; ok {
					didMatch, _ := matchArrayInt(message[key].([]interface{}), value.(int))
					if !didMatch {
						return didMatch, nil
					}
				}
			case int:
				if _, ok := message[key]; ok {
					didMatch := sameish.Sameish(message[key].(int), value.(int))
					if !didMatch {
						return didMatch, nil
					}
				}
				/*			case float64:  //this one is not working, maybe a problem in the sameish lib
							if _, ok := message[key]; ok {
								pretty.Println(message[key].(float64), value.(int))
								didMatch := sameish.Sameish(message[key].(float64), value.(int))
								if !didMatch {
									return didMatch, nil
								}
							} */
			case string:
				if _, ok := message[key]; ok {
					didMatch := sameish.Sameish(message[key].(string), value.(int))
					if !didMatch {
						return didMatch, nil
					}
				}
			default:
				return false, nil
			}
		case float64:
			switch message[key].(type) {
			case float64:
				if _, ok := message[key]; ok {
					didMatch := sameish.Sameish(message[key].(float64), value.(float64))
					if !didMatch {
						return didMatch, nil
					}
				}
			default:
				return false, nil
			}
		case []interface{}:
			switch message[key].(type) {
			case []interface{}:
				if _, ok := message[key]; ok {
					didMatch, _ := matchArrayArray(message[key].([]interface{}), value.([]interface{}))
					if !didMatch {
						return didMatch, nil
					}
				}
			default:
				return false, nil
			}
		case map[string]interface{}:
			switch message[key].(type) {
			case map[string]interface{}:
				if _, ok := message[key]; ok {
					didMatch, _ := matchMapMap(message[key].(map[string]interface{}), value.(map[string]interface{}))
					if !didMatch {
						return didMatch, nil
					}
				}
			default:
				return false, nil
			}
		default:
			return false, nil
		}
	}
	return true, nil
}
