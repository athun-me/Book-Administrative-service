package repoesitory

import (
	"githum.com/athunlal/bookNowAdmin-svc/pkg/domain"
	intereface "githum.com/athunlal/bookNowAdmin-svc/pkg/repository/interface"
	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

// ChangePassword implements intereface.AdminRepo.
func (*adminDatabase) ChangePassword(Admin domain.Admin) error {
	panic("unimplemented")
}

// Create implements intereface.AdminRepo.
func (*adminDatabase) Create(admin domain.Admin) error {
	panic("unimplemented")
}

// FindByAdminEmail implements intereface.AdminRepo.
func (*adminDatabase) FindByAdminEmail(admin domain.Admin) (domain.Admin, error) {
	panic("unimplemented")
}

// FindByAdminName implements intereface.AdminRepo.
func (*adminDatabase) FindByAdminName(user domain.Admin) (domain.Admin, error) {
	panic("unimplemented")
}

func NewAdminRepo(db *gorm.DB) intereface.AdminRepo {
	return &adminDatabase{
		DB: db,
	}
}
