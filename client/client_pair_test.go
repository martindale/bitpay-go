package client_test

import (
	. "github.com/philosodad/bitpay-go/client"
	ku "github.com/philosodad/bitpay-go/key_utils"

	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"os/exec"
)

var _ = Describe("ClientPair", func() {
	It("pairs with the server with a pairing code", func() {
		pm := ku.GeneratePem()
		gopath := os.ExpandEnv("$GOPATH")
		pyloc := gopath + "/helpers/pair_steps.py"
		cmd := exec.Command(pyloc)
		stdout, _ := cmd.StdoutPipe()
		cmd.Start()
		byt, _ := ioutil.ReadAll(stdout)
		code := string(byt)
		apiuri := os.ExpandEnv("$RCROOTADDRESS")
		webClient := Client{ApiUri: apiuri, Insecure: true, Pem: pm}
		token, _ := webClient.PairWithCode(code)
		webClient.Token = token
		Expect(webClient.Token.Facade).To(Equal("pos"))
	})
})
