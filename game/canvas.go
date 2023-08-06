package game

type Canvas struct {
	Content []string
}

func NewBlankCanvas(width, height int) *Canvas {
	content := make([]string, height)
	for i := range content {
		content[i] = repeat(" ", width)
	}
	return &Canvas{content}
}

func (canvas *Canvas) Draw(offsetX, offsetY int, content []string) {
	for y := 0; y < len(content); y++ {
		canvas.Content[y+offsetY] = canvas.Content[y+offsetY][:offsetX] + content[y] + canvas.Content[y+offsetY][offsetX+len(content[y]):]
	}
}

func (canvas *Canvas) Render(OffsetX, offsetY, width, height int) []string {
	if offsetY+height > len(canvas.Content) {
		offsetY = len(canvas.Content) - height
	}
	if offsetY < 0 {
		offsetY = 0
	}

	if OffsetX+width > len(canvas.Content[0]) {
		OffsetX = len(canvas.Content[0]) - width
	}
	if OffsetX < 0 {
		OffsetX = 0
	}

	result := make([]string, height)
	for y := 0; y < height; y++ {
		if y >= len(canvas.Content) {
			result[y] = repeat(" ", width)
			continue
		}
		end := min(OffsetX+width, len(canvas.Content[y+offsetY])-1)
		result[y] = canvas.Content[y+offsetY][OffsetX:end] + repeat(" ", width-end+OffsetX)
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func repeat(char string, times int) string {
	result := ""
	for i := 0; i < times; i++ {
		result += char
	}
	return result
}
