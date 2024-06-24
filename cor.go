package goutils

import (
	"context"
)

// 责任链处理节点接口（English: the interface of the node handler in the *Chain of Responsibility*）
type CorHandler[T any] interface {
	SetNext(handler CorHandler[T]) CorHandler[T]
	Handle(ctx context.Context, data T) error
}

// 责任链基础处理节点（English: the base handler of the node in the *Chain of Responsibility*）
type CorBaseHandler[T any] struct {
	next CorHandler[T]
}

// 设置下一个处理节点（English: set the next node handler）
func (h *CorBaseHandler[T]) SetNext(handler CorHandler[T]) CorHandler[T] {
	h.next = handler
	return handler
}

// 处理数据（English: handle the data）
func (h *CorBaseHandler[T]) Handle(ctx context.Context, data T) error {
	if h.next != nil {
		return h.next.Handle(ctx, data)
	}
	return nil
}
