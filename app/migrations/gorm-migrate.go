package migrations

import (
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/app/configs"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/app/databases"
	homestayData "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/homestays/data"
	userData "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/user/data"
)

func InitialMigration() {
	databases.InitDBMysql(configs.InitConfig()).AutoMigrate(&userData.User{}, &homestayData.Homestay{})
}
