package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	attemptsCount = 10
	minNumber     = 0
	maxNumber     = 300
)

func main() {

	var startGame func()

	a := app.New()
	w := a.NewWindow("Угадайка")
	w.Resize(fyne.NewSize(450, 200))
	w.SetFixedSize(true)

	startGame = func() {
		content := container.NewVBox()
		greetingLabel := widget.NewLabel(fmt.Sprintf("Нажмите кнопку ниже, чтобы компьютер загадал число между %d и %d!", minNumber, maxNumber))

		tryAgainButton := widget.NewButton("Начать заново", func() {
			content.RemoveAll()
			startGame()
		})

		var startButton *widget.Button
		startButton = widget.NewButton("Загадать новое число!", func() {
			lives := attemptsCount
			randomNumber := rand.Intn(maxNumber)
			log.Println("[INFO] | Загаданное число |", randomNumber)

			greetingLabel.SetText(fmt.Sprintf("Компьютер загадал число! Попыток осталось: %d", lives))
			content.Remove(startButton)

			input := widget.NewEntry()
			input.SetPlaceHolder("Введите ваше число: ")
			content.Add(input)

			var tryButton *widget.Button
			tryButton = widget.NewButton("Угадать", func() {

				userInputNumber, err := strconv.Atoi(input.Text)
				if err != nil || userInputNumber > maxNumber || userInputNumber < minNumber {
					lives--
					greetingLabel.SetText(fmt.Sprintf("Неверный ввод. Попыток осталось: %d", lives))
				} else if userInputNumber == randomNumber {
					lives--
					greetingLabel.SetText(fmt.Sprintf("Поздравляю! Вы угадали число за %d попыток.", attemptsCount-lives))
					content.Remove(tryButton)
					content.Add(tryAgainButton)
				} else if lives == 1 {
					greetingLabel.SetText("К сожалению, вы проиграли. Попробуйте заново!")
					content.Remove(tryButton)
					content.Add(tryAgainButton)
				} else {
					lives--
					if userInputNumber > randomNumber {
						greetingLabel.SetText(fmt.Sprintf("Ваше число больше. Попыток осталось: %d", lives))
					} else {
						greetingLabel.SetText(fmt.Sprintf("Ваше число меньше. Попыток осталось: %d", lives))
					}
				}

			})
			content.Add(tryButton)
		})

		content.Add(greetingLabel)
		content.Add(startButton)
		w.SetContent(content)
	}

	startGame()
	w.ShowAndRun()
}
