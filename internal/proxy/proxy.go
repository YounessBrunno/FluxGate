package proxy

import (
	"net/http"
    "log"
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