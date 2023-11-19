package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type response struct {
	Status  bool `json:"status"`
	Message struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"message"`
}

func makeResponse(w http.ResponseWriter, response []byte) error {
	if _, err := w.Write(response); err != nil {
		return err
	}

	return nil
}

func encode(w http.ResponseWriter, object any) error {
	encoder := json.NewEncoder(w)
	fmt.Print(object)
	if err := encoder.Encode(object); err != nil {
		return err
	}

	return nil
}

func decode(r io.Reader, object any) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&object)
	if err != nil {
		return err
	}

	return nil
}
