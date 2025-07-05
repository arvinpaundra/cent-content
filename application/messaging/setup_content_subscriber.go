package messaging

import (
	"context"
	"encoding/json"
	"errors"
	contentcmd "github.com/arvinpaundra/cent/content/application/command/content"
	"github.com/arvinpaundra/cent/content/domain/content/constant"
	"github.com/arvinpaundra/cent/content/domain/content/service"
	contentinfra "github.com/arvinpaundra/cent/content/infrastructure/content"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SetupContentSubscriber struct {
	db     *gorm.DB
	nc     *nats.Conn
	logger *zap.Logger
}

func NewSetupContentSubscriber(db *gorm.DB, nc *nats.Conn, logger *zap.Logger) SetupContentSubscriber {
	return SetupContentSubscriber{
		db:     db,
		nc:     nc,
		logger: logger,
	}
}

func (m SetupContentSubscriber) Subscribe(ctx context.Context) error {
	js, err := jetstream.New(m.nc)
	if err != nil {
		m.logger.Error(err.Error())
		return err
	}

	stream, err := js.Stream(ctx, constant.StreamUser)
	if err != nil && !errors.Is(err, jetstream.ErrStreamNotFound) {
		m.logger.Error(err.Error())
		return err
	}

	if errors.Is(err, jetstream.ErrStreamNotFound) {
		stream, err = js.CreateStream(ctx, jetstream.StreamConfig{
			Name:      constant.StreamUser,
			Retention: jetstream.InterestPolicy,
			Subjects:  []string{constant.EventUserCreated},
		})

		if err != nil {
			m.logger.Error(err.Error())
			return err
		}
	}

	consumer, err := stream.Consumer(ctx, constant.ConsumerSetupContent)
	if err != nil && !errors.Is(err, jetstream.ErrConsumerNotFound) {
		m.logger.Error(err.Error())
		return err
	}

	if errors.Is(err, jetstream.ErrConsumerNotFound) {
		consumer, err = stream.CreateConsumer(ctx, jetstream.ConsumerConfig{
			Name:    constant.ConsumerSetupContent,
			Durable: constant.ConsumerSetupContent,
		})

		if err != nil {
			m.logger.Error(err.Error())
			return err
		}
	}

	consumerContex, err := consumer.Consume(func(msg jetstream.Msg) {
		msg.Ack()

		m.logger.Info("event received", zap.String("event", msg.Subject()), zap.ByteString("data", msg.Data()))

		var command contentcmd.CreateSetupContent

		err := json.Unmarshal(msg.Data(), &command)
		if err != nil {
			m.logger.Error(err.Error(), zap.ByteString("data", msg.Data()))
			return
		}

		svc := service.NewSetupContent(
			contentinfra.NewContentWriterRepository(m.db),
			contentinfra.NewUnitOfWork(m.db),
		)

		err = svc.Exec(ctx, command)
		if err != nil {
			m.logger.Error(err.Error())
			return
		}
	})

	if err != nil {
		return err
	}

	<-consumerContex.Closed()

	m.logger.Info("consumer closed")

	return nil
}
