package addon

import (
	"time"

	"github.com/Incises/go-mitmproxy/proxy"

	log "github.com/sirupsen/logrus"
)

// LogAddon log connection and flow

type LogAddon struct {
	proxy.BaseAddon
}

func (addon *LogAddon) ConnectionConnected(client *proxy.ClientConn) {
	log.Infof("%v client connected\n", client.Conn.RemoteAddr())
}

func (addon *LogAddon) ClientDisconnected(client *proxy.ClientConn) {
	log.Infof("%v client disconnected\n", client.Conn.RemoteAddr())
}

func (addon *LogAddon) ServerConnected(connCtx *proxy.ConnContext) {
	log.Infof("%v server connected %v (%v --> %v)\n",
		connCtx.ClientConn.Conn.RemoteAddr(),
		connCtx.ServerConn.Address,
		connCtx.ClientConn.Conn.LocalAddr(), connCtx.ServerConn.Conn.RemoteAddr())
}

func (addon *LogAddon) ServerDisconnected(connCtx *proxy.ConnContext) {
	log.Infof("%v server disconnected %v (%v --> %v)\n",
		connCtx.ClientConn.Conn.RemoteAddr(),
		connCtx.ServerConn.Address,
		connCtx.ClientConn.Conn.LocalAddr(), connCtx.ServerConn.Conn.RemoteAddr())
}

func (addon *LogAddon) RequestHeaders(flow *proxy.Flow) {
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
		log.Infof("%v %v %v %v %v - %v ms\n",
			flow.ConnContext.ClientConn.Conn.RemoteAddr(),
			flow.Request.Method, flow.Request.URL.String(),
			StatusCode, contentLength, time.Since(start).Milliseconds())
	}()
}
