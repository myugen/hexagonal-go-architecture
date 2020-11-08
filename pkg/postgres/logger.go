package postgres

import (
	"context"

	"github.com/myugen/hexagonal-go-architecture/pkg/logger"

	"github.com/go-pg/pg/v10"
)

type LoggerHook struct {
	Verbose bool
}

func (l *LoggerHook) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (l *LoggerHook) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	query, _ := q.FormattedQuery()
	log := logger.Log()
	var queryLog = log.WithField("module", "database")
	if l.Verbose {
		queryLog = queryLog.WithField("query", string(query))
	}

	if q.Err != nil {
		queryLog.WithField("error", q.Err).Error("database statement error")
	} else if l.Verbose {
		queryLog.Info("database statement executed")
	}

	return nil
}
