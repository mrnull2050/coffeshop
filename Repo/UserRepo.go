package repo

import (
	"coffeshop/model"
	"database/sql"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (ur *UserRepo) Create(um model.UserModel) error {

	_, err := ur.DB.Exec(`INSERT INTO userC (username , password , email , role) VALUES($1,$2,$3,$4)`, um.UserName, um.Password, um.Email, um.Role)
	return err
}
func (ur *UserRepo) GetAll() ([]model.UserModel, error) {
	rows, err := ur.DB.Query(`SELECT id , username , email , role FROM userC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []model.UserModel
	for rows.Next() {
		var u model.UserModel
		rows.Scan(u.ID, u.UserName, u.Email)
		users = append(users, u)
	}
	return users, nil
}
func (ur *UserRepo) GetUserById(id int) (model.UserModel, error) {
	var u model.UserModel
	err := ur.DB.QueryRow(`SELECT username , email , role FROM WHERE id=%1`, id).Scan(u.UserName, u.Email, u.Role)
	return u, err
}
func (ur *UserRepo) Update(u model.UserModel) error {

	_, err := ur.DB.Exec(`UPDATE  userC SET username=$1 , email=$2 , role=$3 , password=$4 WHERE id=$5`, u.UserName, u.Email, u.Role, u.Password, u.ID)
	return err
}
func (ur *UserRepo) Delete(id int) error {
	_, err := ur.DB.Exec(`DELETE FROM userC WHERE id=$1`, id)
	return err
}
