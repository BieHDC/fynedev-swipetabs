package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Tappable = (*TapCard)(nil)

type TapCard struct {
	widget.Card
	onTapped func()
}

func NewTapCard(title, subtitle string, content fyne.CanvasObject, onTapped func()) *TapCard {
	return &TapCard{
		Card: widget.Card{
			Title:    title,
			Subtitle: subtitle,
			Content:  content,
		},
		onTapped: onTapped,
	}
}

func (cc *TapCard) Tapped(pe *fyne.PointEvent) {
	if cc.onTapped != nil {
		cc.onTapped()
	}
}
