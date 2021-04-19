package helpers

import "testing"

func TestParseStringToUint64(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "Test with 12",
			args: args{
				str: "12",
			},
			want: 12,
		},
		{
			name: "Test with twelve",
			args: args{
				str: "twelve",
			},
			want: 0,
		},

		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseStringToUint64(tt.args.str); got != tt.want {
				t.Errorf("ParseStringToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}
