package usb

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/goburrow/serial"
)

// serial.Port interface
// https://github.com/goburrow/serial/blob/master/serial.go#L53

type testReadWriteCloser struct {
	buf *bytes.Buffer
	io.Writer
	io.Reader
	io.Closer
	MockedWrite func(p []byte) (n int, err error)
	MockedOpen  func(*serial.Config) error
}

func (t *testReadWriteCloser) Write(p []byte) (n int, err error) {
	return t.MockedWrite(p)
}

func (t *testReadWriteCloser) Open(s *serial.Config) error {
	return t.MockedOpen(s)
}

// func (t *testReadWriteCloser) Read(p []byte) (n int, err error) {
// 	return 0, nil
// }

// func (t *testReadWriteCloser) Close() error {
// 	return nil
// }

func Test_Write(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		buf := &bytes.Buffer{}
		testStub := &testReadWriteCloser{
			MockedWrite: func(p []byte) (n int, err error) {
				return fmt.Fprint(buf, string(p))
			},
			MockedOpen: func(s *serial.Config) error {
				return nil
			},
		}

		serial := USBSerieal{
			Port: testStub,
		}
		want := "hello"
		err := serial.Write([]byte(want))
		if err != nil {
			t.Errorf("err should be nil: %v", err)
		}

		if want != buf.String() {
			t.Errorf("written strings is invalid. want: %s, got: %s", want, buf.String())
		}
	})
}
