package usecase

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/tiago-g-sales/stress-test-goexpert/internal/model"
)

func Execute(m model.ParameterRequestDTO) {
	timeInit := time.Now()	
	report := model.ReportTotalInfo{}
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(int(m.Requests))	
	var totalLista int64
	
	go request(m, ch )
	for i := int64(1); i <= m.Concurrency; i++ {
	 	go	ExecuteConcurrency(&i,  m.Url, &report, &wg,ch)
	}
	wg.Wait()
	timeFinish := time.Now()
	report.TempoTotal = float64(timeFinish.Sub(timeInit).Seconds())   	

	for _, v := range report.StatusNOK {	
		totalLista += v.Count
	}

	fmt.Printf("Tempo Total Execução--------: %f \n", report.TempoTotal)
	fmt.Printf("Total Requests--------------: %d \n", report.RequestTotal)
	fmt.Printf("Total Requests HTTP 200-----: %d \n", report.Status200)	
	fmt.Printf("Total Requests HTTP outros---: %d \n", totalLista)



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
			report.Status200++	
		} else {
			report.StatusNOK = append(report.StatusNOK, model.StatusNOK{StatusCode: res.StatusCode, Count: 1})			
		}


		if i != 0 {}
		report.RequestTotal++
		wg.Done()

		
	}


}


func request(m model.ParameterRequestDTO,  ch chan int) {
	for i := 0; i <= int(m.Requests); i++ {
		ch <- i
	}
	close(ch)
	
	
}