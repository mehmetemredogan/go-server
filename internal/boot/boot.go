package boot

import (
	"crypto/tls"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/mehmetemredogan/go-server/internal/name"
	"golang.org/x/net/http2"
	"log"
	"net/http"
	"os"
)
func Starter() {
	// Setup: httprouter
	router	:= httprouter.New()

	// Call: Service
	router.GET("/", HomeHandler)
	router.GET("/v1/name/:name", ServiceHandler)

	addr	:= os.Getenv("NETWORK_DOMAIN") + ":" + os.Getenv("NETWORK_PORT")

	config	:= &tls.Config {
		MinVersion:					tls.VersionTLS13,
		CurvePreferences:			[]tls.CurveID {
			tls.CurveP521,
			tls.CurveP384,
			tls.CurveP256,
			tls.X25519,
		},
		PreferServerCipherSuites:	true,
		CipherSuites: []uint16 {
			tls.TLS_AES_128_GCM_SHA256,
			tls.TLS_AES_256_GCM_SHA384,
			tls.TLS_CHACHA20_POLY1305_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		},
	}

	srv		:= http.Server {
		Addr:         addr,
		Handler:      router,
		TLSConfig:    config,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	err := http2.ConfigureServer(&srv, &http2.Server{})
	if err != nil {
		log.Fatalln("ERROR: HTTP/2 Server Configuration Error! - ", err.Error())
	}

	crtFile	:= os.Getenv("CRT")
	keyFile	:= os.Getenv("KEY")

	err		= srv.ListenAndServeTLS(crtFile, keyFile)
	if err != nil {
		log.Fatalln("ERROR: HTTP/2 Server Error! - ", err.Error())
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, err := fmt.Fprintf(w, "Hello World!")
	if err != nil {
		return
	}
}

func ServiceHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name.NameService(w, r, ps.ByName("name"))
}