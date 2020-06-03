package models


type Dept struct {
	DeptId    int    `json:"deptId" gorm:"primary_key;AUTO_INCREMENT"` //部门编码
	ParentId  int    `json:"parentId" gorm:"type:int(11);"`            //上级部门
	DeptPath  string `json:"deptPath" gorm:"type:varchar(255);"`       //
	DeptName  string `json:"deptName"  gorm:"type:varchar(128);"`      //部门名称
	Sort      int    `json:"sort" gorm:"type:int(4);"`                 //排序
	Leader    string `json:"leader" gorm:"type:varchar(128);"`         //负责人
	Phone     string `json:"phone" gorm:"type:varchar(11);"`           //手机
	Email     string `json:"email" gorm:"type:varchar(64);"`           //邮箱
	Status    string `json:"status" gorm:"type:int(1);"`               //状态
	CreateBy  string `json:"createBy" gorm:"type:varchar(64);"`
	UpdateBy  string `json:"updateBy" gorm:"type:varchar(64);"`
	DataScope string `json:"dataScope" gorm:"-"`
	Params    string `json:"params" gorm:"-"`
	Children  []Dept `json:"children" gorm:"-"`
	BaseModel
}

func (Dept) TableName() string {
	return "sys_dept"
}