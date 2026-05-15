package mcfunction

var EventAliasSpec = &Spec{
	Description: "Alias for 'event entity' command. Can be only used in behavior packs for animation and animation controller",
	Overloads: []SpecOverload{
		{
			Parameters: []ParameterSpec{
				{
					Kind:     ParameterKindSelector,
					Name:     "self",
					Literals: []string{"@s"},
				},
				{
					Kind: ParameterKindString,
					Name: "eventName",
					Tags: []string{TagEntityEvent},
				},
			},
		},
	},
}
