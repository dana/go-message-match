package messagematch

import (
	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
	"testing"
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
				"hi": []interface{}{"a", "b"},
			},
		},
	}
	match := map[string]interface{}{
		"Name": "Wednesday",
		"Age":  6,
	}

	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.True(doesMatch)
}

func TestBasicMatchStringAndNotNumber(t *testing.T) {
	assert := assert.New(t)
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
		"Age":  7,
	}

	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.False(doesMatch)
}

func TestBasicNoMessageExist(t *testing.T) {
	assert := assert.New(t)
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
		"NName": "WWednesday",
	}

	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
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
				"hi": []interface{}{"a", "b"},
			},
		},
	}
	match := map[string]interface{}{
		"Name": "WWednesday",
	}

	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.False(doesMatch)
}

func TestBasicMatch(t *testing.T) {
	assert := assert.New(t)
	c := map[string]interface{}{
		"x": "y",
	}
	message := map[string]interface{}{
		"Name": "Wednesday",
		"Age":  6,
		"Parents": map[string]interface{}{
			"bee": "boo",
			"foo": map[string]interface{}{
				"hi": []interface{}{"a", "b", c},
			},
		},
	}
	match := map[string]interface{}{
		"Name": "Wednesday",
	}

	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.True(doesMatch)
	pretty.Print("noop")
}
