package application

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

// App is an application components lifecycle manager
type App struct {
	opts     options
	ctx      context.Context
	cancel   func()
	//instance *registry.ServiceInstance
	log      *log.Logger
}

func New(opts ...Option) *App{
	options := options{
		ctx: context.Background(),
		logger: log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		sigs: []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
	}

	if id, err := uuid.NewUUID(); err == nil {
		options.id = id.String()
	}

	for _, o := range opts {
		o(&options)
	}

	ctx, cancel := context.WithCancel(options.ctx)

	return &App{
		opts:     options,
		ctx:      ctx,
		cancel:   cancel,
		//instance: buildInstance(options),
		log:      options.logger,
	}
}

func (a *App) Run() error{
	a.log.Print("service_id", a.opts.id,
		"service_name", a.opts.name,
		"version", a.opts.version,
	)

	g, ctx := errgroup.WithContext(a.ctx)
	for _, srv := range a.opts.servers {
		srv := srv
		g.Go(func() error {
			<-ctx.Done()
			return srv.Stop()
		})
		g.Go(func() error {
			return srv.Start()
		})
	}

	//linux signal terminal
	c := make(chan os.Signal, 1)
	signal.Notify(c, a.opts.sigs...)

	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				a.Stop()
			}
		}
	})

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (a *App) Stop() error {
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}