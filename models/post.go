package models


type Post struct {
	PostId    int    `gorm:"primary_key;AUTO_INCREMENT" json:"postId"` //岗位编号
	PostName  string `gorm:"type:varchar(128);" json:"postName"`       //岗位名称
	PostCode  string `gorm:"type:varchar(128);" json:"postCode"`       //岗位代码
	Sort      int    `gorm:"type:int(4);" json:"sort"`                 //岗位排序
	Status    string `gorm:"type:int(1);" json:"status"`               //状态
	Remark    string `gorm:"type:varchar(255);" json:"remark"`         //描述
	CreateBy  string `gorm:"type:varchar(128);" json:"createBy"`
	UpdateBy  string `gorm:"type:varchar(128);" json:"updateBy"`
	DataScope string `gorm:"-" json:"dataScope"`
	Params    string `gorm:"-" json:"params"`
	BaseModel
}

func (Post) TableName() string {
	return "sys_post"
}
