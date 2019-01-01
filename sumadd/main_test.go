package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name:"one",args:args{[]int{1,2,3,6},9},want:[]int{2,3}},
		{name:"two",args:args{[]int{5,2,3,6},7},want:[]int{0,1}},
		//{name:"three",args:args{[]int{5},5},want:[]int{0}},
		//{name:"null",args:args{[]int{5},6},want:[]int{}},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.args.nums, tt.args.target)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
