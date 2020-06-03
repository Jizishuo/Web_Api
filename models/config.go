package models


type SysConfig struct {
	ConfigId    int    `json:"configId" gorm:"primary_key;auto_increment;"` //编码
	ConfigName  string `json:"configName" gorm:"type:varchar(128);"`        //参数名称
	ConfigKey   string `json:"configKey" gorm:"type:varchar(128);"`         //参数键名
	ConfigValue string `json:"configValue" gorm:"type:varchar(255);"`       //参数键值
	ConfigType  string `json:"configType" gorm:"type:varchar(64);"`         //是否系统内置
	Remark      string `json:"remark" gorm:"type:varchar(128);"`            //备注
	CreateBy    string `json:"createBy" gorm:"type:varchar(128);"`
	UpdateBy    string `json:"updateBy" gorm:"type:varchar(128);"`
	DataScope   string `json:"dataScope" gorm:"-"`
	Params      string `json:"params"  gorm:"-"`
	BaseModel
}

func (SysConfig) TableName() string {
	return "sys_config"
}
