package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)


type Proxy struct {
	reverseProxy *httputil.ReverseProxy
}

func NewProxy(target *url.URL) (*Proxy, error) {
    rp := &httputil.ReverseProxy{
        Rewrite: func(r *httputil.ProxyRequest) {
            r.SetURL(target)
            r.SetXForwarded()
        },
    }

    return &Proxy{
        reverseProxy: rp,
    }, nil
}


func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
   p.reverseProxy.ServeHTTP(w, r) 
}