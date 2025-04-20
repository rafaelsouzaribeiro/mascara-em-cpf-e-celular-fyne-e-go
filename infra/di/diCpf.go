package di

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/rafaelsouzaribeiro/mascara-em-cpf-e-celular-fyne-e-go/pkg/components/input"
	"github.com/rafaelsouzaribeiro/mascara-em-cpf-e-celular-fyne-e-go/pkg/mask/cpf"
)

func NewCpfDI() *fyne.Container {

	cpf := cpf.NewMask(input.NewEnrty("Digite o CPF"))
	cpfEntry := cpf.SetMask()
	content := container.NewVBox(
		widget.NewLabel("Digite o CPF:"),
		cpfEntry,
	)

	return content
}
