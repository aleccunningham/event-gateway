package controller

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/marjoram/api-glue/pkg/model"
	"github.com/marjoram/api-glue/pkg/provider"

	"github.com/urfave/cli"
)

func getContext(c *cli.Context) context.Context {
	account := c.String("account")
	return context.WithValue(context.Background(), model.ContextKeyAccount, account)
}

func listEvents(c *cli.Context) error {
	// get trigger backend
	eventReaderWriter := backend.NewRedisStore(c.GlobalString("redis"), c.GlobalInt("redis-port"), c.GlobalString("redis-password"), nil, nil)
	// get trigger events
	events, err := eventReaderWriter.GetEvents(getContext(c), c.String("type"), c.String("kind"), c.String("filter"))
	if err != nil {
		return err
	}
	if len(events) == 0 {
		return errors.New("no trigger events found")
	}
	for _, event := range events {
		if c.Bool("quiet") {
			fmt.Println(event.URI)
		} else {
			fmt.Println(event)
		}
	}
	return nil
}

func getEvent(c *cli.Context) error {
	// get trigger backend
	eventReaderWriter := backend.NewStore(c.GlobalString("redis"), c.GlobalInt("redis-port"), c.GlobalString("redis-password"), nil, nil)
	// get trigger events
	event, err := eventReaderWriter.GetEvent(getContext(c), c.Args().First())
	if err != nil {
		return err
	}
	fmt.Println(event)
	return nil
}

func createEvent(c *cli.Context) error {
	// get event provider informer
	eventProvider := provider.NewEventProviderManager(c.GlobalString("config"), c.GlobalBool("skip-monitor"))
	// get trigger backend
	eventReaderWriter := backend.NewStore(c.GlobalString("redis"), c.GlobalInt("redis-port"), c.GlobalString("redis-password"), nil, eventProvider)
	// construct values map
	values := make(map[string]string)
	valueFlag := c.StringSlice("value")
	for _, v := range valueFlag {
		kv := strings.Split(v, "=")
		if len(kv) != 2 {
			return errors.New("invalid value format, must be in form '--value key=val'")
		}
		values[kv[0]] = kv[1]
	}
	// create new event
	event, err := eventReaderWriter.CreateEvent(getContext(c), c.String("type"), c.String("kind"), c.String("secret"), c.String("context"), values)
	if err != nil {
		return err
	}
	// print it out
	fmt.Println("New trigger event successfully created.")
	fmt.Println(event.URI)

	return nil
}

func deleteEvent(c *cli.Context) error {
	// get trigger backend
	eventReaderWriter := backend.NewStore(c.GlobalString("redis"), c.GlobalInt("redis-port"), c.GlobalString("redis-password"), nil, nil)
	// get trigger events
	err := eventReaderWriter.DeleteEvent(getContext(c), c.Args().First(), c.String("context"))
	if err != nil {
		return err
	}
	fmt.Println("Trigger event successfully deleted.")
	return nil
}
