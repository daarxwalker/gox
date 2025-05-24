package datastar

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
	
	"github.com/daarxwalker/gox"
)

func TestBrowser(t *testing.T) {
	t.Run(
		"render custom validity", func(t *testing.T) {
			assert.Equal(
				t,
				`data-custom-validity="$test1 === $test2 ? '' : 'test error'"`,
				gox.Render(CustomValidity("$test1 === $test2 ? '' : 'test error'")),
			)
		},
	)
	t.Run(
		"render intersect", func(t *testing.T) {
			assert.Equal(
				t,
				`data-on-intersect__once="$intersected = true"`,
				gox.Render(Intersect("$intersected = true", WithOnce())),
			)
		},
	)
	t.Run(
		"render interval", func(t *testing.T) {
			assert.Equal(
				t,
				`data-on-interval__duration.500ms="$counter++"`,
				gox.Render(Interval("$counter++", WithDuration("500ms"))),
			)
		},
	)
	t.Run(
		"render any signal change", func(t *testing.T) {
			assert.Equal(
				t,
				`data-on-signal-change="$updates++"`,
				gox.Render(AnySignalChange("$updates++")),
			)
		},
	)
	t.Run(
		"render signal change", func(t *testing.T) {
			assert.Equal(
				t,
				`data-on-signal-change-test="$testUpdates++"`,
				gox.Render(SignalChange("test", "$testUpdates++")),
			)
		},
	)
	t.Run(
		"render persist", func(t *testing.T) {
			assert.Equal(
				t,
				`data-persist__session="test1 test2"`,
				gox.Render(Persist([]string{"test1", "test2"}, WithSession())),
			)
		},
	)
	t.Run(
		"render replace url", func(t *testing.T) {
			assert.Equal(
				t,
				"data-replace-url=\"`/test`\"",
				gox.Render(ReplaceURL("/test")),
			)
		},
	)
	t.Run(
		"render scroll into view", func(t *testing.T) {
			assert.Equal(
				t,
				"data-scroll-into-view__smooth",
				gox.Render(ScrollIntoView(WithSmooth())),
			)
		},
	)
	t.Run(
		"render view transition", func(t *testing.T) {
			assert.Equal(
				t,
				`data-view-transition="$test"`,
				gox.Render(ViewTransition("$test")),
			)
		},
	)
}
