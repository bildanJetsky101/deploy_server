package repositories

import (
	"server_wb/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	FindProfiles() ([]models.Profile, error)
	GetProfile(ID int) (models.Profile, error)
	CreateProfile(user models.Profile) (models.Profile, error)
}

func RepositoryProfile(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProfiles() ([]models.Profile, error) {
	var profiles []models.Profile
	err := r.db.Find(&profiles).Error

	return profiles, err
}

func (r *repository) GetProfile(ID int) (models.Profile, error) {
	var profile models.Profile
	//err := r.db.First(&profile, ID).Error
	err := r.db.Raw("SELECT * FROM profile WHERE user_id=?", ID).Scan(&profile).Error
	return profile, err
}

// Write this code
func (r *repository) CreateProfile(profile models.Profile) (models.Profile, error) {
	err := r.db.Create(&profile).Error // Using Create method

	return profile, err
}
