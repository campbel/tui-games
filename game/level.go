package game

type Offset struct {
	X int
	Y int
}

type Layer struct {
	height  int
	width   int
	content []string
	offset  Offset
}

func NewLayer(content []string) Layer {
	return NewLayerWithOffset(content, Offset{0, 0})
}

func NewLayerWithOffset(content []string, offset Offset) Layer {
	height := len(content)
	width := len(content[0])
	for _, line := range content {
		if len(line) > width {
			width = len(line)
		}
	}
	return Layer{
		height:  height,
		width:   width,
		content: content,
		offset:  offset,
	}
}

func (l Layer) Height() int {
	return l.height
}

func (l Layer) Width() int {
	return l.width
}

func (l Layer) Content() []string {
	return l.content
}

func (l Layer) Offset() Offset {
	return l.offset
}
