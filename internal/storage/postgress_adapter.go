package storage

import (
	"os"

	"github.com/vctaragao/receiver-management-api/internal/storage/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgress struct {
	Db *gorm.DB
}

func NewPostgress() *Postgress {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	if dbPort == "" {
		dbPort = "80"
	}

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("unable to connect to database")
	}

	db.AutoMigrate(&schemas.Receiver{})
	db.AutoMigrate(&schemas.Pix{})

	return &Postgress{
		Db: db,
	}
}
