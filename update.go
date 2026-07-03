package brm

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

// this operation will wait for 5 seconds since its a network call and then cancel
func Update(verbose bool, out io.Writer) error {
	cPath, err := exec.LookPath("go")
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, cPath, "install", "github.com/Vardhanb07/brm/cmd/brm@latest")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "operation failed, please try again after some time")
		return err
	}
	if verbose {
		fmt.Fprintln(out, string(output))
	}
	return nil
}
