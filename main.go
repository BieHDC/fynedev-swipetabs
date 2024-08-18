package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"biehdc.fynedev.swipetabs/swipetabs"
)

func main() {
	fyne.EnsureCustom()
	app := app.New()
	window := app.NewWindow("Swipetabs")
	window.SetPadded(false)
	window.Resize(fyne.NewSize(500, 600))
	window.CenterOnScreen()

	tabs2 := container.NewAppTabs2(true,
		container.NewTabItem("1", widget.NewLabel("item 1")),
		container.NewTabItem("2", randomstuff()),
		container.NewTabItem("3", subswiper()),
	)
	tabs2.SelectIndex(2)

	tabs := swipetabs.NewSwipeTabs(tabs2, func(d swipetabs.Direction) {
		switch d {
		case swipetabs.Previous:
			tabs2.SelectIndex(tabs2.SelectedIndex() - 1)
		case swipetabs.Next:
			tabs2.SelectIndex(tabs2.SelectedIndex() + 1)
		default:
			panic("should not happen")
		}
	})

	window.SetContent(tabs)
	window.ShowAndRun()
}

func randomstuff() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("2 yikes"),
		widget.NewButton("a button", func() {}),
		widget.NewActivity(),
		widget.NewEntry(),
		widget.NewSlider(5, 15),
	)
}

func lb(s string) fyne.CanvasObject {
	return widget.NewLabel(s)
}

func subswiper() fyne.CanvasObject {
	gw := container.NewGridWithColumns(3,
		lb("1"), lb("2"), lb("3"),
		//
		tabberWithSwipe(true, lb("with swipe 1"), lb("yyy"), lb("zzz")),
		tabberWithSwipe(false, lb("with swipe, not hidden"), lb("ccc"), lb("ggg")),
		tabberWithSwipe(true, lb("with swipe 2"), lb("ooo"), lb("ppp")),
		//
		lb("7"), tabber(lb("no swipe"), lb("mmm"), lb("qqq")), lb("9"),
	)
	return gw
}

func tabberWithSwipe(hide bool, a, b, c fyne.CanvasObject) fyne.CanvasObject {
	var tabs2 *container.AppTabs

	// for quick demo
	if hide {
		tabs2 = container.NewAppTabs2(true,
			container.NewTabItem("11", a),
			container.NewTabItem("22", b),
			container.NewTabItem("33", c),
		)
	} else {
		tabs2 = container.NewAppTabs(
			container.NewTabItem("11", a),
			container.NewTabItem("22", b),
			container.NewTabItem("33", c),
		)
	}

	tabs := swipetabs.NewSwipeTabs(tabs2, func(d swipetabs.Direction) {
		switch d {
		case swipetabs.Previous:
			tabs2.SelectIndex(tabs2.SelectedIndex() - 1)
		case swipetabs.Next:
			tabs2.SelectIndex(tabs2.SelectedIndex() + 1)
		default:
			panic("should not happen")
		}
	})

	return tabs
}

func tabber(a, b, c fyne.CanvasObject) fyne.CanvasObject {
	tabs := container.NewAppTabs(
		container.NewTabItem("T1", a),
		container.NewTabItem("T2", b),
		container.NewTabItem("T3", c),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	return tabs
}
