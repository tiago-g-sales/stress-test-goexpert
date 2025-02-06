package usecase

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/tiago-g-sales/stress-test-goexpert/internal/model"
)

func Execute(m model.ParameterRequestDTO) {
	report := model.ReportTotalInfo{}
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(int(m.Requests))	
	
	go request(m, ch )
	for i := int64(1); i <= m.Concurrency; i++ {
	 	go	ExecuteConcurrency(&i,  m.Url, &report, &wg,ch)
	}
	wg.Wait()
	fmt.Printf("Total Requests OK: %d \n", report.StatusOK)	
	fmt.Printf("Total Requests NOK: %d \n", report.StatusNOK)
}




func ExecuteConcurrency(i *int64, m string, report  *model.ReportTotalInfo, wg *sync.WaitGroup,  ch chan int)  {
	fmt.Printf("Executing goroutine %d \n", uint64(*i))

	for i := range ch {
		req, err := http.NewRequest("GET", m, nil)
		if err != nil {
			panic(err)
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		if res.StatusCode == 200 {
			report.StatusOK++	
		} else {
			report.StatusNOK++
		}
		
		if i != 0 {}	
		wg.Done()
	}

}


func request(m model.ParameterRequestDTO,  ch chan int) {
	for i := 0; i <= int(m.Requests); i++ {
		ch <- i
	}
	close(ch)
	
	
}