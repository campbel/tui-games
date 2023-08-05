package game

func Render(relx int, window Window, layers ...Layer) []string {
	result := make([]string, window.Height)
	for _, layer := range layers {
		if window.Height < layer.Height() {
			return []string{"Window is too small"}
		}
		content := layer.Content()

		if relx < 0 {
			relx = 0
		} else if relx > layer.Width()-window.Width {
			relx = layer.Width() - window.Width
		}

		for y := 0; y < len(result); y++ {
			resultIndex := len(result) - y - 1
			contentIndex := len(content) - y - 1
			if y >= len(content) {
				result[resultIndex] = repeat(" ", window.Width)
				continue
			}
			result[resultIndex] = content[contentIndex][relx : relx+window.Width]
		}

	}

	return result
}

func repeat(char string, times int) string {
	result := ""

	for i := 0; i < times; i++ {
		result += char
	}

	return result
}
