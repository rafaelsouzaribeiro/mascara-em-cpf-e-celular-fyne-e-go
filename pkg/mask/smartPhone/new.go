package smartphone

import (
	"fyne.io/fyne/v2/widget"
	"github.com/rafaelsouzaribeiro/mascara-em-cpf-e-celular-fyne-e-go/pkg/mask/types"
)

type Mask struct {
	input *widget.Entry
}

func NewMask(input *widget.Entry) types.IMask {
	return &Mask{
		input: input,
	}
}
