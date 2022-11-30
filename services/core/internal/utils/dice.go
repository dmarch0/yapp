package utils

import (
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

var d_regex *regexp.Regexp

func init() {
	if r, err := regexp.Compile(`^(?P<number>\d+)d(?P<dice>\d+)$`); err == nil {
		d_regex = r
	} else {
		log.Panic("Failed to compile regex")
	}
}

//Pattern of type 3d6+12+2d4
func Roll(pattern string) int {
	return interpret(pattern)
}

func interpret(pattern string) int {
	if before, after, found := strings.Cut(pattern, "+"); found {
		return interpret(before) + interpret(after)
	}
	if before, after, found := strings.Cut(pattern, "-"); found {
		return interpret(before) - interpret(after)
	}

	if d_regex.MatchString(pattern) {
		matches := d_regex.FindStringSubmatch(pattern)
		number_index := d_regex.SubexpIndex("number")
		dice_index := d_regex.SubexpIndex("dice")
		dice_value, dice_error := strconv.Atoi(matches[dice_index])
		if dice_error != nil {
			return 0
		}
		number_value, number_error := strconv.Atoi(matches[number_index])
		if number_error != nil {
			return 0
		}
		result := 0
		for i := 0; i < number_value; i++ {
			result += rand.Intn(dice_value) + 1
		}
		return result
	}

	value, value_error := strconv.Atoi(pattern)
	if value_error != nil {
		return 0
	}
	return value
}
