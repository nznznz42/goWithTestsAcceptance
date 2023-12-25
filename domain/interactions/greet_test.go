package interactions_test

import (
	"testing"

	"goWithTestsAcceptance/domain/interactions"
	"goWithTestsAcceptance/specifications"

	"github.com/alecthomas/assert/v2"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(interactions.Greet),
	)

	t.Run("default", func(t *testing.T) {
		assert.Equal(t, "Hello, World", interactions.Greet(""))
	})
}
