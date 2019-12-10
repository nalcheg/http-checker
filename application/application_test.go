package application

import (
	"log"
	"testing"
	"time"

	"github.com/fasthttp/router"
	"github.com/golang/mock/gomock"
	"github.com/valyala/fasthttp"

	"github.com/nalcheg/http-checker/mocks"
	"github.com/nalcheg/http-checker/types"
)

func TestExample(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedRepository := mocks.NewMockRepositoryInterface(mockCtrl)
	mockedRepository.EXPECT().GetHosts().AnyTimes()

	app, err := NewApplication(mockedRepository, 1)
	if err != nil {
		t.Fatal(err)
	}

	if app.Load() == true {
		t.Error("application initialized in running state")
	}
}

func TestResponseCode(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedRepository := mocks.NewMockRepositoryInterface(mockCtrl)
	mockedRepository.EXPECT().GetHosts().Return([]string{"http://127.0.0.1:8080"}, nil)
	mockedRepository.EXPECT().SaveCheck(gomock.Any()).AnyTimes().Do(func(result types.Result) {
		if result.ResponseCode != 502 {
			t.Error()
		}
	})

	r := router.New()
	r.GET("/", func(ctx *fasthttp.RequestCtx) {
		ctx.Response.SetStatusCode(502)
	})

	go func() {
		log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
	}()

	app, err := NewApplication(mockedRepository, 1)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		if err := app.Start("* * * * * *"); err != nil {
			log.Fatal(err)
		}
	}()

	time.Sleep(2 * time.Second)
	app.chStop <- 1
}
