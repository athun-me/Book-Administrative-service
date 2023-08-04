package repoesitory

import (
	"githum.com/athunlal/bookNowAdmin-svc/pkg/domain"
	intereface "githum.com/athunlal/bookNowAdmin-svc/pkg/repository/interface"
	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

// FindAdminById implements intereface.AdminRepo.
func (r *adminDatabase) FindAdminById(adminid uint) (domain.Admin, error) {
	admin := domain.Admin{}
	result := r.DB.First(&admin, "id = ?", adminid).Error
	return admin, result
}

// DeleteUser implements intereface.AdminRepo.
func (r *adminDatabase) DeleteUser(user domain.Admin) error {
	result := r.DB.Exec("DELETE FROM users WHERE email LIKE ?", user.Email).Error
	return result
}

// ChangePassword implements intereface.AdminRepo.
func (r *adminDatabase) ChangePassword(Admin domain.Admin) error {
	result := r.DB.Model(&Admin).Where("id = ?", 1).Update("password", Admin.Password)
	return result.Error
}

// Create implements intereface.AdminRepo.
func (r *adminDatabase) Create(admin domain.Admin) error {
	result := r.DB.Create(&admin).Error
	return result
}

// FindByAdminEmail implements intereface.AdminRepo.
func (r *adminDatabase) FindByAdminEmail(admin domain.Admin) (domain.Admin, error) {
	result := r.DB.First(&admin, "email LIKE ?", admin.Email).Error
	return admin, result
}

// FindByAdminName implements intereface.AdminRepo.
func (r *adminDatabase) FindByAdminName(admin domain.Admin) (domain.Admin, error) {
	result := r.DB.First(&admin, "username LIKE ?", admin.Username).Error
	return admin, result
}

func NewAdminRepo(db *gorm.DB) intereface.AdminRepo {
	return &adminDatabase{
		DB: db,
	}
}
