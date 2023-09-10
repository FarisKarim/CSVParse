package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
)

type Org struct {
	Index               string
	Organization_ID     string
	Name                string
	Website             string
	Country             string
	Description         string
	Founded             string
	Industry            string
	Number_of_employees string
}

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	if _, err = reader.Read(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	dataCh := make(chan Org)
	errCh := make(chan error)

	var wg sync.WaitGroup

	go func() {
		defer close(dataCh)
		defer close(errCh)

		for {
			record, err := reader.Read()
			if err != nil {
				if err != io.EOF {
					errCh <- err
				}
				break
			}
			wg.Add(1)
			go processRecord(record, dataCh, errCh, &wg)
		}
		wg.Wait()
	}()

	for data := range dataCh {
		fmt.Println("Received data:", data)
	}

	for err := range errCh {
		fmt.Println("Received error:", err)
	}
}

func processRecord(data []string, dataCh chan<- Org, errCh chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	if len(data) != 9 {
		errCh <- fmt.Errorf("invalid record: %v", data)
		return
	}
	dataCh <- Org{Index: data[0], Organization_ID: data[1], Name: data[2], Website: data[3], Country: data[4], Description: data[5], Founded: data[6], Industry: data[7], Number_of_employees: data[8]}
}
