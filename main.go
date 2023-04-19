package main

import (
	"fmt"
	"log"
	"math/rand"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/Knetic/govaluate"
)

func main() {
	// Инициализация приложения
	a := app.New()
	// наше окошко
	w := a.NewWindow("Калькулейтор")

	// поле ввода (считываем формулки)
	entry := ""

	/*
	* работаем с виждетом:
	* Рендерим, ставим Placeholder, вешаем слушатель событй
	 */
	entryDisplay := widget.NewEntry()
	entryDisplay.SetPlaceHolder("Введите число")
	entryDisplay.OnChanged = func(a string) { entry = a }

	valueDisplay := widget.NewLabel("")

	// контейнер для кнопок
	buttonPad := fyne.NewContainerWithLayout(layout.NewGridLayout(5))

	// Добавим кнопочки 1-9
	for i := 1; i < 10; i++ {
		buttonPad.AddObject(changeButton(fmt.Sprint(i), &entry, entryDisplay))
	}
	buttonPad.AddObject(changeButton("0", &entry, entryDisplay))

	btns := []string{"+", "-", "/", "*", "%", "(", ")", "^", "<", ">"}

	// кнопки операций
	for i := range btns {
		buttonPad.AddObject(changeButton(btns[i], &entry, entryDisplay))
	}
	// Ошибки (учим калькулейтор ругаться)
	errz := []string{"АААА, ОШИБКО МАНАМА!(x_x)", "Ты дебил?", "Что ты несешь?", "Я отказываюсь это считать", "Считай сам, я в отпуске", "Считай сам"}

	// Функция "=", посчитает что угодно, если захочет 😎🤖👾
	enter := widget.NewButton("=", func() {
		swear, err := govaluate.NewEvaluableExpression(entry)
		//
		if err != nil {
			valueDisplay.SetText(errz[rand.Intn(len(errz))])
			log.Printf(`Вы не нравитесь калькулятору...
Он говорит - "%s"`, errz[rand.Intn(len(errz))])
			return
		}
		value, err := swear.Evaluate(nil)
		if err != nil {
			valueDisplay.SetText(errz[rand.Intn(len(errz))])
			return
		}
		valueDisplay.SetText(fmt.Sprint(value))
	})

	// Функция очистки поля воода (с формулками, то самое =))
	clear := widget.NewButton("AC", func() {
		entry = ""
		entryDisplay.SetText(entry)
		valueDisplay.SetText("")
	})

	// Отдельная секция для кнопушке меню (Очистить и =)
	submitPad := fyne.NewContainerWithLayout(layout.NewGridLayout(2))
	submitPad.AddObject(clear)
	submitPad.AddObject(enter)

	// волшебная коробочка с виджетами
	w.SetContent(widget.NewVBox(
		entryDisplay,
		valueDisplay,
		buttonPad,
		submitPad,
	))
	// 3... 2... 1... GO!
	// OKAAAAAAAAAAY LET's goooooo
	w.ShowAndRun()
}

// Обработаем события при клике на кнопушки
func changeButton(mod string, entry *string, entryDisplay *widget.Entry) *widget.Button {
	button := widget.NewButton(mod, func() {
		*entry += mod
		entryDisplay.SetText(*entry)

	})
	return button
}
