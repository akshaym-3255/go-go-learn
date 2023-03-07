package main

import "testing"

func Test_test2(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := test2(tt.args.x, tt.args.y)
			if got != tt.want {
				t.Errorf("test2() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("test2() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
