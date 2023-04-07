package time

import (
	"testing"
	"time"
)

func TestToStandardFormat(t *testing.T) {
	type args struct {
		t      *time.Time
		format string
	}
	now := time.Date(2023, time.April, 7, 18, 30, 0, 0, time.Local)
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test One",
			args: args{
				t:      &now,
				format: StandardFormat[1],
			},
			want: "2023-04-07 18:30:00.000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToStandardFormat(tt.args.t, tt.args.format); got != tt.want {
				t.Errorf("ToStandardFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
