package adapters

import (
	"transfeera.backend.developer.test/api/domain"
	"transfeera.backend.developer.test/api/repositories"
)

type GormBeneficiaryRepository struct{}

func NewBeneficiaryRepository() repositories.BeneficiaryRepository {
	return GormBeneficiaryRepository{}
}

func (g GormBeneficiaryRepository) Get(filter string) ([]domain.Beneficiary, error) {
	//return nil, errors.New("something very bad")
	return []domain.Beneficiary{
		{
			Id:             1,
			Status:         "active",
			Name:           "John Doe",
			DocumentNumber: "123.456.789-00",
			Email:          "johndoe@example.com",
			PixKeyType:     "CPF",
			PixKeyValue:    "123.456.789-00",
			Bank:           "NuBank",
			Agency:         "0001",
			Account:        "123456789-0",
		},
		{
			Id:             2,
			Status:         "inactive",
			Name:           "Jane Doe",
			DocumentNumber: "00.000.000/0000-00",
			Email:          "janedoe@example.com",
			PixKeyType:     "CNPJ",
			PixKeyValue:    "00.000.000/0000-00",
			Bank:           "Banco do Brasil",
			Agency:         "0002",
			Account:        "987654321-0",
		},
		{
			Id:             3,
			Status:         "pending",
			Name:           "Bob Smith",
			DocumentNumber: "111.222.333-44",
			Email:          "bobsmith@example.com",
			PixKeyType:     "CPF",
			PixKeyValue:    "111.222.333-44",
			Bank:           "Itau",
			Agency:         "0003",
			Account:        "55555555-5",
		},
	}, nil
}

func (g GormBeneficiaryRepository) Create(beneficiary domain.Beneficiary) (domain.Beneficiary, error) {
	//TODO implement me
	panic("implement me")
}

func (g GormBeneficiaryRepository) Update(beneficiary domain.Beneficiary) (domain.Beneficiary, error) {
	//TODO implement me
	panic("implement me")
}

func (g GormBeneficiaryRepository) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
