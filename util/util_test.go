package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTruncateWithLine(t *testing.T) {
	type args struct {
		str      string
		maxLines int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "maxLines < 1",
			args: args{
				str:      "a\nb\nc\n",
				maxLines: 0,
			},
			want: "a\nb\nc\n",
		},
		{
			name: "lines <= maxLines",
			args: args{
				str:      "a\nb\n",
				maxLines: 3,
			},
			want: "a\nb\n",
		},
		{
			name: "lines > maxLines",
			args: args{
				str:      "a\nb\nc\nd\n",
				maxLines: 3,
			},
			want: "a\nb\nc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TruncateWithLine(tt.args.str, tt.args.maxLines)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSelectLine(t *testing.T) {
	str := `111
222
333
`
	type args struct {
		str  string
		line int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Select 2nd line",
			args: args{
				str:  str,
				line: 2,
			},
			want: "222",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectLine(tt.args.str, tt.args.line); got != tt.want {
				t.Errorf("SelectLine() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestSelectLines(t *testing.T) {
	str := `111
222
333
444
555
`

	type args struct {
		str       string
		startLine int
		endLine   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Select lines",
			args: args{
				str:       str,
				startLine: 2,
				endLine:   4,
			},
			want: "222\n333\n444",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectLines(tt.args.str, tt.args.startLine, tt.args.endLine)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFormatMarkdownForSlack(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no markdown",
			args: args{
				str: "aaa",
			},
			want: "aaa",
		},
		{
			name: "include 1 image",
			args: args{
				str: "aaa ![img1](/foo/img1.png) bbb",
			},
			want: "aaa img1 bbb",
		},
		{
			name: "include 2 images",
			args: args{
				str: "aaa ![img1](/foo/img1.png) bbb ![img2](/foo/img2.png) ccc",
			},
			want: "aaa img1 bbb img2 ccc",
		},
		{
			name: "empty text",
			args: args{
				str: "![](/foo/img1.png)",
			},
			want: "",
		},
		{
			name: "space text",
			args: args{
				str: "![ ](/foo/img1.png)",
			},
			want: " ",
		},
		{
			name: "image with checkbox",
			args: args{
				str: "* [ ] ![img1](/foo/img1.png)",
			},
			want: "* [ ] img1",
		},
		{
			name: "normal link",
			args: args{
				str: "[github](https://github.com/)",
			},
			want: "<https://github.com/|github>",
		},
		{
			name: "url is blank",
			args: args{
				str: "[github]()",
			},
			want: "github",
		},
		{
			name: "text is blank",
			args: args{
				str: "[](https://github.com/)",
			},
			want: "https://github.com/",
		},
		{
			name: "2 links",
			args: args{
				str: "aaa [github](https://github.com/) bbb [twitter](https://twitter.com/) ccc",
			},
			want: "aaa <https://github.com/|github> bbb <https://twitter.com/|twitter> ccc",
		},
		{
			name: "embed image and link",
			args: args{
				str: "aaa ![img1](/foo/img1.png) bbb [github](https://github.com/) ccc",
			},
			want: "aaa img1 bbb <https://github.com/|github> ccc",
		},
		{
			name: "link with checkbox",
			args: args{
				str: "* [ ] [hashdiff](https://github.com/liufengyun/hashdiff): [`0.3.9...0.4.0`](https://github.com/liufengyun/hashdiff/compare/v0.3.9...v0.4.0)",
			},
			want: "* [ ] <https://github.com/liufengyun/hashdiff|hashdiff>: <https://github.com/liufengyun/hashdiff/compare/v0.3.9...v0.4.0|`0.3.9...0.4.0`>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatMarkdownForSlack(tt.args.str)

			assert.Equal(t, tt.want, got)
		})
	}
}
