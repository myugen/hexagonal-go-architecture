package postgres

import (
	"context"

	"github.com/myugen/hexagonal-go-architecture/pkg/logger"

	"github.com/go-pg/pg/v10"
)

type LoggerHook struct {
	verbose bool
	logger  *logger.Logger
}

func (l *LoggerHook) BeforeQuery(ctx context.Context, _ *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (l *LoggerHook) AfterQuery(_ context.Context, q *pg.QueryEvent) error {
	query, _ := q.FormattedQuery()
	var queryLog = l.logger.WithFields(map[string]interface{}{
		"module":    "database",
		"statement": string(query),
	})

	if q.Err != nil {
		queryLog.WithField("error", q.Err).Error("database statement error")
	} else if l.verbose {
		queryLog.Debug("database statement executed")
	}

	return nil
}

func NewLoggerHook(logger *logger.Logger, verbose bool) *LoggerHook {
	return &LoggerHook{verbose, logger}
}
