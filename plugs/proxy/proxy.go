package proxy

import (
	"net/http"

	"golang.org/x/net/context"

	"github.com/AlexanderChen1989/xrest"
)

type proxy struct {
	timeout  int
	location string
	rewrite  struct {
		sourcePatten string
		targetPatten string
		action       string
	}
	proxyHeaders map[string]string // Pass header to backends server
	headers      map[string]string

	proxyHandler http.Handler

	next xrest.Handler
}

func setProxyHeaders(name string, value string) {
}

func addHeaders(name string, value string) {
}

func New(setups ...func(*proxy)) xrest.Plugger {
	p := &proxy{
		location: "/proxy",
		rewrite: &proxy.rewrite{
			sourcePatten: "^/proxy[\\/](.*)$",
			targetPatten: "/$1",
			action:       "break",
		},
		proxyHeaders: map[string]string{
			"Upgrade": "test",
		},
		headers: map[string]string{
			"ETag": "test",
		},
	}

	for _, setup := range setups {
		setup(p)
	}

	// Regiseter Reverse Proxy URL Patten
	// httputil.ReverseProxy()

	p.proxyHandler = http.StripPrefix(
		p.rewrite.sourcePatten,
	)

	return p
}

func (p *proxy) ServeHTTP(ctx context.Context, res http.ResponseWriter, r *http.Request) {
	// Match Reverse Proxy URL Patten
	if strings.Patten(req.URL, p.rewrite.sourcePatten) {
		p.proxyHandler.ServeHTTP(res, req)
		return
	}
	p.next.ServeHTTP(ctx, res, req)
}

func (p *proxy) Plug(h xrest.Handler) xrest.Handler {
	p.next = h
	return p
}
