package handlers

import (
	. "github.com/JasonLovesDoggo/godis/resp"
	"sync"
)

var SETs = map[string]string{}
var SETsMu = sync.RWMutex{}

func Set(args []Value) Value {
	if len(args) != 2 {
		return Value{Typ: "error", Str: "ERR wrong number of arguments for 'set' command"}
	}

	key := args[0].Bulk
	value := args[1].Bulk

	SETsMu.Lock()
	SETs[key] = value
	SETsMu.Unlock()

	return Value{Typ: "string", Str: "OK"}
}
