package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"go.lsp.dev/protocol"
	//"github.com/go-language-server/protocol"
	//"go.lsp.dev/jsonrpc2"
)

type JsonRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	Id      int         `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type JsonResult struct {
	Jsonrpc string                    `json:"jsonrpc"`
	Id      int                       `json:"id"`
	Result  protocol.InitializeResult `json:"result"`
}

type JsonNotification struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

func main() {
	jsonRequest := JsonRequest{
		Jsonrpc: "2.0",
		Id:      1,
		Method:  "initialize",
		Params:  struct{}{},
	}

	jsonNotification := JsonNotification{
		Jsonrpc: "2.0",
		Method:  "initialize",
		Params:  struct{}{},
	}
	params := protocol.InitializeParams{}
	params.ProcessID = 0
	params.RootURI = "file:///home/slzatz/go_fragments"
	params.Capabilities = clientcapabilities
	jsonRequest.Params = params
	b, err := json.Marshal(jsonRequest)
	if err != nil {
		log.Fatal(err)
	}
	s := string(b)
	fmt.Printf("\n\n-------------------------------\n\n")
	fmt.Printf("Sending: %s", s[:40])
	fmt.Printf("\n\n-------------------------------\n\n")

	cmd := exec.Command("gopls", "serve", "-rpc.trace", "-logfile", "/home/slzatz/gopls_log")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("#1")
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("#2")

	header := fmt.Sprintf("Content-Length: %d\r\n\r\n", len(s))
	s = header + s

	//Client sends initialize method and server replies with result (not method): {Capabilities ...)
	io.WriteString(stdin, s)
	fmt.Println("#3")

	//time.Sleep(2 * time.Second)

	//buffer_out0 := bufio.NewReader(stdout)
	buffer_out0 := bufio.NewReaderSize(stdout, 10000)
	p := make([]byte, 10000)
	fmt.Printf("buffer_out0 = %v\n", buffer_out0.Size())
	n, err := buffer_out0.Read(p)
	if err != nil {
		log.Fatal(err)
	}
	fullRead := string(p)
	fmt.Printf("Number of bytes read = %d\n", n)
	fmt.Printf("\n\n-------------------------------\n\n")
	fmt.Printf("Full Read = %s", fullRead)

	idx := strings.Index(fullRead, "\r\n\r\n")
	jsonRead := fullRead[idx+4:]
	fmt.Printf("\n\n-------------------------------\n\n")
	fmt.Printf("jsonRead = %v", jsonRead[:40])
	idx0 := bytes.Index(p, []byte(":")) + 2
	idx = bytes.Index(p, []byte("\r\n\r\n"))
	length, _ := strconv.Atoi(string(p[idx0:idx]))

	//idx = bytes.Index(p, []byte("\r\n\r\n"))
	//bb := p[idx+4 : idx+4+2956]
	bb := p[idx+4 : idx+4+length]
	//idx = bytes.Index(bb, []byte("\x00"))
	//fmt.Printf("\n\nIndex = %v\n\n", idx)
	//var v protocol.InitializeResult
	var v JsonResult
	err = json.Unmarshal(bb, &v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n\n-------------------------------\n\n")
	fmt.Printf("Result = %+v\n", v)
	fmt.Printf("length = %d", length)
	fmt.Printf("\n\n-------------------------------\n\n")

	//fmt.Printf("ServerInfo: %v\n", v.Result.ServerInfo)
	//fmt.Printf("WorkSpace: %v\n", v.Result.Capabilities.Workspace)

	//Client sends notification method:initialized and server replies with notification (no id) method "window/showMessage"
	jsonNotification.Method = "initialized"
	//jsonRequest.Id = 2
	jsonNotification.Params = struct{}{}
	b, err = json.Marshal(jsonNotification)
	if err != nil {
		log.Fatal(err)
	}
	s = string(b)
	header = fmt.Sprintf("Content-Length: %d\r\n\r\n", len(s))
	s = header + s
	io.WriteString(stdin, s)
	fmt.Println("#4")
	pp := make([]byte, 10000)
	fmt.Printf("buffer_out0 = %v\n", buffer_out0.Size())
	n, err = buffer_out0.Read(pp)
	if err != nil {
		log.Fatal(err)
	}
	fullRead = string(pp)
	fmt.Printf("Number of bytes read = %d\n", n)
	fmt.Printf("\n\n-------------------------------\n\n")
	fmt.Printf("Full Read = %s", fullRead)

	// Client sends notification method:did/Open and server replies with notification (no id) method "window/logMessage"
	// It looks like this is a notification and should not have an id
	//jsonMethod.Method = "textDocument/didOpen"
	jsonNotification.Method = "textDocument/didOpen"
	//jsonMethod.Id = 3
	var textParams protocol.DidOpenTextDocumentParams
	textParams.TextDocument.URI = "file:///home/slzatz/go_fragments/main.go"
	textParams.TextDocument.LanguageID = "go"
	textParams.TextDocument.Text = " "
	textParams.TextDocument.Version = 1
	//jsonMethod.Params = textParams
	jsonNotification.Params = textParams
	//b, err = json.Marshal(jsonMethod)
	b, err = json.Marshal(jsonNotification)
	if err != nil {
		log.Fatal(err)
	}
	s = string(b)
	header = fmt.Sprintf("Content-Length: %d\r\n\r\n", len(s))
	s = header + s
	fmt.Printf("\n\n%s\n\n", s)
	io.WriteString(stdin, s)
	ppp := make([]byte, 10000)
	//time.Sleep(2 * time.Second)
	fmt.Println("#5")
	n, err = buffer_out0.Read(ppp)
	if err != nil {
		log.Fatal(err)
	}
	fullRead = string(ppp)
	fmt.Printf("Number of bytes read = %d\n", n)
	fmt.Printf("\n\n-------------------------------\n\n")
	fmt.Printf("Full Read = %s", fullRead)

	z := make([]byte, 10000)
	//i := 0
	var j int32
	fmt.Printf("\n\nEntering for loop\n\n")
	go func() {
		for {
			n, err = buffer_out0.Read(z)
			if err != nil {
				log.Fatalf("\nRead -> %s\n%v", string(z), err)

			}
			fullRead = string(z)
			fmt.Printf("Number of bytes read = %d\n", n)
			fmt.Printf("\n\n-------------------------------\n\n")
			fmt.Printf("Full Read = %s", fullRead)
			idx0 := bytes.Index(z, []byte(":")) + 2
			idx := bytes.Index(z, []byte("\r\n\r\n"))
			length, _ := strconv.Atoi(string(z[idx0:idx]))
			bb := z[idx+4 : idx+4+length]
			//idx = bytes.Index(bb, []byte("\x00"))
			//fmt.Printf("\n\nIndex = %v\n\n", idx)
			//var v protocol.InitializeResult
			var v JsonNotification
			err = json.Unmarshal(bb, &v)
			if err != nil {
				log.Fatalf("\nA -> %s\n%v", string(bb), err)
			}

			fmt.Printf("\n\n-------------------------------\n\n")
			fmt.Printf("params = %+v", v.Params)
			fmt.Printf("\n\n-------------------------------\n\n")

			if v.Method == "textDocument/publishDiagnostics" {
				type JsonPubDiag struct {
					Jsonrpc string                            `json:"jsonrpc"`
					Method  string                            `json:"method"`
					Params  protocol.PublishDiagnosticsParams `json:"params"`
				}
				var vv JsonPubDiag
				err = json.Unmarshal(bb, &vv)

				fmt.Printf("\n\n-----------------------------------------------\n\n")
				fmt.Printf("params = %+v\n", vv.Params)
				fmt.Printf("uri = %+v\n", vv.Params.URI) //file:///home/slzatz/go_fragments/main.go
				for i, d := range vv.Params.Diagnostics {
					fmt.Printf("diagnostics = %+v\n", vv.Params.Diagnostics)
					//fmt.Printf("diagnostics[0] = %+v\n", vv.Params.Diagnostics[0])
					fmt.Printf("diagnostics[%d] = %+v\n", i, d)
					//fmt.Printf("diagnostics[0].Range = %+v\n", vv.Params.Diagnostics[0].Range)                       //{Start:{Line:1 Character:0} End:{Line:1 Character:0}}
					fmt.Printf("diagnostics[%d].Range = %+v\n", i, d.Range) //{Start:{Line:1 Character:0} End:{Line:1 Character:0}}
					//fmt.Printf("diagnostics[0].Range.Start = %+v\n", vv.Params.Diagnostics[0].Range.Start)           //{Line:1 Character:0}
					fmt.Printf("diagnostics[%d].Range.Start = %+v\n", i, d.Range.Start)           //{Line:1 Character:0}
					fmt.Printf("diagnostics[%d].Range.Start.Line = %+v\n", i, d.Range.Start.Line) //1
					//fmt.Printf("diagnostics[0].Range.Start.Line = %+v\n", vv.Params.Diagnostics[0].Range.Start.Line) //1
				}
				if len(vv.Params.Diagnostics) == 0 {
					fmt.Println("Diagnostics was []")
				}
				fmt.Printf("\n\n-----------------------------------------------\n\n")
			}

			//time.Sleep(time.Second)
		}
	}()
	for {
		/*
			i++
			if (i % 10000000) == 0 {
				fmt.Printf("i = %d\n", i)
			}
		*/
		//fmt.Println(i)

		//if false {
		time.Sleep(time.Second * 4)
		j++
		jsonNotification.Method = "textDocument/didChange"
		jsonNotification.Params = protocol.DidChangeTextDocumentParams{
			TextDocument: protocol.VersionedTextDocumentIdentifier{
				TextDocumentIdentifier: protocol.TextDocumentIdentifier{
					URI: "file:///home/slzatz/go_fragments/main.go"},
				Version: j},
			ContentChanges: []protocol.TextDocumentContentChangeEvent{{Text: "package main\nimport \"fmt\"\n func main() {\n fmt.Println(\"hello\"\n}\n"}},
		}
		b, err = json.Marshal(jsonNotification)
		if err != nil {
			log.Fatalf("\n%s\n%v", string(b), err)
		}
		s = string(b)
		header = fmt.Sprintf("Content-Length: %d\r\n\r\n", len(s))
		s = header + s
		fmt.Printf("\n\nSent incorrect: %s\n\n", s)
		io.WriteString(stdin, s)
		//time.Sleep(time.Second)
		//if false {
		time.Sleep(time.Second * 4)
		//i = 0
		j++
		jsonNotification.Method = "textDocument/didChange"
		jsonNotification.Params = protocol.DidChangeTextDocumentParams{
			TextDocument: protocol.VersionedTextDocumentIdentifier{
				TextDocumentIdentifier: protocol.TextDocumentIdentifier{
					URI: "file:///home/slzatz/go_fragments/main.go"},
				Version: j},
			ContentChanges: []protocol.TextDocumentContentChangeEvent{{Text: "package main\nimport \"fmt\"\n func main() {\n fmt.Println(\"hello\")\n}\n"}},
		}
		b, err = json.Marshal(jsonNotification)
		if err != nil {
			log.Fatalf("\n%s\n%v", string(b), err)
		}
		s = string(b)
		header = fmt.Sprintf("Content-Length: %d\r\n\r\n", len(s))
		s = header + s
		fmt.Printf("\n\nSent corrected: %s\n\n", s)
		io.WriteString(stdin, s)

		/*
				file := os.Stdout
				fi, errr := file.Stat()
				if errr != nil {
					log.Fatalf("Couldn't stat file: %v", errr)
				}
				if fi.Size() == 0 {
					//fmt.Println("Nothing to Read")
					continue
				}
					_, err = buffer_out0.Peek(0)
					if err != nil {
						fmt.Println("Nothing to Read")
						continue
					}
			xx := buffer_out0.Buffered()
			if xx == 0 {
				continue
			}
		*/

	}

}
