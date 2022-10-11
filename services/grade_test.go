package services_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gotest/services"
	"testing"
)

func Test_CheckGrade(t *testing.T) {
	type testCase struct {
		nameFunc string
		score    int
		expected string
	}

	cases := []testCase{
		{nameFunc: "Success Grade A", score: 80, expected: "A"},
		{nameFunc: "Success Grade B", score: 70, expected: "B"},
		{nameFunc: "Success Grade C", score: 60, expected: "C"},
		{nameFunc: "Success Grade D", score: 50, expected: "D"},
		{nameFunc: "Success Grade F", score: 0, expected: "F"},
	}

	for _, c := range cases {
		t.Run(c.nameFunc, func(t *testing.T) {
			grade := services.CheckGrade(c.score)
			assert.Equal(t, c.expected, grade)
		})
	}
}

func Benchmark_CheckGrade(b *testing.B) {
	for i := 0; i < b.N; i++ {
		services.CheckGrade(80)
	}
}

func ExampleCheckGrade() {
	grade := services.CheckGrade(80)
	fmt.Println(grade)
	// Output: A
}
