// Patrick Wieschollek <mail@patwie.com>

package main

import (
	"fmt"
	"github.com/patwie/tensorflow-serving-lite/tfslite/router"
	"github.com/patwie/tensorflow-serving-lite/tfslite/serving"
	"log"
	"net/http"
)

func main() {

	config := serving.GetConfiguration("config.yml")
	log.Println("Adding/updating models.")
	for _, m := range config.Models {
		serving.Server.AddModel(m.Path, m.Name, m.Version)
	}

	log.Printf("Exporting HTTP/REST API at: %v ...\n", config.TensorFlowServingLite.RestPort)

	r := router.GetRouter()
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.TensorFlowServingLite.RestPort), r)
	if err != nil {
		panic(err)
	}
}
