package mysqluser

import (
	"app/entity"
	"app/pkg/errmsg"
	"app/pkg/richerror"
	"app/repository/mysql"
	"database/sql"
	"fmt"
	"time"
)

func (d *DB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	const op = "mysql.IsPhoneNumberUnique"
	row := d.adapter.Conn().QueryRow(`select * from users where phone_number= ?`, phoneNumber)
	_, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, richerror.New(op).
			WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected)
	}
	return false, nil
}
func (d *DB) Register(u entity.User) (entity.User, error) {
	res, err := d.adapter.Conn().Exec(`insert into users(name,phone_number,password,role) values(?,?,?,?)`, u.Name, u.PhoneNumber, u.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("can't execute command:%w", err)
	}
	id, _ := res.LastInsertId()
	u.ID = uint(id)
	return u, nil
}
func (d *DB) GetUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	const op = "mysql.GetUserByPhoneNumber"
	row := d.adapter.Conn().QueryRow(`select * from users where phone_number= ?`, phoneNumber)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgNotfound).WithKind(richerror.KindNotFound)
		}
		return entity.User{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return user, nil
}
func (d *DB) GetUserByID(userID uint) (entity.User, error) {
	const op = "mysql.GetUserByID"
	row := d.adapter.Conn().QueryRow(`select * from users where id= ?`, userID)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgNotfound).WithKind(richerror.KindNotFound)
		}
		return entity.User{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindNotFound)
	}
	return user, nil
}
func scanUser(scanner mysql.Scanner) (entity.User, error) {
	var createdAt time.Time
	var user entity.User
	var roleStr string
	err := scanner.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createdAt, &user.Password, &roleStr)
	return user, err
}
