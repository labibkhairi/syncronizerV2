package server

import "prima-integrasi.com/fendiya/syncronizer/internal/controller"

func (server *Server) initializeRoute() {
	// server.Router.HandleFunc("/", controller.Hello).Methods("GET")
	server.Router.HandleFunc("/sync/{awb}", controller.Sync).Methods("GET")
	server.Router.HandleFunc("/sync/bycnotedate/{date}", controller.SyncByCnoteDate).Methods("GET")
	server.Router.HandleFunc("/sync/cmsreturn/{awb}", controller.SyncCmsReturn).Methods("GET")
	server.Router.HandleFunc("/sync/ccc/{awb}", controller.SyncDwhCms).Methods("GET")
}
