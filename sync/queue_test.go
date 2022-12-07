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
