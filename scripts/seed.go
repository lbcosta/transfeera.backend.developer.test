package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"transfeera.backend.developer.test/src/api/v1/repositories/model"
)

func main() {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s timezone=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		"America/Sao_Paulo")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.Beneficiary{})
	beneficiaries := []model.Beneficiary{
		{Status: "Validado", Name: "John Doe", DocumentNumber: "12345678907", Email: "johndoe@example.com", PixKeyType: "CPF", PixKeyValue: "12345678907", Bank: "ABC Bank", Agency: "1234", Account: "56789"},
		{Status: "Validado", Name: "Jane Doe", DocumentNumber: "98765432107", Email: "janedoe@example.com", PixKeyType: "CPF", PixKeyValue: "98765432107", Bank: "XYZ Bank", Agency: "5678", Account: "12345"},
		{Status: "Validado", Name: "Bob Smith", DocumentNumber: "45678912307", Email: "bobsmith@example.com", PixKeyType: "CPF", PixKeyValue: "45678912307", Bank: "DEF Bank", Agency: "2345", Account: "67890"},
		{Status: "Rascunho", Name: "Alice Johnson", DocumentNumber: "78945612307", Email: "alicejohnson@example.com", PixKeyType: "CNPJ", PixKeyValue: "78945612307", Bank: "GHI Bank", Agency: "3456", Account: "78901"},
		{Status: "Rascunho", Name: "David Lee", DocumentNumber: "32165498707", Email: "davidlee@example.com", PixKeyType: "CNPJ", PixKeyValue: "32165498707", Bank: "JKL Bank", Agency: "4567", Account: "89012"},
		{Status: "Validado", Name: "Sarah Brown", DocumentNumber: "65432198707", Email: "sarahbrown@example.com", PixKeyType: "EMAIL", PixKeyValue: "sarahbrown@example.com", Bank: "MNO Bank", Agency: "5678", Account: "90123"},
		{Status: "Validado", Name: "Michael Davis", DocumentNumber: "14785236907", Email: "michaeldavis@example.com", PixKeyType: "CPF", PixKeyValue: "14785236907", Bank: "PQR Bank", Agency: "6789", Account: "01234"},
		{Status: "Rascunho", Name: "Emily Clark", DocumentNumber: "25836914707", Email: "emilyclark@example.com", PixKeyType: "EMAIL", PixKeyValue: "emilyclark@example.com", Bank: "STU Bank", Agency: "7890", Account: "12345"},
		{Status: "Rascunho", Name: "Daniel White", DocumentNumber: "36985214740", Email: "danielwhite@example.com", PixKeyType: "CNPJ", PixKeyValue: "36985214740", Bank: "VWX Bank", Agency: "8901", Account: "23456"},
		{Status: "Validado", Name: "Karen Garcia", DocumentNumber: "96325874140", Email: "karengarcia@example.com", PixKeyType: "CNPJ", PixKeyValue: "96325874140", Bank: "YZA Bank", Agency: "9012", Account: "34567"},
		{Status: "Validado", Name: "Alex Johnson", DocumentNumber: "25814736940", Email: "alexjohnson@example.com", PixKeyType: "CPF", PixKeyValue: "25814736940", Bank: "ABC Bank", Agency: "1234", Account: "56789"},
		{Status: "Validado", Name: "Linda Smith", DocumentNumber: "45612378940", Email: "lindasmith@example.com", PixKeyType: "CNPJ", PixKeyValue: "45612378940", Bank: "DEF Bank", Agency: "2345", Account: "67890"},
		{Status: "Rascunho", Name: "Paul Martin", DocumentNumber: "96374185240", Email: "paulmartin@example.com", PixKeyType: "CPF", PixKeyValue: "96374185240", Bank: "GHI Bank", Agency: "3456", Account: "78901"},
		{Status: "Rascunho", Name: "Laura Taylor", DocumentNumber: "36914725840", Email: "laurataylor@example.com", PixKeyType: "CPF", PixKeyValue: "36914725840", Bank: "JKL Bank", Agency: "4567", Account: "89012"},
		{Status: "Validado", Name: "Timothy Allen", DocumentNumber: "85296374189", Email: "timothyallen@example.com", PixKeyType: "CPF", PixKeyValue: "85296374189", Bank: "MNO Bank", Agency: "5678", Account: "90123"},
		{Status: "Validado", Name: "Timothy Allen", DocumentNumber: "85296374189", Email: "timothyallen@example.com", PixKeyType: "CPF", PixKeyValue: "85296374189", Bank: "MNO Bank", Agency: "5678", Account: "90123"},
		{Status: "Validado", Name: "Rachel Lee", DocumentNumber: "78912345689", Email: "rachellee@example.com", PixKeyType: "CNPJ", PixKeyValue: "78912345689", Bank: "PQR Bank", Agency: "6789", Account: "01234"},
		{Status: "Rascunho", Name: "Grace Davis", DocumentNumber: "96385274189", Email: "gracedavis@example.com", PixKeyType: "EMAIL", PixKeyValue: "gracedavis@example.com", Bank: "YZA Bank", Agency: "9012", Account: "34567"},
		{Status: "Rascunho", Name: "Benjamin King", DocumentNumber: "32178945689", Email: "benjaminking@example.com", PixKeyType: "CPF", PixKeyValue: "32178945689", Bank: "STU Bank", Agency: "7890", Account: "12345"},
		{Status: "Validado", Name: "Mark Wilson", DocumentNumber: "36985214789", Email: "markwilson@example.com", PixKeyType: "CPF", PixKeyValue: "36985214789", Bank: "ABC Bank", Agency: "1234", Account: "56789"},
		{Status: "Rascunho", Name: "Jennifer Lee", DocumentNumber: "14725836989", Email: "jenniferlee@example.com", PixKeyType: "EMAIL", PixKeyValue: "jenniferlee@example.com", Bank: "XYZ Bank", Agency: "5678", Account: "12345"},
		{Status: "Validado", Name: "Thomas Jackson", DocumentNumber: "95175385289", Email: "thomasjackson@example.com", PixKeyType: "CPF", PixKeyValue: "95175385289", Bank: "DEF Bank", Agency: "2345", Account: "67890"},
		{Status: "Rascunho", Name: "Laura Rodriguez", DocumentNumber: "36925814740", Email: "laurarodriguez@example.com", PixKeyType: "CNPJ", PixKeyValue: "36925814740", Bank: "GHI Bank", Agency: "3456", Account: "78901"},
		{Status: "Validado", Name: "Jacob Green", DocumentNumber: "75395185240", Email: "jacobgreen@example.com", PixKeyType: "CNPJ", PixKeyValue: "75395185240", Bank: "JKL Bank", Agency: "4567", Account: "89012"},
		{Status: "Rascunho", Name: "Isabella Perez", DocumentNumber: "85296374102", Email: "isabellaperez@example.com", PixKeyType: "CPF", PixKeyValue: "85296374102", Bank: "MNO Bank", Agency: "5678", Account: "90123"},
		{Status: "Validado", Name: "William Davis", DocumentNumber: "95175385299", Email: "williamdavis@example.com", PixKeyType: "CPF", PixKeyValue: "95175385299", Bank: "PQR Bank", Agency: "6789", Account: "01234"},
		{Status: "Rascunho", Name: "Sophia Thompson", DocumentNumber: "75315985240", Email: "sophiathompson@example.com", PixKeyType: "CPF", PixKeyValue: "75315985240", Bank: "STU Bank", Agency: "7890", Account: "12345"},
		{Status: "Validado", Name: "Oliver Wright", DocumentNumber: "15975325840", Email: "oliverwright@example.com", PixKeyType: "CNPJ", PixKeyValue: "15975325840", Bank: "VWX Bank", Agency: "8901", Account: "23456"},
		{Status: "Rascunho", Name: "Ava Lewis", DocumentNumber: "85274196340", Email: "avalewis@example.com", PixKeyType: "EMAIL", PixKeyValue: "avalewis@example.com", Bank: "YZA Bank", Agency: "9012", Account: "34567"},
		{Status: "Validado", Name: "Andrew Brown", DocumentNumber: "753159852543", Email: "andrewbrown@example.com", PixKeyType: "EMAIL", PixKeyValue: "andrewbrown@example.com", Bank: "ABC Bank", Agency: "1234", Account: "56789"},
		{Status: "Rascunho", Name: "Emma Wilson", DocumentNumber: "852963741543", Email: "emmawilson@example.com", PixKeyType: "EMAIL", PixKeyValue: "emmawilson@example.com", Bank: "XYZ Bank", Agency: "5678", Account: "12345"},
		{Status: "Validado", Name: "Christopher Taylor", DocumentNumber: "753159852543", Email: "christophertaylor@example.com", PixKeyType: "CPF", PixKeyValue: "753159852543", Bank: "DEF Bank", Agency: "2345", Account: "67890"},
		{Status: "Rascunho", Name: "Victoria Perez", DocumentNumber: "963852741543", Email: "victoriaperez@example.com", PixKeyType: "CNPJ", PixKeyValue: "963852741543", Bank: "GHI Bank", Agency: "3456", Account: "78901"},
		{Status: "Validado", Name: "Nicholas Hernandez", DocumentNumber: "753951852543", Email: "nicholashernandez@example.com", PixKeyType: "CNPJ", PixKeyValue: "753951852543", Bank: "JKL Bank", Agency: "4567", Account: "89012"},
		{Status: "Rascunho", Name: "Elizabeth Martinez", DocumentNumber: "963852741543", Email: "elizabethmartinez@example.com", PixKeyType: "CPF", PixKeyValue: "963852741543", Bank: "MNO Bank", Agency: "5678", Account: "90123"},
		{Status: "Validado", Name: "Joshua Thompson", DocumentNumber: "753951852543", Email: "joshuathompson@example.com", PixKeyType: "CPF", PixKeyValue: "753951852543", Bank: "PQR Bank", Agency: "6789", Account: "01234"},
		{Status: "Rascunho", Name: "Natalie Jackson", DocumentNumber: "963852741543", Email: "nataliejackson@example.com", PixKeyType: "EMAIL", PixKeyValue: "nataliejackson@example.com", Bank: "STU Bank", Agency: "7890", Account: "12345"},
		{Status: "Validado", Name: "Daniel White", DocumentNumber: "753951852543", Email: "danielwhite@example.com", PixKeyType: "EMAIL", PixKeyValue: "danielwhite@example.com", Bank: "VWX Bank", Agency: "8901", Account: "12345"},
	}
	db.Create(&beneficiaries)

	connection, _ := db.DB()
	connection.Close()
}
