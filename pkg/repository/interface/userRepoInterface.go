package intereface

import "githum.com/athunlal/bookNowAdmin-svc/pkg/domain"

type AdminRepo interface {
	Create(admin domain.Admin) error
	FindByAdminName(user domain.Admin) (domain.Admin, error)
	FindByAdminEmail(admin domain.Admin) (domain.Admin, error)
	ChangePassword(Admin domain.Admin) error
	DeleteUser(user domain.Admin) error
}
