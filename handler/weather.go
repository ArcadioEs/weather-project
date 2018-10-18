package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	city, _ := r.URL.Query()["city"]
	response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city[0] + "&appid=4b6153148536643a1d5476d94cd07181")
	if err != nil {
		log.Fatal("The HTTP request failed with error\n", err)
		RespondWithCodeAndMessage(http.StatusBadRequest, "Something went wrong, try again later.", w)
		return
	} else {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		data, _ := ioutil.ReadAll(response.Body)
		formattedJson := &bytes.Buffer{}
		if err := json.Indent(formattedJson, data, "", "  "); err != nil {
			log.Fatal("Issue while indenting json occured.", err)
		}
		w.Write([]byte(formattedJson.String()))
	}
}
