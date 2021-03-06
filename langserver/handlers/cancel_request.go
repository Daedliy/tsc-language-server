package handlers

import (
	"context"

	"github.com/creachadair/jrpc2"
	"github.com/sourcegraph/go-lsp"
)

// CancelRequest will be called on "$/cancelRequest"
func CancelRequest(ctx context.Context, req *jrpc2.Request) error {
	var params lsp.CancelParams
	err := req.UnmarshalParams(jrpc2.NonStrict(&params))

	if err != nil {
		return err
	}

	id := params.ID.Str
	jrpc2.CancelRequest(ctx, id)
	return nil
}
