package handlers

import (
	"github.com/rockide/language-server/internal/molang"
	"github.com/rockide/language-server/internal/semtok"
)

var tokenType = map[semtok.Type]bool{
	semtok.TokNamespace:                true,
	semtok.TokEnumMember:               true,
	semtok.TokType:                     true,
	semtok.TokVariable:                 true,
	semtok.TokMethod:                   true,
	semtok.TokMacro:                    true,
	semtok.TokKeyword:                  true,
	semtok.TokComment:                  true,
	semtok.TokString:                   true,
	semtok.TokNumber:                   true,
	semtok.TokOperator:                 true,
	semtok.TokColorDefault:             true,
	semtok.TokColorGreen:               true,
	semtok.TokColorAqua:                true,
	semtok.TokColorRed:                 true,
	semtok.TokColorLightPurple:         true,
	semtok.TokColorYellow:              true,
	semtok.TokColorWhite:               true,
	semtok.TokColorMinecoinGold:        true,
	semtok.TokColorMaterialQuartz:      true,
	semtok.TokColorMaterialIron:        true,
	semtok.TokColorMaterialNetherite:   true,
	semtok.TokColorMaterialRedstone:    true,
	semtok.TokColorMaterialCopper:      true,
	semtok.TokColorMaterialGold:        true,
	semtok.TokColorMaterialEmerald:     true,
	semtok.TokColorMaterialDiamond:     true,
	semtok.TokColorMaterialLapisLazuli: true,
	semtok.TokColorMaterialAmethyst:    true,
	semtok.TokColorMaterialResin:       true,
	semtok.TokColorBlack:               true,
	semtok.TokColorDarkBlue:            true,
	semtok.TokColorDarkGreen:           true,
	semtok.TokColorDarkAqua:            true,
	semtok.TokColorDarkRed:             true,
	semtok.TokColorDarkPurple:          true,
	semtok.TokColorGold:                true,
	semtok.TokColorGray:                true,
	semtok.TokColorDarkGray:            true,
	semtok.TokColorBlue:                true,
}

var tokenModifier = map[semtok.Modifier]bool{}

var molangTokenMap = map[molang.TokenKind]semtok.Type{
	molang.KindNumber:   semtok.TokNumber,
	molang.KindString:   semtok.TokString,
	molang.KindMacro:    semtok.TokMacro,
	molang.KindMethod:   semtok.TokMethod,
	molang.KindPrefix:   semtok.TokType,
	molang.KindKeyword:  semtok.TokKeyword,
	molang.KindOperator: semtok.TokOperator,
	molang.KindParen:    semtok.TokEnumMember,
	molang.KindComma:    semtok.TokOperator,
}
