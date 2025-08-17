package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/a2aproject/a2a-go/a2a"
	"github.com/nats-io/nuid"
	"github.com/siuyin/a2atry/jsonrpc"
	"github.com/siuyin/dflt"
)

func main() {
	url := dflt.EnvString("URL", "http://localhost:8080/")
	log.Printf("URL=%s", url)

	msg := a2a.Message{MessageID: "myID", Parts: []a2a.Part{a2a.TextPart{Kind: "text", Text: "Please tell me the time"}}}
	params := a2a.MessageSendParams{Message: msg}
	pDat, err := json.Marshal(params)
	if err != nil {
		log.Fatal("marshal: ", err)
	}

	rpc := jsonrpc.Request{Message: jsonrpc.Message{MessageIdentifier: jsonrpc.MessageIdentifier{ID: nuid.Next()},
		JSONRPC: "2.0"}, Method: "message/send", Params: pDat}
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	if err := enc.Encode(&rpc); err != nil {
		log.Fatal("json encode: ", err)
	}

	r, err := http.Post(url, "application/json", &b)
	if err != nil {
		log.Fatal("post: ", err)
	}

	defer r.Body.Close()
	dat, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("readall: ", err)
	}

	fmt.Printf("response: %s\n", dat)
}
