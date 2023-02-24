package liza

import (
	"fmt"
	"crypto/tls"
	"github.com/lizrice/secure-connections/utils"
	"crypto/x509"
)

func CertReqFunc(certfile, keyfile string) func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
	c, err := getCert(certfile, keyfile)

	return func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		fmt.Printf("Received TLS Hello asking for %s: sending certificate\n", hello.ServerName)
		if err != nil || certfile == "" {
			fmt.Println("I have no certificate")
		} else {
			err := utils.OutputPEMFile(certfile)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		}
		// utils.Wait()
		return &c, nil
	}
}

func ClientCertReqFunc(certfile, keyfile string) func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
	c, err := getCert(certfile, keyfile)

	return func(certReq *tls.CertificateRequestInfo) (*tls.Certificate, error) {
		fmt.Println("Received certificate request: sending certificate")
		if err != nil || certfile == "" {
			fmt.Println("I have no certificate")
		} else {
			err :=utils.OutputPEMFile(certfile)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		}
		// utils.Wait()
		return &c, nil
	}
}

func CertificateChains(rawCerts [][]byte, chains [][]*x509.Certificate) error {
	if len(chains) > 0 {
		fmt.Println("Verified certificate chain from peer:")

		for _, v := range chains {
			// fmt.Printf("Chain %d:\n", j)
			for i, cert := range v {
				fmt.Printf("  Cert %d:\n", i)
				fmt.Printf(utils.CertificateInfo(cert))
			}
		}
	}
	return nil
}


func getCert(certfile, keyfile string) (c tls.Certificate, err error) {
	if certfile != "" && keyfile != "" {
		c, err = tls.LoadX509KeyPair(certfile, keyfile)
		if err != nil {
			fmt.Printf("Error loading key pair: %v\n", err)
		}
	} else {
		err = fmt.Errorf("I have no certificate")
	}
	return
}