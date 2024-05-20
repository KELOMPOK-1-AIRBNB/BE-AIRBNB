package migrations

import (
	"myTaskApp/app/configs"
	"myTaskApp/app/databases"
	userData "myTaskApp/features/user/data"
)

func InitialMigration() {
	databases.InitDBMysql(configs.InitConfig()).AutoMigrate(&userData.User{})
}
