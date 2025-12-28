package handlers

import (
	"github.com/bmatcuk/doublestar/v4"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/protocol"
	"github.com/rockide/language-server/shared"
)

func difference(pattern shared.Pattern, a []core.Symbol, b []core.Symbol) []core.Symbol {
	s, ok := pattern.ToString()
	if !ok {
		return nil
	}
	res := []core.Symbol{}
	set := mapset.NewThreadUnsafeSet[string]()
	for _, ref := range b {
		if doublestar.MatchUnvalidated("**/"+s, string(ref.URI)) {
			set.Add(ref.Value)
		}
	}
	for _, ref := range a {
		if !set.ContainsOne(ref.Value) {
			res = append(res, ref)
		}
	}
	return res
}

func isInside(r protocol.Range, p protocol.Position) bool {
	return protocol.ComparePosition(p, r.Start) != -1 && protocol.ComparePosition(p, r.End) != 1
}
