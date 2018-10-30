package estor

import (
	"estor/controller"
	"estor/utils"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/ssrs100/conf"
	"github.com/ssrs100/logs"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var (
	log = logs.GetLogger()

	stor_config = "estor.json"
)

// Config struct provides configuration fields for the server.
type Server struct {
	configure *conf.Config
}

var s Server

func (s *Server) RegisterRoutes() *httprouter.Router {
	log.Debug("Setting route info...")

	// Set the router.
	router := httprouter.New()

	// Set router options.
	router.HandleMethodNotAllowed = true
	router.HandleOPTIONS = true
	router.RedirectTrailingSlash = true

	// Set the routes for the application.

	// Route for health check
	router.GET("/es/heart", controller.HealthCheck)

	// Routes for users
	router.GET("/es/v1/users", controller.GetUsers)
	router.POST("/es/v1/users", controller.CreateUser)
	router.DELETE("/es/v1/users/:username", controller.DeleteUser)

	return router
}

var stop = make(chan os.Signal)

// Start sets up and starts the main server application
func Start() error {
	log.Info("Setting up server...")

	basedir := utils.GetAppBaseDir()
	if len(basedir) == 0 {
		log.Error("Evironment APP_BASE_DIR(app installed root path) should be set.")
		os.Exit(1)
	}


	//获取配置信息
	appConfig := filepath.Join(basedir, "conf", stor_config)
	s.configure = conf.LoadFile(appConfig)
	if s.configure == nil {
		errStr := fmt.Sprintf("Can not load %s.", stor_config)
		log.Error(errStr)
		os.Exit(1)
	}

	router := s.RegisterRoutes()
	host := s.configure.GetString("host")
	port := s.configure.GetInt("port")
	server := &http.Server{Addr: host + ":" + strconv.Itoa(port), Handler: router}

	log.Debug("Starting server on port %d", port)

	err := server.ListenAndServeTLS(s.configure.GetString("cert"), s.configure.GetString("key"))
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}

	return nil
}
