package presdel

import (
	"banana/pkg/domain"
	_ "banana/pkg/utils/filesaver"
	_ "path/filepath"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	urlPresentationFilePath   = "/static/presentation/files/"
	urlMcs                    = "input_prezentations/"
	urlPresentationSlidesPath = "/static/presentation/slides/"
)

// /presentation/{id}
func (h *presHandler) getPres(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	presId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, domain.ErrUrlParameter.Error(), http.StatusBadRequest)
		return
	}

	pres, err := h.PresUsecase.GetPres(1, presId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(struct {
		Pres domain.PresApiResponse `json:"pres"`
	}{Pres: pres})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

// /presentation/create
func (h *presHandler) createPres(w http.ResponseWriter, r *http.Request) {
/*
	err := r.ParseMultipartForm(10 * 1024 * 1024) // limit 10Mb
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uploaded, header, err := r.FormFile("presentation")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploaded.Close()

	filename, err := filesaver.UploadFile(uploaded, urlPresentationFilePath, filepath.Ext(header.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
*/
	filename := "test_prez.pptx"
	presId, err := h.PresUsecase.CreatePres(urlMcs + filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	marshalledUs, err := json.Marshal(struct {
		PresId uint64 `json:"presId"`
	}{PresId: presId})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshalledUs)
}
