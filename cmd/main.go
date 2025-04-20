package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/rafaelsouzaribeiro/mascara-em-cpf-e-celular-fyne-e-go/infra/di"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Máscara de Celular")

	content := di.NewSmartPoneDI()
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 100))
	myWindow.ShowAndRun()
}
