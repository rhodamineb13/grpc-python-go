package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"grpccode/configs"
	"grpccode/database/entities"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := "host=" + configs.DBConfig.Host + " " + "port=" + configs.DBConfig.Port + " " + "user=" + configs.DBConfig.User + " " + "password=" + configs.DBConfig.Pass + " " +
		"dbname=" + configs.DBConfig.Name + " "
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&entities.Users{})
}
