package qidong

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
	"shadanGo/conf"
	"shadanGo/shodanAPI/shodan"
	"shadanGo/window"
	"time"
)

var a = app.New()

func updateTime(label *widget.Label) {
	config := conf.Init()
	formatted := conf.PrintCurrentTime(config)
	label.SetText(formatted)
}

func Start() {
	w := a.NewWindow("ShodanEX")
	w.Resize(fyne.NewSize(800, 600))
	menu := window.CreateMenu(a, w)
	w.SetMainMenu(menu)
	label := widget.NewLabel("")
	// Create a single-line entry with initial width of 400
	entry := widget.NewEntry()
	// Create table with first row containing "IP" and "Port" headers
	var btn *widget.Button // Define button variable
	exeBtn := widget.NewButton("export report", func() {
		apiKey, err := window.ReadAPIKeyFromFile()
		if err != nil {
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Error",
				Content: "APIKey not found,\nPlease go to the menu bar's Open->verity APIkey to fill in your shodanAPIKey and verify it.",
			})
		} else {
			s := shodan.New(apiKey)
			hostSearch, err := s.HostSearch(entry.Text)
			err = shodan.SaveToExcel(hostSearch.Matches, "output.xlsx")
			if err != nil {
				fyne.CurrentApp().SendNotification(&fyne.Notification{
					Title:   "Error",
					Content: "Some errors have occurred.\nPlease try again later.",
				})
			}
		}
	})
	btn = widget.NewButton("shodan Search", func() {
		apiKey, err := window.ReadAPIKeyFromFile()
		if err != nil {
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Error",
				Content: "APIKey not found,\nPlease go to the menu bar's Open->verity APIkey to fill in your shodanAPIKey and verify it.",
			})
		} else {
			s := shodan.New(apiKey)
			hostSearch, err := s.HostSearch(entry.Text)
			if err != nil {
				log.Fatal(err)
			}
			table := window.CreateTable(hostSearch)
			// Create a container with table and scrollbar
			scrollContainer := container.NewScroll(table)
			scrollContainer.SetMinSize(fyne.NewSize(800, 400)) // Set the minimum size of the scroll container

			// Create a vertical layout container with label, entry, and scroll container
			vBox := container.NewVBox(
				label,
				entry,
				btn,
				exeBtn,
				scrollContainer)
			// Set the vertical layout container as the window content
			w.SetContent(vBox)
		}
	})
	//Create export button to Excel

	// Create a horizontal layout container
	hBox := container.New(layout.NewHBoxLayout(),
		label,
		layout.NewSpacer(),
		btn,
		exeBtn,
		//table,
	)

	// Create a vertical layout container
	vBox := container.New(layout.NewVBoxLayout(), hBox, entry)
	w.SetContent(vBox)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(label)
		}
	}()
	w.ShowAndRun()
}
