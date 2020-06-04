package models

import orm "Web_Api/database"

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
type MenuPath struct {
	Path string `json:"path"`
}


func (rm *RoleMenu) GetIDS() ([]MenuPath, error) {
	var r []MenuPath
	table := orm.Eloquent.Select("sys_menu.path").Table("sys_role_menu")
	table = table.Joins("left join sys_role on sys_role.role_id=sys_role_menu.role_id")
	table = table.Joins("left join sys_menu on sys_menu.id=sys_role_menu.menu_id")
	table = table.Where("sys_role.role_name = ? and sys_menu.type=1", rm.RoleName)
	if err := table.Find(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}