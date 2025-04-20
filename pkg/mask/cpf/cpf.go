package cpf

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
		if len(numericText) > 3 {
			maskedText = numericText[:3] + "."
		} else {
			maskedText = numericText
		}
		if len(numericText) > 6 {
			maskedText += numericText[3:6] + "."
		} else if len(numericText) > 3 {
			maskedText += numericText[3:]
		}
		if len(numericText) > 9 {
			maskedText += numericText[6:9] + "-"
		} else if len(numericText) > 6 {
			maskedText += numericText[6:]
		}
		if len(numericText) > 9 {
			maskedText += numericText[9:]
		}

		c.input.SetText(maskedText)
		c.input.CursorColumn = len(maskedText)
	}

	return c.input
}
