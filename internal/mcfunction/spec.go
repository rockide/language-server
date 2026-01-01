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
	ParameterKindMapJSON                   // Just a map but with { } wrapping, {key=value,...}
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
	Greedy   bool     // only for KindString
	Suffix   string   // only for KindSuffixedInteger
	MapSpec  *MapSpec // only for KindMap and KindMapJSON
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
		s = "[" + s + "]"
	} else {
		s = "<" + s + ">"
	}
	if p.Kind == ParameterKindSuffixedInteger && p.Suffix != "" {
		s += p.Suffix
	}
	return s
}

type NumberRange struct {
	Min float64
	Max float64
}

type MapSpec struct {
	mapSpec map[string]*ParameterSpec
	keySpec *ParameterSpec
	spec    *ParameterSpec
}

func NewMapSpec(spec map[string]*ParameterSpec) *MapSpec {
	return &MapSpec{
		mapSpec: spec,
	}
}

func NewMapValueSpec(key *ParameterSpec, value *ParameterSpec) *MapSpec {
	return &MapSpec{
		keySpec: key,
		spec:    value,
	}
}

func (m *MapSpec) KeySpec() (*ParameterSpec, bool) {
	if m.keySpec != nil {
		return m.keySpec, true
	}
	return nil, false
}

func (m *MapSpec) ValueSpec(key string) (*ParameterSpec, bool) {
	if m.mapSpec != nil {
		s, ok := m.mapSpec[key]
		return s, ok
	}
	return m.spec, m.spec != nil
}

func (m *MapSpec) Keys() []string {
	if m.mapSpec != nil {
		keys := make([]string, len(m.mapSpec))
		i := 0
		for k := range m.mapSpec {
			keys[i] = k
			i++
		}
		return keys
	}
	return []string{}
}
