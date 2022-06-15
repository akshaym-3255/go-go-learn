package main

import "testing"

func TestAdd_and_mulitply(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
		{
			"test1",
			args{3,4},
			7,
			12,
		},
		{
			"test2",
			args{7,6},
			13,
			42,

		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Add_and_mulitply(tt.args.x, tt.args.y)
			if got != tt.want {
				t.Errorf("Add_and_mulitply() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Add_and_mulitply() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
