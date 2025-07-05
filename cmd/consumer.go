package cmd

import (
	"context"
	"log"
	"time"

	messagingapp "github.com/arvinpaundra/cent/content/application/messaging"
	"github.com/arvinpaundra/cent/content/config"
	"github.com/arvinpaundra/cent/content/core"
	"github.com/arvinpaundra/cent/content/core/logger"
	"github.com/arvinpaundra/cent/content/core/messaging"
	"github.com/arvinpaundra/cent/content/database/sqlpkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Start consumer/subscriber",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnv(".", ".env", "env")

		pgsql := sqlpkg.NewPostgres()

		sqlpkg.NewConnection(pgsql)

		nc := messaging.NewNats(viper.GetString("NATS_URL"))

		go func() {
			subscriber := messagingapp.NewSetupContentSubscriber(
				sqlpkg.GetConnection(),
				nc.GetConnection(),
				logger.NewLogger(viper.GetString("APP_MODE")),
			)

			err := subscriber.Subscribe(context.Background())
			if err != nil {
				log.Fatalf("failed to start subscriber: %s", err.Error())
			}
		}()

		wait := core.GracefulShutdown(context.Background(), 30*time.Second, map[string]func(ctx context.Context) error{
			"consumer": func(_ context.Context) error {
				conn := nc.GetConnection()
				return conn.Drain()
			},
			"postgres": func(_ context.Context) error {
				return pgsql.Close()
			},
		})

		_ = <-wait
	},
}

func init() {
	rootCmd.AddCommand(consumerCmd)
}
