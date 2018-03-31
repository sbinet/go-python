package python

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"sync"
)

func (py *Interpreter) Command(pyfile string, args ...string) *Cmd {
	return &Cmd{py: py, file: pyfile, Args: args}
}

type Cmd struct {
	//Path string // TODO:

	Args []string
	//Env []string // TODO
	//Dir string

	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer

	py   *Interpreter
	file string
	fds  []io.Closer
	wg   sync.WaitGroup
}

func (cmd *Cmd) close() error {
	for _, f := range cmd.fds {
		f.Close()
	}
	cmd.fds = nil
	return nil
}

func (cmd *Cmd) inStream(r io.Reader) (*os.File, error) {
	pr, pw, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	cmd.fds = append(cmd.fds, pw, pr)
	cmd.wg.Add(1)
	go func(r io.Reader) {
		defer cmd.wg.Done()
		defer pw.Close()
		_, _ = io.Copy(pw, r)
	}(cmd.Stdin)
	return pr, nil
}

func (cmd *Cmd) outStream(w io.Writer) (*os.File, error) {
	pr, pw, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	cmd.fds = append(cmd.fds, pw, pr)
	cmd.wg.Add(1)
	go func() {
		defer cmd.wg.Done()
		_, _ = io.Copy(w, pr)
	}()
	return pw, nil
}
func (cmd *Cmd) Run() error {
	defer cmd.close()
	if cmd.Stdin == nil {
		cmd.Stdin = bytes.NewReader(nil)
	}
	if cmd.Stdout == nil {
		// TODO: it's better to open /dev/null
		cmd.Stdout = ioutil.Discard
	}
	if cmd.Stderr == nil {
		cmd.Stderr = ioutil.Discard
	}
	var (
		in, out, er *os.File
		err         error
	)
	if f, ok := cmd.Stdin.(*os.File); ok {
		in = f
	} else {
		in, err = cmd.inStream(cmd.Stdin)
		if err != nil {
			return err
		}
	}
	if f, ok := cmd.Stdout.(*os.File); ok {
		out = f
	} else {
		out, err = cmd.outStream(cmd.Stdout)
		if err != nil {
			return err
		}
	}
	if f, ok := cmd.Stderr.(*os.File); ok {
		er = f
	} else {
		er, err = cmd.outStream(cmd.Stderr)
		if err != nil {
			return err
		}
	}
	if err := cmd.py.SetStdinFile(in); err != nil {
		return err
	}
	if err := cmd.py.SetStdoutFile(out); err != nil {
		return err
	}
	if err := cmd.py.SetStderrFile(er); err != nil {
		return err
	}
	err = cmd.py.RunMain(cmd.file, cmd.Args...)
	cmd.close()
	cmd.wg.Wait()
	return err
}
