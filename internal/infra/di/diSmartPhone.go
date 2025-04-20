package di

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/rafaelsouzaribeiro/mascara-em-cpf-e-celular-fyne-e-go/internal/components/input"
	smartphone "github.com/rafaelsouzaribeiro/mascara-em-cpf-e-celular-fyne-e-go/pkg/mask/smartPhone"
)

func NewSmartPoneDI() *fyne.Container {

	smartphone := smartphone.NewMask(input.NewEnrty("Digite o Número"))
	smartphoneEntry := smartphone.SetMask()
	content := container.NewVBox(
		widget.NewLabel("Digite o Número:"),
		smartphoneEntry,
	)

	return content
}
