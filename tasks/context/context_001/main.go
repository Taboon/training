package main

import (
	"context"
	"fmt"
	"time"
)

// Вводные:
// Функция executeTask может зависнуть.
// В ней не предусмотрен механизм отмены.
// Она не принимает Context или канал с событием отмены как аргумент.

func main() {

	// Задача:
	// Для функции executeTask написать обертку executeTaskWithTimeout.
	// Функция executeTaskWithTimeout принимает аргументом тайм-аут,
	// через который функция executeTask будет отменена.
	// Если executeTask была отменена по тайм-ауту, нужно вернуть ошибку
	timeout := time.Duration(time.Second * 12)
	err := executeTaskWithTimeout(context.Background(), timeout)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("DONE")
}

func executeTask() {
	time.Sleep(1 * time.Second)
}

func executeTaskWithTimeout(ctx context.Context, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan struct{})

	go func() {
		defer close(done)
		executeTask()
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
