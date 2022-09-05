package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DBRepository struct {
	db     *sql.DB
	config *PgConfig
}

func NewDBRepository() *DBRepository {
	return &DBRepository{}
}

func (d *DBRepository) InsertUrl(code string, originalURL string) {
	_, err := d.db.Exec(`INSERT INTO urls ("code", "originalUrl") VALUES ($1, $2)`, code, originalURL)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (d *DBRepository) GetUrl(code string) string {
	var result string
	row := d.db.QueryRow(`SELECT "originalUrl" FROM urls WHERE "code" = $1`, code)
	err := row.Scan(&result)
	if err != nil {
		fmt.Println(err.Error())
	}
	return result
}

func (d *DBRepository) ContainsUrl(originalURL string) (bool, string) {
	var code string
	row := d.db.QueryRow(`SELECT "code" FROM urls WHERE "originalUrl" = $1`, originalURL)
	err := row.Scan(&code)
	if err != nil {
		fmt.Println(err.Error())
	}
	return code != "", code
}

func (d *DBRepository) Connect() error {
	var e error
	d.db, e = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PgHost, cfg.PgPort, cfg.PgUser, cfg.PgPass, cfg.PgBase))
	if e != nil {
		return e
	}
	return nil
}
