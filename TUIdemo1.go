package main

import (
	"github.com/rivo/tview"
	"github.com/xuri/excelize/v2"
	"log"
)

type DataScreen struct {
	Company  *tview.InputField
	JobTitle *tview.InputField
	Location *tview.InputField
	IsUSA    *tview.Checkbox
}

var dataDisplay DataScreen

func main() {
	dataToDisplay := getData("JobData.xlsx")
	firstItem := dataToDisplay[1]
	makeDataDisplay()
	dataDisplay.Company.SetText(firstItem[0])
	dataDisplay.JobTitle.SetText(firstItem[9])
	dataDisplay.Location.SetText(firstItem[4])
	app := tview.NewApplication()
	form := tview.NewForm()
	form.SetBorder(true)
	form.AddFormItem(dataDisplay.Company)
	form.AddFormItem(dataDisplay.JobTitle)
	form.AddFormItem(dataDisplay.Location)
	form.AddFormItem(dataDisplay.IsUSA)
	app.SetRoot(form, true)
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}

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
