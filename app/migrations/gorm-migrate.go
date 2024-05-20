package migrations

import (
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/app/configs"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/app/databases"
	feedbackData "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/feedback/data"
	homestayData "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/homestays/data"
	reservationData "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/reservation/data"
	userData "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/user/data"
)

func InitialMigration() {
	databases.InitDBMysql(configs.InitConfig()).AutoMigrate(&userData.User{}, &homestayData.Homestay{}, &reservationData.Reservation{}, &feedbackData.Feedback{})
}
