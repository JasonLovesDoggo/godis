package main

import (
	"fmt"
	"github.com/JasonLovesDoggo/godis/handlers"
	"github.com/JasonLovesDoggo/godis/resp"
	"net"
	"strings"
)

func main() {
	fmt.Println("Listening on port :6379")

	// Create a new server
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Listen for connections
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	for {
		resper := resp.NewResp(conn)
		value, err := resper.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		if value.Typ != "array" {
			fmt.Println("Invalid request, expected array")
			continue
		}

		if len(value.Array) == 0 {
			fmt.Println("Invalid request, expected array length > 0")
			continue
		}

		command := strings.ToUpper(value.Array[0].Bulk)
		args := value.Array[1:]

		writer := resp.NewWriter(conn)

		handler, ok := handlers.Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			writer.Write(resp.Value{Typ: "string", Str: ""})
			continue
		}

		result := handler(args)
		writer.Write(result)
	}
}
