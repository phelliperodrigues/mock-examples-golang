//go:generate go run github.com/vektra/mockery/v2 --name=DB --with-expecter --filename	db.go

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
	mockDB := new(mocks.DB)
	mockDB.On("Query", "SELECT * FROM users").Return([]byte("fake data"), nil)

	// Usando o mock no código
	data, _ := mockDB.Query("SELECT * FROM users")
	fmt.Println(string(data))
}
