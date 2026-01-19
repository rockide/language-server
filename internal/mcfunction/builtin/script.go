package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Script = &mcfunction.Spec{
	Name:        "script",
	Description: "Script debugger commands.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"debugger"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"listen"},
				},
				{
					Kind: mcfunction.ParameterKindInteger,
					Name: "port",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"debugger"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"connect"},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "host",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "port",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"debugger"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"close"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"profiler"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"start"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"profiler"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"stop"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"diagnostics"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"startcapture"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"diagnostics"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"stopcapture"},
				},
			},
		},
	},
}
