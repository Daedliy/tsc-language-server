package handlers

import (
	"context"

	"github.com/creachadair/jrpc2"
	"github.com/sourcegraph/go-lsp"
	lsctx "pkg.nimblebun.works/tsc-language-server/langserver/context"
	"pkg.nimblebun.works/tsc-language-server/tsc"
)

// TextDocumentCompletion is the callback that runs on the
// "textDocument/completion" method
func TextDocumentCompletion(ctx context.Context, _ *jrpc2.Request) ([]lsp.CompletionItem, error) {
	conf, err := lsctx.Config(ctx)

	if err != nil {
		return []lsp.CompletionItem{}, nil
	}

	completions := tsc.GetCompletions(conf)

	return completions, nil
}
