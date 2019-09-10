package golac // import "github.com/vulcangz/golac"

import (
	"bytes"
	"context"
	"errors"
	"os/exec"
	log "github.com/sirupsen/logrus"
)

// LocalExec is responsible to run PaddleHub LAC process.
type LocalExec struct {
	LacCmd    string
	LacArgs   []string
	Command   string
	ClassPath string

	Module string
	Option string

	ctx context.Context
}

// NewLocalExec returns a pointer of LocalExec
func NewLocalExec(ctx context.Context) *LocalExec {
	if ctx == nil {
		ctx = context.Background()
	}
	return &LocalExec{
		LacCmd:  "hub",
		LacArgs: []string{},
		Command: "run",

		Module: "lac",
		Option: "--input_text", // or --input_file

		ctx: ctx,
	}
}

// Run Use PaddleHub lexical analysis model LAC for word segmentation.
func (c *LocalExec) Run(text string) (response string, err error) {
	if text == "" {
		err := errors.New("ERROR: The number of values in input file is inconsistent with expectations.")
		return "", err
	}
	// build arguments
	args := c.LacArgs

	if c.Command != "" {
		args = append(args, c.Command)
	}

	if c.Module != "" {
		args = append(args, c.Module)
	}

	if c.Option == "--input_text" {
		args = append(args,
			"--input_text",
			text,
		)
	} else if c.Option == "--input_file" {
		args = append(args,
			"--input_file",
			text,
		)
	} else {
		err := errors.New("ERROR: Invalid command option.")
		log.Warnf("Failed to execute command. [%s]", err.Error())
		return "", err
	}

	log.Debugf("Run command [%s]", args)
	// execute command
	cmd := exec.CommandContext(c.ctx, c.LacCmd, args...)

	log.Debugf("Run command [%s]", strings.Join(cmd.Args, " "))

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	if err := cmd.Run(); err != nil {
		err := errors.New(err.Error() + ": " + stderr.String())
		log.Warnf("Failed to execute command. [%s]", err.Error())
		return "", err
	}

	response = stdout.String()

	if stdout.Len() > 0 {
		log.Debugf("Success to execute command. [%s]", stdout.String())
	} else {
		log.Debugf("Success to execute command.")
	}

	return response, nil
}
