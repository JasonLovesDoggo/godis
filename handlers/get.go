package handlers

import "github.com/JasonLovesDoggo/godis/resp"

func Get(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'get' command"}
	}

	key := args[0].Bulk

	SETsMu.RLock()
	value, ok := SETs[key]
	SETsMu.RUnlock()

	if !ok {
		return resp.Value{Typ: "null"}
	}

	return resp.Value{Typ: "bulk", Bulk: value}
}
