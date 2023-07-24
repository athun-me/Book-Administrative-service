package interfaces

import "githum.com/athunlal/bookNowAdmin-svc/pkg/domain"

type AdminUseCase interface {
	Login(Admin domain.Admin) (domain.Admin, error)
	ChangePassword(Admin domain.Admin) error
}
