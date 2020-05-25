package fibonacci

import "testing"

func TestSumOfEvenValuedTerms(t *testing.T) {
	type args struct {
		maxSum uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "Sum of even fibonacci numbers below 10",
			args: args{maxSum: 10},
			want: 10,
		},
		{
			name: "Sum of even fibonacci numbers below 4,000,000",
			args: args{maxSum: 4000000},
			want: 4613732,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfEvenValuedTerms(tt.args.maxSum); got != tt.want {
				t.Errorf("SumOfEvenValuedTerms() = %v, want %v", got, tt.want)
			}
		})
	}
}
