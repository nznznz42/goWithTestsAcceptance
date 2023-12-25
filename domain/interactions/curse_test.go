package interactions_test

import (
	"testing"

	"goWithTestsAcceptance/domain/interactions"
	"goWithTestsAcceptance/specifications"
)

func TestCurse(t *testing.T) {
	specifications.CurseSpecification(
		t,
		specifications.CurseAdapter(interactions.Curse),
	)
}
