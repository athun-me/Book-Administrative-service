package domain

type Admin struct {
	Id       uint   `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Username string `json:"adminname" validate:"required,min=8,max=24"`
	Password string `json:"password" validate:"required,min=8,max=16"`
	Email    string `json:"email" validate:"email,required"`
}
