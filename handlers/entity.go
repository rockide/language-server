package handlers

import (
	"slices"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/jsonc"
	"github.com/rockide/language-server/internal/sliceutil"
)

func Entity(workspace *core.Workspace) *JsonHandler {
	propertyTests := []string{"bool_property", "enum_property", "float_property", "int_property"}

	filterPaths := sliceutil.FlatMap([]string{
		"**/filters",
		"minecraft:ageable/interact_filters",
		"minecraft:anger_level/nuisance_filter",
		"minecraft:angry/broadcast_filters",
		"minecraft:area_attack/entity_filter",
		"minecraft:behavior.knockback_roar/damage_filters",
		"minecraft:behavior.knockback_roar/knockback_filters",
		"minecraft:behavior.move_to_block/target_block_filters",
		"minecraft:behavior.nap/can_nap_filters",
		"minecraft:behavior.nap/wake_mob_exceptions",
		"minecraft:behavior.stalk_and_pounce_on_target/stuck_blocks",
		"minecraft:block_sensor/sources",
		"minecraft:breedable/love_filters",
		"minecraft:celebrate_hunt/celebration_targets",
		"minecraft:conditional_bandwidth_optimization/conditional_values",
		"minecraft:entity_sensor/event_filters",
		"minecraft:entity_sensor/subsensors/*/event_filters",
		"minecraft:mob_effect/entity_filter",
		"minecraft:trail/spawn_filter",
	}, func(path string) []string {
		return []string{
			"minecraft:entity/components/" + path + "/**",
			"minecraft:entity/component_groups/*/" + path + "/**",
			"minecraft:entity/events/*/" + path + "/**",
		}
	})

	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.BehaviorPattern("/entities/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:       Definition,
				Binding:    workspace.EntityId,
				Path:       []core.JsonPath{core.JsonValue("minecraft:entity/description/identifier")},
				FilterDiff: true,
			},
			{
				Kind:    References,
				Binding: workspace.EntityId,
				Path: sliceutil.FlatMap([]string{
					"minecraft:addrider/entity_type",
					"minecraft:addrider/riders/*/entity_type",
					"minecraft:behavior.follow_mob/preferred_actor_type",
					"minecraft:behavior.mingle/mingle_partner_type",
					"minecraft:breedable/breeds_with/baby_type",
					"minecraft:breedable/breeds_with/*/baby_type",
					"minecraft:breedable/breeds_with/mate_type",
					"minecraft:breedable/breeds_with/*/mate_type",
					"minecraft:spawn_entity/entities/*/spawn_entity",
					"minecraft:transformation/into",
				}, func(value string) []core.JsonPath {
					return []core.JsonPath{
						core.JsonValue("minecraft:entity/components/" + value),
						core.JsonValue("minecraft:entity/component_groups/*/" + value),
					}
				}),
			},
			{
				Kind:       Definition,
				Binding:    workspace.AnimationAlias,
				Path:       []core.JsonPath{core.JsonKey("minecraft:entity/description/animations/*")},
				FilterDiff: true,
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.AnimationAlias.Definition.GetFrom(ctx.URI)
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					res := workspace.AnimationAlias.References.GetFrom(ctx.URI)
					set := mapset.NewThreadUnsafeSet[string]()
					for _, symbol := range workspace.AnimationId.References.GetFrom(ctx.URI) {
						if !set.ContainsOne(symbol.Value) {
							set.Add(symbol.Value)
							res = append(res, workspace.AnimationAlias.References.Get(symbol.Value)...)
						}
					}
					return res
				},
			},
			{
				Kind:    References,
				Binding: workspace.AnimationAlias,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:entity/description/scripts/animate/*"),
					core.JsonKey("minecraft:entity/description/scripts/animate/*/*"),
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.AnimationAlias.Definition.GetFrom(ctx.URI)
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					res := workspace.AnimationAlias.References.GetFrom(ctx.URI)
					set := mapset.NewThreadUnsafeSet[string]()
					for _, symbol := range workspace.AnimationId.References.GetFrom(ctx.URI) {
						if !set.ContainsOne(symbol.Value) {
							set.Add(symbol.Value)
							res = append(res, workspace.AnimationAlias.References.Get(symbol.Value)...)
						}
					}
					return res
				},
			},
			{
				Kind:    References,
				Binding: workspace.AnimationId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:entity/description/animations/*")},
				ScopeKey: func(ctx *JsonContext) string {
					return ctx.NodeValue
				},
			},
			{
				Kind:       Definition,
				Binding:    workspace.EntityProperty,
				Path:       []core.JsonPath{core.JsonKey("minecraft:entity/description/properties/*")},
				FilterDiff: true,
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.EntityProperty.Definition.GetFrom(ctx.URI)
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.EntityProperty.References.GetFrom(ctx.URI)
				},
			},
			{
				Kind:    References,
				Binding: workspace.EntityProperty,
				Path:    []core.JsonPath{core.JsonKey("minecraft:entity/events/**/set_property/*")},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.EntityProperty.Definition.GetFrom(ctx.URI)
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.EntityProperty.References.GetFrom(ctx.URI)
				},
			},
			{
				Kind:    References,
				Binding: workspace.EntityProperty,
				Path: sliceutil.Map(filterPaths, func(path string) core.JsonPath {
					return core.JsonValue(path + "/domain")
				}),
				Matcher: func(ctx *JsonContext) bool {
					parent := ctx.GetParentNode()
					test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
					if test == nil {
						return false
					}
					if value, ok := test.Value.(string); ok {
						return slices.Contains(propertyTests, value)
					}
					return false
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					parent := ctx.GetParentNode()
					subject := jsonc.FindNodeAtLocation(parent, jsonc.Path{"subject"})
					if subject == nil || subject.Value == "self" {
						return workspace.EntityProperty.Definition.GetFrom(ctx.URI)
					}
					return workspace.EntityProperty.Definition.Get()
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.EntityProperty.References.GetFrom(ctx.URI)
				},
			},
			{
				Kind:       Definition,
				Binding:    workspace.EntityPropertyValue,
				Path:       []core.JsonPath{core.JsonValue("minecraft:entity/description/properties/*/values/*")},
				FilterDiff: true,
				ScopeKey: func(ctx *JsonContext) string {
					if id, ok := ctx.GetPath()[3].(string); ok {
						return id
					}
					return defaultScope
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					if id, ok := ctx.GetPath()[3].(string); ok {
						return workspace.EntityPropertyValue.Definition.GetFrom(ctx.URI, id)
					}
					return nil
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					if id, ok := ctx.GetPath()[3].(string); ok {
						return workspace.EntityPropertyValue.References.GetFrom(ctx.URI, id)
					}
					return nil
				},
			},
			{
				Kind:    References,
				Binding: workspace.EntityPropertyValue,
				Path:    []core.JsonPath{core.JsonValue("minecraft:entity/description/properties/*/default")},
				Matcher: func(ctx *JsonContext) bool {
					parent := ctx.GetParentNode()
					propertyType := jsonc.FindNodeAtLocation(parent, jsonc.Path{"type"})
					return propertyType != nil && propertyType.Value == "enum"
				},
				ScopeKey: func(ctx *JsonContext) string {
					if id, ok := ctx.GetPath()[3].(string); ok {
						return id
					}
					return defaultScope
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					if id, ok := ctx.GetPath()[3].(string); ok {
						return workspace.EntityPropertyValue.Definition.GetFrom(ctx.URI, id)
					}
					return nil
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					if id, ok := ctx.GetPath()[3].(string); ok {
						return workspace.EntityPropertyValue.References.GetFrom(ctx.URI, id)
					}
					return nil
				},
			},
			{
				Kind:    References,
				Binding: workspace.EntityPropertyValue,
				Path:    []core.JsonPath{core.JsonValue("minecraft:entity/events/**/set_property/*")},
				ScopeKey: func(ctx *JsonContext) string {
					path := ctx.GetPath()
					if id, ok := path[len(path)-1].(string); ok {
						return id
					}
					return defaultScope
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					path := ctx.GetPath()
					if id, ok := path[len(path)-1].(string); ok {
						return workspace.EntityPropertyValue.Definition.GetFrom(ctx.URI, id)
					}
					return nil
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					path := ctx.GetPath()
					if id, ok := path[len(path)-1].(string); ok {
						return workspace.EntityPropertyValue.References.GetFrom(ctx.URI, id)
					}
					return nil
				},
			},
			{
				Kind:    References,
				Binding: workspace.EntityPropertyValue,
				Path: sliceutil.Map(filterPaths, func(path string) core.JsonPath {
					return core.JsonValue(path + "/value")
				}),
				Matcher: func(ctx *JsonContext) bool {
					parent := ctx.GetParentNode()
					test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
					return test != nil && test.Value == "enum_property"
				},
				ScopeKey: func(ctx *JsonContext) string {
					parent := ctx.GetParentNode()
					domain := jsonc.FindNodeAtLocation(parent, jsonc.Path{"domain"})
					if domain == nil {
						return defaultScope
					}
					if id, ok := domain.Value.(string); ok {
						return id
					}
					return defaultScope
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					parent := ctx.GetParentNode()
					domain := jsonc.FindNodeAtLocation(parent, jsonc.Path{"domain"})
					if domain == nil {
						return nil
					}
					id, ok := domain.Value.(string)
					if !ok {
						return nil
					}
					subject := jsonc.FindNodeAtLocation(parent, jsonc.Path{"subject"})
					if subject == nil || subject.Value == "self" {
						return workspace.EntityPropertyValue.Definition.GetFrom(ctx.URI, id)
					}
					return workspace.EntityPropertyValue.Definition.Get(id)
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					parent := ctx.GetParentNode()
					domain := jsonc.FindNodeAtLocation(parent, jsonc.Path{"domain"})
					if domain == nil {
						return nil
					}
					id, ok := domain.Value.(string)
					if !ok {
						return nil
					}
					subject := jsonc.FindNodeAtLocation(parent, jsonc.Path{"subject"})
					if subject == nil || subject.Value == "self" {
						return workspace.EntityPropertyValue.References.GetFrom(ctx.URI, id)
					}
					return workspace.EntityPropertyValue.References.Get(id)
				},
			},
			{
				Kind:       Definition,
				Binding:    workspace.EntityComponentGroup,
				Path:       []core.JsonPath{core.JsonKey("minecraft:entity/component_groups/*")},
				FilterDiff: true,
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.EntityComponentGroup.Definition.GetFrom(ctx.URI)
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.EntityComponentGroup.References.GetFrom(ctx.URI)
				},
			},
			{
				Kind:    References,
				Binding: workspace.EntityComponentGroup,
				Path:    []core.JsonPath{core.JsonValue("minecraft:entity/events/**/component_groups/*")},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.EntityComponentGroup.Definition.GetFrom(ctx.URI)
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.EntityComponentGroup.References.GetFrom(ctx.URI)
				},
			},
			{
				Kind:       Definition,
				Binding:    workspace.EntityEvent,
				Path:       []core.JsonPath{core.JsonKey("minecraft:entity/events/*")},
				FilterDiff: true,
				ScopeKey: func(ctx *JsonContext) string {
					root := ctx.GetRootNode()
					node := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:entity", "description", "identifier"})
					if node != nil {
						if identifier, ok := node.Value.(string); ok {
							return identifier
						}
					}
					return defaultScope
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					root := ctx.GetRootNode()
					node := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:entity", "description", "identifier"})
					if node != nil {
						if identifier, ok := node.Value.(string); ok {
							return workspace.EntityEvent.Definition.Get(identifier)
						}
					}
					return workspace.EntityEvent.Definition.GetFrom(ctx.URI)
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					root := ctx.GetRootNode()
					node := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:entity", "description", "identifier"})
					if node != nil {
						if identifier, ok := node.Value.(string); ok {
							return workspace.EntityEvent.References.Get(identifier)
						}
					}
					return workspace.EntityEvent.References.GetFrom(ctx.URI)
				},
			},
			{
				Kind:    References,
				Binding: workspace.EntityEvent,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:entity/components/**/event"),
					core.JsonValue("minecraft:entity/component_groups/**/event"),
					core.JsonValue("minecraft:entity/events/**/trigger"),
					core.JsonValue("minecraft:entity/events/**/trigger/event"),
				},
				ScopeKey: func(ctx *JsonContext) string {
					parent := ctx.GetParentNode()
					target := jsonc.FindNodeAtLocation(parent, jsonc.Path{"target"})
					if target == nil || target.Value == "self" {
						root := ctx.GetRootNode()
						node := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:entity", "description", "identifier"})
						if node != nil {
							if identifier, ok := node.Value.(string); ok {
								return identifier
							}
						}
					}
					return defaultScope
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					parent := ctx.GetParentNode()
					target := jsonc.FindNodeAtLocation(parent, jsonc.Path{"target"})
					if target == nil || target.Value == "self" {
						root := ctx.GetRootNode()
						node := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:entity", "description", "identifier"})
						if node != nil {
							if identifier, ok := node.Value.(string); ok {
								return workspace.EntityEvent.Definition.Get(identifier)
							}
						}
					}
					return workspace.EntityEvent.Definition.Get()
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					parent := ctx.GetParentNode()
					target := jsonc.FindNodeAtLocation(parent, jsonc.Path{"target"})
					if target == nil || target.Value == "self" {
						root := ctx.GetRootNode()
						node := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:entity", "description", "identifier"})
						if node != nil {
							if identifier, ok := node.Value.(string); ok {
								return workspace.EntityEvent.References.Get(identifier)
							}
						}
					}
					return workspace.EntityEvent.References.Get()
				},
			},
			{
				Kind:    References,
				Binding: workspace.EntityEvent,
				Path: sliceutil.FlatMap([]string{
					"minecraft:block_sensor/on_break/*/on_block_broken",
					"minecraft:rideable/on_rider_enter_event",
					"minecraft:rideable/on_rider_exit_event",
				}, func(value string) []core.JsonPath {
					return []core.JsonPath{
						core.JsonValue("minecraft:entity/components/" + value),
						core.JsonValue("minecraft:entity/component_groups/*/" + value),
					}
				}),
				ScopeKey: func(ctx *JsonContext) string {
					root := ctx.GetRootNode()
					node := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:entity", "description", "identifier"})
					if node != nil {
						if identifier, ok := node.Value.(string); ok {
							return identifier
						}
					}
					return defaultScope
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					root := ctx.GetRootNode()
					node := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:entity", "description", "identifier"})
					if node != nil {
						if identifier, ok := node.Value.(string); ok {
							return workspace.EntityEvent.Definition.Get(identifier)
						}
					}
					return workspace.EntityEvent.Definition.GetFrom(ctx.URI)
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					root := ctx.GetRootNode()
					node := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:entity", "description", "identifier"})
					if node != nil {
						if identifier, ok := node.Value.(string); ok {
							return workspace.EntityEvent.References.Get(identifier)
						}
					}
					return workspace.EntityEvent.References.GetFrom(ctx.URI)
				},
			},
			{
				Kind:    References,
				Binding: workspace.EntityEvent,
				Path:    []core.JsonPath{core.JsonValue("minecraft:entity/components/**/spawn_event")},
				ScopeKey: func(ctx *JsonContext) string {
					parent := ctx.GetParentNode()
					node := jsonc.FindNodeAtLocation(parent, jsonc.Path{"spawn_entity"})
					if node != nil {
						if identifier, ok := node.Value.(string); ok {
							return identifier
						}
					}
					return defaultScope
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					parent := ctx.GetParentNode()
					node := jsonc.FindNodeAtLocation(parent, jsonc.Path{"spawn_entity"})
					if node != nil {
						if identifier, ok := node.Value.(string); ok {
							return workspace.EntityEvent.Definition.Get(identifier)
						}
					}
					return workspace.EntityEvent.Definition.Get()
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					parent := ctx.GetParentNode()
					node := jsonc.FindNodeAtLocation(parent, jsonc.Path{"spawn_entity"})
					if node != nil {
						if identifier, ok := node.Value.(string); ok {
							return workspace.EntityEvent.References.Get(identifier)
						}
					}
					return workspace.EntityEvent.References.Get()
				},
			},
			{
				Kind:    Definition,
				Binding: workspace.EntityFamily,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:entity/components/minecraft:type_family/family/*"),
					core.JsonValue("minecraft:entity/component_groups/*/minecraft:type_family/family/*"),
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return nil
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return slices.Concat(workspace.EntityFamily.Definition.Get(), workspace.EntityFamily.References.Get())
				},
			},
			{
				Kind:    References,
				Binding: workspace.EntityFamily,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:entity/components/minecraft:rideable/family_types/*"),
					core.JsonValue("minecraft:entity/component_groups/*/minecraft:rideable/family_types/*"),
				},
			},
			{
				Kind:    References,
				Binding: workspace.EntityFamily,
				Path: sliceutil.Map(filterPaths, func(path string) core.JsonPath {
					return core.JsonValue(path + "/value")
				}),
				Matcher: func(ctx *JsonContext) bool {
					parent := ctx.GetParentNode()
					test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
					return test != nil && (test.Value == "is_family" || test.Value == "is_vehicle_family")
				},
			},
			{
				Kind:    References,
				Binding: workspace.ItemId,
				Path: sliceutil.FlatMap([]string{
					"minecraft:behavior.avoid_block/target_blocks/*",
					"minecraft:behavior.eat_block/eat_and_replace_block_pairs/*/eat_block",
					"minecraft:behavior.eat_block/eat_and_replace_block_pairs/*/replace_block",
					"minecraft:behavior.float_tempt/items/*",
					"minecraft:behavior.jump_to_block/forbidden_blocks/*",
					"minecraft:behavior.jump_to_block/preferred_blocks/*",
					"minecraft:behavior.lay_egg/egg_type",
					"minecraft:behavior.lay_egg/target_blocks/*",
					"minecraft:behavior.move_to_block/target_blocks/*",
					"minecraft:behavior.raid_garden/blocks/*",
					"minecraft:behavior.random_search_and_dig/target_blocks/*",
					"minecraft:block_sensor/on_break/*/block_list/*",
					"minecraft:break_blocks/breakable_blocks/*",
					"minecraft:breathable/breathe_blocks/*",
					"minecraft:breathable/non_breathe_blocks/*",
					"minecraft:breedable/environment_requirements/blocks",
					"minecraft:breedable/environment_requirements/*/blocks",
					"minecraft:buoyant/liquid_blocks/*",
					"minecraft:home/home_block_list/*",
					"minecraft:inside_block_notifier/block_list/*/block/name",
					"minecraft:navigation.climb/blocks_to_avoid/*",
					"minecraft:navigation.climb/blocks_to_avoid/*/name",
					"minecraft:navigation.float/blocks_to_avoid/*",
					"minecraft:navigation.float/blocks_to_avoid/*/name",
					"minecraft:navigation.fly/blocks_to_avoid/*",
					"minecraft:navigation.fly/blocks_to_avoid/*/name",
					"minecraft:navigation.generic/blocks_to_avoid/*",
					"minecraft:navigation.generic/blocks_to_avoid/*/name",
					"minecraft:navigation.hover/blocks_to_avoid/*",
					"minecraft:navigation.hover/blocks_to_avoid/*/name",
					"minecraft:navigation.swim/blocks_to_avoid/*",
					"minecraft:navigation.swim/blocks_to_avoid/*/name",
					"minecraft:navigation.walk/blocks_to_avoid/*",
					"minecraft:navigation.walk/blocks_to_avoid/*/name",
					"minecraft:preferred_path/preferred_path_blocks/blocks/*",
					"minecraft:preferred_path/preferred_path_blocks/blocks/*/name",
					"minecraft:trail/block_type",
					"minecraft:transformation/delay/block_types/*",
				}, func(value string) []core.JsonPath {
					return []core.JsonPath{
						core.JsonValue("minecraft:entity/components/" + value),
						core.JsonValue("minecraft:entity/component_groups/*/" + value),
					}
				}),
				ScopeKey: func(ctx *JsonContext) string {
					return "block"
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.ItemId.Definition.Get("block")
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.ItemId.References.Get("block")
				},
			},
			{
				Kind:    References,
				Binding: workspace.ItemId,
				Path: sliceutil.FlatMap([]string{
					"minecraft:ageable/drop_items/*",
					"minecraft:ageable/feed_items",
					"minecraft:ageable/feed_items/*",
					"minecraft:ageable/feed_items/*/item",
					"minecraft:ageable/feed_items/*/result_item",
					"minecraft:behavior.beg/items/*",
					"minecraft:behavior.charge_held_item/items/*",
					"minecraft:behavior.pickup_items/excluded_items/*",
					"minecraft:behavior.snacking/items/*",
					"minecraft:behavior.tempt/items/*",
					"minecraft:boostable/boost_items/*/item",
					"minecraft:boostable/boost_items/*/replace_item",
					"minecraft:breedable/breed_items",
					"minecraft:breedable/breed_items/*",
					"minecraft:breedable/breed_items/*/item",
					"minecraft:breedable/breed_items/*/result_item",
					"minecraft:bribeable/bribe_items/*",
					"minecraft:equippable/slots/*/accepted_items/*",
					"minecraft:equippable/slots/*/item",
					"minecraft:giveable/triggers/items/*",
					"minecraft:giveable/triggers/*/items/*",
					"minecraft:healable/items/*/item",
					"minecraft:healable/items/*/result_item",
					"minecraft:interact/interactions/transform_to_item",
					"minecraft:interact/interactions/*/transform_to_item",
					"minecraft:item_controllable/control_items",
					"minecraft:item_controllable/control_items/*",
					"minecraft:shareables/items/*/craft_into",
					"minecraft:shareables/items/*/item",
					"minecraft:spawn_entity/entities/spawn_item",
					"minecraft:spawn_entity/entities/*/spawn_item",
					"minecraft:tameable/tame_items",
					"minecraft:tameable/tame_items/*",
					"minecraft:tameable/tame_items/*/item",
					"minecraft:tameable/tame_items/*/result_item",
					"minecraft:tamemount/auto_reject_items/*/item",
					"minecraft:tamemount/feed_items/*/item",
					"minecraft:trusting/trust_items/*",
				}, func(value string) []core.JsonPath {
					return []core.JsonPath{
						core.JsonValue("minecraft:entity/components/" + value),
						core.JsonValue("minecraft:entity/component_groups/*/" + value),
					}
				}),
			},
			{
				Kind:    References,
				Binding: workspace.ItemId,
				Path: sliceutil.Map(filterPaths, func(path string) core.JsonPath {
					return core.JsonValue(path + "/value")
				}),
				Matcher: func(ctx *JsonContext) bool {
					parent := ctx.GetParentNode()
					test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
					return test != nil && test.Value == "has_equipment"
				},
			},
			{
				Kind:    References,
				Binding: workspace.ItemId,
				Path: sliceutil.Map(filterPaths, func(path string) core.JsonPath {
					return core.JsonValue(path + "/value")
				}),
				ScopeKey: func(ctx *JsonContext) string {
					return "block"
				},
				Matcher: func(ctx *JsonContext) bool {
					parent := ctx.GetParentNode()
					test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
					return test != nil && test.Value == "is_block"
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.ItemId.Definition.Get("block")
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.ItemId.References.Get("block")
				},
			},
			{
				Kind:    References,
				Binding: workspace.ItemTag,
				Path: sliceutil.Map(filterPaths, func(path string) core.JsonPath {
					return core.JsonValue(path + "/value")
				}),
				Matcher: func(ctx *JsonContext) bool {
					parent := ctx.GetParentNode()
					test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
					return test != nil && test.Value == "has_equipment_tag"
				},
			},
			{
				Path: sliceutil.FlatMap([]string{
					"minecraft:loot/table",
					"minecraft:behavior.sneeze/loot_table",
					"minecraft:barter/barter_table",
					"minecraft:interact/interactions/add_items/table",
					"minecraft:interact/interactions/*/add_items/table",
					"minecraft:interact/interactions/spawn_items/table",
					"minecraft:interact/interactions/*/spawn_items/table",
				}, func(value string) []core.JsonPath {
					return []core.JsonPath{
						core.JsonValue("minecraft:entity/components/" + value),
						core.JsonValue("minecraft:entity/component_groups/*/" + value),
					}
				}),
				DisableRename: true,
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.LootTablePath.Get()
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return nil
				},
			},
			{
				Path: sliceutil.FlatMap([]string{
					"minecraft:trade_table/table",
					"minecraft:economy_trade_table/table",
				}, func(value string) []core.JsonPath {
					return []core.JsonPath{
						core.JsonValue("minecraft:entity/components/" + value),
						core.JsonValue("minecraft:entity/component_groups/*/" + value),
					}
				}),
				DisableRename: true,
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.TradeTablePath.Get()
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return nil
				},
			},
			{
				Kind:    References,
				Binding: workspace.BiomeTag,
				Path: sliceutil.Map(filterPaths, func(path string) core.JsonPath {
					return core.JsonValue(path + "/value")
				}),
				Matcher: func(ctx *JsonContext) bool {
					parent := ctx.GetParentNode()
					test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
					return test != nil && test.Value == "has_biome_tag"
				},
			},
			{
				Kind:    References,
				Binding: workspace.BiomeId,
				Path: sliceutil.Map(filterPaths, func(path string) core.JsonPath {
					return core.JsonValue(path + "/value")
				}),
				Matcher: func(ctx *JsonContext) bool {
					parent := ctx.GetParentNode()
					test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
					return test != nil && test.Value == "is_biome"
				},
			},
			{
				Kind:    References,
				Binding: workspace.Lang,
				Path: sliceutil.FlatMap([]string{
					"minecraft:equippable/slots/*/interact_text",
					"minecraft:interact/interactions/interact_text",
					"minecraft:interact/interactions/*/interact_text",
					"minecraft:is_dyeable/interact_text",
					"minecraft:rideable/interact_text",
				}, func(value string) []core.JsonPath {
					return []core.JsonPath{
						core.JsonValue("minecraft:entity/components/" + value),
						core.JsonValue("minecraft:entity/component_groups/*/" + value),
					}
				}),
			},
		},
		MolangLocations: slices.Concat(
			[]core.JsonPath{
				core.JsonValue("minecraft:entity/description/properties/*/default"),
				core.JsonValue("minecraft:entity/description/scripts/animate/*/*"),
				core.JsonValue("minecraft:entity/events/**/set_property/*"),
			},
			sliceutil.FlatMap([]string{
				"minecraft:ambient_sound_interval/event_names/*/condition",
				"minecraft:anger_level/on_increase_sounds/*/condition",
				"minecraft:behavior.eat_block/success_chance",
				"minecraft:experience_reward/on_bred",
				"minecraft:experience_reward/on_death",
				"minecraft:heartbeat/interval",
				"minecraft:navigation.climb/blocks_to_avoid/*/tags",
				"minecraft:navigation.float/blocks_to_avoid/*/tags",
				"minecraft:navigation.fly/blocks_to_avoid/*/tags",
				"minecraft:navigation.generic/blocks_to_avoid/*/tags",
				"minecraft:navigation.hover/blocks_to_avoid/*/tags",
				"minecraft:navigation.swim/blocks_to_avoid/*/tags",
				"minecraft:navigation.walk/blocks_to_avoid/*/tags",
				"minecraft:projectile/on_hit/impact_damage/filter",
				"minecraft:rideable/seats/*/rotate_rider_by",
				"minecraft:rideable/seats/rotate_rider_by",
			}, func(value string) []core.JsonPath {
				return []core.JsonPath{
					core.JsonValue("minecraft:entity/components/" + value),
					core.JsonValue("minecraft:entity/component_groups/*/" + value),
				}
			}),
		),
		CommandEntry: JsonCommandEntry{
			Handler:   EmbedCommand(workspace),
			SlashMode: SlashNever,
			Path: []core.JsonPath{
				core.JsonValue("minecraft:entity/events/**/queue_command/command"),
				core.JsonValue("minecraft:entity/events/**/queue_command/command/*"),
			},
		},
	}
}
