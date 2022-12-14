package initialize

import (
	"log"
	"os"
	"time"

	"github.com/beslow/goblog/config"
	"github.com/getsentry/sentry-go"
)

// fetch sentry dsn order:
// 1. from ENV
// 2. from config file
func GetSentryDsn() string {
	dsn := os.Getenv("SENTRY_DSN")
	if dsn == "" {
		if config.Sentry.Enable {
			dsn = config.Sentry.Dsn
		}
	}

	return dsn
}

func init() {
	dsn := GetSentryDsn()

	if dsn == "" {
		return
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn: dsn,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})

	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)
}
