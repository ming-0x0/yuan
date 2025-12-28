package command

import (
	"context"
	"fmt"
	"strings"

	"github.com/ming-0x0/yuan/pkg/logger"
)

func Execute[C any](cmd Command[C], logger logger.Logger) Command[C] {
	return defaultCommandLoggingDecorator[C]{
		cmd:    cmd,
		logger: logger,
	}
}

type Command[C any] interface {
	Handle(ctx context.Context, cmd C) error
}

type defaultCommandLoggingDecorator[C any] struct {
	cmd    Command[C]
	logger logger.Logger
}

func (d defaultCommandLoggingDecorator[C]) Handle(ctx context.Context, cmd C) (err error) {
	cmdName := generateCommandName(cmd)

	logger := d.logger.WithFields(
		"command", cmdName,
		"command_body", fmt.Sprintf("%#v", cmd),
	)

	logger.Debug("Executing command")
	defer func() {
		if err == nil {
			logger.Info("Command executed successfully")
		} else {
			logger.Error("Failed to execute command", "error", err)
		}
	}()

	return d.cmd.Handle(ctx, cmd)
}

func generateCommandName(cmd any) string {
	return strings.Split(fmt.Sprintf("%T", cmd), ".")[1]
}
