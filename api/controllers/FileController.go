package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/aristotelesFerreira/crud_go_lang/api/responses"
)

func (server *Server) Uploadfile(w http.ResponseWriter, r *http.Request) {
	// Maximum MB supported
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("myFile")

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	defer file.Close()

	// Create File
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	tempFile.Write(fileBytes)

	responses.JSON(w, http.StatusOK, "Successfully Uploaded File")

}
