package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"golang.org/x/crypto/acme/autocert"

	"github.com/Traliaa/go_back_2/api/server/handler"
)

//
//type Server struct {
//}
//
//func NewServer() *Server {
//	return &Server{}
//}

type Server struct {
	Addr string

	srv    http.Server
	srvTLS http.Server
	log    *zap.Logger
	router *chi.Mux
}

func NewServer(info handler.VersionInfo, Addr string) *Server {

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache("cert-cache"),
		HostPolicy: autocert.HostWhitelist("localhost"),
	}
	s := &Server{
		//log: &log,
		Addr: Addr,
	}
	s.srv = http.Server{
		Addr:    ":" + Addr,
		Handler: handler.NewHandler(info),
	}
	s.srvTLS = http.Server{
		Addr:    ":443",
		Handler: handler.NewHandler(info),
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}
	return s
}

func (s *Server) Serve(ctx context.Context) {

	go func() {
		fmt.Printf("start server on port: %s", s.Addr)
		err := s.srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	//go func() {
	//	fmt.Println("start TLS server on port: 443", s.Addr)
	//	err := s.srvTLS.ListenAndServe()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}()

	<-ctx.Done()
	s.Stop(ctx)

}

func (s *Server) Stop(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

	err := s.srv.Shutdown(ctx)
	if err != nil {
		//s.log.Error(err.Error())
	}
	cancel()

}
