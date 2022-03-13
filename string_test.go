package golibrary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicateStr(t *testing.T) {
	type args struct {
		strSlice []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{strSlice: []string{"1", "1", "2", "3"}}, []string{"1", "2", "3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, RemoveDuplicateStr(tt.args.strSlice), "RemoveDuplicateStr(%v)", tt.args.strSlice)
		})
	}
}
