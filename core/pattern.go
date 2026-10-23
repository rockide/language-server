package core

type Pattern string

func BehaviorPattern(pattern string) Pattern {
	return Pattern("b" + pattern)
}

func ResourcePattern(pattern string) Pattern {
	return Pattern("r" + pattern)
}

func (p Pattern) PackDir(workspace *Workspace) (dir string, ok bool) {
	switch p[0] {
	case 'b':
		return workspace.BehaviorPack, workspace.BehaviorPack != ""
	case 'r':
		return workspace.ResourcePack, workspace.ResourcePack != ""
	default:
		panic("invalid pattern")
	}
}

func (p Pattern) Resolve(workspace *Workspace) (path string, ok bool) {
	dir, ok := p.PackDir(workspace)
	if !ok {
		return
	}
	return dir + string(p[1:]), true
}
