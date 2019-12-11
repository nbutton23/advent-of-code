package six

import "strings"

var OrbitMap = map[string]*Body{}

type Body struct {
	ID        string
	Childeren []*Body
	Parent    *Body

	orbitCount int

	routeMap map[string]int
}

func NewBody(ID string, ParentID string) *Body {
	p, ok := OrbitMap[ParentID]
	if !ok {
		// TODO: no panic
		panic("Could not find parent")
	}
	b := &Body{
		ID:        ID,
		Childeren: make([]*Body, 0),
		Parent:    p,

		orbitCount: -1,
		routeMap:   make(map[string]int),
	}

	p.AddChild(b)

	return b
}

func (b *Body) DirectAndIndirectOrbitCount() int {
	if b.orbitCount > 0 {
		return b.orbitCount
	}
	if b.Parent == nil {
		return 0
	}
	// My orbit count is parents plus one
	b.orbitCount = b.Parent.DirectAndIndirectOrbitCount() + 1
	return b.orbitCount
}

func (b *Body) AddChild(c *Body) {
	b.Childeren = append(b.Childeren, c)
}

func (b *Body) RouteTo(ID string, from string) (route bool, steps int) {

	if b.ID == ID {
		return true, 0
	}

	// Cache
	if c, ok := b.routeMap[ID]; ok && c > 0 {
		return true, c
	}

	if ok, count := b.checkRoutFromChilderen(ID); ok {
		return ok, count
	}

	if b.Parent.ID != from {
		if ok, c := b.Parent.RouteTo(ID, b.ID); ok {
			b.routeMap[ID] = c + 1
			return true, c + 1
		}
	}

	b.routeMap[ID] = -1
	return false, 0
}

func (b *Body) checkRoutFromChilderen(ID string) (bool, int) {
	for _, c := range b.Childeren {
		if ok, count := c.RouteTo(ID, b.ID); ok {
			b.routeMap[ID] = count + 1
			return ok, count + 1
		}
	}

	return false, 0
}

func ProcessOrbitMap(mapStr string) {
	// Rest map
	OrbitMap = make(map[string]*Body)
	// Make COM
	com := &Body{
		ID:        "COM",
		Childeren: make([]*Body, 0),

		orbitCount: -1,
		routeMap:   make(map[string]int),
	}

	OrbitMap["COM"] = com

	mapSplice := strings.Split(mapStr, "\n")

	for _, orbit := range mapSplice {
		p, c := splitOrbit(orbit)

		if _, ok := OrbitMap[p]; !ok {
			pBody := &Body{
				ID:        p,
				Childeren: make([]*Body, 0),

				orbitCount: -1,
				routeMap:   make(map[string]int),
			}
			OrbitMap[p] = pBody
		}

		if cBody, ok := OrbitMap[c]; ok {
			if pBody, ok := OrbitMap[p]; ok {
				cBody.Parent = pBody
				pBody.AddChild(cBody)
			} else {
				// Im not sure if this is a good idea. . .
				mapSplice = append(mapSplice, orbit)
			}
		}

		b := NewBody(c, p)

		OrbitMap[c] = b
	}
}

func GetTotalOrbitCount(oMap map[string]*Body) int {
	tcount := 0
	for _, b := range oMap {
		tcount += b.DirectAndIndirectOrbitCount()
	}

	return tcount
}

func splitOrbit(orbit string) (p string, c string) {
	// TODO bounds check
	o := strings.Split(orbit, ")")
	return o[0], o[1]
}
