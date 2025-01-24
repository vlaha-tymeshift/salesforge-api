package main

import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net/http"
	"os"
	"os/signal"
	"salesforge-api/internal/api"
	"salesforge-api/internal/config"
	"salesforge-api/internal/persistence"
	"salesforge-api/internal/psql"
	"salesforge-api/internal/service"
	"syscall"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	var l *zap.Logger

	// Config.
	cfg, err := config.LoadConfig(os.DirFS("."))
	if err != nil {
		log.Fatal("failed to load config", zap.Error(err))
	}

	// Logging.
	if cfg.Logger.Format == "json" {
		l = zap.New(zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			zapcore.AddSync(os.Stdout),
			zap.NewAtomicLevelAt(parseLogLevel(cfg.Logger.Level)),
		))
	} else {
		l = zap.New(zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			zapcore.AddSync(os.Stdout),
			zap.NewAtomicLevelAt(parseLogLevel(cfg.Logger.Level)),
		))
	}
	defer l.Sync()
	l.Info("logger initialized")
	defer l.Info("server exiting")

	// Database.
	db, err := psql.New(cfg.MySql)
	if err != nil {
		l.Fatal("failed to connect to database", zap.Error(err))
	}
	l.Info("connected to database")
	defer db.Close()

	// Services.
	sequenceRepository := persistence.NewSequenceRepository(db)
	sequenceService := service.NewSequenceService(sequenceRepository)

	server := api.NewServer(cfg.Server, sequenceService, db, l)
	// Start server.
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			l.Fatal("server failed", zap.Error(err))
		}
	}()

	// Listen to interrupts.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, interruptSignals...)
	sig := <-sigChan
	l.Info("shutting down", zap.String("signal", sig.String()))
	os.Exit(0)
}

func parseLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
