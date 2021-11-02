package main

import (
	"feeder-service/internal/shared/domain/config"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numberRunes = []rune("123456789")
var dashRune = []rune("-")

func RandInvalid() string {
	b := make([]rune, rand.Intn(9))
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandValid() string {
	b := make([]rune, 9)
	for i := range b {
		if i < 4 {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		} else if i == 4 {
			b[i] = dashRune[0]
		} else {
			b[i] = numberRunes[rand.Intn(len(numberRunes))]
		}
	}
	return string(b)
}

func main() {
	config := config.LoadConfig()
	fmt.Println("Connecting to server on port " + fmt.Sprint(config.Server.Port))

	valid := flag.Bool("valid", false, "")
	flag.Parse()

	for {
		conn, err := net.Dial("tcp", ":"+fmt.Sprint(config.Server.Port))
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			os.Exit(1)
		}

		var input string
		if *valid {
			input = RandValid()
		} else {
			input = RandInvalid()
		}
		time.Sleep(2 * time.Second)

		fmt.Fprintf(conn, input+"\n")
	}

}
