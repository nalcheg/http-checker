package application

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/nalcheg/http-checker/repository"
	"github.com/nalcheg/http-checker/types"
	"github.com/robfig/cron/v3"
	"go.uber.org/atomic"
)

type Application struct {
	atomic.Bool

	repo  repository.RepositoryInterface
	hosts []string
	chIn  chan request
	chOut chan types.Result
}

type request struct {
	wg   *sync.WaitGroup
	host string
}

func NewApplication(repo repository.RepositoryInterface, wc int64) (*Application, error) {
	app := Application{}

	chIn := make(chan request)
	chOut := make(chan types.Result)
	app.chIn = chIn
	app.chOut = chOut

	app.repo = repo

	app.Store(false)

	for i := int64(0); i < wc; i++ {
		go app.worker(app.chIn, app.chOut)
	}

	hosts, err := repo.GetHosts()
	if err != nil {
		return nil, err
	}
	app.hosts = hosts

	return &app, nil
}

func (a *Application) check() {
	if a.Load() == true {
		log.Print("ALARM !!! prevous check() is running, needed to figure out with this") // TODO maybe use log.Fatal
		return
	}
	a.Store(true)

	wg := sync.WaitGroup{}
	for _, host := range a.hosts {
		wg.Add(1)
		a.chIn <- request{
			wg:   &wg,
			host: host,
		}
	}
	wg.Wait()

	a.Store(false)
}

func (a *Application) worker(chIn chan request, chOut chan types.Result) {
	checkHost := func(host string) types.Result {
		r := types.Result{Host: host}

		start := time.Now()

		client := http.Client{Timeout: time.Second * 10}
		resp, err := client.Get(host)
		if err != nil {
			r.ResponseCode = 0
		} else {
			r.ResponseCode = resp.StatusCode
		}

		r.ResponseTime = float64(time.Since(start).Microseconds()) / 1000000
		r.Time = time.Now()

		return r
	}

	for {
		select {
		case request := <-chIn:
			r := checkHost(request.host)
			r.Wg = request.wg
			chOut <- r
		}
	}
}

func (a *Application) Start(cronString string) error {
	if cronString == "" {
		cronString = "*/10 * * * * *"
	}

	c := cron.New(cron.WithSeconds())
	if _, err := c.AddFunc(cronString, func() {
		a.check()
	}); err != nil {
		return err
	}
	c.Start()

	go func() {
		for {
			select {
			case r := <-a.chOut:
				if err := a.repo.SaveCheck(r); err != nil {
					log.Print(err)
				}

				r.Wg.Done()
				log.Printf("%s checked, code - %d, time - %f", r.Host, r.ResponseCode, r.ResponseTime)
			}
		}
	}()

	return nil
}
