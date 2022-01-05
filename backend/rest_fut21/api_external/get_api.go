package apiexternal

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"github.com/JuanMillan7818/FUT21/db"
	"github.com/JuanMillan7818/FUT21/rest_fut21/models"
)

type ApiExternal struct {
	players models.Fut21
}

func (list ApiExternal) GetFUT21() models.Fut21 {
	return list.players
}

func (api *ApiExternal) LoadData(page int, wait *sync.WaitGroup, channel <-chan int) {
	clientRequest := http.Client{}

	request := api.createRequest(page)

	response, err := clientRequest.Do(request)
	if api.errorsManagament(err) {
		panic(err)
	}

	data, err := ioutil.ReadAll(response.Body)
	if api.errorsManagament(err) {
		panic(err)
	}

	content := api.parsePlayer(data)
	createDB(content)
	api.players.Pages = content.Pages

	<-channel
	defer response.Body.Close()
	defer wait.Done()
}

func (api *ApiExternal) createRequest(page int) *http.Request {
	request, err := http.NewRequest(
		"GET",
		"https://www.easports.com/fifa/ultimate-team/api/fut/item?page="+strconv.Itoa(page),
		nil)
	if api.errorsManagament(err) {
		panic(err)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	return request
}

func (api *ApiExternal) parsePlayer(data []byte) models.Fut21 {
	var model models.Fut21
	json.Unmarshal(data, &model)
	return model
}

func (api *ApiExternal) errorsManagament(err error) bool {
	switch err {
	case nil:
		return false
	default:
		return true
	}
}

func createDB(data models.Fut21) {
	db.SaveInDB(data.List)
}
