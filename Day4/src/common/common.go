package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	Key    string `json:"key"`
	Count  int    `json:"count"`
	Amount int    `json:"amount"`
}

type SuccessResponse struct {
	Thanks string `json:"thanks"`
	Change int    `json:"change"`
}

type PoorResponse struct {
	Message string `json:"message"`
	Amount  int    `json:"amount"`
}

type FailureResponse struct {
	Message string `json:"message"`
}

func SuccessInit(w http.ResponseWriter, d int, str string) {
	s := SuccessResponse{Thanks: str, Change: d}
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(s)
}

func PoorInit(w http.ResponseWriter, d int) {
	p := PoorResponse{Message: fmt.Sprintf("You need %d more money", -d), Amount: -d}
	w.WriteHeader(http.StatusPaymentRequired) // 402
	json.NewEncoder(w).Encode(p)
}

func FailureInit(w http.ResponseWriter, s string) {
	f := FailureResponse{Message: s}
	w.WriteHeader(http.StatusBadRequest) // 400
	json.NewEncoder(w).Encode(f)
}

func SuccessPrint(body io.ReadCloser) {
	s := SuccessResponse{}
	json.NewDecoder(body).Decode(&s)
	fmt.Println(s.Thanks)
	if s.Change != 0 {
		fmt.Printf(" your cange: %d\n", s.Change)
	} else {
		fmt.Printf(("\n"))
	}
}

func FailurePrint(body io.ReadCloser) {
	f := FailureResponse{}
	json.NewDecoder(body).Decode(&f)
	fmt.Printf("%s\n", f.Message)
}

func PoorPrint(body io.ReadCloser) {
	p := PoorResponse{}
	json.NewDecoder(body).Decode(&p)
	fmt.Printf("%s, bum\n", p.Message)
}
