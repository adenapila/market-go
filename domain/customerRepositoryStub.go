package domain

type CustomerRepositoryStub struct {
	customer []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customer, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Ade", City: "Tangerang", Zipcode: "15418", DateofBirth: "2000-12-21", Status: "1"},
		{Id: "1002", Name: "Napila", City: "Jakarta", Zipcode: "15215", DateofBirth: "2000-01-20", Status: "1"},
	}
	return CustomerRepositoryStub{customer: customers}
}
