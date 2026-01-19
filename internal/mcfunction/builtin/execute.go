package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var ifUnless = []string{"if", "unless"}

var Execute = &mcfunction.Spec{
	Name:        "execute",
	Description: "Executes a command on behalf of one or more entities.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "subcommand",
					Literals: []string{"as"},
					Tags:     []string{mcfunction.TagExecuteChain},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "origin",
				},
				{
					Kind: mcfunction.ParameterKindChainedCommand,
					Name: "chainedCommand",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "subcommand",
					Literals: []string{"at"},
					Tags:     []string{mcfunction.TagExecuteChain},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "origin",
				},
				{
					Kind: mcfunction.ParameterKindChainedCommand,
					Name: "chainedCommand",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "subcommand",
					Literals: []string{"in"},
					Tags:     []string{mcfunction.TagExecuteChain},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "dimension",
					Literals: []string{"overworld", "nether", "the_end"},
					Tags:     []string{mcfunction.TagDimensionId},
				},
				{
					Kind: mcfunction.ParameterKindChainedCommand,
					Name: "chainedCommand",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "subcommand",
					Literals: []string{"positioned"},
					Tags:     []string{mcfunction.TagExecuteChain},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "position",
				},
				{
					Kind: mcfunction.ParameterKindChainedCommand,
					Name: "chainedCommand",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "subcommand",
					Literals: []string{"positioned"},
					Tags:     []string{mcfunction.TagExecuteChain},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "secondary subcommand",
					Literals: []string{"as"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "origin",
				},
				{
					Kind: mcfunction.ParameterKindChainedCommand,
					Name: "chainedCommand",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "subcommand",
					Literals: []string{"rotated"},
					Tags:     []string{mcfunction.TagExecuteChain},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "yaw",
					Tags: []string{mcfunction.TagYaw},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "pitch",
					Tags: []string{mcfunction.TagPitch},
				},
				{
					Kind: mcfunction.ParameterKindChainedCommand,
					Name: "chainedCommand",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "subcommand",
					Literals: []string{"rotated"},
					Tags:     []string{mcfunction.TagExecuteChain},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "secondary subcommand",
					Literals: []string{"as"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "origin",
				},
				{
					Kind: mcfunction.ParameterKindChainedCommand,
					Name: "chainedCommand",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "subcommand",
					Literals: []string{"facing"},
					Tags:     []string{mcfunction.TagExecuteChain},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "position",
				},
				{
					Kind: mcfunction.ParameterKindChainedCommand,
					Name: "chainedCommand",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "subcommand",
					Literals: []string{"facing"},
					Tags:     []string{mcfunction.TagExecuteChain},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "secondary subcommand",
					Literals: []string{"entity"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "origin",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "anchor",
					Literals: []string{"eyes", "feet"},
				},
				{
					Kind: mcfunction.ParameterKindChainedCommand,
					Name: "chainedCommand",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "subcommand",
					Literals: []string{"align"},
					Tags:     []string{mcfunction.TagExecuteChain},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "axes",
					Literals: []string{"xyz", "xy", "xz", "yz", "x", "y", "z"},
				},
				{
					Kind: mcfunction.ParameterKindChainedCommand,
					Name: "chainedCommand",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "subcommand",
					Literals: []string{"anchored"},
					Tags:     []string{mcfunction.TagExecuteChain},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "anchored",
					Literals: []string{"eyes", "feet"},
				},
				{
					Kind: mcfunction.ParameterKindChainedCommand,
					Name: "chainedCommand",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "subcommand",
					Literals: ifUnless,
					Tags:     []string{mcfunction.TagExecuteChain},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "secondary subcommand",
					Literals: []string{"block"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "position",
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "block",
					Tags: []string{mcfunction.TagBlockId},
				},
				{
					Kind:     mcfunction.ParameterKindChainedCommand,
					Name:     "chainedCommand",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "subcommand",
					Literals: ifUnless,
					Tags:     []string{mcfunction.TagExecuteChain},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "secondary subcommand",
					Literals: []string{"block"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "position",
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "block",
					Tags: []string{mcfunction.TagBlockId},
				},
				{
					Kind: mcfunction.ParameterKindMap,
					Name: "blockStates",
					Tags: []string{mcfunction.TagBlockState},
				},
				{
					Kind:     mcfunction.ParameterKindChainedCommand,
					Name:     "chainedCommand",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "subcommand",
					Literals: ifUnless,
					Tags:     []string{mcfunction.TagExecuteChain},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "secondary subcommand",
					Literals: []string{"blocks"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "begin",
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "end",
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "destination",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "scan mode",
					Literals: []string{"masked", "all"},
				},
				{
					Kind:     mcfunction.ParameterKindChainedCommand,
					Name:     "chainedCommand",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "subcommand",
					Literals: ifUnless,
					Tags:     []string{mcfunction.TagExecuteChain},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "secondary subcommand",
					Literals: []string{"entity"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "target",
				},
				{
					Kind:     mcfunction.ParameterKindChainedCommand,
					Name:     "chainedCommand",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "subcommand",
					Literals: ifUnless,
					Tags:     []string{mcfunction.TagExecuteChain},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "secondary subcommand",
					Literals: []string{"score"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "target",
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "objective",
					Tags: []string{mcfunction.TagScoreboardObjectiveId},
				},
				{
					Kind: mcfunction.ParameterKindLiteral,
					Name: "operation",
					Literals: []string{
						"<",
						"<=",
						">",
						">=",
					},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "source",
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "objective",
					Tags: []string{mcfunction.TagScoreboardObjectiveId},
				},
				{
					Kind:     mcfunction.ParameterKindChainedCommand,
					Name:     "chainedCommand",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "subcommand",
					Literals: ifUnless,
					Tags:     []string{mcfunction.TagExecuteChain},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "secondary subcommand",
					Literals: []string{"score"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "target",
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "objective",
					Tags: []string{mcfunction.TagScoreboardObjectiveId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "matches",
					Literals: []string{"matches"},
				},
				{
					Kind: mcfunction.ParameterKindRange,
					Name: "range",
				},
				{
					Kind:     mcfunction.ParameterKindChainedCommand,
					Name:     "chainedCommand",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "subcommand",
					Literals: []string{"run"},
					Tags:     []string{mcfunction.TagExecuteChain},
				},
				{
					Kind: mcfunction.ParameterKindCommand,
					Name: "command",
				},
			},
		},
	},
}
