package graph

import (
	"encoding/json"
	"example/graph/model"
	"io/ioutil"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Todos    []*model.Todo `json:"todos"`
	Users    []*model.User `json:"users"`
	Relation []struct {
		From int `json:"from"`
		To   int `json:"to"`
	} `json:"relation"`
}

func (r *Resolver) Init() {
	readFile, err := ioutil.ReadFile("file/db.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(readFile, r); err != nil {
		panic(err)
	}
}
