package controllers

import (
	"net/http"
)

// RegisterUser creates a new user in database
func (c *AppContext) RegisterUser(w http.ResponseWriter, r *http.Request) {
	c.Logger.Errorln("ERROR LOG")
	// panic("PANIC")
	// shouldBlock := r.URL.Query().Get("block")
	// fmt.Println(shouldBlock)
	// user := &models.User{}
	// if err := json.NewDecoder(r.Body).Decode(user); err != nil {
	// 	log.Println("Unable to decode request body", err.Error())
	// 	return
	// }

	// c.DB.AddUser(&models.User{}, shouldBlock)

	// if err := c.DB.Save(user).Error; err != nil {
	// 	log.Println("Unable to register", err.Error())
	// 	return
	// }

	v := r.URL.Query()
	w.Write([]byte(v.Get("q")))
}
