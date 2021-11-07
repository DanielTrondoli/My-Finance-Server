package actuator

import (
	"encoding/json"
	"net/http"

	"github.com/DanielTrondoli/My-Finance-Server/model/helthboddy"
)

// Health blablabal
func Health(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var resp = helthboddy.HelthBoddy{
		Status: "alive",
	}

	_ = json.NewEncoder(w).Encode(resp)

}
