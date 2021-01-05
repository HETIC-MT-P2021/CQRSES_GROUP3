package models

import (
	"github.com/edwinvautier/go-boilerplate/models"
	"testing"
)

func TestValidateCustomer(t *testing.T) {
	type args struct {
		customer *models.CustomerForm
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty name", 
			args: args{
				&models.CustomerForm{
					Name: "",
					Password: "tesT12345",
					Email: "bob@gmail.com",
				},
			},
			wantErr: true,
		},
		{
			name: "wrong email", 
			args: args{
				&models.CustomerForm{
					Name: "Bob",
					Password: "tesT12345",
					Email: "bobgmail.com",
				},
			},
			wantErr: true,
		},
		{
			name: "wrong password", 
			args: args{
				&models.CustomerForm{
					Name: "Bob",
					Password: "test",
					Email: "bob@gmail.com",
				},
			},
			wantErr: true,
		},
		{
			name: "correct customer", 
			args: args{
				&models.CustomerForm{
					Name: "Bob",
					Password: "tesT12345",
					Email: "bob@gmail.com",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := models.ValidateCustomer(tt.args.customer); (err != nil) != tt.wantErr {
				t.Errorf("ValidateCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
