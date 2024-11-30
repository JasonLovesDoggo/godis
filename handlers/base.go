package handlers

import "github.com/JasonLovesDoggo/godis/resp"

var Handlers = map[string]func([]resp.Value) resp.Value{
	"PING": Ping,
	"SET":  Set,
	"GET":  Get,
}
