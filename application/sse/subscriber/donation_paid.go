package subscriber

import (
	"context"
	"encoding/json"
	"errors"

	contentcmd "github.com/arvinpaundra/cent/content/application/command/content"
	"github.com/arvinpaundra/cent/content/application/sse"
	"github.com/arvinpaundra/cent/content/domain/content/constant"
	"github.com/arvinpaundra/cent/content/domain/content/service"
	contentinfra "github.com/arvinpaundra/cent/content/infrastructure/content"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DonationPaid struct {
	db     *gorm.DB
	nc     *nats.Conn
	logger *zap.Logger
}

func NewDonationPaid(db *gorm.DB, nc *nats.Conn, logger *zap.Logger) DonationPaid {
	return DonationPaid{
		db:     db,
		nc:     nc,
		logger: logger,
	}
}

func (s DonationPaid) Subscribe(ctx context.Context) error {
	js, err := jetstream.New(s.nc)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}

	stream, err := js.Stream(ctx, constant.StreamDonation)
	if err != nil && !errors.Is(err, jetstream.ErrStreamNotFound) {
		s.logger.Error(err.Error())
		return err
	}

	if errors.Is(err, jetstream.ErrStreamNotFound) {
		stream, err = js.CreateStream(ctx, jetstream.StreamConfig{
			Name:      constant.StreamDonation,
			Retention: jetstream.InterestPolicy,
			Subjects:  []string{constant.EventDonationPaid},
		})

		if err != nil {
			s.logger.Error(err.Error())
			return err
		}
	}

	consumer, err := stream.Consumer(ctx, constant.ConsumerShowDonation)
	if err != nil && !errors.Is(err, jetstream.ErrConsumerNotFound) {
		s.logger.Error(err.Error())
		return err
	}

	if errors.Is(err, jetstream.ErrConsumerNotFound) {
		consumer, err = stream.CreateConsumer(ctx, jetstream.ConsumerConfig{
			Name:    constant.ConsumerShowDonation,
			Durable: constant.ConsumerShowDonation,
		})

		if err != nil {
			s.logger.Error(err.Error())
			return err
		}
	}

	consumerCtx, err := consumer.Consume(func(msg jetstream.Msg) {
		msg.Ack()

		s.logger.Info("event received", zap.String("event", msg.Subject()), zap.ByteString("data", msg.Data()))

		svc := service.NewShowDonation(
			contentinfra.NewContentReaderRepository(s.db),
		)

		var command contentcmd.ShowDonation

		err := json.Unmarshal(msg.Data(), &command)
		if err != nil {
			s.logger.Error(err.Error(), zap.ByteString("data", msg.Data()))
			return
		}

		res, err := svc.Exec(context.Background(), command)
		if err != nil {
			s.logger.Error(err.Error())
			return
		}

		data, err := sse.EncodeJSON(res)
		if err != nil {
			s.logger.Error(err.Error())
			return
		}

		err = sse.Clients.Send(res.UserKey, data)
		if err != nil {
			s.logger.Error(err.Error())
			return
		}
	})

	if err != nil {
		return err
	}

	<-consumerCtx.Closed()

	s.logger.Info("consumer closed")

	return nil
}
