package app

import (
	"banner/internal/config"
	endpointsbaner "banner/internal/endpoint/endpointsBaner"
	"banner/internal/endpoint/midllware"
	"banner/internal/logger"
	psql "banner/internal/repo/Db"
	servicebanner "banner/internal/service/serviceBanner"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type App struct {
	logger  *logrus.Logger
	mux     *mux.Router
	address string
}

func New() *App {

	var app App
	app.logger = logger.New()

	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatal(err)
	}

	db, err := psql.InitDb(app.logger)
	if err != nil {
		app.logger.Fatal(err)
	}

	cfg := config.ServNew()

	app.address = cfg.Host + ":" + cfg.Port

	s := servicebanner.InitServiceBanner(app.logger, db)

	h := endpointsbaner.Init(app.logger, s)

	m := midllware.Init(app.logger)

	app.mux = mux.NewRouter()

	app.mux.Handle("/user_banner", m.ReqID(m.Aut(h.GetBanner))).Methods("GET")
	app.mux.Handle("/banner", m.ReqID(m.Aut(h.Filter))).Methods("GET")
	app.mux.Handle("/banner", m.ReqID(m.Aut(h.Insert))).Methods("POST")
	app.mux.Handle("/banner/{id}", m.ReqID(m.Aut(h.Update))).Methods("PATCH")
	app.mux.Handle("/banner/{id}", m.ReqID(m.Aut(h.Delete))).Methods("DELETE")

	return &app

}

func (a *App) Run() {
	a.logger.Info("Listening on: ", a.address)
	a.logger.Fatal(http.ListenAndServe(a.address, a.mux))

}
