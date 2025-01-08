package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"runtime/debug"

	"github.com/pires/go-proxyproto"
	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"
	"github.com/soheilhy/cmux"
	"github.com/xmapst/logx"

	"github.com/xmapst/go-hotfix"
	"github.com/xmapst/go-hotfix/example/handler"
	"github.com/xmapst/go-hotfix/example/symbols"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			logx.Errorln(err, string(debug.Stack()))
		}
	}()
	// Create the main listener.
	ln, err := net.Listen("tcp", ":3333")
	if err != nil {
		logx.Fatal(err)
	}
	// Create a proxyproto listener.
	ln = &proxyproto.Listener{Listener: ln}

	// Create a cmux.
	mux := cmux.New(ln)

	// Match connections in order:
	// then HTTP, and otherwise Go TCP.
	httpL := mux.Match(cmux.HTTP1Fast())
	// Any means anything that is not yet matched.
	tcpL := mux.Match(cmux.Any())

	go func() {
		logx.Infof("starting HTTP server on %s", httpL.Addr())
		if err = http.Serve(httpL, handler.New()); err != nil {
			logx.Fatalln(err)
		}
	}()

	go func() {
		logx.Infof("starting TCP server on %s", tcpL.Addr())
		if err = telnet.Serve(tcpL, telnetHandler()); err != nil {
			logx.Fatalln(err)
		}
	}()

	if err = mux.Serve(); err != nil {
		logx.Fatalln(err)
	}
}

func telnetHandler() telnet.Handler {
	shellHandler := telsh.NewShellHandler()
	var commands = map[string]string{
		"help":   "Display a list of available commands.",
		"hotfix": "Load a hotfix script and apply the specified function.",
		"exit":   "Exit the shell.",
	}
	// Register the "help" command.
	if err := shellHandler.RegisterHandlerFunc("help", func(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
		_, _ = stdout.Write([]byte("Available commands:\n"))
		for command, desc := range commands {
			_, _ = stdout.Write([]byte(fmt.Sprintf("\t%s - %s\n", command, desc)))
		}
		_, _ = oi.LongWrite(stdout, []byte("\r\n"))
		return nil
	}); err != nil {
		logx.Fatalln(err)
	}

	// Register the "hotfix" command.
	if err := shellHandler.RegisterHandlerFunc("hotfix", func(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
		if len(args) < 2 {
			_, _ = stdout.Write([]byte("hotfix 选项 脚本路径 脚本函数名\n"))
			_, _ = stdout.Write([]byte("例如: hotfix patch/test_yaegi_01.patch patch.PatchTest()\n"))
			_, _ = oi.LongWrite(stdout, []byte("\r\n"))
			return nil
		}
		_, err := hotfix.ApplyFunc(args[0], args[1], symbols.Symbols)
		if err != nil {
			_, _ = oi.LongWriteString(stdout, err.Error())
			return nil
		}
		_, _ = oi.LongWriteString(stdout, "hotfix success\r\n")
		logx.Info("hotfix success")
		return nil
	}); err != nil {
		logx.Fatalln(err)
	}
	return shellHandler
}
