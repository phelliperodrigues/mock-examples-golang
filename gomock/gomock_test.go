package gomock__test

import (
	"github.com/golang/mock/gomock"
	"testing"
)

type DB interface {
	Query(query string) ([]byte, error)
}

func TestMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// Gerando o mock
	mockDB := NewMockDB(ctrl)
	mockDB.EXPECT().Query("SELECT * FROM users").Return([]byte("fake data"), nil)

	// Usando o mock no código
	data, _ := mockDB.Query("SELECT * FROM users")
	if string(data) != "fake data" {
		t.Errorf("Dados devem ser iguais a 'fake data', mas recebidos '%s'", string(data))
	}
}
