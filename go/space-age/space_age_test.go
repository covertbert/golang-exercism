package space

import (
	"math"
	"testing"
)

func TestAge(t *testing.T) {
	const precision = 0.01
	for _, tc := range testCases {
		actual := Age(tc.seconds, tc.planet)
		if math.IsNaN(actual) || math.Abs(actual-tc.expected) > precision {
			t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", tc.description, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkAge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Age(tc.seconds, tc.planet)
		}
	}
}

func Test_getPlanetYearInEarthYears(t *testing.T) {
	type args struct {
		planet Planet
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Earth",
			args: args{
				planet: "Earth",
			},
			want: 1,
		},
		{
			name: "Mercury",
			args: args{
				planet: "Mercury",
			},
			want: 0.2408467,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPlanetYearInEarthYears(tt.args.planet); got != tt.want {
				t.Errorf("getPlanetYearInEarthYears() = %v, want %v", got, tt.want)
			}
		})
	}
}
