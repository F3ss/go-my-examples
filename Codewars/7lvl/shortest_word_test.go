package main

import "testing"

func TestFindShort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindShort(tt.args.s); got != tt.want {
				t.Errorf("FindShort() = %v, want %v", got, tt.want)
			}
		})
	}
}
