package stream

import (
	"fmt"
	pb "github.com/dukhyungkim/libharago/gen/go/proto"
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

func (s *Client) PublishAction(request *pb.ActionRequest) error {
	msg, err := proto.Marshal(request)
	if err != nil {
		return err
	}

	subject := fmt.Sprintf("harago.action")
	if err = s.nc.Publish(subject, msg); err != nil {
		return err
	}
	return nil
}

func (s *Client) ClamResponse() error {
	if _, err := s.nc.QueueSubscribe("handago.response", "harago", func(msg *nats.Msg) {
		log.Println("Subject:", msg.Subject)
		log.Println("Data:", string(msg.Data))
	}); err != nil {
		return err
	}
	return nil
}
