package main

import "go.lsp.dev/protocol"

var clientcapabilities = protocol.ClientCapabilities{
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
