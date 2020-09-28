package database

import (
	"database/sql"
)

// type DB struct {
// 	Db *sql.DB
// }

// Config of database
type Config struct {
	DSN string
	// Host     string `env:"POSTGRES_HOST,required"`
	// Port     int    `env:"POSTGRES_HOST,required"`
	// User     string `env:"POSTGRES_HOST,required"`
	// Password string `env:"POSTGRES_HOST,required"`
	// DBName   string `env:"POSTGRES_HOST,required"`
	// SSLMode  bool
}

// func (cfg Config) String() string {
// 	return fmt.Sprintf("Host=%s Port=%d User=%s Password=%s DBName=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
// }

func (cfg Config) validate() error {
	// if cfg.Host == "" || cfg.Port == 0 || cfg.User == "" ||
	// 	cfg.Password == "" || cfg.DBName == "" {
	// 	return fmt.Errorf(
	// 		"All fields must be set (%s)", cfg)
	// }

	return nil

}

//postgres://username:password@url.com:5432/dbName
//go through
//https://medium.com/@beld_pro/postgres-with-golang-3b788d86f2ef

//Starts Database

//New represents the database logic
func New(dbDSN string) (*sql.DB, error) {

	// if err := cfg.validate(); err != nil {
	// 	return nil, err
	// }
	// db, err := sql.Open("postgres", fmt.Sprintf(
	// 	"user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", cfg.User, cfg.Password, cfg.DBName, cfg.Host, cfg.Port))
	db, err := sql.Open("postgres", dbDSN)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
