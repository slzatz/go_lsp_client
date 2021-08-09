package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
	"time"

	"go.lsp.dev/protocol"
	//"github.com/go-language-server/protocol"
	//"go.lsp.dev/jsonrpc2"
)

type JsonMethod struct {
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

func main() {
	jsonMethod := JsonMethod{
		Jsonrpc: "2.0",
		Id:      1,
		Method:  "initialize",
		Params:  struct{}{},
	}

	params := protocol.InitializeParams{}
	params.ProcessID = 0
	params.RootURI = "file:///"
	params.Capabilities = clientcapabilities
	jsonMethod.Params = params
	b, err := json.Marshal(jsonMethod)
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

	//s := `{"jsonrpc": "2.0", "id": 0, "method": "initialize", "params": {"processId": 0, "rootPath": null, "rootUri": "file:///", "initializationOptions": null, "capabilities": {"offsetEncoding": ["utf-8"], "textDocument": {"codeAction": {"dynamicRegistration": true}, "codeLens": {"dynamicRegistration": true}, "colorProvider": {"dynamicRegistration": true}, "completion": {"completionItem": {"commitCharactersSupport": true, "documentationFormat": ["markdown", "plaintext"], "snippetSupport": true}, "completionItemKind": {"valueSet": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25]}, "contextSupport": true, "dynamicRegistration": true}, "definition": {"dynamicRegistration": true}, "documentHighlight": {"dynamicRegistration": true}, "documentLink": {"dynamicRegistration": true}, "documentSymbol": {"dynamicRegistration": true, "symbolKind": {"valueSet": [1, 2, 3, 4, 5, 6, 7, 8, 9,10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26]}}, "formatting": {"dynamicRegistration": true}, "hover": {"contentFormat": ["markdown", "plaintext"], "dynamicRegistration": true}, "implementation": {"dynamicRegistration": true}, "onTypeFormatting": {"dynamicRegistration": true}, "publishDiagnostics": {"relatedInformation": true}, "rangeFormatting": {"dynamicRegistration": true}, "references": {"dynamicRegistration": true}, "rename": {"dynamicRegistration": true}, "signatureHelp": {"dynamicRegistration": true, "signatureInformation": {"documentationFormat": ["markdown", "plaintext"]}}, "synchronization": {"didSave": true, "dynamicRegistration": true, "willSave": true, "willSaveWaitUntil": true}, "typeDefinition": {"dynamicRegistration": true}}, "workspace": {"applyEdit": true, "configuration": true, "didChangeConfiguration": {"dynamicRegistration": true}, "didChangeWatchedFiles": {"dynamicRegistration": true}, "executeCommand": {"dynamicRegistration": true}, "symbol": {"dynamicRegistration": true, "symbolKind": {"valueSet": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26]}}, "workspaceEdit": {"documentChanges": true}, "workspaceFolders": true}}, "trace": "off", "workspaceFolders": [{"name": "listmanager", "uri": "file:///"}]}}`

	header := fmt.Sprintf("Content-Length: %d\r\n\r\n", len(s))
	s = header + s
	//Client sends initialize method
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

	idx = bytes.Index(p, []byte("\r\n\r\n"))
	bb := p[idx+4 : idx+4+2956]
	idx = bytes.Index(bb, []byte("\x00"))
	fmt.Printf("\n\nIndex = %v\n\n", idx)
	//var v protocol.InitializeResult
	var v JsonResult
	err = json.Unmarshal(bb, &v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n\n-------------------------------\n\n")
	fmt.Printf("Result = %+v", v)
	fmt.Printf("\n\n-------------------------------\n\n")

	fmt.Printf("ServerInfo: %v\n", v.Result.ServerInfo)
	fmt.Printf("WorkSpace: %v\n", v.Result.Capabilities.Workspace)

	//Client sends initialized method
	jsonMethod.Method = "initialized"
	jsonMethod.Id = 2
	jsonMethod.Params = struct{}{}
	b, err = json.Marshal(jsonMethod)
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

	// Client sends did/Open method
	jsonMethod.Method = "textDocument/didOpen"
	jsonMethod.Id = 3
	var textParams protocol.DidOpenTextDocumentParams
	textParams.TextDocument.URI = "file:///home/slzatz/go_fragments/main.go"
	textParams.TextDocument.LanguageID = "go"
	textParams.TextDocument.Text = " "
	textParams.TextDocument.Version = 1
	jsonMethod.Params = textParams
	b, err = json.Marshal(jsonMethod)
	if err != nil {
		log.Fatal(err)
	}
	s = string(b)
	header = fmt.Sprintf("Content-Length: %d\r\n\r\n", len(s))
	s = header + s
	fmt.Printf("\n\n%s\n\n", s)
	io.WriteString(stdin, s)
	ppp := make([]byte, 10000)
	time.Sleep(2 * time.Second)
	fmt.Println("#5")
	n, err = buffer_out0.Read(ppp)
	if err != nil {
		log.Fatal(err)
	}
	fullRead = string(ppp)
	fmt.Printf("Number of bytes read = %d\n", n)
	fmt.Printf("\n\n-------------------------------\n\n")
	fmt.Printf("Full Read = %s", fullRead)
}
