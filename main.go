package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
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

var organizations []Org // In-memory data store
var mu sync.Mutex       // Mutex for thread-safety

func main() {
	r := gin.Default()

	r.GET("/organizations", func(c *gin.Context) {
		mu.Lock()
		defer mu.Unlock()

		c.JSON(200, organizations)
	})

	r.GET("/search", searchHandler)
	r.POST("/update", updateHandler)

	loadOrganizations("data.csv") // Read and process the CSV when the app starts

	corsMiddleware := cors.Default()
	http.ListenAndServe(":8080", corsMiddleware.Handler(r))
}

func searchHandler(c *gin.Context) {
	name := c.Query("name")
	country := c.Query("country")

	results := searchOrganizations(name, country)
	c.JSON(200, results)
}

func updateHandler(c *gin.Context) {
	// Reload the organizations slice by re-reading the CSV and processing it
	loadOrganizations("data.csv")
	c.JSON(200, gin.H{"status": "updated"})
}

func loadOrganizations(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var tempOrgs []Org

	if _, err = reader.Read(); err != nil { // Skip header line
		fmt.Println("Error:", err)
		return
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		if len(record) != 9 {
			fmt.Println("Invalid record:", record)
			continue
		}
		tempOrgs = append(tempOrgs, Org{
			Index:               record[0],
			Organization_ID:     record[1],
			Name:                record[2],
			Website:             record[3],
			Country:             record[4],
			Description:         record[5],
			Founded:             record[6],
			Industry:            record[7],
			Number_of_employees: record[8],
		})
	}

	mu.Lock()
	organizations = tempOrgs // Update the global slice
	mu.Unlock()
}

func searchOrganizations(name, country string) []Org {
	mu.Lock()
	defer mu.Unlock()

	var results []Org
	for _, org := range organizations {
		if strings.Contains(strings.ToLower(org.Name), strings.ToLower(name)) && strings.Contains(strings.ToLower(org.Country), strings.ToLower(country)) {
			results = append(results, org)
		}
	}
	return results
}
