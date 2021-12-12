package application

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/wilorios/microservice-template-go/internal/configurations"
	"github.com/wilorios/microservice-template-go/internal/routes"
)

// Instance defines an application.
type Instance struct {
	Name          string
	configuration configurations.Application
}

// Event contains an application event.
type Event struct {
	Message string
	Error   error
}

// New creates a new application.
func New(args []string) *Instance {
	return &Instance{}
}

// Start initialize and start the Instance
func (i *Instance) Start() error {
	log.Println("level", "INFO", "msg", "starting application")

	confError := i.loadConfiguration()
	if confError != nil {
		panic(confError)
	}
	log.Println("level", "DEBUG", "msg", "application configuration", "parameters", i.configuration)

	eventStream := make(chan Event)
	i.listenToOSSignal(eventStream)
	i.startWebServer(eventStream)

	eventMessage := <-eventStream
	fmt.Println(
		"level", "INFO",
		"msg", "ending server",
		"event", eventMessage.Message,
	)

	if eventMessage.Error != nil {
		fmt.Println(
			"level", "ERROR",
			"msg", "ending server with error",
			"error", eventMessage.Error,
		)
		return eventMessage.Error
	}
	return nil
}

func (i *Instance) listenToOSSignal(eventStream chan<- Event) {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		osSignal := fmt.Sprintf("%s", <-c)
		event := Event{
			Message: osSignal,
		}
		eventStream <- event
	}()
}

// startWebServer starts the web server.
func (i *Instance) startWebServer(eventStream chan<- Event) {
	go func() {
		routes.SetupRouter(i.configuration)
		eventStream <- Event{
			Message: "web server was ended",
		}
	}()
}

func (i *Instance) loadConfiguration() error {
	applicationSetUp, err := configurations.Load()
	if err != nil {
		log.Println("level", "ERROR", "msg", "application setup could not be loaded", "error", err)
		return errors.New("application setup could not be loaded")
	}
	fmt.Println("setup ", applicationSetUp)
	i.configuration = applicationSetUp
	return nil
}
