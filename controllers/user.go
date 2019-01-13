package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/praveen001/quest-server/models"
)

// RegisterUser creates a new user in database
func (c *AppContext) RegisterUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		log.Println("Unable to decode request body", err.Error())
		return
	}

	if err := c.DB.Save(user).Error; err != nil {
		log.Println("Unable to register", err.Error())
		return
	}

	w.Write([]byte("Test"))
}
