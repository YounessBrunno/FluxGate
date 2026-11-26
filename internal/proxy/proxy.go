package proxy

import (
	"net/http"
    "log"
	"net/http/httputil"
	"net/url"
    "net"
    "time"

)


type Proxy struct {
	reverseProxy *httputil.ReverseProxy
}

func NewProxy(target *url.URL) (*Proxy, error) {

    transport := http.DefaultTransport.(*http.Transport).Clone()
    transport.DialContext = (&net.Dialer{
        Timeout:   3 * time.Second,
        KeepAlive: 30 * time.Second,
    }).DialContext
    transport.TLSHandshakeTimeout = 3 * time.Second
    transport.ResponseHeaderTimeout = 2 * time.Second
    transport.IdleConnTimeout = 60 * time.Second
    transport.ExpectContinueTimeout = 1 * time.Second

    rp := &httputil.ReverseProxy{
        Transport: transport,
        Rewrite: func(r *httputil.ProxyRequest) {
            r.SetURL(target)
            r.SetXForwarded()
        },
        ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
            log.Printf(
                "backend request failed: method=%s path=%s error=%v",
                r.Method,
                r.URL.Path,
                err,
            )

            http.Error(
                w,
                "The backend service is unavailable.",
                http.StatusBadGateway,
            )
       },
    }

    return &Proxy{
        reverseProxy: rp,
    }, nil
}


func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
   p.reverseProxy.ServeHTTP(w, r) 
}