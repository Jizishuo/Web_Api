package models


//sys_role_dept
type SysRoleDept struct {
	RoleId int `gorm:"type:int(11)"`
	DeptId int `gorm:"type:int(11)"`
}

func (SysRoleDept) TableName() string {
	return "sys_role_dept"
}