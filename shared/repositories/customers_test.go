package repositories

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/models"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/repositories"
)

func TestFindCustomerByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil{
		t.fatal("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("SELECT customers WHERE email").WillReturnResult(sqlmock.NewResult(1,1))
	mock.ExpectCommit()

	type args struct {
		customer *models.Customer
		repository *repositories.Repository
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test with email success",
			args: args{
				customer: {
					Email: "contact.jason.gauvin@gmail.com"
				},
				repository: {
					Db: db
				}
			},
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.args.repository.FindCustomerByEmail(tt.args.customer); (err != nil) != tt.wantErr {
				t.Errorf("FindCustomerByEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
