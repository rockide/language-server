package mcfunction

type Spec struct {
	Name        string
	Aliases     []string
	Description string
	Overloads   []SpecOverload
}

type SpecOverload struct {
	Parameters []ParameterSpec
}

type ParameterKind int8

const (
	ParameterKindUnknown     ParameterKind = iota
	ParameterKindLiteral                   // fixed set of string literals
	ParameterKindString                    // arbitrary string
	ParameterKindNumber                    // allows float
	ParameterKindInteger                   // strictly integer
	ParameterKindBoolean                   // true or false
	ParameterKindSelector                  // @p, @a, @r, @s, @e, @n, @initiator
	ParameterKindSelectorArg               // selector arguments
	ParameterKindMap                       // [key=value,...]
	ParameterKindJSON                      // JSON object or array
	ParameterKindVector2                   // x y or rotX rotY
	ParameterKindVector3                   // x y z
	// Misc
	ParameterKindRange           // start..end
	ParameterKindSuffixedInteger // int with suffix like 10L, 10D, 10S, 10T.
	ParameterKindWildcardInteger // * or number
	ParameterKindChainedCommand  // chained command
	ParameterKindCommand         // commands
	ParameterKindRawMessage
	ParameterKindItemNbt        // item NBT data
	ParameterKindRelativeNumber // ~number or ^number (part of vector)
	ParameterKindMapPair        // key=value (only as part of map)
)

type Tag string

const (
	TagAimAssistId           = "aim_assist_id"
	TagBiomeId               = "biome_id"
	TagBlockId               = "block_id"
	TagBlockState            = "block_state"
	TagCameraId              = "camera_id"
	TagClientAnimationId     = "client_animation_id"
	TagDialogueId            = "dialogue_id"
	TagDimensionId           = "dimension_id"
	TagEntityEvent           = "entity_event"
	TagEntityId              = "entity_id"
	TagExecuteChain          = "execute_chain"
	TagTypeFamilyId          = "type_family_id"
	TagFeatureId             = "feature_id"
	TagFeatureRuleId         = "feature_rule_id"
	TagFogId                 = "fog_id"
	TagFunctionFile          = "function_file"
	TagItemId                = "item_id"
	TagJigsawId              = "jigsaw_id"
	TagJigsawTemplatePoolId  = "jigsaw_template_pool_id"
	TagLootTableFile         = "loot_table_file"
	TagMolang                = "molang"
	TagMusicId               = "music_id"
	TagParticleId            = "particle_id"
	TagPitch                 = "pitch"
	TagProvidedFogId         = "provided_fog_id"
	TagRecipeId              = "recipe_id"
	TagScoreboardObjectiveId = "scoreboard_objective_id"
	TagScriptEventId         = "script_event_id"
	TagSoundId               = "sound_definition_id"
	TagStructureFile         = "structure_file"
	TagTagId                 = "tag"
	TagTickingAreaId         = "ticking_area_id"
	TagYaw                   = "yaw"
)

type ParameterSpec struct {
	Kind     ParameterKind
	Name     string
	Optional bool
	Literals []string     // only for ParameterKindLiteral
	Range    *NumberRange // For number, integer, and range
	Tags     []string
	Greedy   bool   // only for KindString
	Suffix   string // only for KindSuffixedInteger
}

func (p ParameterSpec) ToString() string {
	s := p.Name
	if len(p.Literals) == 1 && p.Kind == ParameterKindLiteral {
		if !p.Optional {
			return p.Literals[0]
		}
		s = p.Literals[0]
	}
	if p.Optional {
		return "[" + s + "]"
	}
	return "<" + s + ">"
}

type NumberRange struct {
	Min float64
	Max float64
}
