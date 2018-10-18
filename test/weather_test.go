package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jmoiron/jsonq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/weather-project/handler"
)

func TestCreateOrderSuccess(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/weather", handler.GetWeather).Methods(http.MethodGet)
	ts := httptest.NewServer(router)
	defer ts.Close()

	// when
	res, err := http.Get(fmt.Sprintf("%s/weather?city=Gliwice", ts.URL))
	require.NoError(t, err)

	// then
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "application/json;charset=UTF-8", res.Header.Get("Content-Type"))

	data, _ := ioutil.ReadAll(res.Body)
	formattedJson := &bytes.Buffer{}
	if err := json.Indent(formattedJson, data, "", "  "); err != nil {
		log.Fatal("Issue while indenting json occured.", err)
	}
	jsonMap := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(formattedJson.String()))
	dec.Decode(&jsonMap)
	jq := jsonq.NewQuery(jsonMap)
	result, _ := jq.String("name")
	fmt.Println(formattedJson)
	assert.Equal(t, "Gliwice", result)
}
