package driver

import (
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/proto"
	nats "github.com/nats-io/go-nats"
	"github.com/nats-rpc/nrpc"
)

// DriverServer is the interface that providers of the service
// Driver should implement.
type DriverServer interface {
	Tentative(ctx context.Context, req TentativeRequest) (resp TentativeResponse, err error)
	Incentive(ctx context.Context, req IncentiveRequest) (resp IncentiveResponse, err error)
}

// DriverHandler provides a NATS subscription handler that can serve a
// subscription using a given DriverServer implementation.
type DriverHandler struct {
	ctx     context.Context
	workers *nrpc.WorkerPool
	nc      nrpc.NatsConn
	server  DriverServer
}

func NewDriverHandler(ctx context.Context, nc nrpc.NatsConn, s DriverServer) *DriverHandler {
	return &DriverHandler{
		ctx:    ctx,
		nc:     nc,
		server: s,
	}
}

func NewDriverConcurrentHandler(workers *nrpc.WorkerPool, nc nrpc.NatsConn, s DriverServer) *DriverHandler {
	return &DriverHandler{
		workers: workers,
		nc:      nc,
		server:  s,
	}
}

func (h *DriverHandler) Subject() string {
	return "driver.Driver.>"
}

func (h *DriverHandler) Handler(msg *nats.Msg) {
	var ctx context.Context
	if h.workers != nil {
		ctx = h.workers.Context
	} else {
		ctx = h.ctx
	}
	request := nrpc.NewRequest(ctx, h.nc, msg.Subject, msg.Reply)
	// extract method name & encoding from subject
	_, _, name, tail, err := nrpc.ParseSubject(
		"driver", 0, "Driver", 0, msg.Subject)
	if err != nil {
		log.Printf("DriverHanlder: Driver subject parsing failed: %v", err)
		return
	}

	request.MethodName = name
	request.SubjectTail = tail

	// call handler and form response
	var immediateError *nrpc.Error
	switch name {
	case "Tentative":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("TentativeHanlder: Tentative subject parsing failed: %v", err)
			break
		}
		var req TentativeRequest
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("TentativeHandler: Tentative request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.server.Tentative(ctx, req)
				if err != nil {
					return nil, err
				}
				return &innerResp, err
			}
		}
	case "Incentive":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("IncentiveHanlder: Incentive subject parsing failed: %v", err)
			break
		}
		var req IncentiveRequest
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("IncentiveHandler: Incentive request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.server.Incentive(ctx, req)
				if err != nil {
					return nil, err
				}
				return &innerResp, err
			}
		}
	default:
		log.Printf("DriverHandler: unknown name %q", name)
		immediateError = &nrpc.Error{
			Type: nrpc.Error_CLIENT,
			Message: "unknown name: " + name,
		}
	}
	if immediateError == nil {
		if h.workers != nil {
			// Try queuing the request
			if err := h.workers.QueueRequest(request); err != nil {
				log.Printf("nrpc: Error queuing the request: %s", err)
			}
		} else {
			// Run the handler synchronously
			request.RunAndReply()
		}
	}

	if immediateError != nil {
		if err := request.SendReply(nil, immediateError); err != nil {
			log.Printf("DriverHandler: Driver handler failed to publish the response: %s", err)
		}
	} else {
	}
}

type DriverClient struct {
	nc      nrpc.NatsConn
	PkgSubject string
	Subject string
	Encoding string
	Timeout time.Duration
}

func NewDriverClient(nc nrpc.NatsConn) *DriverClient {
	return &DriverClient{
		nc:      nc,
		PkgSubject: "driver",
		Subject: "Driver",
		Encoding: "protobuf",
		Timeout: 5 * time.Second,
	}
}

func (c *DriverClient) Tentative(req TentativeRequest) (resp TentativeResponse, err error) {

	subject := c.PkgSubject + "." + c.Subject + "." + "Tentative"

	// call
	err = nrpc.Call(&req, &resp, c.nc, subject, c.Encoding, c.Timeout)
	if err != nil {
		return // already logged
	}

	return
}

func (c *DriverClient) Incentive(req IncentiveRequest) (resp IncentiveResponse, err error) {

	subject := c.PkgSubject + "." + c.Subject + "." + "Incentive"

	// call
	err = nrpc.Call(&req, &resp, c.nc, subject, c.Encoding, c.Timeout)
	if err != nil {
		return // already logged
	}

	return
}

type Client struct {
	nc      nrpc.NatsConn
	defaultEncoding string
	defaultTimeout time.Duration
	pkgSubject string
	Driver *DriverClient
}

func NewClient(nc nrpc.NatsConn) *Client {
	c := Client{
		nc: nc,
		defaultEncoding: "protobuf",
		defaultTimeout: 5*time.Second,
		pkgSubject: "driver",
	}
	c.Driver = NewDriverClient(nc)
	return &c
}

func (c *Client) SetEncoding(encoding string) {
	c.defaultEncoding = encoding
	if c.Driver != nil {
		c.Driver.Encoding = encoding
	}
}

func (c *Client) SetTimeout(t time.Duration) {
	c.defaultTimeout = t
	if c.Driver != nil {
		c.Driver.Timeout = t
	}
}