package light

import (
	"github.com/go-logr/logr"
	"github.com/ptrszlgy/twamp/pkg/client"
)

var log logr.Logger

func SetLogger(logger logr.Logger) {
	log = logger.WithName("light")
}

func Run() error {
	c := client.NewClient()
	conn, err := c.Connect("localhost:862")
	if err != nil {
		log.Error(err, "could not connect to TWAMP loopback server")
		return err
	}

	session, err := conn.CreateSession(client.TwampSessionConfig{ReceiverPort: 862, SenderPort: 20862, Padding: 0})
	if err != nil {
		log.Error(err, "could not create TWAMP-Light session")
		return err
	}

	test, err := session.CreateTest()
	if err != nil {
		log.Error(err, "")
		return err
	}

	done := make(chan bool, 1)
	defer close(done)

	ping, err := test.Ping(5, 1, done)
	if err != nil {
		return err
	}

	for _, res := range ping.Results {
		res.PrintResults()
	}

	return nil
}
