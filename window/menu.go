package window

import (
	"bufio"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
	"shadanGo/shodanAPI/shodan"
	"strconv"
)

func saveAPIKey(apiKey string) error {
	file, err := os.Create("key.conf")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(apiKey)
	if err != nil {
		return err
	}

	return nil
}

func ReadAPIKeyFromFile() (string, error) {
	file, err := os.Open("key.conf")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var apiKey string
	for scanner.Scan() {
		apiKey = scanner.Text()
	}

	return apiKey, nil
}

func CreateMenu(App fyne.App, win fyne.Window) *fyne.MainMenu {
	fileMenu := fyne.NewMenu("Open",
		fyne.NewMenuItem("verity APIkey", func() {
			// Add action for New
			newWindow := App.NewWindow("verity your shodan apikey")
			apiKey, _ := ReadAPIKeyFromFile()
			input := widget.NewEntry()
			input.SetText(apiKey)
			submit := widget.NewButton("Submit", func() {
				apikey := input.Text
				//verity
				s := shodan.New(apikey)
				info, err := s.APIInfo()
				if err != nil {
					fyne.CurrentApp().SendNotification(&fyne.Notification{
						Title:   "Error",
						Content: "APIKey can't use",
					})
				} else {
					content := "Query Credits: " + strconv.Itoa(info.QueryCredits) + "\n" +
						"Scan Credits: " + strconv.Itoa(info.ScanCredits)
					fyne.CurrentApp().SendNotification(&fyne.Notification{
						Title:   "Success",
						Content: content,
					})
					err := saveAPIKey(apikey)
					if err != nil {
						return
					} else {
						fyne.CurrentApp().SendNotification(&fyne.Notification{
							Title:   "Save Success",
							Content: "APIKey save to key.conf",
						})
					}
				}
			})

			content := container.NewVBox(
				widget.NewLabel("Input:"),
				input,
				submit,
			)
			newWindow.SetContent(content)
			newWindow.Resize(fyne.NewSize(300, 200))
			newWindow.Show()
		}),
		//fyne.NewMenuItem("Open", func() {
		//	// Add action for Open
		//}),
		//fyne.NewMenuItem("Save", func() {
		//	// Add action for Save
		//}),
		//fyne.NewMenuItem("Quit", func() {
		//	win.Close()
		//}),
	)

	return fyne.NewMainMenu(
		fileMenu,
	)
}
