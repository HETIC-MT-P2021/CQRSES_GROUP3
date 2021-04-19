package services

import "testing"

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
			name:" Test with email success",
			args: args{
				email: "contact.jason gauvin@gmail.com",
			},
			wantErr: true,

		},
		{
			name:" Test with email failure",
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
