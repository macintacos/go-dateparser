package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// zuNameAtStart is read before any test runs, because -count reruns tests in
// one process, where zu is already built.
var zuNameAtStart string

func TestMain(m *testing.M) {
	zuNameAtStart = zu_Locale.Name
	m.Run()
}

func TestGetLocaleData_BuildsLocaleOnFirstRequest(t *testing.T) {
	assert.Empty(t, zuNameAtStart)

	locale, ok := GetLocaleData("zu")

	assert.True(t, ok)
	assert.Equal(t, "zu", locale.Name)
	assert.NotNil(t, locale.RxCombined)
}

func TestGetLocaleData_BuildsSubLocaleWithItsParent(t *testing.T) {
	locale, ok := GetLocaleData("fr-CA")

	assert.True(t, ok)
	assert.Equal(t, "fr-CA", locale.Name)
	assert.Equal(t, fr_Locale.RxCombined, locale.RxCombined)
	assert.NotEmpty(t, fr_Locale.Name)
}
