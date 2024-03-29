package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		elems2 := make([]int, 0, l.Len())
		for i2 := l.Front(); i2 != nil; i2 = i2.Next {
			elems2 = append(elems2, i2.Value.(int))
		}
		require.Equal(t, []int{80, 60, 40, 10, 30, 50, 70}, elems2)

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]

		elems1 := make([]int, 0, l.Len())
		for i1 := l.Front(); i1 != nil; i1 = i1.Next {
			elems1 = append(elems1, i1.Value.(int))
		}
		require.Equal(t, []int{80, 60, 40, 10, 30, 50, 70}, elems1)

		l.MoveToFront(l.Back()) // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
	t.Run("yet another test", func(t *testing.T) {
		l := NewList()
		l.PushBack(10)
		l.PushFront(20)
		l.PushFront(30)

		require.Equal(t, 3, l.Len()) // [30, 20, 10]

		l.MoveToFront(l.Back())

		require.Equal(t, 10, l.Front().Value) // [10, 30, 20]
	})
}
