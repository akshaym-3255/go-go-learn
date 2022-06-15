package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{3, 4},
			12,
		},
		{
			"test2",
			args{7, 6},
			42,
		},
	}

	assert := assert.New(t)
	for _, tt := range tests {
		assert.Equal(Multiply(tt.args.x,tt.args.y), tt.want)
	}
}
