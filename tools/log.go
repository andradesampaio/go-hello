package tools

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func RegistraLogs(site string, status bool) {

	arquivo, err := os.OpenFile("./log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}


func ImprimirLogs() {

	arquivo, err := ioutil.ReadFile("./log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}