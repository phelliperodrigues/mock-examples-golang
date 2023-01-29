package mockery_test

import (
	"fmt"
	"testing"

	"github.com/vektra/mockery/mockery"
)

type DB interface {
	Query(query string) ([]byte, error)
}

func TestMock(t *testing.T) {
	// Gerando o mock
	generator, err := mockery.NewGenerator()
	if err != nil {
		t.Fatalf("Não foi possível criar o gerador de mock: %s", err)
	}
	err = generator.Start("./mocks")
	if err != nil {
		t.Fatalf("Não foi possível iniciar o gerador de mock: %s", err)
	}
	err = generator.Generate("DB")
	if err != nil {
		t.Fatalf("Não foi possível gerar o mock: %s", err)
	}

	mockDB := new(mocks.DB)
	mockDB.On("Query", "SELECT * FROM users").Return([]byte("fake data"), nil)

	// Usando o mock no código
	data, _ := mockDB.Query("SELECT * FROM users")
	fmt.Println(string(data))
}
