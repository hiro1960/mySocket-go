package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	connection, error := net.Dial("tcp", "localhost:10000")

	if error != nil {
		panic(error)
	}

	// Ctrl-Cでプログラム終了時にCloseする。
	defer connection.Close()

	// 無限loopで実行し続ける
	for {
		sendMessage(connection)
	}
}

func sendMessage(connection net.Conn) {
	fmt.Print("> ")

	stdin := bufio.NewScanner(os.Stdin)
	if stdin.Scan() == false {
		fmt.Println("Ciao ciao!")
		return
	}

	_, error := connection.Write([]byte(stdin.Text()))

	if error != nil {
		panic(error)
	}

	var response = make([]byte, 4*1024)
	_, error = connection.Read(response)

	if error != nil {
		panic(error)
	}

	fmt.Printf("Server> %s \n", response)

	// 再起呼び出しする必要はない。Stackがどんどん深くなっていくだけ。外部を無限loopにした方がいい。
	// sendMessage(connection)
}
