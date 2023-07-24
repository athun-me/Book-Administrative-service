package usecas

import (
	"githum.com/athunlal/bookNowAdmin-svc/pkg/domain"
	intereface "githum.com/athunlal/bookNowAdmin-svc/pkg/repository/interface"
	usecas "githum.com/athunlal/bookNowAdmin-svc/pkg/usecase/interface"
)

type adminUseCase struct {
	Repo intereface.AdminRepo
}

// ChangePassword implements interfaces.AdminUseCase.
func (*adminUseCase) ChangePassword(Admin domain.Admin) error {
	panic("unimplemented")
}

// Create implements interfaces.AdminUseCase.
func (*adminUseCase) Create(admin domain.Admin) error {
	panic("unimplemented")
}

// FindByAdminEmail implements interfaces.AdminUseCase.
func (*adminUseCase) FindByAdminEmail(admin domain.Admin) (domain.Admin, error) {
	panic("unimplemented")
}

// FindByAdminName implements interfaces.AdminUseCase.
func (*adminUseCase) FindByAdminName(user domain.Admin) (domain.Admin, error) {
	panic("unimplemented")
}

func NewAdminUseCase(repo intereface.AdminRepo) usecas.AdminUseCase {
	return &adminUseCase{
		Repo: repo,
	}
}
