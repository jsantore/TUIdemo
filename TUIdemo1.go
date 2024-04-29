package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/xuri/excelize/v2"
	"log"
)

type DataScreen struct {
	Company     *tview.InputField
	JobTitle    *tview.InputField
	Location    *tview.InputField
	IsUSA       *tview.Checkbox
	currentItem int
}

var dataDisplay DataScreen

func main() {
	dataToDisplay := getData("JobData.xlsx")

	firstItem := dataToDisplay[dataDisplay.currentItem]
	makeDataDisplay()
	dataDisplay.Company.SetText(firstItem[0])
	dataDisplay.JobTitle.SetText(firstItem[9])
	dataDisplay.Location.SetText(firstItem[4])
	dataDisplay.IsUSA.SetChecked(firstItem[3] == "us")
	app := tview.NewApplication()
	form := tview.NewForm()
	form.SetBorder(true)
	form.AddFormItem(dataDisplay.Company)
	form.AddFormItem(dataDisplay.JobTitle)
	form.AddFormItem(dataDisplay.Location)
	form.AddFormItem(dataDisplay.IsUSA)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			app.Stop()
		}
		return event
	})

	moveToPrev := func() {
		dataDisplay.currentItem--
		if dataDisplay.currentItem < 0 {
			dataDisplay.currentItem = 0
		}
		currentItem := dataToDisplay[dataDisplay.currentItem]
		dataDisplay.Company.SetText(currentItem[0])
		dataDisplay.JobTitle.SetText(currentItem[9])
		dataDisplay.Location.SetText(currentItem[4])
		dataDisplay.IsUSA.SetChecked(currentItem[3] == "us")
	}
	moveToNext := func() {
		dataDisplay.currentItem++
		if dataDisplay.currentItem >= len(dataToDisplay) {
			dataDisplay.currentItem = len(dataToDisplay) - 1
		}
		currentItem := dataToDisplay[dataDisplay.currentItem]
		dataDisplay.Company.SetText(currentItem[0])
		dataDisplay.JobTitle.SetText(currentItem[9])
		dataDisplay.Location.SetText(currentItem[4])
		dataDisplay.IsUSA.SetChecked(currentItem[3] == "us")
	}
	form.AddButton("Previous", moveToPrev)
	form.AddButton("Next", moveToNext)
	form.AddButton("Quit", app.Stop)
	form.SetBorder(true)
	form.SetButtonBackgroundColor(tcell.ColorDarkGoldenrod)
	app.SetRoot(form, true)
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func updateUI() {}

func makeDataDisplay() {
	dataDisplay.Company = tview.NewInputField()
	dataDisplay.Company.SetLabel("Company:")
	dataDisplay.Company.SetFieldWidth(30)
	dataDisplay.JobTitle = tview.NewInputField()
	dataDisplay.JobTitle.SetLabel("Job Title:")
	dataDisplay.JobTitle.SetFieldWidth(80)
	dataDisplay.Location = tview.NewInputField()
	dataDisplay.Location.SetLabel("Location:")
	dataDisplay.Location.SetFieldWidth(40)
	dataDisplay.IsUSA = tview.NewCheckbox()
	dataDisplay.IsUSA.SetLabel("Is US Job:")

}

// the getData function from the cache Demo
func getData(fileName string) [][]string {
	excelFile, err := excelize.OpenFile(fileName)
	defer excelFile.Close()
	if err != nil {
		log.Fatalln("couldn't open file", err)
	}
	all_rows, err := excelFile.GetRows("Comp490 Jobs")
	if err != nil {
		log.Fatalln(err)
	}
	return all_rows
}
