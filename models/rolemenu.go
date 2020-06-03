package models


type RoleMenu struct {
	RoleId   int  `gorm:"type:int(11)"`
	MenuId   int  `gorm:"type:int(11)"`
	RoleName string `gorm:"type:varchar(128)"`
	CreateBy string `gorm:"type:varchar(128)"`
	UpdateBy string `gorm:"type:varchar(128)"`
}

func (RoleMenu) TableName() string {
	return "sys_role_menu"
}