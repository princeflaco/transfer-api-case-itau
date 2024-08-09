package response

import (
	"net/http"
	"transfer-api/adapter/response/dto"
)

func Accepted(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
}

func Ok(w http.ResponseWriter, body []byte) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err := w.Write(body)
	if err != nil {
		InternalServerError(w, err)
	}
}

func Created(w http.ResponseWriter, body *[]byte) {
	w.WriteHeader(http.StatusCreated)
	if body != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, err := w.Write(*body)
		if err != nil {
			InternalServerError(w, err)
		}
	}
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func NotFound(w http.ResponseWriter, err error) {
	Error(w, 404, err)
}

func BadRequest(w http.ResponseWriter, err error) {
	Error(w, 400, err)
}

func InternalServerError(w http.ResponseWriter, err error) {
	Error(w, 500, err)
}

func Error(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	response, _ := dto.NewErrorDTO(err.Error()).ToBytes()
	_, err = w.Write(response)
	if err != nil {
		panic(err)
	}
}
