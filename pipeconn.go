package pipeconn

import (
	"io"
	"os"
	"os/exec"
)

type ServerPipeConn struct{}

func (s *ServerPipeConn) Read(p []byte) (int, error) {
	return os.Stdin.Read(p)
}

func (s *ServerPipeConn) Write(p []byte) (int, error) {
	return os.Stdout.Write(p)
}

func (s *ServerPipeConn) Close() error {
	os.Stdin.Close()
	os.Stdout.Close()
	return nil
}

type ClientPipeConn struct {
	pipeIn  io.ReadCloser
	pipeOut io.WriteCloser
}

func (s *ClientPipeConn) Read(p []byte) (int, error) {
	return s.pipeIn.Read(p)
}

func (s *ClientPipeConn) Write(p []byte) (int, error) {
	return s.pipeOut.Write(p)
}

func (s *ClientPipeConn) Close() error {
	s.pipeIn.Close()
	s.pipeOut.Close()
	return nil
}

//NewServerPipeConn 返回由 stdin、stdout 构成的 io.ReadWriteCloser
func NewServerPipeConn() io.ReadWriteCloser {
	return new(ServerPipeConn)
}

//NewClientPipeConn 运行 prog， 返回管道构成的 io.ReadWriteCloser
func NewClientPipeConn(prog string, args ...string) (io.ReadWriteCloser, error) {
	cmd := exec.Command(prog, args...)
	pipeIn, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	pipeOut, err := cmd.StdinPipe()
	if err != nil {
		pipeIn.Close()
		return nil, err
	}
	err = cmd.Start()
	if err != nil {
		pipeIn.Close()
		pipeOut.Close()
		return nil, err
	}
	return &ClientPipeConn{pipeIn: pipeIn, pipeOut: pipeOut}, nil
}
