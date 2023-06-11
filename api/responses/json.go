package responses

import (
	"encoding/json"
	"github.com/ecrespo/goAPIrest/api/utils/logs"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	logger := logs.GetLogger()
	if err != nil {
		//fmt.Fprintf(w, "%s", err.Error())
		logger.Error().Msg(err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
