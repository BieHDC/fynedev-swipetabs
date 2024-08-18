package swipetabs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Declare conformity with Widget interface.
var _ fyne.Widget = (*SwipeTabs)(nil)

// Declare conformity with Draggable interface.
var _ fyne.Draggable = (*SwipeTabs)(nil)

// fixme:
type SwipeTabs struct {
	widget.BaseWidget
	//
	onScroll    func(Direction)
	framebuffer fyne.CanvasObject
	dragstart   *fyne.DragEvent
	dragend     *fyne.DragEvent
}

type Direction int

const (
	Previous Direction = iota
	Next
)

func (t *SwipeTabs) Dragged(de *fyne.DragEvent) {
	t.dragend = de

	if t.dragstart == nil {
		inset := t.Size().Width / 10.0
		posx := de.PointEvent.Position.X
		// if we started dragging inside the outer border, we care
		if posx < inset || t.Size().Width-inset > posx {
			t.dragstart = de
		}
	}
}

func (t *SwipeTabs) DragEnd() {
	if t.dragend == nil || t.dragstart == nil {
		t.dragstart = nil
		return
	}

	startx := t.dragstart.Position.X
	endx := t.dragend.Position.X

	insetleft := (t.Size().Width / 10.0) * 2 //fixme should we *2?
	if startx < insetleft && endx > insetleft {
		t.onScroll(Previous)
	}

	insetright := t.Size().Width - insetleft
	if startx > insetright && endx < insetright {
		t.onScroll(Next)
	}

	t.dragstart = nil
	t.dragend = nil
}

func NewSwipeTabs(item fyne.CanvasObject, onScroll func(Direction)) *SwipeTabs {
	t := &SwipeTabs{}
	t.BaseWidget.ExtendBaseWidget(t)
	t.framebuffer = item
	t.onScroll = onScroll
	return t
}

// CreateRenderer is a private method to Fyne which links this widget to its renderer
//
// Implements: fyne.Widget
func (t *SwipeTabs) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(t.framebuffer)
}
