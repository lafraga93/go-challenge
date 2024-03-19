package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Trucker struct {
	UUID    string  `json:"uuid"`
	Name    string  `json:"name"`
	Banking Banking `json:"banking,omitempty"`
}

type Banking struct {
	Bank          string `json:"bank"`
	Agency        string `json:"agency"`
	AccountNumber string `json:"account_number"`
}

func main() {
	r := setupRouter()
	r.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/truckers", func(c *gin.Context) {
		truckers := GetTruckers()

		for _, trucker := range truckers {
			trucker.Banking = GetBanking(trucker.UUID)
		}

		c.JSON(http.StatusOK, truckers)
	})

	return r
}

func GetTruckers() []*Trucker {
	resp, _ := http.Get("https://go-challenge-truckers.free.beeceptor.com/api/v1/")

	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	var truckers []*Trucker
	json.Unmarshal(body, &truckers)

	return truckers
}

func GetBanking(truckerUuid string) Banking {
	resp, _ := http.Get(fmt.Sprintf("https://go-challenge-banking.free.beeceptor.com/api/v1/trucker/%s", truckerUuid))

	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	var banking Banking
	json.Unmarshal(body, &banking)

	return banking
}
