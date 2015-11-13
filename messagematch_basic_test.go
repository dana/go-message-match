package messagematch

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

//func Match(message map[string]interface{}, match map[string]interface{}) (bool, error) {

func TestBasicMatchStringAndNumber(t *testing.T) {
	assert := assert.New(t)
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
		"Age": 6,
	}

	doesMatch, sendErr := Match(message, match)
	assert.Nil(sendErr)
	assert.True(doesMatch)
}

func TestBasicNoMessageExist(t *testing.T) {
	assert := assert.New(t)
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
		"NName": "WWednesday",
	}

	doesMatch, sendErr := Match(message, match)
	assert.Nil(sendErr)
	assert.False(doesMatch)
}
func TestBasicNoMatch(t *testing.T) {
	assert := assert.New(t)
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
		"Name": "WWednesday",
	}

	doesMatch, sendErr := Match(message, match)
	assert.Nil(sendErr)
	assert.False(doesMatch)
}

func TestBasicMatch(t *testing.T) {
	assert := assert.New(t)
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
	assert.Nil(sendErr)
	assert.True(doesMatch)
	fmt.Println("noop")
}
