package proxy

import ("net/http"
        
		 _ "github.com/YounessBrunno/FluxGate/internal/proxy"
		)
		


func main() {
    
	http.ListenAndServe(":8080", nil)
}