package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create a new Fyne application
	myApp := app.New()

	// Create a new window
	myWindow := myApp.NewWindow("Name List")

	// Create a list widget to display the names
	list := widget.NewList(
		func() int {
			return len(getNames())
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(index int, item fyne.CanvasObject) {
			if label, ok := item.(*widget.Label); ok {
				label.SetText(getNames()[index])
			}
		},
	)

	// Create a container to hold the list widget
	listContainer := container.NewVBox(
		widget.NewLabel("Name List"),
		list,
	)

	// Set the content of the window to the list container
	myWindow.SetContent(listContainer)

	// Start a goroutine to periodically update the names
	go func() {
		for {
			time.Sleep(2 * time.Second) // Wait for 2 seconds
			myApp.QueueUpdate(func() {
				list.Refresh() // Refresh the list to fetch updated names
			})
		}
	}()

	// Show the window and run the application
	myWindow.ShowAndRun()
}

// Function to fetch names from an external source (dummy implementation)
func getNames() []string {
	// Simulating fetching names from an external source
	names := []string{
		"John",
		"Jane",
		"Michael",
		"Sara",
		"David",
	}

	// Adding a timestamp to the names to demonstrate updates
	for i := range names {
		names[i] = fmt.Sprintf("%s - %s", names[i], time.Now().Format("15:04:05"))
	}

	return names
}

