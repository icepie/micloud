package main

import (
	"log"

	"github.com/icepie/micloud/xiaomiio"
)

func main() {

	miioUser := xiaomiio.NewXiaoMiio("fff", "fff")
	err := miioUser.Login()
	if err != nil {
		panic(err)
	}

	log.Println(miioUser.Client.Cookies)

}
