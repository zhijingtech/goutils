package main

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ConcreteHandlerA struct {
	CorBaseHandler
	requests []string
}

func (h *ConcreteHandlerA) Handle(_ context.Context, data any) error {
	if data == "A" {
		h.requests = append(h.requests, data.(string))
		return nil
	}
	return h.CorBaseHandler.Handle(context.Background(), data)
}

type ConcreteHandlerB struct {
	CorBaseHandler
	requests []string
}

func (h *ConcreteHandlerB) Handle(_ context.Context, data any) error {
	if data == "B" {
		h.requests = append(h.requests, data.(string))
		return nil
	}
	return h.CorBaseHandler.Handle(context.Background(), data)
}

type ConcreteHandlerC struct {
	CorBaseHandler
	requests []string
}

func (h *ConcreteHandlerC) Handle(_ context.Context, data any) error {
	if data == "C" {
		h.requests = append(h.requests, data.(string))
		return errors.ErrUnsupported
	}
	return nil
}

func TestCOR(t *testing.T) {
	// Create the chain of responsibility
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}
	handlerC := &ConcreteHandlerC{}

	handlerA.SetNext(handlerB).SetNext(handlerC)

	// Send requests through the chain
	requests := []string{"A", "B", "C"}
	err := handlerA.Handle(context.Background(), requests[0])
	assert.NoError(t, err)
	assert.Equal(t, []string{"A"}, handlerA.requests)
	assert.Equal(t, []string(nil), handlerB.requests)
	assert.Equal(t, []string(nil), handlerC.requests)
	err = handlerA.Handle(context.Background(), requests[1])
	assert.NoError(t, err)
	assert.Equal(t, []string{"A"}, handlerA.requests)
	assert.Equal(t, []string{"B"}, handlerB.requests)
	assert.Equal(t, []string(nil), handlerC.requests)
	err = handlerA.Handle(context.Background(), requests[2])
	assert.EqualError(t, err, errors.ErrUnsupported.Error())
	assert.Equal(t, []string{"A"}, handlerA.requests)
	assert.Equal(t, []string{"B"}, handlerB.requests)
	assert.Equal(t, []string{"C"}, handlerC.requests)
}
