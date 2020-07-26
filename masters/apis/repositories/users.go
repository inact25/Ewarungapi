package repositories

import (
	"database/sql"
	"github.com/inact25/E-WarungApi/masters/apis/models"
	"log"
)

type UserRepoImpl struct {
	db *sql.DB
}

func (u UserRepoImpl) OAuth(userModels *models.SelfUserModels) ([]*models.OAuth, error) {
	log.Print("r : ", userModels)
	var dataOAuth []*models.OAuth
	data := models.OAuth{}
	query := OAuth
	row := u.db.QueryRow(query, userModels.UserName, userModels.UserPassword)
	if err := row.Scan(&data.UserID, &data.UserLevel); err != nil {
		return nil, err
	}
	dataOAuth = append(dataOAuth, &data)

	log.Print("Data r: ", dataOAuth)

	return dataOAuth, nil
}

func InitUserRepoImpl(db *sql.DB) UsersRepo {
	return &UserRepoImpl{db}

}
