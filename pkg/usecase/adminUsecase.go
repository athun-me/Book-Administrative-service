package usecas

import (
	"errors"

	"githum.com/athunlal/bookNowAdmin-svc/pkg/domain"
	intereface "githum.com/athunlal/bookNowAdmin-svc/pkg/repository/interface"
	usecas "githum.com/athunlal/bookNowAdmin-svc/pkg/usecase/interface"
	"githum.com/athunlal/bookNowAdmin-svc/pkg/utils"
)

type adminUseCase struct {
	Repo intereface.AdminRepo
}

// ValidateJwtAdmin implements interfaces.AdminUseCase.
func (use *adminUseCase) ValidateJwtAdmin(adminid uint) (domain.Admin, error) {
	admin, err := use.Repo.FindAdminById(adminid)
	if err != nil {
		return admin, errors.New("Unauthorized User")
	}
	return admin, nil

}

// ChangePassword implements interfaces.AdminUseCase.
func (use *adminUseCase) ChangePassword(Admin domain.Admin) error {
	// user.Password = utils.HashPassword(user.Password)
	Admin.Password = utils.HashPassword(Admin.Password)
	err := use.Repo.ChangePassword(Admin)
	if err != nil {
		return errors.New("Could not change the password")
	}
	return nil
}

// Login implements interfaces.AdminUseCase.
func (use *adminUseCase) Login(Admin domain.Admin) (domain.Admin, error) {
	var adminDetatils domain.Admin
	var err error
	if Admin.Username != "" {
		adminDetatils, err = use.Repo.FindByAdminName(Admin)
		if err != nil {
			return adminDetatils, errors.New("User not found")
		}
	} else if Admin.Email != "" {
		adminDetatils, err = use.Repo.FindByAdminEmail(Admin)
		if err != nil {
			return adminDetatils, errors.New("User not found")
		}
	}
	if !utils.VerifyPassword(Admin.Password, adminDetatils.Password) {
		return adminDetatils, errors.New("Password is not matched or worg")
	}
	return adminDetatils, nil
}

func NewAdminUseCase(repo intereface.AdminRepo) usecas.AdminUseCase {
	return &adminUseCase{
		Repo: repo,
	}
}
