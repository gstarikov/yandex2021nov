package main

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func baseTest(t *testing.T, days calcDays) {
	tests := []struct {
		X, K   int
		tasks  []int
		result int
	}{
		{
			X:      5,
			K:      10,
			tasks:  []int{1, 2, 3, 4, 5, 6},
			result: 10,
		},
		{
			X:      7,
			K:      12,
			tasks:  []int{5, 22, 17, 13, 8},
			result: 27,
		},
	}

	for _, test := range tests {
		cpyTasks := make([]int, len(test.tasks))
		copy(cpyTasks, test.tasks)
		res := days(test.X, test.K, cpyTasks)
		assert.Equalf(t, test.result, res, "test(%+v) res(%d) %+v", test, res, GetFunctionName(days))
	}
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
