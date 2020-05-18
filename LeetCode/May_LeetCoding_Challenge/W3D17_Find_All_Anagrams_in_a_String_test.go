package May_LeetCoding_Challenge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test__W3D17__findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{"TestCase01", args{s: "baa", p: "aa"}, []int{1}},
		{"TestCase02", args{s: "abab", p: "ab"}, []int{0, 1, 2}},
		{"TestCase03", args{s: "cbaebabacd", p: "abc"}, []int{0, 6}},
		{"TestCase04", args{s: "", p: "a"}, nil},
		{"TestCase05", args{s: "aaaaaaaaaa", p: "aaaaaaaaaaaaa"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, findAnagrams(tt.args.s, tt.args.p))
		})
	}
}
