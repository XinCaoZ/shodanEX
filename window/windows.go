package window

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"shadanGo/shodanAPI/shodan"
)

func CreateTable(hostSearch *shodan.HostSearch) *widget.Table {
	table := widget.NewTable(
		func() (int, int) {
			// Return the number of matches and columns
			return len(hostSearch.Matches) + 1, 3 // Add 1 to include header row
		},
		func() fyne.CanvasObject {
			// Return a label for each cell
			return widget.NewLabel("")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			// Update cell content with search results
			if id.Row == 0 { // Header row
				switch id.Col {
				case 0:
					cell.(*widget.Label).SetText("IP")
				case 1:
					cell.(*widget.Label).SetText("Port")
				case 2:
					cell.(*widget.Label).SetText("Domains")
				}
			} else { // Data rows
				switch id.Col {
				case 0:
					cell.(*widget.Label).SetText(hostSearch.Matches[id.Row-1].IPString)
				case 1:
					cell.(*widget.Label).SetText(fmt.Sprintf("%d", hostSearch.Matches[id.Row-1].Port))
				case 2:
					if len(hostSearch.Matches[id.Row-1].Domains) > 0 {
						cell.(*widget.Label).SetText(hostSearch.Matches[id.Row-1].Domains[0])
					} else {
						cell.(*widget.Label).SetText("No Domains")
					}
				}
			}
		})

	// Set column names
	table.SetColumnWidth(0, 200)
	table.SetColumnWidth(1, 200) // Adjust the width of the second column to fit the window
	table.SetColumnWidth(2, 200)

	return table
}
