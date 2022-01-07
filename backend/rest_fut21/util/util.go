package util

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/JuanMillan7818/FUT21/db"
	api "github.com/JuanMillan7818/FUT21/rest_fut21/api_external"
)

func load(workers int) {
	script := api.ApiExternal{}

	results := make(chan int, workers)
	results <- 1

	var wait sync.WaitGroup
	wait.Add(1)

	script.LoadData(1, &wait, results)

	limitPages := script.GetFUT21().Pages
	for i := 2; i <= limitPages; i++ {
		//fmt.Println("Leer:", i)
		results <- 1
		wait.Add(1)
		go script.LoadData(i, &wait, results)
	}
	wait.Wait()
}

func Start() {
	db.ConnectedDB("")
	if os.Getenv("REWRITEDATA") == "1" {
		workers, err := strconv.Atoi(os.Getenv("WORKERS"))
		if err != nil {
			log.Panic(err)
		} else {
			load(workers)
		}
	}
}
