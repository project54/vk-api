package vkapi

type Keyboard struct {
	OneTime bool       `json:"one_time"`
	Buttons [][]Button `json:"buttons"`
}

type Button struct {
	Action Action      `json:"action"`
	Color  ColorButton `json:"color"`
}

type Action struct {
	Type  string `json:"type"`
	Label string `json:"label"`
}

type ColorButton string

const (
	Negative ColorButton = "negative"
	Positive ColorButton = "positive"
	Default  ColorButton = "default"
	Primary  ColorButton = "primary"
)

func CreateButtons(buttons ...[]Button) [][]Button {
	length := len(buttons)
	if length > 8 {
		panic("Amount rows are more 8.")
	}

	return buttons
}

func CreateButtonsInRow(textAndColorButtons map[string]ColorButton) []Button {
	length := len(textAndColorButtons)
	if length > 4 {
		panic("Amount buttons in the row are more 4.")
	}

	var buttons []Button

	for text, color := range textAndColorButtons {
		button := new(Button)
		button.Action.Type = "text"
		button.Action.Label = text
		button.Color = color

		buttons = append(buttons, *button)
	}

	return buttons
}
