package composite

import (
	"github.com/askft/go-behave/core"
)

// Selector creates a new selector node.
func Selector(children ...core.Node) core.Node {
	base := core.NewComposite("Selector", children)
	return &selector{Composite: base}
}

// selector ...
type selector struct {
	*core.Composite
}

// Enter ...
func (s *selector) Enter(ctx *core.Context) {
	s.Composite.CurrentChild = 0
}

// Tick ...
func (s *selector) Tick(ctx *core.Context) core.Status {
	for {
		status := core.Update(s.Children[s.CurrentChild], ctx)
		if status != core.StatusFailure {
			return status
		}
		s.Composite.CurrentChild++
		if s.CurrentChild >= len(s.Children) {
			return core.StatusFailure
		}
	}
}

// Leave ...
func (s *selector) Leave(ctx *core.Context) {}
