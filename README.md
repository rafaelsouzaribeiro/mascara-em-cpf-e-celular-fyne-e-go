<h1>Aplicando máscara em CPF e CELULAR em golang e fyne.</h1><br />
<p>Chamadas no main.go de CELULAR para CELULAR:</p><br />
```go
	myApp := app.New()
	myWindow := myApp.NewWindow("Máscara de CPF")

	content := di.NewCpfDI()
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 100))
	myWindow.ShowAndRun()

 ```
 <br />

```go
 	myApp := app.New()
	myWindow := myApp.NewWindow("Máscara de Celular")

	content := di.NewSmartPoneDI()
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 100))
	myWindow.ShowAndRun()
 ```
