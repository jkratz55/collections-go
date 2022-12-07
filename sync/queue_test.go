package sync

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewBlockingQueue(t *testing.T) {
	queue := NewBlockingQueue[string](10)
	assert.Equal(t, 10, queue.capacity)
	assert.Equal(t, 0, queue.Size())
}

func TestBlockingQueue_Offer(t *testing.T) {
	q := NewBlockingQueue[string](10)
	assert.NoError(t, q.Offer(context.Background(), "a"))
	assert.NoError(t, q.Offer(context.Background(), "b"))
	assert.NoError(t, q.Offer(context.Background(), "c"))
	assert.NoError(t, q.Offer(context.Background(), "d"))
	assert.NoError(t, q.Offer(context.Background(), "e"))
	assert.NoError(t, q.Offer(context.Background(), "f"))
	assert.NoError(t, q.Offer(context.Background(), "g"))
	assert.NoError(t, q.Offer(context.Background(), "h"))
	assert.NoError(t, q.Offer(context.Background(), "i"))
	assert.NoError(t, q.Offer(context.Background(), "j"))

	// This should block and timeout resulting in an error
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	assert.ErrorIs(t, q.Offer(ctx, "j"), context.DeadlineExceeded)
	cancel()
}

func TestBlockingQueue_TryOffer(t *testing.T) {
	q := NewBlockingQueue[string](10)
	assert.True(t, q.TryOffer("1"))
	assert.True(t, q.TryOffer("2"))
	assert.True(t, q.TryOffer("3"))
	assert.True(t, q.TryOffer("4"))
	assert.True(t, q.TryOffer("5"))
	assert.True(t, q.TryOffer("6"))
	assert.True(t, q.TryOffer("7"))
	assert.True(t, q.TryOffer("8"))
	assert.True(t, q.TryOffer("9"))
	assert.True(t, q.TryOffer("10"))
	assert.False(t, q.TryOffer("11")) // This one will fail to add because Queue is at capacity
}

func TestBlockingQueue_Poll(t *testing.T) {
	q := NewBlockingQueue[int](5)

	go func() {
		for i := 1; i <= 5; i++ {
			q.data <- i
		}
	}()

	res, err := q.Poll(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, res, 1)

	res, err = q.Poll(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, res, 2)

	res, err = q.Poll(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, res, 3)

	res, err = q.Poll(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, res, 4)

	res, err = q.Poll(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, res, 5)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	res, err = q.Poll(ctx)
	assert.ErrorIs(t, err, context.DeadlineExceeded)
}

func TestBlockingQueue_TryPoll(t *testing.T) {
	q := NewBlockingQueue[int](5)

	val, ok := q.TryPoll()
	assert.False(t, ok)

	for i := 1; i <= 5; i++ {
		q.data <- i
	}

	val, ok = q.TryPoll()
	assert.True(t, ok)
	assert.Equal(t, 1, val)

	val, ok = q.TryPoll()
	assert.True(t, ok)
	assert.Equal(t, 2, val)

	val, ok = q.TryPoll()
	assert.True(t, ok)
	assert.Equal(t, 3, val)

	val, ok = q.TryPoll()
	assert.True(t, ok)
	assert.Equal(t, 4, val)

	val, ok = q.TryPoll()
	assert.True(t, ok)
	assert.Equal(t, 5, val)

	val, ok = q.TryPoll()
	assert.False(t, ok)
}

func TestBlockingQueue_Capacity(t *testing.T) {
	q := NewBlockingQueue[int](10)
	assert.Equal(t, 10, q.Capacity())
}

func TestBlockingQueue_CapacityRemaining(t *testing.T) {
	q := NewBlockingQueue[int](10)
	assert.Equal(t, 10, q.CapacityRemaining())

	q.data <- 1
	q.data <- 2
	q.data <- 3
	q.data <- 4
	q.data <- 5

	assert.Equal(t, 5, q.CapacityRemaining())
}

func TestBlockingQueue_Empty(t *testing.T) {
	q := NewBlockingQueue[int](10)
	assert.True(t, q.Empty())

	q.data <- 1

	assert.False(t, q.Empty())
}

func TestBlockingQueue_Size(t *testing.T) {
	q := NewBlockingQueue[int](10)
	assert.Equal(t, 0, q.Size())

	q.data <- 1
	q.data <- 2
	q.data <- 3
	q.data <- 4
	q.data <- 5

	assert.Equal(t, 5, q.Size())
}

func TestBlockingQueue_Clear(t *testing.T) {
	q := NewBlockingQueue[int](10)
	assert.Equal(t, 0, len(q.data))
	q.Clear()
	assert.Equal(t, 0, len(q.data))

	q.data <- 1
	q.data <- 2
	q.data <- 3
	q.data <- 4
	q.data <- 5
	assert.Equal(t, 5, len(q.data))
	q.Clear()
	assert.Equal(t, 0, len(q.data))
}
