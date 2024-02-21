package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Spence156/goapi/api"
	"github.com/Spence156/goapi/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	var params = api.UserDetailsParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	t0 := time.Now()
	log.Info("Processing GetUserDetails Request for: ", params.Username)

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.UserDetails = (*database).GetUserDetails(params.Username)
	//	tokenDetails = (*database).GetUserDetails(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.UserResponse{
		Username: (*tokenDetails).Username,
		FullName: (*tokenDetails).FullName,
		Country:  (*tokenDetails).Country,
		Email:    (*tokenDetails).Email,
		Code:     http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	log.Info("Completed GetUserDetails Request for: ", params.Username, " in: ", time.Since(t0))
}
