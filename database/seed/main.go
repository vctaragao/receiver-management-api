package main

import (
	"fmt"
	"math/rand"

	"github.com/brianvoe/gofakeit"
	"github.com/dimiro1/faker"
	"github.com/google/uuid"
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
	"github.com/vctaragao/receiver-management-api/internal/storage"
	"github.com/vctaragao/receiver-management-api/internal/storage/schemas"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Connecting to database...")
	repo := storage.NewPostgress()
	fmt.Println("Database suceffuly connected!")

	cleanDatabase(repo.Db)

	fmt.Println("Seeding...")
	seedInBatch(repo.Db, 30)
	fmt.Println("Database succesfully seeded!")

}

func cleanDatabase(db *gorm.DB) {
	db.Unscoped().Delete(&schemas.Pix{}, "1=1")
	db.Unscoped().Delete(&schemas.Receiver{}, "1=1")
}

func seedInBatch(db *gorm.DB, size int) {
	brFaker, _ := faker.NewForLocale("pt-br")
	gofakeit.Seed(0)
	for i := 0; i < size; i++ {
		cpfCnpj := brFaker.BrazilCPF()
		if rand.Intn(2) == 1 {
			cpfCnpj = brFaker.BrazilCNPJ()
		}

		status := entity.GetValidReciverStatus()[rand.Intn(2)]

		receiver := schemas.Receiver{
			CorporateName: gofakeit.Name(),
			Email:         gofakeit.Email(),
			CpfCnpj:       cpfCnpj,
			Status:        status,
		}

		t := entity.GetValidTypes()[rand.Intn(4)]

		key := ""
		switch t {
		case entity.EMAIL:
			key = gofakeit.Email()
		case entity.CPF:
			key = brFaker.BrazilCPF()
		case entity.CNPJ:
			key = brFaker.BrazilCNPJ()
		case entity.PHONE:
			key = brFaker.CellPhoneNumber()
		case entity.RANDOM_KEY:
			key = uuid.New().String()
		}

		pix := schemas.Pix{
			Type: t,
			Key:  key,
		}

		receiver.Pix = append(receiver.Pix, pix)

		db.Create(&receiver)
	}
}
