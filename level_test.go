package log_test

import (
	. "github.com/bukowa/log"
)

import "testing"

func Test_isLogged(t *testing.T) {
	type args struct {
		desired Level
		logged  Level
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				desired: DebugLevel,
				logged:  InfoLevel,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				desired: InfoLevel,
				logged:  DebugLevel,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				desired: FatalLevel,
				logged:  ErrorLevel,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLogged(tt.args.desired, tt.args.logged); got != tt.want {
				t.Errorf("isLogged() = %v, want %v", got, tt.want)
			}
		})
	}
}
