package id_test

import (
	"fmt"
	"testing"

	"github.com/goexl/id"
)

func TestSnowflake(test *testing.T) {
	generator := id.New().Snowflake().Node(1).Build().Build()
	for count := 0; count < 100; count++ {
		next := generator.Next()
		fmt.Println(next.Value(), next.String().Build().Format())
	}
}

func TestAutoincrement(test *testing.T) {
	generator := id.New().Autoincrement().Build().Build()
	for count := 0; count < 100; count++ {
		next := generator.Next()
		fmt.Println(next.Value(), next.String().Build().Format())
	}
}
