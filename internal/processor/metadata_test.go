package processor

import "testing"

type MockStorage struct {
	Called bool
}

func (m *MockStorage) Save(key string, data []byte) {
	m.Called = true
}

func TestProcessImage(t *testing.T) {
	mock := &MockStorage{}

	processImage(1, mock)

	if !mock.Called {
		t.Error("Expected Save to be called")
	}
}
