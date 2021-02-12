package main

import (
	"fmt"

	"github.com/goyalk72/JobBazar/backend/api"
	"github.com/pkg/errors"
)

func main() {
	err := RunServer()
	if err != nil {
		fmt.Println(err)
	}
}

func RunServer() error {

	cfg := api.DefaultConfig()
	dbclient, err := api.DatabaseConnect()
	if err != nil {
		return errors.Wrap(err, "ERROR : STARTING SERVER")
	}
	server := api.NewServer(cfg, dbclient)
	err = server.Start()
	if err != nil {
		return errors.Wrap(err, "ERROR : STARTING SERVER")
	}

	return nil
}
