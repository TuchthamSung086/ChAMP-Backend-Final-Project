package main_test

import (
	"ChAMP-Backend-Final-Project/initializers"
	"ChAMP-Backend-Final-Project/models"
	"ChAMP-Backend-Final-Project/routes"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var r *gin.Engine

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	r = routes.SetupRouter()
}

func TestTrue(t *testing.T) {
	assert.True(t, true, "True is true!")
}

func TestGetAPI(t *testing.T) {
	mockResponse := `{"message":"pong"}`
	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData), "that's not pong!")
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPostAPI(t *testing.T) {
	list := models.List{
		Title: "TestTitle",
	}
	jsonValue, _ := json.Marshal(list)
	req, _ := http.NewRequest("POST", "/list", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	//assert.Equal(t, http.StatusCreated, w.Code)
}

func testAPI(t *testing.T, httpType string, path string, jsonData string) string {
	// Declare a models.List struct with JSON data
	var list models.List
	var task models.Task
	// If list
	if strings.Contains(path, "list") {
		err := json.Unmarshal([]byte(jsonData), &list)
		if err != nil {
			fmt.Println("Error:", err)
			return "ERROR"
		}
	} else if strings.Contains(path, "task") {
		err := json.Unmarshal([]byte(jsonData), &task)
		if err != nil {
			fmt.Println("Error:", err)
			return "ERROR"
		}
	}
	jsonValue, _ := json.Marshal(list)
	req, _ := http.NewRequest(httpType, path, bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)
	return string(responseData)
}

func TestPlayground(t *testing.T) {
	testAPI(t, "DELETE", "/dev/clearDB", ``) // clear DB
	testAPI(t, "POST", "/list", `{"title":"TitlePOOH"}`)
	type Result struct {
		Title string
	}
	var result Result
	initializers.DB.Raw("SELECT id, name, age FROM users WHERE id = ?", 3).Scan(&result)
	// res := testAPI(t,"GET","/lists","")
	// assert.Equal(t, http.StatusCreated, w.Code)
}
