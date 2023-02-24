package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"local/common"
	"local/liza"
	"net/http"
	"os/exec"
	"math/rand"
	"time"
	"os"
)


var candy = map[string]int{
	"CE": 10,
	"AA": 15,
	"NT": 17,
	"DE": 21,
	"YR": 23,
}

func getServer() (s *http.Server, err error) {
	data, err := ioutil.ReadFile("../ca/minica.pem")
	if err != nil {
		return s, err
	}
	cp, err := x509.SystemCertPool()
	if err != nil {
		return s, err
	}
	cp.AppendCertsFromPEM(data)
	s = &http.Server{
		Addr: ":8080",
		TLSConfig: &tls.Config{
			ClientCAs:             cp,
			ClientAuth:            tls.RequireAndVerifyClientCert,
			GetCertificate:        liza.CertReqFunc("cert.pem", "key.pem"),
			VerifyPeerCertificate: liza.CertificateChains,
		},
	}
	return
}

func main() {
	srv, err := getServer()
	if err != nil {
		fmt.Println(err, " rrr")
		return
	}
	http.HandleFunc("/", myHandler)
	err = srv.ListenAndServeTLS("", "")
	if err != nil {
		fmt.Println(err, " rrr")
		return
	}

}

func myHandler(w http.ResponseWriter, r *http.Request) {
	var req common.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest) // 400
		return
	}

	fmt.Println("Key:", req.Key)
	fmt.Println("Count:", req.Count)
	fmt.Println("Amount:", req.Amount)

	w.Header().Set("Content-Type", "application/json")
	if candy[req.Key] == 0 {
		common.FailureInit(w, "Invalid key :(\nKeys: CE, AA, NT, DE, YR")
	} else if req.Amount < 0 {
		common.FailureInit(w, "Invalid amount < 0")
	} else if req.Count < 1 {
		common.FailureInit(w, "Invalid count < 1")
	} else {
		dist := req.Amount - req.Count*candy[req.Key]
		if dist >= 0 {
			cmd :=exec.Command("gcc", "C/cow.c", "-o", "cow")
			_,err := cmd.Output()
			if err != nil {
				fmt.Println(err)
				return
			}
			cmd = exec.Command("./cow", randomString())
			str,err := cmd.Output()
			if err != nil {
				fmt.Println(err)
				return
			}
			err = os.Remove("cow")
			if err != nil {
				fmt.Println(err)
				return
			}
			common.SuccessInit(w, dist, string(str))
		} else {
			common.PoorInit(w, dist)
		}
	}
}

func randomString() string {
	const consonants = "bcdfghjklmnpqrstvwxz"
	const vowels = "aeiouy"
	rand.Seed(time.Now().UnixNano())
	l := 2 + rand.Intn(15)
	f := rand.Intn(2)
	str := make([]byte, l)
	if f == 0 {
		str[0] = consonants[rand.Intn(len(consonants))] - 'a' + 'A'
	} else {
		str[0] = vowels[rand.Intn(len(vowels))] - 'a' + 'A'
	}
	for k := 1; k < l - 1; k++ {
		if k % 5 == 0 && k != l - 2 {
			str[k] = ' '
		} else if (f + k) % 2 == 1 {
			str[k] = vowels[rand.Intn(len(vowels))]
		} else {
			str[k] = consonants[rand.Intn(len(consonants))]
		}
	}
	str[l - 1] = '!'
	return string(str)
}
