# go-mitmproxy

`go-mitmproxy` is a Go-based tool similar to [mitmproxy](https://mitmproxy.org/) that enables man-in-the-middle attacks, allowing for the interception, monitoring, and modification of HTTP/HTTPS traffic.

## Key features

- Intercepts and displays HTTP/HTTPS traffic details through a [web interface](#web-interface).
- Extensible via a [plugin mechanism](#adding-functionality-by-developing-plugins) with various event hooks available in the [examples](./examples) directory.
- Compatible with [mitmproxy](https://mitmproxy.org/) for HTTPS certificate handling, storing certificates in the `~/.mitmproxy` folder. If a root certificate is already trusted from previous `mitmproxy` use, `go-mitmproxy` can utilize it directly.
- Supports Map Remote and Map Local functionalities.
- Provides HTTP/2 support.
- Additional features can be found in the [configuration documentation](#additional-parameters).

## Unsupported features
- Only supports manual proxy configuration on the client side; transparent proxy mode is not supported.
- WebSocket protocol parsing is currently not supported.

For a detailed explanation of the differences between manual proxy configuration and transparent proxy mode, please refer to the mitmproxy documentation for the Python version: [How mitmproxy works](https://docs.mitmproxy.org/stable/concepts-howmitmproxyworks/). Currently, go-mitmproxy supports "Explicit HTTP" and "Explicit HTTPS" as described in the article.

## Command Line Tool

### Installation

```bash
go install github.com/Incises/go-mitmproxy/cmd/go-mitmproxy@latest
```

### Usage
To start the go-mitmproxy proxy server, use the following command:

```bash
go-mitmproxy
```

By default, the HTTP proxy listens on port 9080, and the web interface is accessible on port 9081.

To intercept HTTPS traffic, you need to install the generated certificate after the first startup. The certificate is automatically created and stored at `~/.mitmproxy/mitmproxy-ca-cert.pem`. For installation instructions, refer to the Python mitmproxy documentation: [About Certificates](https://docs.mitmproxy.org/stable/concepts-certificates/).

### Additional Parameters

You can use the following command to view more options for go-mitmproxy:

```bash
go-mitmproxy -h
```
### Command Line Options

Here are the available command line options for `go-mitmproxy`:

- `-addr [string]`: Proxy listen address (default `":9080"`).
- `-allow_hosts [value]`: A list of allowed hosts.
- `-cert_path [string]`: Path to generate certificate files.
- `-debug [int]`: Debug mode: `1` - print debug log, `2` - show debug form.
- `-f [string]`: Read configuration from a file by passing the file path of a JSON configuration file.
- `-ignore_hosts [value]`: A list of ignored hosts.
- `-map_local [string]`: Map local configuration filename.
- `-map_remote [string]`: Map remote configuration filename.
- `-ssl_insecure`: Do not verify upstream server SSL/TLS certificates.
- `-upstream [string]`: Upstream proxy.
- `-upstream_cert`: Connect to upstream server to look up certificate details (default `true`).
- `-version`: Show `go-mitmproxy` version.
- `-web_addr [string]`: Web interface listen address (default `":9081"`).

## Using go-mitmproxy as a Library

You can import `go-mitmproxy` as a package to develop custom functionalities. 

### Simple Example

```go
package main

import (
	"log"

	"github.com/Incises/go-mitmproxy/proxy"
)

func main() {
	opts := &proxy.Options{
		Addr:              ":9080",
		StreamLargeBodies: 1024 * 1024 * 5,
	}

	p, err := proxy.NewProxy(opts)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(p.Start())
}
```

### Adding Functionality by Developing Plugins

Refer to the [examples](./examples) for adding your own plugins by implementing the `AddAddon` method.

The following are the currently supported event nodes:

```go
type Addon interface {
	// A client has connected to mitmproxy. Note that a connection can correspond to multiple HTTP requests.
	ClientConnected(*ClientConn)

	// A client connection has been closed (either by us or the client).
	ClientDisconnected(*ClientConn)

	// Mitmproxy has connected to a server.
	ServerConnected(*ConnContext)

	// A server connection has been closed (either by us or the server).
	ServerDisconnected(*ConnContext)

	// The TLS handshake with the server has been completed successfully.
	TlsEstablishedServer(*ConnContext)

	// HTTP request headers were successfully read. At this point, the body is empty.
	Requestheaders(*Flow)

	// The full HTTP request has been read.
	Request(*Flow)

	// HTTP response headers were successfully read. At this point, the body is empty.
	Responseheaders(*Flow)

	// The full HTTP response has been read.
	Response(*Flow)

	// Stream request body modifier
	StreamRequestModifier(*Flow, io.Reader) io.Reader

	// Stream response body modifier
	StreamResponseModifier(*Flow, io.Reader) io.Reader
}
```

## Web Interface

Access the web interface at http://localhost:9081/ using your web browser.

### Features

- Detailed view of HTTP/HTTPS requests
- Formatted preview for JSON requests/responses
- Binary mode for viewing response bodies
- Advanced filtering capabilities
- Request breakpoint functionality

### Screenshot Examples

![Web Interface Example 1](./assets/web-1.png)

![Web Interface Example 2](./assets/web-2.png)

![Web Interface Example 3](./assets/web-3.png)

## License

[MIT License](./LICENSE)
