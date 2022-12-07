package id_test

import (
	"fmt"
	"github.com/goexl/id"
	"testing"
)

func TestSnowflake(test *testing.T) {
	generator := id.New().Snowflake(1).Build()
	for count := 0; count < 100; count++ {
		fmt.Println(generator.Next().String(id.Base76()))
	}
}

func TestAutoincrement(test *testing.T) {
	generator := id.New().Autoincrement().Build()
	for count := 0; count < 100; count++ {
		fmt.Println(generator.Next().String(id.Base76()))
	}
}
