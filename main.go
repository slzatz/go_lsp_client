package main

import (
	"encoding/json"
	"fmt"
	"log"

	"go.lsp.dev/protocol"
	//"github.com/go-language-server/protocol"
	//"go.lsp.dev/jsonrpc2"
)

type JSONRPC struct {
	Jsonrpc string      `json:"jsonrpc"`
	Id      int         `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

func main() {
	jsonMethod := JSONRPC{
		Jsonrpc: "2.0",
		Id:      1,
		Method:  "initialize",
		Params:  struct{}{},
	}

	/*
		b, err := json.Marshal(jsonMethod)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\n\n-------------------------------\n\n")
		fmt.Printf("%s", string(b))
		fmt.Printf("\n\n-------------------------------\n\n")
	*/

	params := protocol.InitializeParams{}
	params.ProcessID = 0
	params.RootURI = "file:///"

	/*
		b, err = json.Marshal(params)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\n\n-------------------------------\n\n")
		fmt.Printf("%s", string(b))
		fmt.Printf("\n\n-------------------------------\n\n")
	*/

	/*
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

		s := `{"jsonrpc": "2.0", "id": 0, "method": "initialize", "params": {"processId": 0, "rootPath": null, "rootUri": "file:///", "initializationOptions": null, "capabilities": {"offsetEncoding": ["utf-8"], "textDocument": {"codeAction": {"dynamicRegistration": true}, "codeLens": {"dynamicRegistration": true}, "colorProvider": {"dynamicRegistration": true}, "completion": {"completionItem": {"commitCharactersSupport": true, "documentationFormat": ["markdown", "plaintext"], "snippetSupport": true}, "completionItemKind": {"valueSet": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25]}, "contextSupport": true, "dynamicRegistration": true}, "definition": {"dynamicRegistration": true}, "documentHighlight": {"dynamicRegistration": true}, "documentLink": {"dynamicRegistration": true}, "documentSymbol": {"dynamicRegistration": true, "symbolKind": {"valueSet": [1, 2, 3, 4, 5, 6, 7, 8, 9,10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26]}}, "formatting": {"dynamicRegistration": true}, "hover": {"contentFormat": ["markdown", "plaintext"], "dynamicRegistration": true}, "implementation": {"dynamicRegistration": true}, "onTypeFormatting": {"dynamicRegistration": true}, "publishDiagnostics": {"relatedInformation": true}, "rangeFormatting": {"dynamicRegistration": true}, "references": {"dynamicRegistration": true}, "rename": {"dynamicRegistration": true}, "signatureHelp": {"dynamicRegistration": true, "signatureInformation": {"documentationFormat": ["markdown", "plaintext"]}}, "synchronization": {"didSave": true, "dynamicRegistration": true, "willSave": true, "willSaveWaitUntil": true}, "typeDefinition": {"dynamicRegistration": true}}, "workspace": {"applyEdit": true, "configuration": true, "didChangeConfiguration": {"dynamicRegistration": true}, "didChangeWatchedFiles": {"dynamicRegistration": true}, "executeCommand": {"dynamicRegistration": true}, "symbol": {"dynamicRegistration": true, "symbolKind": {"valueSet": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26]}}, "workspaceEdit": {"documentChanges": true}, "workspaceFolders": true}}, "trace": "off", "workspaceFolders": [{"name": "listmanager", "uri": "file:///"}]}}`

		header := fmt.Sprintf("Content-Length: %d\r\n\r\n", len(s))
		s = header + s
		io.WriteString(stdin, s)
		fmt.Println("#3")

		time.Sleep(2 * time.Second)

		//buffer_out0 := bufio.NewReader(stdout)
		buffer_out0 := bufio.NewReaderSize(stdout, 10000)
		p := make([]byte, 10000)
		fmt.Printf("buffer_out0 = %v\n", buffer_out0.Size())
		n, err := buffer_out0.Read(p)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("n = %d", n)
		fmt.Printf("Read = %s", string(p))
	*/

	/*
		for {

			bytes, _, err := buffer_out0.ReadLine()
			fmt.Printf("Length = %d\n", len(bytes))
			rows = append(rows, string(bytes))
			if err == io.EOF {
				break
			}
			if len(bytes) == 0 {
				break
			}
		}
		fmt.Printf("rows = %q\n\n", rows)
		//buffer_out1 := bufio.NewReader(stdout)
		for {
			fmt.Println("#4")

			bytes, _, err := buffer_out0.ReadLine()
			fmt.Println("#5")
			rows = append(rows, string(bytes))
			fmt.Printf("rows = %q\n\n", rows)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("Error: %v", err)
				break
			}
			if len(bytes) > 100 {
				break
			}
		}
		fmt.Printf("rows = %q\n\n", rows)
	*/
	//fmt.Printf("%q", s)
	//var clientCapabilities protocol.ClientCapabilities
	//var textDocument *protocol.TextDocumentClientCapabilities
	//var workspace protocol.WorkspaceClientCapabilities
	//fmt.Printf("%v\n", c)

	/*
		clientcapabilities := protocol.ClientCapabilities{
			Workspace: &protocol.WorkspaceClientCapabilities{
				ApplyEdit: true,
				WorkspaceEdit: &protocol.WorkspaceClientCapabilitiesWorkspaceEdit{
					DocumentChanges:    true,
					FailureHandling:    "FailureHandling",
					ResourceOperations: []string{"ResourceOperations"},
				},
				DidChangeConfiguration: &protocol.DidChangeConfigurationWorkspaceClientCapabilities{
					DynamicRegistration: true,
				},
				DidChangeWatchedFiles: &protocol.DidChangeWatchedFilesWorkspaceClientCapabilities{
					DynamicRegistration: true,
				},
				Symbol: &protocol.WorkspaceSymbolClientCapabilities{
					DynamicRegistration: true,
					SymbolKind: &protocol.SymbolKindCapabilities{
						ValueSet: []protocol.SymbolKind{
							protocol.SymbolKindFile,
							protocol.SymbolKindModule,
							protocol.SymbolKindNamespace,
							protocol.SymbolKindPackage,
							protocol.SymbolKindClass,
							protocol.SymbolKindMethod,
						},
					},
				},
				ExecuteCommand: &protocol.ExecuteCommandClientCapabilities{
					DynamicRegistration: true,
				},
				WorkspaceFolders: true,
				Configuration:    true,
			},
			TextDocument: &protocol.TextDocumentClientCapabilities{
				Synchronization: &protocol.TextDocumentSyncClientCapabilities{
					DynamicRegistration: true,
					WillSave:            true,
					WillSaveWaitUntil:   true,
					DidSave:             true,
				},
				Completion: &protocol.CompletionTextDocumentClientCapabilities{
					DynamicRegistration: true,
					CompletionItem: &protocol.CompletionTextDocumentClientCapabilitiesItem{
						SnippetSupport:          true,
						CommitCharactersSupport: true,
						DocumentationFormat: []protocol.MarkupKind{
							protocol.PlainText,
							protocol.Markdown,
						},
						DeprecatedSupport: true,
						PreselectSupport:  true,
					},
					CompletionItemKind: &protocol.CompletionTextDocumentClientCapabilitiesItemKind{
						ValueSet: []protocol.CompletionItemKind{protocol.CompletionItemKindText},
					},
					ContextSupport: true,
				},
				Hover: &protocol.HoverTextDocumentClientCapabilities{
					DynamicRegistration: true,
					ContentFormat: []protocol.MarkupKind{
						protocol.PlainText,
						protocol.Markdown,
					},
				},
				SignatureHelp: &protocol.SignatureHelpTextDocumentClientCapabilities{
					DynamicRegistration: true,
					SignatureInformation: &protocol.TextDocumentClientCapabilitiesSignatureInformation{
						DocumentationFormat: []protocol.MarkupKind{
							protocol.PlainText,
							protocol.Markdown,
						},
					},
				},
				Declaration: &protocol.DeclarationTextDocumentClientCapabilities{
					DynamicRegistration: true,
					LinkSupport:         true,
				},
				Definition: &protocol.DefinitionTextDocumentClientCapabilities{
					DynamicRegistration: true,
					LinkSupport:         true,
				},
				TypeDefinition: &protocol.TypeDefinitionTextDocumentClientCapabilities{
					DynamicRegistration: true,
					LinkSupport:         true,
				},
				Implementation: &protocol.ImplementationTextDocumentClientCapabilities{
					DynamicRegistration: true,
					LinkSupport:         true,
				},
				References: &protocol.ReferencesTextDocumentClientCapabilities{
					DynamicRegistration: true,
				},
				DocumentHighlight: &protocol.DocumentHighlightClientCapabilities{
					DynamicRegistration: true,
				},
				DocumentSymbol: &protocol.DocumentSymbolClientCapabilities{
					DynamicRegistration: true,
					SymbolKind: &protocol.SymbolKindCapabilities{
						ValueSet: []protocol.SymbolKind{
							protocol.SymbolKindFile,
							protocol.SymbolKindModule,
							protocol.SymbolKindNamespace,
							protocol.SymbolKindPackage,
							protocol.SymbolKindClass,
							protocol.SymbolKindMethod,
						},
					},
					HierarchicalDocumentSymbolSupport: true,
				},
				CodeAction: &protocol.CodeActionClientCapabilities{
					DynamicRegistration: true,
					CodeActionLiteralSupport: &protocol.CodeActionClientCapabilitiesLiteralSupport{
						CodeActionKind: &protocol.CodeActionClientCapabilitiesKind{
							ValueSet: []protocol.CodeActionKind{
								protocol.QuickFix,
								protocol.Refactor,
								protocol.RefactorExtract,
								protocol.RefactorRewrite,
								protocol.Source,
								protocol.SourceOrganizeImports,
							},
						},
					},
				},
				CodeLens: &protocol.CodeLensClientCapabilities{
					DynamicRegistration: true,
				},
				DocumentLink: &protocol.DocumentLinkClientCapabilities{
					DynamicRegistration: true,
				},
				ColorProvider: &protocol.DocumentColorClientCapabilities{
					DynamicRegistration: true,
				},
				Formatting: &protocol.DocumentFormattingClientCapabilities{
					DynamicRegistration: true,
				},
				RangeFormatting: &protocol.DocumentRangeFormattingClientCapabilities{
					DynamicRegistration: true,
				},
				OnTypeFormatting: &protocol.DocumentOnTypeFormattingClientCapabilities{
					DynamicRegistration: true,
				},
				PublishDiagnostics: &protocol.PublishDiagnosticsClientCapabilities{
					RelatedInformation: true,
				},
				Rename: &protocol.RenameClientCapabilities{
					DynamicRegistration: true,
					PrepareSupport:      true,
				},
				FoldingRange: &protocol.FoldingRangeClientCapabilities{
					DynamicRegistration: true,
					RangeLimit:          uint32(5),
					LineFoldingOnly:     true,
				},
				SelectionRange: &protocol.SelectionRangeClientCapabilities{
					DynamicRegistration: true,
				},
				CallHierarchy: &protocol.CallHierarchyClientCapabilities{
					DynamicRegistration: true,
				},
				SemanticTokens: &protocol.SemanticTokensClientCapabilities{
					DynamicRegistration: true,
					Requests: protocol.SemanticTokensWorkspaceClientCapabilitiesRequests{
						Range: true,
						Full:  true,
					},
					TokenTypes:     []string{"test", "tokenTypes"},
					TokenModifiers: []string{"test", "tokenModifiers"},
					Formats: []protocol.TokenFormat{
						protocol.TokenFormatRelative,
					},
					OverlappingTokenSupport: true,
					MultilineTokenSupport:   true,
				},
				LinkedEditingRange: &protocol.LinkedEditingRangeClientCapabilities{
					DynamicRegistration: true,
				},
				Moniker: &protocol.MonikerClientCapabilities{
					DynamicRegistration: true,
				},
			},
			Window: &protocol.WindowClientCapabilities{
				WorkDoneProgress: true,
				ShowMessage: &protocol.ShowMessageRequestClientCapabilities{
					MessageActionItem: &protocol.ShowMessageRequestClientCapabilitiesMessageActionItem{
						AdditionalPropertiesSupport: true,
					},
				},
				ShowDocument: &protocol.ShowDocumentClientCapabilities{
					Support: true,
				},
			},
			General: &protocol.GeneralClientCapabilities{
				RegularExpressions: &protocol.RegularExpressionsClientCapabilities{
					Engine:  "ECMAScript",
					Version: "ES2020",
				},
				Markdown: &protocol.MarkdownClientCapabilities{
					Parser:  "marked",
					Version: "1.1.0",
				},
			},
			Experimental: "testExperimental",
		}
	*/
	/*
		b, err := json.Marshal(clientcapabilities)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("-------------------------------\n\n")
		fmt.Printf("%s", string(b))
		fmt.Printf("-------------------------------\n\n")
	*/

	params.Capabilities = clientcapabilities
	jsonMethod.Params = params
	/*
		b, err = json.Marshal(params)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\n\n-------------------------------\n\n")
		fmt.Printf("%s", string(b))
		fmt.Printf("\n\n-------------------------------\n\n")
	*/
	b, err := json.Marshal(jsonMethod)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n\n-------------------------------\n\n")
	fmt.Printf("%s", string(b))
	fmt.Printf("\n\n-------------------------------\n\n")
}
