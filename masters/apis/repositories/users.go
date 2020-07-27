package repositories

import (
	"database/sql"
	"github.com/inact25/E-WarungApi/masters/apis/models"
	"log"
)

type UserRepoImpl struct {
	db *sql.DB
}

func (u UserRepoImpl) GetSelfUser(userModels *models.SelfUserModels) ([]*models.SelfUserModels, error) {
	var result []*models.SelfUserModels
	data := models.SelfUserModels{}
	query := GetSelfUsersQuery
	row := u.db.QueryRow(query, userModels.UserID)
	if err := row.Scan(&data.UserID, &data.UserName, &data.UserEmail, &data.UserPassword, &data.UserLevel); err != nil {
		log.Fatal(err)
		return nil, err
	}
	result = append(result, &data)
	return result, nil
}

func (u UserRepoImpl) OAuth(userModels *models.SelfUserModels) ([]*models.OAuth, error) {
	var dataOAuth []*models.OAuth
	data := models.OAuth{}
	query := OAuth
	row := u.db.QueryRow(query, userModels.UserName, userModels.UserPassword)
	if err := row.Scan(&data.UserID, &data.UserLevel); err != nil {
		return nil, err
	}
	dataOAuth = append(dataOAuth, &data)
	return dataOAuth, nil
}

func (u UserRepoImpl) GetAllUsers() ([]*models.UserModels, error) {
	var dataUsers []*models.UserModels
	query := GetAllUsersQuery
	data, err := u.db.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for data.Next() {
		user := models.UserModels{}
		err := data.Scan(&user.UserID, &user.UserName, &user.UserEmail, &user.UserLevel)
		if err != nil {
			log.Fatal(err)
			return nil, err

		}
		dataUsers = append(dataUsers, &user)
	}
	return dataUsers, nil
}

func InitUserRepoImpl(db *sql.DB) UsersRepo {
	return &UserRepoImpl{db}

}
