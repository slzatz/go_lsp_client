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

var jsonNotification = JsonNotification{
	Jsonrpc: "2.0",
	Method:  "initialize",
	Params:  struct{}{},
}

func main() {
	jsonRequest := JsonRequest{
		Jsonrpc: "2.0",
		Id:      1,
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

	//z := make([]byte, 10000)
	fmt.Printf("\n\nEntering for loop\n\n")
	diagnostics := make(chan []protocol.Diagnostic)
	go receiveDiagnostics(diagnostics)
	go readMessages2(buffer_out0, diagnostics)

	// below create some files to test diagnostics
	var j int32
	for i := 0; i < 2; i++ {
		time.Sleep(time.Second * 2)
		text := "package main\nimport \"fmt\"\n func main() {\n fmt.Println(\"hello\"\n}\n"
		j++
		sendDidChangeNotification(&stdin, text, j)

		time.Sleep(time.Second * 2)
		text = "package main\nimport \"fmt\"\n func main() {\n fmt.Println(\"hello\")\n}\n"
		j++
		sendDidChangeNotification(&stdin, text, j)
	}

	// tell server the file is closed
	jsonNotification.Method = "textDocument/didClose"
	var closeParams protocol.DidCloseTextDocumentParams
	closeParams.TextDocument.URI = "file:///home/slzatz/go_fragments/main.go"
	jsonNotification.Params = closeParams
	b, err = json.Marshal(jsonNotification)
	if err != nil {
		log.Fatal(err)
	}
	s = string(b)
	header = fmt.Sprintf("Content-Length: %d\r\n\r\n", len(s))
	s = header + s
	fmt.Printf("\n\n%s\n\n", s)
	io.WriteString(stdin, s)

	// shutdown request sent to server
	shutdownRequest := JsonRequest{
		Jsonrpc: "2.0",
		Id:      2,
		Method:  "shutdown",
		Params:  nil,
	}
	b, err = json.Marshal(shutdownRequest)
	if err != nil {
		log.Fatal(err)
	}
	s = string(b)
	header = fmt.Sprintf("Content-Length: %d\r\n\r\n", len(s))
	s = header + s
	io.WriteString(stdin, s)
	fmt.Printf("\n\n%s\n\n", s)

	// exit notification semt to server
	jsonNotification.Method = "exit"
	jsonNotification.Params = nil
	b, err = json.Marshal(jsonNotification)
	if err != nil {
		log.Fatal(err)
	}
	s = string(b)
	header = fmt.Sprintf("Content-Length: %d\r\n\r\n", len(s))
	s = header + s
	fmt.Printf("\n\n%s\n\n", s)
	io.WriteString(stdin, s)

	time.Sleep(3 * time.Second)
}

func sendDidChangeNotification(stdinp *io.WriteCloser, text string, j int32) {
	jsonNotification.Method = "textDocument/didChange"
	jsonNotification.Params = protocol.DidChangeTextDocumentParams{
		TextDocument: protocol.VersionedTextDocumentIdentifier{
			TextDocumentIdentifier: protocol.TextDocumentIdentifier{
				URI: "file:///home/slzatz/go_fragments/main.go"},
			Version: j},
		ContentChanges: []protocol.TextDocumentContentChangeEvent{{Text: text}},
	}
	b, err := json.Marshal(jsonNotification)
	if err != nil {
		log.Fatalf("\n%s\n%v", string(b), err)
	}
	s := string(b)
	header := fmt.Sprintf("Content-Length: %d\r\n\r\n", len(s))
	s = header + s
	fmt.Printf("\n\nSent incorrect: %s\n\n", s)
	io.WriteString(*stdinp, s)

}

func receiveDiagnostics(diagnostics chan []protocol.Diagnostic) { //? []protocol.Diagnostics
	for {
		dd := <-diagnostics
		fmt.Printf("\n\n-----------------------------------------------\n\n")
		fmt.Printf("Diagnostics = %+v\n", dd)
		for i, d := range dd {
			fmt.Printf("Diagnostics = %+v\n", dd)
			fmt.Printf("Diagnostics[%d] = %+v\n", i, d)
			fmt.Printf("Diagnostics[%d].Range = %+v\n", i, d.Range)                                //{Start:{Line:1 Character:0} End:{Line:1 Character:0}}
			fmt.Printf("Diagnostics[%d].Range.Start = %+v\n", i, d.Range.Start)                    //{Line:1 Character:0}
			fmt.Printf("Diagnostics[%d].Range.Start.Line = %v\n", i, d.Range.Start.Line)           //uint32
			fmt.Printf("Diagnostics[%d].Range.Start.Character = %v\n", i, d.Range.Start.Character) //uint32
			fmt.Printf("Diagnostics[%d].Message = %s\n", i, d.Message)                             //1
		}
		if len(dd) == 0 {
			fmt.Println("Diagnostics was []")
		}
		fmt.Printf("\n\n-----------------------------------------------\n\n")
	}
}

func readMessages(reader *bufio.Reader, diagnostics chan []protocol.Diagnostic) {
	// note if more than one jsonrpc message is read at one time; only dealing with first
	var bb []byte
	z := make([]byte, 10000)
	//reader := bufio.NewReaderSize(*stdoutp, 10000)
	for {
		n, err := reader.Read(z)
		if err == io.EOF {
			fmt.Printf("\n\nGot EOF presumably from shutdown\n\n")
			break
		}
		if err != nil {
			log.Fatalf("\nRead -> %s\n%v", string(z), err)
		}
		//fullRead := string(z)
		bb = z[:n]
		fmt.Printf("\n\n-------------------------------\n\n")
		fmt.Printf("ReadMessages: Number of bytes read = %d\n", n)
		//fmt.Printf("ReadMessages: Full Read = %s", fullRead)
		fmt.Printf("ReadMessages: Full Read = %s", string(bb))
		idx0 := bytes.Index(z, []byte(":")) + 2
		idx := bytes.Index(z, []byte("\r\n\r\n"))
		length, _ := strconv.Atoi(string(z[idx0:idx]))

		// not sure it's common but saw an instance where only got header info
		if n < 30 {
			//bb = z[:n]
			fmt.Printf("\n!!!Partial - Only got %q\n\n", string(bb))
			n, err = reader.Read(z)
			if err == io.EOF {
				fmt.Printf("\n\nGot EOF presumably from shutdown\n\n")
				break
			}
			if err != nil {
				log.Fatalf("\nRead -> %s\n%v", string(z), err)
			}
			bb = append(bb, z...)
		}

		//bb = z[idx+4 : idx+4+length]
		bb = bb[idx+4 : idx+4+length]
		var v JsonNotification
		err = json.Unmarshal(bb, &v)
		if err != nil {
			log.Fatalf("\nA -> %s\n%v", string(bb), err)
		}

		/*
			fmt.Printf("\n\n-------------------------------\n\n")
			fmt.Printf("params = %+v", v.Params)
			fmt.Printf("\n\n-------------------------------\n\n")
		*/

		if v.Method == "textDocument/publishDiagnostics" {
			type JsonPubDiag struct {
				Jsonrpc string                            `json:"jsonrpc"`
				Method  string                            `json:"method"`
				Params  protocol.PublishDiagnosticsParams `json:"params"`
			}
			var vv JsonPubDiag
			err = json.Unmarshal(bb, &vv)

			fmt.Printf("\n\n+++++++++++++++++++++++++++++++++++++++++++++++\n\n")
			fmt.Printf("params = %+v\n", vv.Params)
			fmt.Printf("uri = %+v\n", vv.Params.URI) //file:///home/slzatz/go_fragments/main.go
			fmt.Printf("\n\n+++++++++++++++++++++++++++++++++++++++++++++++\n\n")
			diagnostics <- vv.Params.Diagnostics
		}

		//time.Sleep(time.Second)
	}
}

func readMessages2(reader *bufio.Reader, diagnostics chan []protocol.Diagnostic) {
	// note if more than one jsonrpc message is read at one time; only dealing with first
	bb := make([]byte, 10000)
	z := make([]byte, 10000)
	//reader := bufio.NewReaderSize(*stdoutp, 10000)
	for {
		//z = z[:0]
		n, err := reader.Read(z)
		if err == io.EOF {
			fmt.Printf("\n\nGot EOF presumably from shutdown\n\n")
			break
		}
		if err != nil {
			log.Fatalf("\nRead -> %s\n%v", string(z), err)
		}

		fmt.Printf("\n\n-------------------------------\n\n")
		fmt.Printf("ReadMessages: Number of bytes read = %d\n", n)
		fmt.Printf("ReadMessages: Full Read = %s", string(z[:n]))

		var v JsonNotification
		var length int
		var idx int
		var idx0 int
		if string(z[:7]) == "Content" {
			idx0 = bytes.Index(z, []byte(":")) + 2
			idx = bytes.Index(z, []byte("\r\n\r\n"))
			length, _ = strconv.Atoi(string(z[idx0:idx]))
			if (idx + 4 + length) == n {
				err = json.Unmarshal(z[idx+4:n], &v)
				if err != nil {
					log.Fatalf("\nA -> %s\n%v", string(z), err)
				}
				bb = bb[0:]
			} else if (idx + 4 + length) > n {
				//bb = z[:n]
				copy(bb, z[:n])
				bb = bb[:n]
				fmt.Printf("\n### len(b) = %d", len(bb))
				fmt.Printf("\n###CONTINUATION!!! bb = %s\n", string(bb))
				continue
			} else if (idx + 4 + length) < n {
				//bb = z[idx+4+length : n]
				copy(bb, z[idx+4+length:n])
				fmt.Printf("\n### len(b) = %d", len(bb))
				fmt.Printf("\n###CONTINUATION!!! bb = %s\n", string(bb))
				err = json.Unmarshal(z[idx+4:idx+4+length], &v)
				if err != nil {
					log.Fatalf("\nB -> %s\n%v", string(z), err)
				}
			}
		} else {
			z = append(bb, z...)
			fmt.Printf("\n!!!CONTINUATION!!! bb = %s\n", string(bb))
			fmt.Printf("!!!CONTINUATION!!! z = %s\n", string(z[:n+len(bb)]))
			idx0 = bytes.Index(z, []byte(":")) + 2
			idx = bytes.Index(z, []byte("\r\n\r\n"))
			length, _ = strconv.Atoi(string(z[idx0:idx]))
			if (idx + 4 + length) == (n + len(bb)) {
				err = json.Unmarshal(z[idx+4:n+len(bb)], &v)
				if err != nil {
					log.Fatalf("\nC -> %s\n%v", string(z), err)
				}
				bb = bb[0:]
			} else if (idx + 4 + length) > (n + len(bb)) {
				copy(bb, z[:n+len(bb)])
				bb = bb[:n+len(bb)]
				continue
			} else if (idx + 4 + length) < (n + len(bb)) {
				fmt.Printf("idx + 4 + length = %d; n + len(bb) = %d", idx+4+length, len(bb))
				err = json.Unmarshal(z[idx+4:idx+4+length], &v)
				if err != nil {
					log.Fatalf("\nD -> %s\n%v", string(z), err)
				}
				copy(bb, z[idx+4+length:n+len(bb)])
				bb = bb[:n+len(bb)-(idx+4+length)]
			}
		}

		if v.Method == "textDocument/publishDiagnostics" {
			type JsonPubDiag struct {
				Jsonrpc string                            `json:"jsonrpc"`
				Method  string                            `json:"method"`
				Params  protocol.PublishDiagnosticsParams `json:"params"`
			}
			var vv JsonPubDiag
			err = json.Unmarshal(z[idx+4:idx+4+length], &vv)

			fmt.Printf("\n\n+++++++++++++++++++++++++++++++++++++++++++++++\n\n")
			fmt.Printf("params = %+v\n", vv.Params)
			fmt.Printf("uri = %+v\n", vv.Params.URI) //file:///home/slzatz/go_fragments/main.go
			fmt.Printf("\n\n+++++++++++++++++++++++++++++++++++++++++++++++\n\n")
			diagnostics <- vv.Params.Diagnostics
		}

		//time.Sleep(time.Second)
	}
}
