package repository

import (
	"database/sql"
	"fmt"
	"log"
	"tatar/book/.gen/mydb/model"
	"tatar/book/.gen/mydb/table"

	"github.com/go-jet/jet/v2/mysql"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{
		db: db,
	}
}

func (r *MySQLUserRepository) GetUserByUUID(id string) (*model.User, error) {
	stmt := mysql.SELECT(
		table.User.ID,
		table.User.Name,
		table.User.Email,
	).FROM(table.User).WHERE(table.User.UUID.EQ(mysql.String(id)))
	var user model.User
	err := stmt.Query(r.db, &user)
	if err != nil {
		return nil, fmt.Errorf("error querying user: %w", err)
	}
	
	return &user, nil
}

func (r *MySQLUserRepository) GetUsers(limit int, offset int) ([]*model.User, error) {
	stmt := mysql.SELECT(
		table.User.ID,
		table.User.Name,
		table.User.Email,
		table.User.UUID,
	).FROM(
		table.User,
	).LIMIT(int64(limit)).
		OFFSET(int64(offset))

	var users []*model.User
	err := stmt.Query(r.db, &users)
	if err != nil {
		return nil, fmt.Errorf("error listing users: %w", err)
	}
	return users, err
}

func (r *MySQLUserRepository) GetUserByEmail(email string) (*model.User, error) {
	stmt := mysql.SELECT(
		table.User.ID,
		table.User.Name,
		table.User.Email,
	).FROM(table.User).WHERE(table.User.Email.EQ(mysql.String(email)))
	var user model.User
	err := stmt.Query(r.db, &user)
	if err != nil {
		return nil, fmt.Errorf("error querying user: %w", err)
	}
	return &user, nil
}

func (r *MySQLUserRepository) GetUserByName(name string) (*model.User, error) {
	stmt := mysql.SELECT(
		table.User.ID,
		table.User.Name,
		table.User.Email,
	).FROM(table.User).WHERE(table.User.Name.EQ(mysql.String(name)))
	var user model.User
	err := stmt.Query(r.db, &user)
	if err != nil {
		return nil, fmt.Errorf("error querying user: %w", err)
	}
	return &user, nil
}

func (r *MySQLUserRepository) CreateUser(user *model.User) (*model.User, error) {
	stmt := table.User.INSERT(
		table.User.Name,
		table.User.Email,
		table.User.UUID,
	).VALUES(
		mysql.String(user.Name),
		mysql.String(user.Email),
		mysql.String(user.UUID),
	)
	_, err := stmt.Exec(r.db)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("error inserting user: %w", err)
	}
	return user, err
}

func (r *MySQLUserRepository) UpdateUser(user *model.User) (*model.User, error) {
	stmt := table.User.UPDATE(
		table.User.Name,
		table.User.Email,
	).SET(
		mysql.String(user.Name),
		mysql.String(user.Email),
	).WHERE(
		table.User.ID.EQ(mysql.Int32(user.ID)),
	)
	_, err := stmt.Exec(r.db)
	if err != nil {
		return user, fmt.Errorf("error updating user: %w", err)
	}
	return user, nil
}

func (r *MySQLUserRepository) DeleteUser(id string) error {
	stmt := table.User.DELETE().WHERE(table.User.UUID.EQ(mysql.String(id)))
	_, err := stmt.Exec(r.db)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}
