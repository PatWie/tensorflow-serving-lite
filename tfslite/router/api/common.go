// Patrick Wieschollek <mail@patwie.com>

package api

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/patwie/tensorflow-serving-lite/tfslite/router/render"
	"github.com/patwie/tensorflow-serving-lite/tfslite/serving"
	"log"
	"net/http"
	"strconv"
)

// Default homepage.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	render.WriteText(w, "active")
}

type Request struct {
	Instances []float32
}

// curl -d '{"instances": [1.0]}' -X POST http://localhost:3000/v1/models/simple_model:predict
func PredictHandler(w http.ResponseWriter, r *http.Request) {

	version, err := strconv.Atoi(chi.URLParam(r, "version"))
	log.Println(version)
	if err != nil {
		http.Error(w, "400 Bad Request - version does not exists", http.StatusBadRequest)
	}
	modelname := chi.URLParam(r, "modelname")
	log.Println(modelname)

	model, err := serving.Server.GetModel(modelname)
	if err != nil {
		http.Error(w, "400 Bad Request - model does not exists", http.StatusBadRequest)
	}

	signature := chi.URLParam(r, "signature")
	log.Println(signature)

	if signature == "predict" {
		decoder := json.NewDecoder(r.Body)

		var t Request
		err = decoder.Decode(&t)

		if err != nil {
			panic(err)
		}

		render.WriteJSON(w, render.H{
			"result": model.Predict(signature, t.Instances),
		})
	} else {
		http.Error(w, "400 Bad Request - signature does not exists", http.StatusBadRequest)

	}

}
