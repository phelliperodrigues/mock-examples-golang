package interfaces

import (
	"testing"
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

func TestMock(t *testing.T) {
	// Gerando o mock
	mockDB := &DBTest{}

	// Usando o mock no código
	data, _ := mockDB.Query("SELECT * FROM users")
	if string(data) != "fake data" {
		t.Errorf("Dados devem ser iguais a 'fake data', mas recebidos '%s'", string(data))
	}
}
