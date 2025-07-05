package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	restrouter "github.com/arvinpaundra/cent/content/application/rest/router"
	sserouter "github.com/arvinpaundra/cent/content/application/sse/router"
	"github.com/arvinpaundra/cent/content/application/sse/subscriber"
	"github.com/arvinpaundra/cent/content/config"
	"github.com/arvinpaundra/cent/content/core"
	"github.com/arvinpaundra/cent/content/core/logger"
	"github.com/arvinpaundra/cent/content/core/messaging"
	"github.com/arvinpaundra/cent/content/core/validator"
	"github.com/arvinpaundra/cent/content/database/sqlpkg"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var restPort string

var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "Start rest server",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnv(".", ".env", "env")

		pgsql := sqlpkg.NewPostgres()

		sqlpkg.NewConnection(pgsql)

		nc := messaging.NewNats(viper.GetString("NATS_URL"))

		g := gin.New()

		restrouter.Register(g, sqlpkg.GetConnection(), validator.NewValidator())

		subs := subscriber.NewDonationPaid(sqlpkg.GetConnection(), nc.GetConnection(), logger.NewLogger(viper.GetString("APP_MODE")))
		go subs.Subscribe(context.Background())

		sserouter.Register(g, sqlpkg.GetConnection(), nc.GetConnection(), logger.NewLogger(viper.GetString("APP_MODE")))

		srv := http.Server{
			Addr:    fmt.Sprintf(":%s", restPort),
			Handler: g,
		}

		go func() {
			if err := srv.ListenAndServe(); err != http.ErrServerClosed {
				log.Fatalf("failed to start server: %s", err.Error())
			}
		}()

		wait := core.GracefulShutdown(context.Background(), 30*time.Second, map[string]func(ctx context.Context) error{
			"rest-server": func(_ context.Context) error {
				return srv.Close()
			},
			"postgres": func(_ context.Context) error {
				return pgsql.Close()
			},
			"nats": func(_ context.Context) error {
				return nc.Close()
			},
		})

		_ = <-wait
	},
}

func init() {
	restCmd.Flags().StringVarP(&restPort, "port", "p", "8070", "bind rest server to port")
	rootCmd.AddCommand(restCmd)
}
