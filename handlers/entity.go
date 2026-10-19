package handlers

import (
	"slices"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/jsonc"
	"github.com/rockide/language-server/internal/sliceutil"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var Entity = &JsonHandler{
	Pattern: shared.EntityGlob,
	Entries: []JsonEntry{
		{
			Kind:       Source,
			Binding:    stores.EntityId,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:entity/description/identifier")},
			FilterDiff: true,
		},
		{
			Kind:    References,
			Binding: stores.EntityId,
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
			}, func(value string) []shared.JsonPath {
				return []shared.JsonPath{
					shared.JsonValue("minecraft:entity/components/" + value),
					shared.JsonValue("minecraft:entity/component_groups/*/" + value),
				}
			}),
		},
		{
			Kind:       Source,
			Binding:    stores.AnimationAlias,
			Path:       []shared.JsonPath{shared.JsonKey("minecraft:entity/description/animations/*")},
			FilterDiff: true,
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				res := stores.AnimationAlias.References.GetFrom(ctx.URI)
				set := mapset.NewThreadUnsafeSet[string]()
				for _, symbol := range stores.AnimationId.References.GetFrom(ctx.URI) {
					if !set.ContainsOne(symbol.Value) {
						set.Add(symbol.Value)
						res = append(res, stores.AnimationAlias.References.Get(symbol.Value)...)
					}
				}
				return res
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.AnimationAlias.Source.GetFrom(ctx.URI)
			},
		},
		{
			Kind:    References,
			Binding: stores.AnimationAlias,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:entity/description/scripts/animate/*"),
				shared.JsonKey("minecraft:entity/description/scripts/animate/*/*"),
			},
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.AnimationAlias.Source.GetFrom(ctx.URI)
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				res := stores.AnimationAlias.References.GetFrom(ctx.URI)
				set := mapset.NewThreadUnsafeSet[string]()
				for _, symbol := range stores.AnimationId.References.GetFrom(ctx.URI) {
					if !set.ContainsOne(symbol.Value) {
						set.Add(symbol.Value)
						res = append(res, stores.AnimationAlias.References.Get(symbol.Value)...)
					}
				}
				return res
			},
		},
		{
			Kind:    References,
			Binding: stores.AnimationId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:entity/description/animations/*")},
			ScopeKey: func(ctx *JsonContext) string {
				return ctx.NodeValue
			},
		},
		{
			Kind:       Source,
			Binding:    stores.EntityProperty,
			Path:       []shared.JsonPath{shared.JsonKey("minecraft:entity/description/properties/*")},
			FilterDiff: true,
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityProperty.References.GetFrom(ctx.URI)
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityProperty.Source.GetFrom(ctx.URI)
			},
		},
		{
			Kind:    References,
			Binding: stores.EntityProperty,
			Path:    []shared.JsonPath{shared.JsonKey("minecraft:entity/events/**/set_property/*")},
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityProperty.Source.GetFrom(ctx.URI)
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityProperty.References.GetFrom(ctx.URI)
			},
		},
		{
			Kind:    References,
			Binding: stores.EntityProperty,
			Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
				return shared.JsonValue(path + "/domain")
			}),
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
				if test == nil {
					return false
				}
				if value, ok := test.Value.(string); ok {
					return slices.Contains(shared.PropertyTests, value)
				}
				return false
			},
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				parent := ctx.GetParentNode()
				subject := jsonc.FindNodeAtLocation(parent, jsonc.Path{"subject"})
				if subject == nil || subject.Value == "self" {
					return stores.EntityProperty.Source.GetFrom(ctx.URI)
				}
				return stores.EntityProperty.Source.Get()
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityProperty.References.GetFrom(ctx.URI)
			},
		},
		{
			Kind:       Source,
			Binding:    stores.EntityPropertyValue,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:entity/description/properties/*/values/*")},
			FilterDiff: true,
			ScopeKey: func(ctx *JsonContext) string {
				if id, ok := ctx.GetPath()[3].(string); ok {
					return id
				}
				return defaultScope
			},
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				if id, ok := ctx.GetPath()[3].(string); ok {
					return stores.EntityPropertyValue.References.GetFrom(ctx.URI, id)
				}
				return nil
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				if id, ok := ctx.GetPath()[3].(string); ok {
					return stores.EntityPropertyValue.Source.GetFrom(ctx.URI, id)
				}
				return nil
			},
		},
		{
			Kind:    References,
			Binding: stores.EntityPropertyValue,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:entity/description/properties/*/default")},
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
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				if id, ok := ctx.GetPath()[3].(string); ok {
					return stores.EntityPropertyValue.Source.GetFrom(ctx.URI, id)
				}
				return nil
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				if id, ok := ctx.GetPath()[3].(string); ok {
					return stores.EntityPropertyValue.References.GetFrom(ctx.URI, id)
				}
				return nil
			},
		},
		{
			Kind:    References,
			Binding: stores.EntityPropertyValue,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:entity/events/**/set_property/*")},
			ScopeKey: func(ctx *JsonContext) string {
				path := ctx.GetPath()
				if id, ok := path[len(path)-1].(string); ok {
					return id
				}
				return defaultScope
			},
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				path := ctx.GetPath()
				if id, ok := path[len(path)-1].(string); ok {
					return stores.EntityPropertyValue.Source.GetFrom(ctx.URI, id)
				}
				return nil
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				path := ctx.GetPath()
				if id, ok := path[len(path)-1].(string); ok {
					return stores.EntityPropertyValue.References.GetFrom(ctx.URI, id)
				}
				return nil
			},
		},
		{
			Kind:    References,
			Binding: stores.EntityPropertyValue,
			Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
				return shared.JsonValue(path + "/value")
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
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
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
					return stores.EntityPropertyValue.Source.GetFrom(ctx.URI, id)
				}
				return stores.EntityPropertyValue.Source.Get(id)
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
					return stores.EntityPropertyValue.References.GetFrom(ctx.URI, id)
				}
				return stores.EntityPropertyValue.References.Get(id)
			},
		},
		{
			Kind:       Source,
			Binding:    stores.EntityComponentGroup,
			Path:       []shared.JsonPath{shared.JsonKey("minecraft:entity/component_groups/*")},
			FilterDiff: true,
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityComponentGroup.References.GetFrom(ctx.URI)
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityComponentGroup.Source.GetFrom(ctx.URI)
			},
		},
		{
			Kind:    References,
			Binding: stores.EntityComponentGroup,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:entity/events/**/component_groups/*")},
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityComponentGroup.Source.GetFrom(ctx.URI)
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityComponentGroup.References.GetFrom(ctx.URI)
			},
		},
		{
			Kind:       Source,
			Binding:    stores.EntityEvent,
			Path:       []shared.JsonPath{shared.JsonKey("minecraft:entity/events/*")},
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
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				root := ctx.GetRootNode()
				node := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:entity", "description", "identifier"})
				if node != nil {
					if identifier, ok := node.Value.(string); ok {
						return stores.EntityEvent.References.Get(identifier)
					}
				}
				return stores.EntityEvent.References.GetFrom(ctx.URI)
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				root := ctx.GetRootNode()
				node := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:entity", "description", "identifier"})
				if node != nil {
					if identifier, ok := node.Value.(string); ok {
						return stores.EntityEvent.Source.Get(identifier)
					}
				}
				return stores.EntityEvent.Source.GetFrom(ctx.URI)
			},
		},
		{
			Kind:    References,
			Binding: stores.EntityEvent,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:entity/components/**/event"),
				shared.JsonValue("minecraft:entity/component_groups/**/event"),
				shared.JsonValue("minecraft:entity/events/**/trigger"),
				shared.JsonValue("minecraft:entity/events/**/trigger/event"),
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
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				parent := ctx.GetParentNode()
				target := jsonc.FindNodeAtLocation(parent, jsonc.Path{"target"})
				if target == nil || target.Value == "self" {
					root := ctx.GetRootNode()
					node := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:entity", "description", "identifier"})
					if node != nil {
						if identifier, ok := node.Value.(string); ok {
							return stores.EntityEvent.Source.Get(identifier)
						}
					}
				}
				return stores.EntityEvent.Source.Get()
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				parent := ctx.GetParentNode()
				target := jsonc.FindNodeAtLocation(parent, jsonc.Path{"target"})
				if target == nil || target.Value == "self" {
					root := ctx.GetRootNode()
					node := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:entity", "description", "identifier"})
					if node != nil {
						if identifier, ok := node.Value.(string); ok {
							return stores.EntityEvent.References.Get(identifier)
						}
					}
				}
				return stores.EntityEvent.References.Get()
			},
		},
		{
			Kind:    References,
			Binding: stores.EntityEvent,
			Path: sliceutil.FlatMap([]string{
				"minecraft:block_sensor/on_break/*/on_block_broken",
				"minecraft:rideable/on_rider_enter_event",
				"minecraft:rideable/on_rider_exit_event",
			}, func(value string) []shared.JsonPath {
				return []shared.JsonPath{
					shared.JsonValue("minecraft:entity/components/" + value),
					shared.JsonValue("minecraft:entity/component_groups/*/" + value),
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
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				root := ctx.GetRootNode()
				node := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:entity", "description", "identifier"})
				if node != nil {
					if identifier, ok := node.Value.(string); ok {
						return stores.EntityEvent.Source.Get(identifier)
					}
				}
				return stores.EntityEvent.Source.GetFrom(ctx.URI)
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				root := ctx.GetRootNode()
				node := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:entity", "description", "identifier"})
				if node != nil {
					if identifier, ok := node.Value.(string); ok {
						return stores.EntityEvent.References.Get(identifier)
					}
				}
				return stores.EntityEvent.References.GetFrom(ctx.URI)
			},
		},
		{
			Kind:    References,
			Binding: stores.EntityEvent,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:entity/components/**/spawn_event")},
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
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				parent := ctx.GetParentNode()
				node := jsonc.FindNodeAtLocation(parent, jsonc.Path{"spawn_entity"})
				if node != nil {
					if identifier, ok := node.Value.(string); ok {
						return stores.EntityEvent.Source.Get(identifier)
					}
				}
				return stores.EntityEvent.Source.Get()
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				parent := ctx.GetParentNode()
				node := jsonc.FindNodeAtLocation(parent, jsonc.Path{"spawn_entity"})
				if node != nil {
					if identifier, ok := node.Value.(string); ok {
						return stores.EntityEvent.References.Get(identifier)
					}
				}
				return stores.EntityEvent.References.Get()
			},
		},
		{
			Kind:    Source,
			Binding: stores.EntityFamily,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:entity/components/minecraft:type_family/family/*"),
				shared.JsonValue("minecraft:entity/component_groups/*/minecraft:type_family/family/*"),
			},
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(stores.EntityFamily.Source.Get(), stores.EntityFamily.References.Get())
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Kind:    References,
			Binding: stores.EntityFamily,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:entity/components/minecraft:rideable/family_types/*"),
				shared.JsonValue("minecraft:entity/component_groups/*/minecraft:rideable/family_types/*"),
			},
		},
		{
			Kind:    References,
			Binding: stores.EntityFamily,
			Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
				return shared.JsonValue(path + "/value")
			}),
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
				return test != nil && (test.Value == "is_family" || test.Value == "is_vehicle_family")
			},
		},
		{
			Kind:    References,
			Binding: stores.ItemId,
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
			}, func(value string) []shared.JsonPath {
				return []shared.JsonPath{
					shared.JsonValue("minecraft:entity/components/" + value),
					shared.JsonValue("minecraft:entity/component_groups/*/" + value),
				}
			}),
			ScopeKey: func(ctx *JsonContext) string {
				return "block"
			},
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get("block")
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get("block")
			},
		},
		{
			Kind:    References,
			Binding: stores.ItemId,
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
			}, func(value string) []shared.JsonPath {
				return []shared.JsonPath{
					shared.JsonValue("minecraft:entity/components/" + value),
					shared.JsonValue("minecraft:entity/component_groups/*/" + value),
				}
			}),
		},
		{
			Kind:    References,
			Binding: stores.ItemId,
			Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
				return shared.JsonValue(path + "/value")
			}),
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
				return test != nil && test.Value == "has_equipment"
			},
		},
		{
			Kind:    References,
			Binding: stores.ItemId,
			Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
				return shared.JsonValue(path + "/value")
			}),
			ScopeKey: func(ctx *JsonContext) string {
				return "block"
			},
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
				return test != nil && test.Value == "is_block"
			},
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get("block")
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get("block")
			},
		},
		{
			Kind:    References,
			Binding: stores.ItemTag,
			Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
				return shared.JsonValue(path + "/value")
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
			}, func(value string) []shared.JsonPath {
				return []shared.JsonPath{
					shared.JsonValue("minecraft:entity/components/" + value),
					shared.JsonValue("minecraft:entity/component_groups/*/" + value),
				}
			}),
			DisableRename: true,
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.LootTablePath.Get()
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Path: sliceutil.FlatMap([]string{
				"minecraft:trade_table/table",
				"minecraft:economy_trade_table/table",
			}, func(value string) []shared.JsonPath {
				return []shared.JsonPath{
					shared.JsonValue("minecraft:entity/components/" + value),
					shared.JsonValue("minecraft:entity/component_groups/*/" + value),
				}
			}),
			DisableRename: true,
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.TradeTablePath.Get()
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Kind:    References,
			Binding: stores.BiomeTag,
			Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
				return shared.JsonValue(path + "/value")
			}),
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
				return test != nil && test.Value == "has_biome_tag"
			},
		},
		{
			Kind:    References,
			Binding: stores.BiomeId,
			Path: sliceutil.Map(shared.FilterPaths, func(path string) shared.JsonPath {
				return shared.JsonValue(path + "/value")
			}),
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
				return test != nil && test.Value == "is_biome"
			},
		},
		{
			Kind:    References,
			Binding: stores.Lang,
			Path: sliceutil.FlatMap([]string{
				"minecraft:equippable/slots/*/interact_text",
				"minecraft:interact/interactions/interact_text",
				"minecraft:interact/interactions/*/interact_text",
				"minecraft:is_dyeable/interact_text",
				"minecraft:rideable/interact_text",
			}, func(value string) []shared.JsonPath {
				return []shared.JsonPath{
					shared.JsonValue("minecraft:entity/components/" + value),
					shared.JsonValue("minecraft:entity/component_groups/*/" + value),
				}
			}),
		},
	},
	MolangLocations: slices.Concat(
		[]shared.JsonPath{
			shared.JsonValue("minecraft:entity/description/properties/*/default"),
			shared.JsonValue("minecraft:entity/description/scripts/animate/*/*"),
			shared.JsonValue("minecraft:entity/events/**/set_property/*"),
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
		}, func(value string) []shared.JsonPath {
			return []shared.JsonPath{
				shared.JsonValue("minecraft:entity/components/" + value),
				shared.JsonValue("minecraft:entity/component_groups/*/" + value),
			}
		}),
	),
	CommandEntry: JsonCommandEntry{
		Handler:   EmbedCommand,
		SlashMode: SlashNever,
		Path: []shared.JsonPath{
			shared.JsonValue("minecraft:entity/events/**/queue_command/command"),
			shared.JsonValue("minecraft:entity/events/**/queue_command/command/*"),
		},
	},
}
