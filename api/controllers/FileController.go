package controllers

import (
	"io"
	"net/http"
	"os"

	"github.com/aristotelesFerreira/crud_go_lang/api/responses"
)

func (server *Server) Uploadfile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	defer file.Close()

	// Create File
	dst, err := os.Create(handler.Filename)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)

		return
	}
	responses.JSON(w, http.StatusOK, "Successfully Uploaded File")

}
