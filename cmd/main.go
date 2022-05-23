package main

import (
	"log"

	"github.com/icepie/micloud"
	"github.com/icepie/micloud/xiaomiio"
)

func main() {

	miioUser := xiaomiio.NewXiaoMiio("1213", "123456")
	err := miioUser.Login()
	if err != nil {
		panic(err)
	}

	log.Println(micloud.GenNonce())
	log.Println(miioUser.SecurityToken)

	test, _ := micloud.GetSignedNonce("7AL5MZR1mL4R8UmaEn8QZA==", "mWX0lu1kqZENc6QB")

	print(test)

	// test = gen_signed_nonce("7AL5MZR1mL4R8UmaEn8QZA==", "mWX0lu1kqZENc6QB")

	log.Println(micloud.GenSignature("123", test, "mWX0lu1kqZENc6QB", "123"))

}
