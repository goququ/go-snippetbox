package psql

import (
	"database/sql"
	"errors"

	"github.com/goququ/snippetbox/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created)
    VALUES($1, $2, $3, now())`

	if _, err = m.DB.Exec(stmt, name, email, string(hashedPassword)); err != nil {
		return err
	}

	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	var (
		ID             int
		hashedPassword []byte
	)
	stmt := "SELECT id, hashed_password FROM users WHERE email = $1 AND active = TRUE"

	row := m.DB.QueryRow(stmt, email)
	if err := row.Scan(&ID, &hashedPassword); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	return ID, nil
}

func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
