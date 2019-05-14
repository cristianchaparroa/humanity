package repositories

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func TestFindByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()

	gormDB, _ := gorm.Open("postgres", db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var test = []struct {
		ID       string
		email    string
		pass     string
		nickname string
	}{
		{"9dcb45b5-6845-46a9-93bf-6deb63266f66", "mlopez@gmail.com", "12345", "mlopez"},
		{"9dcb45b5-6845-46a9-93bf-6deb63266f61", "santiagocastro@gmail.com", "12345", "scastro"},
	}

	ar := NewAccountRepository(gormDB)

	for _, tc := range test {

		rows := sqlmock.NewRows([]string{"id", "email", "password", "nickname"}).
			AddRow(tc.ID, tc.email, tc.pass, tc.nickname)

		mock.ExpectQuery("^SELECT (.+) FROM \"accounts\" (.+)$").WillReturnRows(rows)

		acc, err := ar.FindByEmail(tc.email)

		if err != nil {
			t.Error(err)
		}
		if acc.ID != tc.ID {
			t.Errorf("Expected %v, but get %v", tc.ID, acc.ID)
		}

		if acc.Email != tc.email {
			t.Errorf("Expected %v, but get %v", tc.email, acc.Email)
		}

		if acc.Password != tc.pass {
			t.Errorf("Expected %v, but get %v", tc.pass, acc.Password)
		}
	}

}
