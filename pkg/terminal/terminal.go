package terminal

import (
	"bytes"
	"context"
	"github.com/gorilla/websocket"
	"github.com/yahahaff/rapide/pkg/handleerror"
	"github.com/yahahaff/rapide/pkg/logger"
	sshPkg "github.com/yahahaff/rapide/pkg/ssh"
	"golang.org/x/crypto/ssh"
	"os"
)

type Terminal struct {
	opts     Options
	ip       string
	ws       *websocket.Conn
	stdin    *os.File
	stdinr   *os.File
	conn     *ssh.Client
	session  *sshPkg.SshConn
	inited   int32
	cancelFn context.CancelFunc
}

type Options struct {
	Addr     string `json:"addr"`
	User     string `json:"user"`
	Password string `json:"password"`
	Cols     int    `json:"cols"`
	Rows     int    `json:"rows"`
}

func NewTerminal(ws *websocket.Conn, opts Options) *Terminal {
	return &Terminal{opts: opts, ws: ws}
}

func (t *Terminal) Run() {
	var err error
	t.conn, err = sshPkg.NewSshClient(t.opts.Addr, t.opts.User, t.opts.Password)
	if handleerror.WsHandleError(t.ws, err) {
		return
	}
	defer func() {
		t.conn.Close()
	}()
	//startTime := time.Now()
	t.session, err = sshPkg.NewSshConn(t.opts.Cols, t.opts.Rows, t.conn)

	if handleerror.WsHandleError(t.ws, err) {
		return
	}
	defer func() {
		t.session.Close()
	}()

	quitChan := make(chan bool, 3)

	var logBuff = new(bytes.Buffer)

	go t.session.ReceiveWsMsg(t.ws, logBuff, quitChan)
	go t.session.SendComboOutput(t.ws, quitChan)
	go t.session.SessionWait(quitChan)

	<-quitChan
	logger.DebugString("websoket", "webssh", "websocket finished")
}
