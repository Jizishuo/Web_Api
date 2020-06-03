package models


type DictType struct {
	DictId    int    `gorm:"primary_key;AUTO_INCREMENT" json:"dictId"`
	DictName  string `gorm:"type:varchar(128);" json:"dictName"` //字典名称
	DictType  string `gorm:"type:varchar(128);" json:"dictType"` //字典类型
	Status    string `gorm:"type:int(1);" json:"status"`         //状态
	DataScope string `gorm:"-" json:"dataScope"`                 //
	Params    string `gorm:"-" json:"params"`                    //
	CreateBy  string `gorm:"type:varchar(11);" json:"createBy"`  //创建者
	UpdateBy  string `gorm:"type:varchar(11);" json:"updateBy"`  //更新者
	Remark    string `gorm:"type:varchar(255);" json:"remark"`   //备注
	BaseModel
}

func (DictType) TableName() string {
	return "sys_dict_type"
}