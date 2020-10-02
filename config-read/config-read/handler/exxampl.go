package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	exxampl "exxampl/proto"
)

type Exxampl struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Exxampl) Call(ctx context.Context, req *exxampl.Request, rsp *exxampl.Response) error {
	log.Info("Received Exxampl.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Exxampl) Stream(ctx context.Context, req *exxampl.StreamingRequest, stream exxampl.Exxampl_StreamStream) error {
	log.Infof("Received Exxampl.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&exxampl.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Exxampl) PingPong(ctx context.Context, stream exxampl.Exxampl_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&exxampl.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
