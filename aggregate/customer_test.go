package aggregate_test

import (
	"github.com/MaksimDzhangirov/DDD-and-go/aggregate"
	"testing"
)

func TestCustomer_NewCustomer(t *testing.T) {
	// Создаём необходимую нам структуру данных для тестового случая
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}
	// Создаём новые тестовые случаи
	testCases := []testCase{
		{
			test:        "Empty Name validation",
			name:        "",
			expectedErr: aggregate.ErrInvalidPerson,
		}, {
			test:        "Valid Name",
			name:        "Percy Bolmer",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		// Запускаем тесты
		t.Run(tc.test, func(t *testing.T) {
			// Создаём нового клиента
			_, err := aggregate.NewCustomer(tc.name)
			// Проверяем, соответствует ли ошибка ожидаемой
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
