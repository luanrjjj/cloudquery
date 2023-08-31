package client

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"os"
	"testing"
	"time"
)

type TestOptions struct {
	SkipEmptyJsonB bool
}

func MockTestHelper(t *testing.T, table *schema.Table, builder func(t *testing.T, ctrl *gomock.Controller) Services, options TestOptions) {
	version := "vDev"
	t.Helper()
	ctrl := gomock.NewController(t)
	table.IgnoreInTests = false
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	services := builder(t, ctrl)

	c := Client{
		logger:   l,
		Services: &services,
	}

	sched := scheduler.NewScheduler(scheduler.WithLogger(l))

	tables := schema.Tables{table}
	if err := transformers.TransformTables(tables); err != nil {
		t.Fatal(err)
	}

	messages, err := sched.SyncAll(context.Background(), c, tables)
	if err != nil {
		t.Fatalf("failed to sync: %v", err)
	}
	plugin.ValidateNoEmptyColumns()
}
