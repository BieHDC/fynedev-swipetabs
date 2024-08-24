package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"biehdc.fynedev.swipetabs/swipetabs"
)

// good enough for quick dialog popping on mobile instead of printf
var window fyne.Window

func main() {
	fyne.EnsureCustom()
	app := app.New()
	window = app.NewWindow("Swipetabs - Fake Messenger")
	window.SetPadded(false)
	window.Resize(fyne.NewSize(500, 600))
	window.CenterOnScreen()

	window.SetContent(makeChannelList())
	window.ShowAndRun()
}

func makeChannelList() fyne.CanvasObject {
	var channels *container.Scroll
	disp := container.NewStack()

	// fixme do we want to make this a push/pull stack maybe?
	setView := func(co fyne.CanvasObject) {
		if co != nil {
			// if the caller wants us to set a new view
			disp.Objects = []fyne.CanvasObject{co}
		} else {
			// or it asks us to reset the view
			disp.Objects = []fyne.CanvasObject{channels}
		}
		disp.Refresh()
	}

	var channelWidgets []fyne.CanvasObject
	for i := range 30 {
		channelWidgets = append(channelWidgets, makeChannelWidget(fmt.Sprintf("Channel %d", i), setView))
	}

	channels = container.NewVScroll(container.NewVBox(channelWidgets...))
	setView(nil)

	return disp
}

func makeChannelWidget(channelname string, setView func(fyne.CanvasObject)) fyne.CanvasObject {
	messageText := widget.NewLabelWithStyle(channelname, fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	showChannelName := NewTapCard("", "", messageText, func() {
		resetView := func() { setView(nil) }
		content := append(
			[]fyne.CanvasObject{widget.NewButton("Back", resetView)},
			makeLotsOfFakeMessages(channelname, resetView)...,
		)
		setView(container.NewVScroll(container.NewVBox(content...)))
	})

	actions := container.NewGridWithColumns(3,
		widget.NewButton("LEAVE", func() { dialog.ShowInformation("LEAVE", fmt.Sprintf("leave for %s\n", channelname), window) }),
		widget.NewButton("ARCHIVE", func() { dialog.ShowInformation("ARCHIVE", fmt.Sprintf("archive for %s\n", channelname), window) }),
		widget.NewButton("OTHER", func() { dialog.ShowInformation("OTHER", fmt.Sprintf("other for %s\n", channelname), window) }),
	)

	channelView := container.NewStack(showChannelName) //default

	showChannelActions := NewTapCard("", "", container.NewGridWithColumns(2, messageText, actions), func() {
		channelView.Objects = []fyne.CanvasObject{showChannelName}
		channelView.Refresh()
	})

	return swipetabs.NewSwipeTabs(channelView, func(d swipetabs.Direction) {
		switch d {
		case swipetabs.Previous:
			channelView.Objects = []fyne.CanvasObject{showChannelName}
			channelView.Refresh()
		case swipetabs.Next:
			channelView.Objects = []fyne.CanvasObject{showChannelActions}
			channelView.Refresh()
		default:
			panic("should not happen")
		}
	})
}

func makeLotsOfFakeMessages(s string, resetView func()) []fyne.CanvasObject {
	var obj []fyne.CanvasObject

	for i := range 100 {
		w := widget.NewLabel(fmt.Sprintf("%s fake msg nr %d", s, i))

		handler := swipetabs.NewSwipeTabs(w, func(d swipetabs.Direction) {
			switch d {
			case swipetabs.Previous:
				resetView()
			case swipetabs.Next:
				dialog.ShowInformation("DELTE MESSAGE", fmt.Sprintf("delete invoked on %d\n", i), window)
			default:
				panic("should not happen ether")
			}
		})

		obj = append(obj, handler)
	}

	return obj
}
