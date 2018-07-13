package util

import (
	"testing"
	"fmt"
)

func TestParseWebConf(t *testing.T) {
	result, err := parseWebConf()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

}

func TestSql(t *testing.T) {
	parseWebConf()
	database, name, password := Database()
	dataSource := fmt.Sprintf("%s:%s@/%s?charset=utf8&loc=Asia FShanghai&parseTime=true", name, password, database)
	fmt.Println(dataSource)
}
