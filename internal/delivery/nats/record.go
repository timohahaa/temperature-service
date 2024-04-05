package nats

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"github.com/timohahaa/temperature-service/internal/entity"
	"github.com/timohahaa/temperature-service/internal/service"
)

type message struct {
	Timestamp string  `json:"timestamp"`
	Value     float32 `json:"value"`
}

type recordConsumer struct {
	recordService    service.RecordService
	natsUrl          string
	subject          string
	concurentWorkers int
	bufferSize       int
	log              *logrus.Logger
}

func NewRecordConsumer(recordService service.RecordService, natsUrl, subject string, concurentWorkers, bufferSize int, log *logrus.Logger) *recordConsumer {
	return &recordConsumer{
		recordService:    recordService,
		natsUrl:          natsUrl,
		subject:          subject,
		concurentWorkers: concurentWorkers,
		bufferSize:       bufferSize,
		log:              log,
	}
}

func (c *recordConsumer) StartRecordConsumer(ctx context.Context) error {
	nc, err := nats.Connect(c.natsUrl)
	if err != nil {
		return err
	}
	defer nc.Close()

	msgChan := make(chan *nats.Msg, c.bufferSize)

	sub, err := nc.ChanSubscribe(c.subject, msgChan)
	if err != nil {
		return err
	}
	defer func() {
		_ = sub.Unsubscribe()
		close(msgChan)
	}()

	nc.SetErrorHandler(func(_ *nats.Conn, _ *nats.Subscription, err error) {
		c.log.Errorf("error receiving message: %v", err)
	})

	nc.SetDisconnectErrHandler(func(_ *nats.Conn, err error) {
		c.log.Errorf("disconected: %v", err)
	})

	nc.SetClosedHandler(func(_ *nats.Conn) {
		c.log.Info("closed")
	})

	var wg sync.WaitGroup
	wg.Add(c.concurentWorkers)

	for _ = range c.concurentWorkers {
		go func() {
			for natsMsg := range msgChan {
				var msg message
				_ = json.Unmarshal(natsMsg.Data, &msg)
				c.recordService.SaveRecord(ctx, entity.Record{Timestamp: msg.Timestamp, Value: msg.Value})
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return nil
}
