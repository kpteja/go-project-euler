package multiples

import "testing"

func TestCalculateSumOfMultiples(t *testing.T) {
	type args struct {
		multiplesOf []int
		from        int
		to          int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Multiples of 3 and 5 below 10",
			args: args{
				multiplesOf: []int{3, 5},
				from:        1,
				to:          10,
			},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateSumOfMultiples(tt.args.multiplesOf, tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("CalculateSumOfMultiples() = %v, want %v", got, tt.want)
			}
		})
	}
}
