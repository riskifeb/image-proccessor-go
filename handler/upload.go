package handler

import (
	"encoding/base64"
	"io"
	"net/http"

	"github.com/riskifeb/compresGambar/library"
)

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse, _ := library.JsonEncode(map[string]interface{}{
			"response": false,
			"code":     http.StatusBadRequest,
			"message":  "request method not allowed"},
			true)
		w.Header().Add("Contet-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorResponse))
		return
	}

	data, errReadBody := io.ReadAll(r.Body)
	if errReadBody != nil {
		errorResponse, _ := library.JsonEncode(map[string]interface{}{
			"response": false,
			"code":     http.StatusBadRequest,
			"message":  errReadBody.Error()},
			true)
		w.Header().Add("Contet-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorResponse))
		return
	}

	decodedRequestBody, errorDecodedRequestBody := library.JsonDecode(string(data))
	if errorDecodedRequestBody != nil {
		errorResponse, _ := library.JsonEncode(map[string]interface{}{
			"response": false,
			"code":     http.StatusBadRequest,
			"message":  errorDecodedRequestBody.Error()},
			true)
		w.Header().Add("Contet-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorResponse))
		return
	}

	result, errorImageProcess := library.ImageProcessor(decodedRequestBody["picture"].(string))
	if errorImageProcess != nil {
		errorResponse, _ := library.JsonEncode(map[string]interface{}{
			"response": false,
			"code":     http.StatusBadRequest,
			"message":  errorImageProcess.Error()},
			true)
		w.Header().Add("Contet-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorResponse))
		return
	}

	imageResult := base64.StdEncoding.EncodeToString(result.Bytes())

	okResponse, _ := library.JsonEncode(map[string]interface{}{
		"response": true,
		"code":     http.StatusOK,
		"message":  imageResult},
		true)
	w.Header().Add("Contet-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(okResponse))
	return
}
