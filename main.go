package main

import (
	"fmt"

	//"github.com/go-language-server/protocol"
	"go.lsp.dev/protocol"
)

func main() {
	var c protocol.Client

	var clientCapabilities protocol.ClientCapabilities
	var textDocument *protocol.TextDocumentClientCapabilities
	var workspace protocol.WorkspaceClientCapabilities
	fmt.Printf("%v\n", c)

	clientCapabilities = protocol.ClientCapabilities{
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
						PlainText,
						Markdown,
					},
					DeprecatedSupport: true,
					PreselectSupport:  true,
				},
				CompletionItemKind: &CompletionTextDocumentClientCapabilitiesItemKind{
					ValueSet: []CompletionItemKind{CompletionItemKindText},
				},
				ContextSupport: true,
			},
			Hover: &HoverTextDocumentClientCapabilities{
				DynamicRegistration: true,
				ContentFormat: []MarkupKind{
					PlainText,
					Markdown,
				},
			},
			SignatureHelp: &SignatureHelpTextDocumentClientCapabilities{
				DynamicRegistration: true,
				SignatureInformation: &TextDocumentClientCapabilitiesSignatureInformation{
					DocumentationFormat: []MarkupKind{
						PlainText,
						Markdown,
					},
				},
			},
			Declaration: &DeclarationTextDocumentClientCapabilities{
				DynamicRegistration: true,
				LinkSupport:         true,
			},
			Definition: &DefinitionTextDocumentClientCapabilities{
				DynamicRegistration: true,
				LinkSupport:         true,
			},
			TypeDefinition: &TypeDefinitionTextDocumentClientCapabilities{
				DynamicRegistration: true,
				LinkSupport:         true,
			},
			Implementation: &ImplementationTextDocumentClientCapabilities{
				DynamicRegistration: true,
				LinkSupport:         true,
			},
			References: &ReferencesTextDocumentClientCapabilities{
				DynamicRegistration: true,
			},
			DocumentHighlight: &DocumentHighlightClientCapabilities{
				DynamicRegistration: true,
			},
			DocumentSymbol: &DocumentSymbolClientCapabilities{
				DynamicRegistration: true,
				SymbolKind: &SymbolKindCapabilities{
					ValueSet: []SymbolKind{
						SymbolKindFile,
						SymbolKindModule,
						SymbolKindNamespace,
						SymbolKindPackage,
						SymbolKindClass,
						SymbolKindMethod,
					},
				},
				HierarchicalDocumentSymbolSupport: true,
			},
			CodeAction: &CodeActionClientCapabilities{
				DynamicRegistration: true,
				CodeActionLiteralSupport: &CodeActionClientCapabilitiesLiteralSupport{
					CodeActionKind: &CodeActionClientCapabilitiesKind{
						ValueSet: []CodeActionKind{
							QuickFix,
							Refactor,
							RefactorExtract,
							RefactorRewrite,
							Source,
							SourceOrganizeImports,
						},
					},
				},
			},
			CodeLens: &CodeLensClientCapabilities{
				DynamicRegistration: true,
			},
			DocumentLink: &DocumentLinkClientCapabilities{
				DynamicRegistration: true,
			},
			ColorProvider: &DocumentColorClientCapabilities{
				DynamicRegistration: true,
			},
			Formatting: &DocumentFormattingClientCapabilities{
				DynamicRegistration: true,
			},
			RangeFormatting: &DocumentRangeFormattingClientCapabilities{
				DynamicRegistration: true,
			},
			OnTypeFormatting: &DocumentOnTypeFormattingClientCapabilities{
				DynamicRegistration: true,
			},
			PublishDiagnostics: &PublishDiagnosticsClientCapabilities{
				RelatedInformation: true,
			},
			Rename: &RenameClientCapabilities{
				DynamicRegistration: true,
				PrepareSupport:      true,
			},
			FoldingRange: &FoldingRangeClientCapabilities{
				DynamicRegistration: true,
				RangeLimit:          uint32(5),
				LineFoldingOnly:     true,
			},
			SelectionRange: &SelectionRangeClientCapabilities{
				DynamicRegistration: true,
			},
			CallHierarchy: &CallHierarchyClientCapabilities{
				DynamicRegistration: true,
			},
			SemanticTokens: &SemanticTokensClientCapabilities{
				DynamicRegistration: true,
				Requests: SemanticTokensWorkspaceClientCapabilitiesRequests{
					Range: true,
					Full:  true,
				},
				TokenTypes:     []string{"test", "tokenTypes"},
				TokenModifiers: []string{"test", "tokenModifiers"},
				Formats: []TokenFormat{
					TokenFormatRelative,
				},
				OverlappingTokenSupport: true,
				MultilineTokenSupport:   true,
			},
			LinkedEditingRange: &LinkedEditingRangeClientCapabilities{
				DynamicRegistration: true,
			},
			Moniker: &MonikerClientCapabilities{
				DynamicRegistration: true,
			},
		},
		Window: &WindowClientCapabilities{
			WorkDoneProgress: true,
			ShowMessage: &ShowMessageRequestClientCapabilities{
				MessageActionItem: &ShowMessageRequestClientCapabilitiesMessageActionItem{
					AdditionalPropertiesSupport: true,
				},
			},
			ShowDocument: &ShowDocumentClientCapabilities{
				Support: true,
			},
		},
		General: &GeneralClientCapabilities{
			RegularExpressions: &RegularExpressionsClientCapabilities{
				Engine:  "ECMAScript",
				Version: "ES2020",
			},
			Markdown: &MarkdownClientCapabilities{
				Parser:  "marked",
				Version: "1.1.0",
			},
		},
		Experimental: "testExperimental",
	}
}
