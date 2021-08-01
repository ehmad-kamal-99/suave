package redisbloom

import (
	"os"
	"testing"
)

func Test_client_AddUsername(t *testing.T) {
	os.Setenv("DB_HOST", "suave-db")
	client := NewClient()

	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "add username in redis-bloom",
			args: args{
				username: "johndoe123",
			},
			wantErr: false,
		},
		{
			name: "add username in redis-bloom",
			args: args{
				username: "alexscott998",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := client.AddUsername(tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("AddUsername() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_CheckUsername(t *testing.T) {
	os.Setenv("DB_HOST", "suave-db")
	client := NewClient()

	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "success - check username in redis-bloom",
			args: args{
				username: "johndoe123",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "fail - check username in redis-bloom",
			args: args{
				username: "alexscotti9",
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.CheckUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckUsername() got = %v, want %v", got, tt.want)
			}
		})
	}
}
