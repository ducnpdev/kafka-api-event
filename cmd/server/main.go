package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"event-tracking/config"
	"event-tracking/internal/api"
	"event-tracking/internal/facade"
	"event-tracking/internal/repository/kafka"
	"event-tracking/internal/repository/postgres"
	"event-tracking/internal/usecase"
	awsPkg "event-tracking/pkg/aws"
	dbPkg "event-tracking/pkg/db"
	"event-tracking/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Server ...
type Server struct {
	httpServer *http.Server
	router     *gin.Engine
	cfg        *config.Config
}

// NewServer construct server
func NewServer(ctx context.Context) (*Server, error) {
	router := gin.New()
	s := &Server{
		router: router,
	}
	return s, nil
}

func (s *Server) initSwagger(ctx context.Context) {
	docs.SwaggerInfo.BasePath = "/"
	s.router.GET("/cmd/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (s *Server) initCors(ctx context.Context) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{
		"*",
		"Origin",
		"Content-Length",
		"Content-Type",
		"Authorization",
	}
	s.router.Use(cors.New(corsConfig))
}

// set mode gin debug or releases
func (s *Server) ginMode() {
	if mode := os.Getenv("MODE"); mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
}

// Init server
func (s *Server) Init(ctx context.Context) error {
	var err error

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Default().Println("load env error", err)
		panic(err)
	}
	var (
		once   sync.Once
		onceIn sync.Once
	)
	once.Do(func() {
		logger.Newlogger(cfg.Logger.Mode, cfg.Logger.Level, cfg.Logger.Encoding, cfg.ServiceName)
	})

	s.cfg = cfg
	s.ginMode()
	s.initCors(ctx)

	onceIn.Do(func() {
		err = s.Singleton(ctx)
		if err != nil {
			logger.GLogger.Fatalf("func init singleton error %s", err)
		}
	})

	return nil
}

func (s *Server) ListenHTTP() error {

	address := fmt.Sprintf(":%d", s.cfg.Http.Port)

	s.httpServer = &http.Server{
		Handler: s.router,
		Addr:    address,
	}

	zap.S().Info("Start server at port %d", s.cfg.Http.Port)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Singleton(ctx context.Context) error {
	var (
		errDB      error
		postgresDB *gorm.DB
	)

	postgresDB, errDB = dbPkg.InitPostgres(s.cfg.Postgres)
	if errDB != nil {
		logger.GLogger.Panicf("InitPostgres error %s", errDB)
	}

	kafkaNotificationWriter := awsPkg.KafkaWriter(awsPkg.KafkaWriterProperty{
		Brokers:   s.cfg.Kafka.KafkaWriter.BrokerAddress,
		Topic:     s.cfg.Kafka.KafkaWriter.Topic,
		BatchSize: s.cfg.Kafka.KafkaWriter.BatchSize,
	})
	kafkaTranHistReader := awsPkg.KafkaReader(awsPkg.KafkaReaderProperty{
		Brokers: s.cfg.Kafka.KafkaReader.BrokerAddress,
		Topic:   s.cfg.Kafka.KafkaReader.Topic,
		GroupID: s.cfg.Kafka.KafkaReader.GroupID,
		Enable:  s.cfg.Kafka.KafkaReader.Enable,
	})

	kakfaWriter := kafka.NewWriterRepository(s.cfg, kafkaNotificationWriter)
	kafkaReader := kafka.NewReaderRepository(kafkaTranHistReader)

	healthCheckRepo := postgres.NewHealthCheckRepository(postgresDB)

	healthCheckFc := facade.NewHealthCheckFacade(s.cfg, healthCheckRepo)
	eventFacade := facade.NewEventFacade(s.cfg, kafkaReader, kakfaWriter)

	healthCheckUc := usecase.NewHealthCheckUsecase(s.cfg, healthCheckFc)
	eventUsecase := usecase.NewEventUseCase(s.cfg, eventFacade)

	useCase := &usecase.UseCase{
		HealthCheckCase: healthCheckUc,
		EventUseCase:    eventUsecase,
	}

	s.Router(useCase)
	return nil
}

func (s *Server) Router(
	useCase *usecase.UseCase,
) error {
	healthyGroup := s.router.Group("/v1")
	healthy := api.NewHealthHandler(s.cfg, useCase.HealthCheckCase)
	healthy.HealthRouter(healthyGroup)

	groupEvent := s.router.Group("event/v1")

	eventAPI := api.NewEventHandler(useCase.EventUseCase)
	eventAPI.EventRouter(groupEvent)

	return nil
}

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	s, err := NewServer(ctx)
	if err != nil {
		log.Default().Println("New server panic")
		panic(err)
	}
	s.initSwagger(ctx)

	err = s.Init(ctx)
	if err != nil {
		logger.GLogger.Fatalf("min init func error %s", err)
	}
	logger.GLogger.Infof("Start project http ok at port %d", s.cfg.Http.Port)
	go func() {
		if err := s.ListenHTTP(); err != nil {
			logger.GLogger.Fatalf("ListenHTTP  err %s", err)
		}
	}()

	term := make(chan os.Signal, 1)
	//signal.Notify(term, unix.SIGTERM)
	signal.Notify(term, os.Interrupt)

	select {
	case <-term:
	}
	time.Sleep(time.Second * time.Duration(s.cfg.SignalShutDown))
	logger.GLogger.Info("Shutting down program")
}
