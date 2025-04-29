package database

import (
	"database/sql"
	"fmt"
	"tatar/book/repository"

	_ "github.com/go-sql-driver/mysql"
)

type DBManager struct {
	db             *sql.DB
	UserRepository repository.UserRepository
}

func NewDBManager(dsn string) (*DBManager, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("cannot open database connection: %w", err)
	}

	// ตรวจสอบการเชื่อมต่อ
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("cannot ping database: %w", err)
	}

	// เปลี่ยน schema ด้วย SQL command
	_, err = db.Exec("USE mydb")
	if err != nil {
		return nil, fmt.Errorf("failed to use schema: %w", err)
	}

	// สร้าง repositories (ต้องมีฟังก์ชัน NewMySQLUserRepository ใน package repository)
	userRepo := repository.NewMySQLUserRepository(db)

	return &DBManager{
		db:             db,
		UserRepository: userRepo,
	}, nil
}

func (m *DBManager) Close() error {
	return m.db.Close()
}

func (m *DBManager) BeginTx() (*sql.Tx, error) {
	return m.db.Begin()
}
