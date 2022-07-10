package database

import (
	"CoopeLunch-api/domain"
	"CoopeLunch-api/tools"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) SighUp(user domain.TUserInsert) (resUser domain.TUserResponse, err error) {
	row, err := repo.Query(
		"SELECT EXISTS(SELECT * FROM users WHERE Email = ?)",
		user.Email,
	)
	defer row.Close()
	var exists bool
	for row.Next() {
		err = row.Scan(&exists)

		if err != nil {
			return
		}
	}
	if exists {
		err = domain.ErrExistingEmail
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		fmt.Println("passwordHash", err)
		return
	}
	exe, err := repo.Execute(
		"INSERT INTO users (FirstName, LastName, Email, PasswordHash) VALUES (?, ?, ?, ?)",
		user.FirstName, user.LastName, user.Email, string(passwordHash),
	)
	if err != nil {
		fmt.Println("SignUpDBError", err)
		return
	}

	rawId, err := exe.LastInsertId()
	if err != nil {
		return
	}
	resUser.ID = int(rawId)

	resUser.Token, err = tools.CreateJwtToken(resUser.ID)
	if err != nil {
		return
	}

	return
}

func (repo *UserRepository) LoginUser(user domain.TLoginUser) (resUser domain.TUserResponse, err error) {
	row, err := repo.Query(
		"SELECT ID, Email, PasswordHash FROM users WHERE Email = ?",
		user.Email,
	)
	if err != nil {
		return
	}
	defer row.Close()

	var u domain.TLoginUserDBResponse

	for row.Next() {
		err = row.Scan(
			&u.ID, &u.Email, &u.PasswordHash,
		)
		if err != nil {
			panic(err.Error())
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(user.Password))

	if err != nil {
		err = domain.ErrIncorrectPassword
		return
	}

	resUser.ID = u.ID
	resUser.Token, err = tools.CreateJwtToken(u.ID)
	if err != nil {
		return
	}

	return
}
