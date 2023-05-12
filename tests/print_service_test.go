package test

import (
	"testing"

	"github.com/Ramadani354/tiket_museum/app/models"
	"github.com/Ramadani354/tiket_museum/app/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository
type MockPrintRepository struct {
	mock.Mock
}

func (m *MockPrintRepository) CreatePrint(print *models.Print) error {
	args := m.Called(print)
	return args.Error(0)
}

func TestCreatePrint(t *testing.T) {
	// Membuat instance mock repository
	mockRepo := new(MockPrintRepository)

	// Mengatur hasil panggilan mock repository
	mockRepo.On("CreatePrint", mock.AnythingOfType("*models.Print")).Return(nil)

	// Membuat instance printService dengan menggunakan mock repository
	printService := services.NewPrintService(mockRepo)

	// Memanggil fungsi yang akan diuji
	ticketID := uint(3)
	tanggalCetak := "2023-05-12"
	print, err := printService.CreatePrint(ticketID, tanggalCetak)

	// Memeriksa hasil pengujian
	assert.NoError(t, err)
	assert.NotNil(t, print)

	// Memeriksa apakah metode CreatePrint pada mock repository dipanggil
	mockRepo.AssertCalled(t, "CreatePrint", mock.AnythingOfType("*models.Print"))
}
