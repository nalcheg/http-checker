package application

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nalcheg/http-checker/mocks"
)

func TestExample(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedRepository := mocks.NewMockRepositoryInterface(mockCtrl)
	mockedRepository.EXPECT().GetHosts().AnyTimes()

	app, err := NewApplication(mockedRepository, 10)
	if err != nil {
		t.Fatal(err)
	}

	if app.Load() == true {
		t.Error("application initialized in running state")
	}
}
