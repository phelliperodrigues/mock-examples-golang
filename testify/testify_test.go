package testify__test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DB interface {
	Query(query string) ([]byte, error)
}

type DBTest struct {
	mock.Mock
}

func (db *DBTest) Query(query string) ([]byte, error) {
	args := db.Called(query)
	return args.Get(0).([]byte), args.Error(1)
}

func TestMock(t *testing.T) {
	// Gerando o mock
	mockDB := new(DBTest)
	mockDB.On("Query", "SELECT * FROM users").Return([]byte("fake data"), nil)

	// Usando o mock no código
	data, _ := mockDB.Query("SELECT * FROM users")
	assert.Equal(t, []byte("fake data"), data, "Dados devem ser iguais a 'fake data'")
}
