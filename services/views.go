package services

import (
	"github.com/abwhop/portal_models/models"
)

func ConvertViews(viewsAPI *models.ViewsAPI) (*models.ViewsDB, error) {
	if viewsAPI == nil {
		return nil, nil
	}
	usersDB, err := ConvertUsers(viewsAPI.Users)
	if err != nil {
		usersDB = nil
	}
	return &models.ViewsDB{
		Count: viewsAPI.Count,
		Users: usersDB,
	}, nil
}
