package main_test

import (
	"ChAMP-Backend-Final-Project/controllers"
	"ChAMP-Backend-Final-Project/initializers"
	"ChAMP-Backend-Final-Project/models"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func TestTrue(t *testing.T) {
	assert.True(t, true, "True is true!")
}

func TestGetAPI(t *testing.T) {
	mockResponse := `{"message":"pong"}`
	r := SetUpRouter()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData), "that's not pong!")
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPostAPI(t *testing.T) {
	r := SetUpRouter()
	r.POST("/list", controllers.ListCreate)
	list := models.List{
		Title: "TestTitle",
	}
	jsonValue, _ := json.Marshal(gin.H{"title": list.Title})
	req, _ := http.NewRequest("POST", "/list", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	//assert.Equal(t, http.StatusCreated, w.Code)
}
