package user

import "context"

func (r *repository) FindUserByID(ctx context.Context, id int) (user *User, err error) {
	err = r.DB.Where("id = ?", id).First(&user).Error
	return
}

func (r *repository) FindUserPreferenceByID(ctx context.Context, id int) (userPreference *UserPreference, err error) {
	err = r.DB.Where("id = ?", id).First(&userPreference).Error
	return
}

func (r *repository) UpsertUserPreference(ctx context.Context, userPreference *UserPreference) (err error) {
	err = r.DB.Save(userPreference).Error
	return
}
