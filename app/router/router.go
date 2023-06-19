package router

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"interview/app/controllers"
	"interview/app/router/middlewares"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type CoreConfigs struct {
	Name string
	Port int
	Host string
}

type handlerFunc func(w http.ResponseWriter, r *http.Request)

type endpoint struct {
	path    string
	method  string
	handler handlerFunc
}

type RouteGroup struct {
	group      string
	middleware []mux.MiddlewareFunc
	endpoints  []endpoint
}

func getRouteGroup() [2]RouteGroup {
	routeGroup := [2]RouteGroup{
		{
			group:      "/",
			middleware: nil,
			endpoints: []endpoint{
				{path: "/", method: "GET", handler: controllers.AppController.Get},
				{path: "/health", method: "GET", handler: controllers.AppController.Health},
			},
		},
		{
			group:      "/api/",
			middleware: nil,
			endpoints:  []endpoint{},
		},
	}

	return routeGroup
}

func InitializeRouter(configs CoreConfigs) error {
	router := mux.NewRouter()
	err := initializeRouter(router, configs)
	if err != nil {
		return err
	}
	return nil
}

func initializeRouter(router *mux.Router, configs CoreConfigs) error {
	router.Use(middlewares.CorsMiddleware)

	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedOrigins: []string{"*"},
	})

	c.Handler(router)

	registerRoutes(router)

	fmt.Println("Routes initialized Successfully")

	router.NotFoundHandler = router.NewRoute().HandlerFunc(controllers.AppController.NotFoundHandler).GetHandler()

	s := &http.Server{
		Addr:    getAddressForServer(configs),
		Handler: router,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Sever started on port: ", configs.Port)
	}()

	fmt.Println("Server started on port: ", configs.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGTSTP)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5000)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		fmt.Println("Server shutdown with error")
	}

	fmt.Println("Server Shut Down Successfully")

	return nil
}

func registerRoutes(router *mux.Router) {
	for _, routeGroup := range getRouteGroup() {
		registerRouteGroup(routeGroup, router)
	}
}

func registerRouteGroup(routeGroup RouteGroup, router *mux.Router) {
	route := router.PathPrefix(routeGroup.group).Subrouter()
	route.Use(routeGroup.middleware...)
	for _, endpoint := range routeGroup.endpoints {
		route.HandleFunc(endpoint.path, endpoint.handler).Methods(endpoint.method)
	}
}

func getAddressForServer(configs CoreConfigs) string {
	addr := ":" + strconv.Itoa(configs.Port)
	return addr
}
