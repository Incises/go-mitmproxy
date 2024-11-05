package addon

import (
	"time"

	"github.com/Incises/go-mitmproxy/proxy"

	"github.com/sirupsen/logrus"
)

// Logger - log connection and flow

type Logger struct {
	proxy.BaseAddon
}

func (addon *Logger) ClientConnected(client *proxy.ClientConn) {
	logrus.Infof("%v client connected\n", client.Conn.RemoteAddr())
}

func (addon *Logger) ClientDisconnected(client *proxy.ClientConn) {
	logrus.Infof("%v client disconnected\n", client.Conn.RemoteAddr())
}

func (addon *Logger) ServerConnected(connCtx *proxy.ConnContext) {
	logrus.Infof("%v server connected %v (%v --> %v)\n",
		connCtx.ClientConn.Conn.RemoteAddr(),
		connCtx.ServerConn.Address,
		connCtx.ClientConn.Conn.LocalAddr(), connCtx.ServerConn.Conn.RemoteAddr())
}

func (addon *Logger) ServerDisconnected(connCtx *proxy.ConnContext) {
	logrus.Infof("%v server disconnected %v (%v --> %v)\n",
		connCtx.ClientConn.Conn.RemoteAddr(),
		connCtx.ServerConn.Address,
		connCtx.ClientConn.Conn.LocalAddr(), connCtx.ServerConn.Conn.RemoteAddr())
}

func (addon *Logger) RequestHeaders(flow *proxy.Flow) {
	start := time.Now()
	go func() {
		<-flow.Done()
		var StatusCode int
		if flow.Response != nil {
			StatusCode = flow.Response.StatusCode
		}
		var contentLength int
		if flow.Response != nil && flow.Response.Body != nil {
			contentLength = len(flow.Response.Body)
		}
		logrus.Infof("%v %v %v %v %v - %v ms\n",
			flow.ConnContext.ClientConn.Conn.RemoteAddr(),
			flow.Request.Method, flow.Request.URL.String(),
			StatusCode, contentLength, time.Since(start).Milliseconds())
	}()
}

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logrus.SetLevel(logrus.InfoLevel)
}
