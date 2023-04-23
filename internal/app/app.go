package app

import (
	"context"
	"sync"
	"time"

	blog_memory_repo "github.com/avelex/blog/internal/blog/repo/memory"
	blog_repo "github.com/avelex/blog/internal/blog/repo/mongodb"
	blog_usecase "github.com/avelex/blog/internal/blog/usecase"
	"github.com/avelex/blog/internal/config"
	http_router "github.com/avelex/blog/internal/controller/http"
	"github.com/gofiber/template/html"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type app struct {
	config *config.Config
	logger *zap.SugaredLogger
}

func NewApp(logger *zap.SugaredLogger, cfg *config.Config) *app {
	return &app{
		logger: logger,
		config: cfg,
	}
}

func (a *app) Run(ctx context.Context) error {
	closer := newCloser()

	var blogRepository blog_usecase.BlogRepository
	if !a.config.Debug {
		ctxInit, cancel := context.WithTimeout(ctx, a.config.InitTimeout)
		defer cancel()

		a.logger.Debug("try connect to mongo_db...")
		mongoClient, err := mongo.Connect(ctxInit, options.Client().ApplyURI(a.config.MongoURI))
		if err != nil {
			return err
		}

		if err = mongoClient.Ping(ctxInit, nil); err != nil {
			return err
		}
		closer.Add(mongoClient.Disconnect)

		a.logger.Debug("successfully connected to mongo_db")

		db := mongoClient.Database("avelex")

		blogRepository = blog_repo.NewBlogRepository(db)
	} else {
		blogRepository = blog_memory_repo.NewRepository()
	}

	blogUsecase := blog_usecase.NewUsecase(a.logger, blogRepository)

	engine := html.New("templates", ".html")

	fiberApp := fiber.New(fiber.Config{
		WriteTimeout:          time.Second * 10,
		ReadTimeout:           time.Second * 10,
		DisableStartupMessage: true,
		Views:                 engine,
	})

	closer.Add(fiberApp.ShutdownWithContext)

	http_router.NewRouter(fiberApp, blogUsecase)

	wg := sync.WaitGroup{}
	wg.Add(1)

	errChan := make(chan error, 1)

	go func() {
		defer wg.Done()

		address := a.config.Host + ":" + a.config.HttpPort

		a.logger.Infof("http server start listen %v", address)
		if err := fiberApp.Listen(address); err != nil {
			a.logger.Errorf("http server error: %v", err)
			errChan <- err
		}
	}()

	select {
	case <-ctx.Done():
		a.logger.Info("stop signal")
	case <-errChan:
	}

	a.logger.Info("starting release resources")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), a.config.ShutdownTimeout)
	defer cancel()

	if err := closer.Close(ctxShutdown); err != nil {
		return err
	}

	wg.Wait()
	close(errChan)

	a.logger.Info("shutdown app gracefully")

	return nil
}
