package repositories

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestFindByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// expected
	accountID := "9dcb45b5-6845-46a9-93bf-6deb63266f67"
	email := "usertest@gmail.com"
	pass := "12345"
	nickname := "usert"

	rows := sqlmock.NewRows([]string{"id", "email", "password", "nickname"}).
		AddRow(accountID, email, pass, nickname)

	mock.ExpectQuery("^SELECT (.+) FROM account").WillReturnRows(rows)

	ar := NewAccountRepository(db)
	acc, err := ar.FindByEmail("usertest@gmail.com")

	if err != nil {
		t.Error(err)
	}
	if acc.ID != accountID {
		t.Errorf("Expected %v, but get %v", accountID, acc.ID)
	}

	if acc.Email != email {
		t.Errorf("Expected %v, but get %v", email, acc.Email)
	}

	if acc.Password != pass {
		t.Errorf("Expected %v, but get %v", pass, acc.Password)
	}

}
