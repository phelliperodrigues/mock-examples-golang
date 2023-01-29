package ginkgo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type DB interface {
	Query(query string) ([]byte, error)
}

type DBTest struct{}

func (db *DBTest) Query(query string) ([]byte, error) {
	if query == "SELECT * FROM users" {
		return []byte("fake data"), nil
	}
	return nil, nil
}

var _ = Describe("TestMain", func() {
	It("Deve retornar os dados fake corretos", func() {
		// Gerando o mock
		mockDB := &DBTest{}
		// Usando o mock no código
		data, _ := mockDB.Query("SELECT * FROM users")
		Expect(string(data)).To(Equal("fake data"))
	})
})
