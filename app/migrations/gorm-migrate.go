package migrations

import (
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/app/configs"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/app/databases"
	userData "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/user/data"
)

func InitialMigration() {
	databases.InitDBMysql(configs.InitConfig()).AutoMigrate(&userData.User{})
}
