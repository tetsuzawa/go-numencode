package main

import (
	"reflect"
	"testing"
)

func Test_encode(t *testing.T) {
	type args struct {
		s     string
		dtype string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "short",
			args: args{
				s:     "32767",
				dtype: "short",
			},
			want:    []byte{255, 127},
			wantErr: false,
		},
		{
			name: "float",
			args: args{
				s:     "1.6777215e+07",
				dtype: "float",
			},
			want:    []byte{255, 255, 127, 75},
			wantErr: false,
		},
		{
			name: "double",
			args: args{
				s:     "9.007199e+15",
				dtype: "float",
			},
			want:    []byte{0, 0, 0, 90},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encode(tt.args.s, tt.args.dtype)
			if (err != nil) != tt.wantErr {
				t.Errorf("encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encode() got = %v, want %v", got, tt.want)
			}
			got2, err := decode(got, tt.args.dtype)
			if (err != nil) != tt.wantErr {
				t.Errorf("decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got2, tt.args.s) {
				t.Errorf("decode() got = %v, want %v", got2, tt.args.s)
			}
		})
	}
}
