package proxy

import ("net/http"
         "net/url")


func main() {
    
	http.ListenAndServe(":8080", nil)
}