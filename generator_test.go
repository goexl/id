package id_test

import (
	"fmt"
	"testing"

	"github.com/goexl/id"
)

func TestSnowflake(test *testing.T) {
	generator := id.New().Snowflake(1).Build()
	for count := 0; count < 100; count++ {
		next := generator.Next()
		fmt.Println(next.Value(), next.String(id.Base76()))
	}
}

func TestAutoincrement(test *testing.T) {
	generator := id.New().Autoincrement().Build()
	for count := 0; count < 100; count++ {
		fmt.Println(generator.Next().String(id.Base76()))
	}
}
