package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"
)

func markdown(ctx context.Context, input io.Reader, output io.Writer) error {
	stderr := &bytes.Buffer{}

	cmd := exec.CommandContext(ctx, markdownProgram)
	cmd.Stdin = input
	cmd.Stdout = output
	cmd.Stderr = stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run %q: %s", markdownProgram, stderr.String())
	}

	return nil
}
