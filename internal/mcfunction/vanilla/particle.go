package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var Particle = &mcfunction.Spec{
	Name:        "particle",
	Description: "Creates a particle emitter",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindString,
					Name: "effect",
					Tags: []string{mcfunction.TagParticleId},
				},
				{
					Kind:     mcfunction.ParameterKindVector3,
					Name:     "position",
					Optional: true,
				},
			},
		},
	},
}
