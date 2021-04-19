package services

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: " Test with email success",
			args: args{
				email: "contact.jason gauvin@gmail.com",
			},
			wantErr: true,
		},
		{
			name: " Test with email failure",
			args: args{
				email: "contact.jason.gauvin@gmail.com",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GenerateToken(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v", err)
				return
			}
		})
	}
}

func TestDecodeToken(t *testing.T) {
	type args struct {
		tokenString string
	}
	tokenTest, err := GenerateToken("contact.jason.gauvin@gmail.com")
	if err != nil {
		t.Fail()
	}
	tests := []struct {
		name    string
		args    args
		want   *Claims
		wantErr bool
	}{
		{
			name: " Test with token success",
			args: args{
				tokenString: tokenTest,
			},
			want: &Claims{
				Email: "contact.jason.gauvin@gmail.com",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, claims, err := DecodeToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if claims.Email != tt.want.Email {
				t.Errorf("DecodeToken() claims = %v, want %v", claims, tt.want)
			}
		})
	}
}
