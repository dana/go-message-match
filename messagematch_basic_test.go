package messagematch

import (
	//	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayFullMatchLengthMismatchMatchLonger(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"a": []interface{}{1, 2, 3},
	}
	match := map[string]interface{}{
		"a": []interface{}{1, 2, 3, 4},
	}
	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.False(doesMatch)
}
func TestArrayFullMisMatch(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"a": []interface{}{1, 2, 4},
	}
	match := map[string]interface{}{
		"a": []interface{}{1, 2, 3},
	}
	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.False(doesMatch)
}
func TestArrayFullMatchLengthMismatchMessageLonger(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"a": []interface{}{1, 2, 3, 4},
	}
	match := map[string]interface{}{
		"a": []interface{}{1, 2, 3},
	}
	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.True(doesMatch)
}
func TestStringMap(t *testing.T) {
	assert := assert.New(t)
	internalMap := map[string]interface{}{
		"other": "side",
	}
	message := map[string]interface{}{
		"a": internalMap,
	}
	match := map[string]interface{}{
		"a": "b",
	}
	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.False(doesMatch)
}
func TestArrayContainsMisMatchIntString(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"a": []interface{}{"z", 2, 3},
	}
	match := map[string]interface{}{
		"a": 2,
	}
	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.True(doesMatch)
}

//ported from https://github.com/dana/perl-Message-Match/blob/master/t/basic.t
//not nested
func TestSimplestPossible(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"a": "b",
	}
	match := map[string]interface{}{
		"a": "b",
	}
	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.True(doesMatch)
}
func TestExtraStuff(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"a": "b",
		"c": "d",
	}
	match := map[string]interface{}{
		"a": "b",
	}
	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.True(doesMatch)
}
func TestSimpleMiss(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"a": "b",
		"c": "d",
	}
	match := map[string]interface{}{
		"e": "f",
	}
	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.False(doesMatch)
}
func TestSimplestPossibleMultiMatchRequired(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"a": "b",
		"c": "d",
	}
	match := map[string]interface{}{
		"a": "b",
		"c": "d",
	}
	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.True(doesMatch)
}

//nested
func TestSimplestNested(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"x": map[string]interface{}{
			"y": "z",
		},
	}
	match := map[string]interface{}{
		"x": map[string]interface{}{
			"y": "z",
		},
	}
	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.True(doesMatch)
}
func TestSimplestNestedWithExtraStuff(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"a": "b",
		"x": map[string]interface{}{
			"y": "z",
		},
	}
	match := map[string]interface{}{
		"x": map[string]interface{}{
			"y": "z",
		},
	}
	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.True(doesMatch)
}
func TestMultipleMatchesRequiredNested(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"a": "b",
		"x": map[string]interface{}{
			"y": "z",
		},
	}
	match := map[string]interface{}{
		"x": map[string]interface{}{
			"y": "z",
		},
		"a": "b",
	}
	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.True(doesMatch)
}

//array in message, scalar in match: checks membership
func TestArrayContains(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"a": []interface{}{1, 2, 3},
	}
	match := map[string]interface{}{
		"a": 2,
	}
	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.True(doesMatch)
}
func TestArrayDoesNotContain(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"a": []interface{}{1, 2, 3},
	}
	match := map[string]interface{}{
		"a": 5,
	}
	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.False(doesMatch)
}

//array on both sides: full recursion
func TestArrayFullMatch(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"a": []interface{}{1, 2, 3},
	}
	match := map[string]interface{}{
		"a": []interface{}{1, 2, 3},
	}
	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.True(doesMatch)
}
func TestNestedArrayFullMatch(t *testing.T) {
	assert := assert.New(t)
	firstSub := map[string]interface{}{
		"a": "b",
	}
	secondSub := map[string]interface{}{
		"a": "b",
	}
	message := map[string]interface{}{
		"a": []interface{}{firstSub, 2, 3},
	}
	match := map[string]interface{}{
		"a": []interface{}{secondSub, 2, 3},
	}
	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.True(doesMatch)
}

//regex
func SkipTestSimplestRegex(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"a": "forefoot",
	}
	match := map[string]interface{}{
		"a": " special/foo/",
	}
	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.True(doesMatch)
}
func SkipTestSimplestRegexFailure(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"a": "forefoot",
	}
	match := map[string]interface{}{
		"a": " special/smurf/",
	}
	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.False(doesMatch)
}
func SkipRegexFailureForCaseSensitivity(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"a": "forefoot",
	}
	match := map[string]interface{}{
		"a": " special/FOO/",
	}
	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.False(doesMatch)
}
func SkipRegexPassForCaseSensitivity(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"a": "forefoot",
	}
	match := map[string]interface{}{
		"a": " special/FOO/i",
	}
	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.True(doesMatch)
}

/* Initial set */
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
func SkipTestBasicIntFloat(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"a": 1.0,
	}
	match := map[string]interface{}{
		"a": 1,
	}
	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.True(doesMatch)
}
func TestBasicIntStringFloat(t *testing.T) {
	assert := assert.New(t)
	message := map[string]interface{}{
		"a": "1.0",
	}
	match := map[string]interface{}{
		"a": 1,
	}
	doesMatch, matchErr := Match(message, match)
	assert.Nil(matchErr)
	assert.True(doesMatch)
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
}
