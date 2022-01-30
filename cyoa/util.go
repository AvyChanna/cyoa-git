package cyoa

import (
	"errors"
	"fmt"
	"strings"
)

func ParseTags(tag string) ([]int, error) {
	if tag == "" {
		return nil, errors.New("tag is empty")
	}
	splitTag := strings.Split(tag, ".")
	fmt.Println(splitTag)
	return make([]int, 0), nil
}
