package util

import (
	"testing"
	"fmt"
)

func TestParseWebConf(t *testing.T) {
	result, err := ParseWebConf()
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

}
