package test

import (
	"fmt"
	"sync"
	"time"

	pkg_http "github.com/BMokarzel/stress-test/pkg/http"
	logger "github.com/BMokarzel/stress-test/pkg/logger"
)

type Test struct {
	Logger *logger.Logger
	Client pkg_http.Client
}

type Result struct {
	Status       int
	ErrorMessage string
}

func New(logger *logger.Logger, url string) (*Test, error) {
	return &Test{
		Logger: logger,
		Client: *pkg_http.New(url),
	}, nil
}

func (t *Test) Run(url string, r, c int) error {

	var wg sync.WaitGroup

	jobs := make(chan int, r)
	results := make(chan Result, r)

	for x := 0; x < c; x++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			t.Worker(r, jobs, results)
		}()
	}

	start := time.Now()

	for j := 0; j < r; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	var errorMessages []string

	statusCount := make(map[int]int)
	for res := range results {
		statusCount[res.Status]++
		errorMessages = append(errorMessages, res.ErrorMessage)
	}

	fmt.Println("\n---- Relatório ----")
	fmt.Printf("\nRequests totais: %d\n", r)
	fmt.Printf("Concorrência: %d\n", c)
	fmt.Printf("Tempo total de execução: %v\n", time.Since(start))
	fmt.Printf("\nDistribuição de status codes: \n\n")
	fmt.Printf("HTTP 200: %d\n", statusCount[200])
	for code, count := range statusCount {
		if code == 0 {
			fmt.Printf("Erros: %d\n", count)
			continue
		}
		if code == 200 {
			continue
		}
		fmt.Printf("HTTP %d: %d\n", code, count)
	}

	if statusCount[0] != 0 {
		err := t.Logger.Batch(errorMessages)
		if err != nil {
			fmt.Printf("Erro ao enviar os logs das requisições para o arquivo. Error: %s\n", err)
		}
	}

	return nil
}

func (t *Test) Worker(requests int, jobs <-chan int, result chan<- Result) {
	for range jobs {
		code, err := t.Client.Call()
		if err != nil {
			result <- Result{Status: 0, ErrorMessage: err.Error()}
			continue
		}
		result <- Result{Status: code}
	}
}
