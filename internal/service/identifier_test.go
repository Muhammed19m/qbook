package service

import (
	"testing"

	assert "github.com/Muhammed19m/qbook/internal/domain/assert"
)

func NewService() *Identifier {
	return &Identifier{}
}

func Test_Identifier(t *testing.T) {
	t.Run("первый id = 0", func(t *testing.T) {
		ider := NewService()
		assert.Equal(t, 0, ider.CurrenID())
	})
	t.Run("последующие id увеличваются на 1 каждый раз когда вызываем GetID", func(t *testing.T) {
		ider := NewService()
		assert.Equal(t, 0, ider.CurrenID())
		assert.Equal(t, 1, ider.GetID())
		assert.Equal(t, 2, ider.GetID())
		assert.Equal(t, 3, ider.GetID())
	})
	t.Run("CurrenID не изменяет значение, а только возвращает", func(t *testing.T) {
		ider := NewService()
		assert.Equal(t, 1, ider.GetID())
		assert.Equal(t, 2, ider.GetID())
		assert.Equal(t, 2, ider.CurrenID())
		assert.Equal(t, 2, ider.CurrenID())
		assert.Equal(t, 2, ider.CurrenID())
		assert.Equal(t, 2, ider.CurrenID())
	})
}
