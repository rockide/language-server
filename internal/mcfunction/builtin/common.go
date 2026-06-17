package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var blockStates = mcfunction.ParameterSpec{
	Name: "blockStates",
	Kind: mcfunction.ParameterKindMap,
	MapSpec: mcfunction.NewMapValueSpec(&mcfunction.ParameterSpec{
		Kind: mcfunction.ParameterKindString,
		Tags: []string{mcfunction.TagBlockState},
	}, nil),
}
