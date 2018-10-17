package bob

import (
	"testing"
)

func TestHey(t *testing.T) {
	for _, tt := range testCases {
		actual := Hey(tt.input)
		if actual != tt.expected {
			msg := `
	ALICE (%s): %q
	BOB: %s

	Expected Bob to respond: %s`
			t.Fatalf(msg, tt.description, tt.input, actual, tt.expected)
		}
	}
}

func BenchmarkHey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			Hey(tt.input)
		}
	}
}

func Test_isAllUpperCase(t *testing.T) {
	type args struct {
		remark string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "BOOM",
			args: args{
				remark: "BOOM",
			},
			want: true,
		},
		{
			name: "BomM",
			args: args{
				remark: "BomM",
			},
			want: false,
		},
		{
			name: "1,2,3",
			args: args{
				remark: "1,2,3",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAllUpperCase(tt.args.remark); got != tt.want {
				t.Errorf("isAllUpperCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAllLetters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Is all letters",
			args: args{
				s: "asdasklh",
			},
			want: true,
		},
		{
			name: "Is not letters",
			args: args{
				s: "1,2,3",
			},
			want: false,
		},
		{
			name: "Is not letters",
			args: args{
				s: "1, 2, 3 GO!",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAllLetters(tt.args.s); got != tt.want {
				t.Errorf("isAllLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
