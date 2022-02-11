package stream

import (
	pbAct "github.com/dukhyungkim/libharago/gen/go/proto/action"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
	"harago/config"
	"log"
	"strings"
	"time"
)

type Client struct {
	nc      *nats.Conn
	timeout time.Duration
}

func NewStreamClient(cfg *config.Nats) (*Client, error) {
	nc, err := nats.Connect(strings.Join(cfg.Servers, ","),
		nats.UserInfo(cfg.Username, cfg.Password))
	if err != nil {
		return nil, err
	}

	return &Client{nc: nc, timeout: cfg.Timeout}, nil
}

func (s *Client) Close() {
	s.nc.Close()
}

func (s *Client) PublishAction(subject string, request *pbAct.ActionRequest) error {
	msg, err := proto.Marshal(request)
	if err != nil {
		return err
	}

	if err := s.nc.Publish(subject, msg); err != nil {
		return err
	}
	return nil
}

type ResponseHandler func(message *pbAct.ActionResponse)

func (s *Client) ClamResponse(handler ResponseHandler) error {
	if _, err := s.nc.QueueSubscribe("handago.response", "harago", func(msg *nats.Msg) {
		var message pbAct.ActionResponse
		if err := proto.Unmarshal(msg.Data, &message); err != nil {
			log.Println(err)
			return
		}
		handler(&message)
	}); err != nil {
		return err
	}
	return nil
}
