package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"local/common"
	"local/liza"
	"net/http"
)

// -k AA -c 2 -m 50
func handleFlags() (common.Request, error) {
	var req common.Request
	flag.StringVar(&req.Key, "k", "", "accepts two-letter abbreviation for the candy type")
	flag.IntVar(&req.Count, "c", -1, "count of candy to buy")
	flag.IntVar(&req.Amount, "m", -1, "amount of money you \"gave to machine\"")
	flag.Parse()
	if req.Key == "" || req.Count == -1 || req.Amount == -1 {
		return req, errors.New("expected -k <name> -c <count>, -m <amount>")
	}
	return req, nil
}

func getClient() (c *http.Client, err error) {
	data, err := ioutil.ReadFile("../ca/minica.pem")
	if err != nil {
		return c, err
	}
	cp, err := x509.SystemCertPool()
	if err != nil {
		return c, err
	}
	cp.AppendCertsFromPEM(data)
	tlsConfig := &tls.Config{
		RootCAs:               cp,
		GetClientCertificate:  liza.ClientCertReqFunc("cert.pem", "key.pem"),
		VerifyPeerCertificate: liza.CertificateChains,
	}
	c = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
	return
}

func main() {
	client, err := getClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	buff, err := handleFlags()
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := json.Marshal(buff)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest("POST", "https://localhost:8080/", bytes.NewBuffer(b))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer resp.Body.Close()
	output((resp))
}

func output(r *http.Response) {
	fmt.Println("Response status:", r.Status)
	switch r.StatusCode {
	case 201:
		common.SuccessPrint(r.Body)
	case 400:
		common.FailurePrint(r.Body)
	case 402:
		common.PoorPrint(r.Body)
	}
}
