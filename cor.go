package main

import (
	"context"
)

// 责任链处理节点接口（English: the interface of the node handler in the *Chain of Responsibility*）
type CorHandler interface {
	SetNext(handler CorHandler) CorHandler
	Handle(ctx context.Context, data any) error
}

// 责任链基础处理节点（English: the base handler of the node in the *Chain of Responsibility*）
type CorBaseHandler struct {
	next CorHandler
}

// 设置下一个处理节点（English: set the next node handler）
func (h *CorBaseHandler) SetNext(handler CorHandler) CorHandler {
	h.next = handler
	return handler
}

// 处理数据（English: handle the data）
func (h *CorBaseHandler) Handle(ctx context.Context, data any) error {
	if h.next == nil {
		return nil
	}
	return h.next.Handle(ctx, data)
}
