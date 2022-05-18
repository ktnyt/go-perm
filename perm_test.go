package perm

import (
	"reflect"
	"testing"
)

func TestPerm(t *testing.T) {
	a := []string{"foo", "bar", "baz"}
	p := Permute(a)
	exp := [][]string{
		{"foo", "bar", "baz"},
		{"bar", "foo", "baz"},
		{"baz", "foo", "bar"},
		{"foo", "baz", "bar"},
		{"bar", "baz", "foo"},
		{"baz", "bar", "foo"},
	}
	for i := 0; p.Next(); i++ {
		if !reflect.DeepEqual(a, exp[i]) {
			t.Errorf("Permutation %d = %v, want %v", i+1, a, exp)
		}
	}
}
