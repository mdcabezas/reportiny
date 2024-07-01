package test

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	report "github.com/mdcabezas/reportiny/internal/report"
	repositoryMock "github.com/mdcabezas/reportiny/internal/report/mock"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := repositoryMock.NewMockIRepository(ctrl)

	data := report.RepositoryModel{
		ID:          "1",
		Image:       "http://via.placeholder.com/640x360",
		Description: "Description-1",
		User:        "gonzalomunoz@outlook.com",
		Latitude:    -35.290073,
		Longitude:   -71.270594,
	}

	dataArray := []report.RepositoryModel{
		{
			Image:       "http://via.placeholder.com/640x360",
			Description: "Primer elemento",
			User:        "gonzalomunoz@outlook.com",
			Latitude:    -35.290073,
			Longitude:   -71.270594,
		},
		{
			Image:       "http://via.placeholder.com/640x360",
			Description: "Segundo elemento",
			User:        "jperez@gmail.com",
			Latitude:    -35.000000,
			Longitude:   -71.000000,
		},
	}

	t.Run("Create Item Success", func(t *testing.T) {
		mockRepository.EXPECT().Create(data).Return(data, nil)
		resp, err := mockRepository.Create(data)

		assert.Nil(t, err)
		assert.Equal(t, resp, data)
	})

	t.Run("Read All Items Success", func(t *testing.T) {
		mockRepository.EXPECT().ReadAll().Return(dataArray)
		resp := mockRepository.ReadAll()

		assert.Equal(t, resp, dataArray)
	})

	t.Run("Update Item Success", func(t *testing.T) {
		mockRepository.EXPECT().Update(data.ID, data).Return(nil)
		err := mockRepository.Update(data.ID, data)

		assert.Nil(t, err)
	})

	t.Run("Delete Item Success", func(t *testing.T) {
		mockRepository.EXPECT().Delete(gomock.Any()).Return(nil)
		err := mockRepository.Delete(data.ID)
		assert.Nil(t, err)
	})

}
