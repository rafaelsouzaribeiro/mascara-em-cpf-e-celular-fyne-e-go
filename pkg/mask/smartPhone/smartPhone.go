package smartphone

import (
	"fyne.io/fyne/v2/widget"
	onlydigits "github.com/rafaelsouzaribeiro/mascara-em-cpf-e-celular-fyne-e-go/pkg/mask/onlyDigits"
)

func (c *Mask) SetMask() *widget.Entry {
	c.input.OnChanged = func(text string) {

		numericText := onlydigits.SetOnlyDigits(text)

		if len(numericText) > 11 {
			numericText = numericText[:11]
		}

		var maskedText string
		switch len(numericText) {
		case 0:
			maskedText = ""
		case 1, 2:
			maskedText = "(" + numericText
		case 3, 4, 5, 6:
			maskedText = "(" + numericText[:2] + ") " + numericText[2:]
		case 7, 8, 9, 10:
			maskedText = "(" + numericText[:2] + ") " + numericText[2:7]
			if len(numericText) > 7 {
				maskedText += "-" + numericText[7:]
			}
		case 11:
			maskedText = "(" + numericText[:2] + ") " + numericText[2:7] + "-" + numericText[7:]
		}

		if c.input.Text != maskedText {
			c.input.SetText(maskedText)
			c.input.CursorColumn = len(maskedText)
		}
	}

	return c.input
}
