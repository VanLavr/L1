package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// adapted class
type StandartLibraryRequestHandler struct{}

// adapted method
func (s *StandartLibraryRequestHandler) HandleHttpRequest(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Path
	_, err := fmt.Fprintf(w, "hello from: %s", data)
	if err != nil {
		log.Fatal(err)
	}
}

// adaptee class
type EchoRequestHandler struct{}

// adaptee method
func (e *EchoRequestHandler) HandleHTTP(c echo.Context) error {
	data := c.Request().URL.Path
	return c.HTML(http.StatusOK, fmt.Sprintf("<p>hello from: %s</p>", data))
}

// target for adaptee to adapt
type Target interface {
	HandleHttpRequest(w http.ResponseWriter, r *http.Request)
}

// certain adapter class
type EchoToStandartLibraryHandlerAdapter struct {
	*EchoRequestHandler
}

// target method
func (e *EchoToStandartLibraryHandlerAdapter) HandleHttpRequest(w http.ResponseWriter, r *http.Request) {
	echo := echo.New()
	c := echo.NewContext(r, w)
	if err := e.HandleHTTP(c); err != nil {
		log.Fatal(err)
	}
}

// constructor for certain adapter (returns Target interface to implement targeted methods)
func NewEchoToStandartLibraryHandlerAdapter(e *EchoRequestHandler) Target {
	return &EchoToStandartLibraryHandlerAdapter{EchoRequestHandler: e}
}

func main() {
	mux := http.NewServeMux()

	// adapted to
	stdHandler := StandartLibraryRequestHandler{}
	// adaptee
	echoHandler := EchoRequestHandler{}

	// adapter usage
	echoTostd := NewEchoToStandartLibraryHandlerAdapter(&echoHandler)

	mux.HandleFunc("GET /std/lib", stdHandler.HandleHttpRequest)
	mux.HandleFunc("GET /echo/adapted", echoTostd.HandleHttpRequest)

	log.Println("server is running")
	if err := http.ListenAndServe(":8989", mux); err != nil {
		log.Fatal(err)
	}
}
