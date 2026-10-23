package core

type RockideOptions struct {
	ProjectPaths *struct {
		BehaviorPack string `json:"behaviorPack,omitzero"`
		ResourcePack string `json:"resourcePack,omitzero"`
	} `json:"projectPaths,omitempty"`
}
