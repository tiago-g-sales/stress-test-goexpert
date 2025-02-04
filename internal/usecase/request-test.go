package usecase

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/tiago-g-sales/stress-test-goexpert/internal/model"
)




func Execute(model model.ParameterRequestDTO) {
	var wg sync.WaitGroup


	fmt.Printf("URL: %s \n" , model.Url)
	fmt.Printf("Requests: %d \n", model.Requests)	
	fmt.Printf("Concurrency: %d \n", model.Concurrency)	


	concorrencia := make(chan int, model.Concurrency)



	for i := int64(0); i < model.Requests; i++ {
		wg.Add(1)
		concorrencia <- 1	
		go func (){

				defer wg.Done()
				defer func() { <-concorrencia }()
				
				req, err := http.NewRequest("GET", model.Url, nil)
				if err != nil {
					panic(err)
				}
				res, err := http.DefaultClient.Do(req)
				if err != nil {
					panic(err)
				}
				if res.StatusCode == 200 {
					print("ok")
				} else {
					print("nok")
				}
			
				defer res.Body.Close()

			}()
		}

	wg.Wait()
}