package database

import "Web_Api/tools/config"

func Setup()  {
	dbType := config.DatabaseConfig.Database
	if dbType == "mysql" {
		var db = new(Mysql)
		db.Setup()
	}
}
