package main

import (
	"context"
	"io"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

func markdown(ctx context.Context, input io.Reader, output io.Writer) error {
	source, err := io.ReadAll(input)
	if err != nil {
		return err
	}

	markdown := goldmark.New(
		goldmark.WithExtensions(
			extension.Table,
		),
	)
	return markdown.Convert(source, output)
}
