package toolkit

import (
	"strconv"
	"testing"
)

func TestStrArrEqual(t *testing.T) {
	cases := []struct {
		Arr1     []string
		Arr2     []string
		Excepted bool
	}{
		{[]string{}, []string{}, true},
		{[]string{"1", "2"}, []string{"1", "2"}, true},
		{[]string{"1"}, []string{"1", "2"}, false},
		{[]string{"1", "2"}, []string{"1"}, false},
		{[]string{"1", "2", "3", "1"}, []string{"1", "2", "3", "1"}, true},
		{[]string{"1", "2", "3", "1"}, []string{"1", "2", "1", "3"}, false},
		{[]string{"1", "3", "1"}, []string{"1", "2", "3", "1"}, false},
		{[]string{"1", "2", "3", "1"}, []string{"1", "2"}, false},
	}
	for _, row := range cases {
		if StrArrEqual(row.Arr1, row.Arr2) != row.Excepted {
			t.Fatalf("StrArrEqual Arr1 %v ,Arr2 %v ,except %v ,but %v got", row.Arr1, row.Arr2, row.Excepted, StrArrEqual(row.Arr1, row.Arr2))
		}
	}
}

func TestIndexStrArr(t *testing.T) {
	cases := []struct {
		Str      string
		Arr      []string
		Excepted int
	}{
		{"0", []string{"1", "2"}, -1},
		{"1", []string{}, -1},
		{"1", []string{"1", "2"}, 0},
	}

	for i, row := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if index := IndexStrArr(row.Str, row.Arr); index != row.Excepted {
				t.Fatalf("%s index in %v excepted to be %d ,but %d got", row.Str, row.Arr, row.Excepted, index)
			}
		})
	}
}

func TestDeduplicateStrArrWithMap(t *testing.T) {
	cases := []struct {
		Arr      []string
		Excepted []string
	}{
		{[]string{"1", "2", "2", "1"}, []string{"1", "2"}},
		{[]string{""}, []string{""}},
		{[]string{"1", "1", "1"}, []string{"1"}},
		{[]string{"1", "2", "1", "2", "1"}, []string{"1", "2"}},
		{[]string{"1", "2", "1", "2", "1", "3"}, []string{"1", "2", "3"}},
	}

	for i, row := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if res := DeduplicateStrArrWithMap(row.Arr); !StrArrEqual(res, row.Excepted) {
				t.Fatalf("DeduplicateStrArrWithMap arr: %v ,excepted %v ,but %v got", row.Arr, row.Excepted, res)
			} else {
				//log.Printf("[PASSED] DeduplicateStrArrWithMap arr: %v ,excepted %v , %v got", row.Arr, row.Excepted, res)
			}
		})
	}
}

func TestStrArraysIntersect(t *testing.T) {
	cases := []struct {
		Arr1     []string
		Arr2     []string
		Excepted []string
	}{
		{[]string{}, []string{}, []string{}},
		{[]string{"1", "2"}, []string{}, []string{}},
		{[]string{"1", "2"}, []string{"2", "3"}, []string{"2"}},
		{[]string{"1", "2"}, []string{"3", "4"}, []string{}},
		{[]string{"1", "2"}, []string{"2", "4", "1"}, []string{"1", "2"}},
		{[]string{"1", "2"}, []string{"2", "3", "4"}, []string{"2"}},
	}

	for _, row := range cases {
		if res := StrArraysIntersect(row.Arr1, row.Arr2); !StrArrEqual(res, row.Excepted) {
			t.Fatalf("StrArraysIntersect arr1: %v   arr2: %v , excepted to be %v ,but %v got", row.Arr1, row.Arr2, row.Excepted, res)
		}
	}
}
