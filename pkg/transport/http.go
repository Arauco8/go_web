package transport

import (
	"context"
	"net/http"
)

type Transport interface {
	Server(
		enpoint Endpoint,
		decode func(ctx context.Context, r *http.Request) (interface{}, error),
		encode func(ctx context.Context, w http.ResponseWriter, response interface{}) error,
		encodeError func(ctx context.Context, w http.ResponseWriter, err error) error,
	)
}

type Endpoint func(ctx context.Context, request interface{}) (interface{}, error) // Endpoint is a function that takes a context and a request, and returns a response and an error

type transport struct {
	w   http.ResponseWriter
	r   *http.Request
	ctx context.Context
}

func NewTransport(w http.ResponseWriter, r *http.Request, ctx context.Context) Transport {
	return &transport{
		w:   w,
		r:   r,
		ctx: ctx,
	}
}

func (t *transport) Server( //middleware function that handles the request and response
	enpoint Endpoint,
	decode func(ctx context.Context, r *http.Request) (interface{}, error),
	encode func(ctx context.Context, w http.ResponseWriter, response interface{}) error,
	encodeError func(ctx context.Context, w http.ResponseWriter, err error) error,
) {
	data, err := decode(t.ctx, t.r)
	if err != nil {
		encodeError(t.ctx, t.w, err)
		return
	}

	response, err := enpoint(t.ctx, data)
	if err != nil {
		encodeError(t.ctx, t.w, err)
		return
	}

	if err := encode(t.ctx, t.w, response); err != nil {
		encodeError(t.ctx, t.w, err)
		return
	}
}
