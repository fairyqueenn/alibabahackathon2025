package user

import "context"

func (svc *service) GetUserByID(ctx context.Context, id int) (user *User, err error) {
	return svc.repo.FindUserByID(ctx, id)
}

func (svc *service) GetUserPreferenceByID(ctx context.Context, id int) (userPreference *UserPreference, err error) {
	return svc.repo.FindUserPreferenceByID(ctx, id)
}

func (svc *service) UpdateUserPreferenceRequest(ctx context.Context, id int, req UpdateUserPreferenceRequest) (err error) {
	userPreference, err := svc.repo.FindUserPreferenceByID(ctx, id)
	if err != nil {
		return
	}

	userPreference.IsHalal = req.IsHalal
	userPreference.IsVegan = req.IsVegan
	userPreference.IsDiabetes = req.IsDiabetes
	userPreference.IsHypertension = req.IsHypertension
	userPreference.IsCholesterol = req.IsCholesterol
	userPreference.IsGout = req.IsGout
	userPreference.IsKidney = req.IsKidney
	userPreference.IsCeliac = req.IsCeliac
	userPreference.IsLactoseIntolerance = req.IsLactoseIntolerance
	userPreference.IsAutoimmune = req.IsAutoimmune
	userPreference.IsThyroid = req.IsThyroid
	userPreference.IsObesity = req.IsObesity
	userPreference.IsDigestiveIssues = req.IsDigestiveIssues

	return svc.repo.UpsertUserPreference(ctx, userPreference)
}
