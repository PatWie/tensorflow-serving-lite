// Patrick Wieschollek <mail@patwie.com>

package serving

import (
	"errors"
	"github.com/golang/protobuf/proto"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	pb "github.com/tensorflow/tensorflow/tensorflow/go/core/protobuf"
	"io/ioutil"
	"log"
)

// Information about a model
type Model struct {
	Path       string
	Name       string
	Version    int
	SavedModel *tf.SavedModel
	Info       *pb.SavedModel
}

// One Servable can have multiple models
type Servable struct {
	Models []*Model
}

// Server Supports one servable
type TensorflowServer struct {
	Servable Servable
}

var Server TensorflowServer

func (m *Model) Predict(signature string, val interface{}) interface{} {

	inputData, err := tf.NewTensor(val)

	signature_def := m.Info.GetMetaGraphs()[0].GetSignatureDef()["serving_default"]

	var inputs = map[tf.Output]*tf.Tensor{}
	for _, i := range signature_def.Inputs {
		name := i.GetName()
		name_len := len(name)
		name = name[:name_len-2]
		op := m.SavedModel.Graph.Operation(name).Output(0)
		inputs[op] = inputData
	}

	var outputs = []tf.Output{}
	for _, i := range signature_def.Outputs {
		name := i.GetName()
		name_len := len(name)
		name = name[:name_len-2]
		op := m.SavedModel.Graph.Operation(name).Output(0)
		outputs = append(outputs, op)
	}

	result, err := m.SavedModel.Session.Run(inputs, outputs, nil)

	if err != nil {
		panic(err)
	}

	return result[0].Value()
}

func (tfs *TensorflowServer) Initialize() {
	tfs.Servable = Servable{}
}

func (tfs *TensorflowServer) AddModel(path string, name string, version int) {
	var err error

	log.Printf("(Re-)adding model: simple_model %v\n", name)

	in, err := ioutil.ReadFile("/tmp/simple_model/1/saved_model.pb")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	model_pb := &pb.SavedModel{}
	if err := proto.Unmarshal(in, model_pb); err != nil {
		log.Fatalln("Failed to parse address model_pb:", err)
	}

	m := &Model{Path: path, Name: name, Version: version}

	m.SavedModel, err = tf.LoadSavedModel(m.Path, []string{"serve"}, nil)
	m.Info = model_pb
	if err != nil {
		panic(err)
	}

	initOp := m.SavedModel.Graph.Operation("legacy_init_op")
	_, err = m.SavedModel.Session.Run(nil, nil, []*tf.Operation{initOp})
	if err != nil {
		panic(err)
	}

	tfs.Servable.Models = append(tfs.Servable.Models, m)
}

func (tfs *TensorflowServer) GetModel(name string) (*Model, error) {
	for _, model := range Server.Servable.Models {
		if model.Name == name {
			return model, nil
		}
	}
	return &Model{}, errors.New("model not found")
}
