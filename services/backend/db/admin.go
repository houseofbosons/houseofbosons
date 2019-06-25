package db

import (
	"database/sql"
)

// Admin ...
type Admin struct {
	Email    string `json:"email"`
	GoogleID string `json:"id"`
}

// Insert ...
func (a *Admin) Insert() (sql.Result, error) {
	str := `
	INSERT INTO admins (email, google_id)
	VALUES ($1, $2)`

	return PG.DB.Exec(str, a.Email, a.GoogleID)
}

// Query ...
func (a *Admin) Query() error {
	str := `SELECT email, google_id FROM admins WHERE email=$1;`

	row := PG.DB.QueryRow(str, a.Email)

	return row.Scan(&a.Email, &a.GoogleID)
}