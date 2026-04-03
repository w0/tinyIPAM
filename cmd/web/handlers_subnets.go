package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/w0/tinyIPAM/internal/database"
)

type subnet struct {
	Id            int    `json:"id,omitempty"`
	Name          string `json:"name"`
	NetworkPrefix string `json:"network_prefix"`
	Cidr          int    `json:"cidr"`
}

func (app *application) newSubnet(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var s subnet
	if err := decoder.Decode(&s); err != nil {
		app.logger.Debug("failed to decode subnet", "error", err)
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err := app.database.NewSubnet(r.Context(), database.NewSubnetParams{
		Name:          s.Name,
		NetworkPrefix: s.NetworkPrefix,
		Cidr:          int64(s.Cidr),
	})

	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			app.clientError(w, http.StatusConflict)
			return
		}
		app.serverError(w, r, err)
		return
	}
}

func (app *application) getSubnet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	s, err := app.database.GetSubnet(r.Context(), intId)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(s)
}
