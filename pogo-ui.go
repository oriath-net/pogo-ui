package main

import (
	"fmt"
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/oriath-net/pogo/ggpk"
)

var app struct {
	mw        *walk.MainWindow
	treeView  *walk.TreeView
	treeModel ggpkTreeModel

	ggpkFilename string
	ggpk         *ggpk.File
}

func main() {
	var err error

	err = MainWindow{
		Title:    "pogo-ui",
		MinSize:  Size{320, 200},
		Size:     Size{1024, 600},
		Layout:   VBox{},
		AssignTo: &app.mw,
		MenuItems: []MenuItem{
			Menu{
				Text: "&File",
				Items: []MenuItem{
					Action{
						Text:        "&Open GGPK...",
						Shortcut:    Shortcut{walk.ModControl, walk.KeyO},
						OnTriggered: do_open,
					},
					Action{
						Text:        "&Quit",
						Shortcut:    Shortcut{walk.ModControl, walk.KeyQ},
						OnTriggered: do_quit,
					},
				},
			},
			Menu{
				Text: "&Edit",
				Items: []MenuItem{
					Action{
						Text:     "&Undo",
						Enabled:  false,
						Shortcut: Shortcut{Modifiers: walk.ModControl, Key: walk.KeyZ},
					},
					Separator{},
					Action{
						Text:     "Cu&t",
						Shortcut: Shortcut{walk.ModControl, walk.KeyX},
						Enabled:  false,
					},
					Action{
						Text:     "&Copy",
						Shortcut: Shortcut{walk.ModControl, walk.KeyC},
						// FIXME: action?
					},
					Action{
						Text:     "&Paste",
						Shortcut: Shortcut{walk.ModControl, walk.KeyV},
						Enabled:  false,
					},
				},
			},
			Menu{
				Text: "&Help",
				Items: []MenuItem{
					Action{
						Text:        "&About pogo-ui...",
						OnTriggered: do_about,
					},
				},
			},
		},

		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TreeView{
						AssignTo: &app.treeView,
						Model:    &app.treeModel,
					},
					CustomWidget{},
				},
			},
		},
	}.Create()
	if err != nil {
		log.Fatal(err)
	}

	do_open()

	app.mw.Run()
}

func do_about() {
	walk.MsgBox(
		app.mw,
		"About pogo-ui",
		"https://github.com/oriath-net/pogo-ui",
		walk.MsgBoxIconInformation,
	)
}

func do_quit() {
	app.mw.Close()
}

func do_open() {
	dlg := &walk.FileDialog{
		Title:    "Select a GGPK file",
		Filter:   "GGPK archive (*.ggpk)|*.ggpk",
		FilePath: "C:\\Program Files (x86)\\Grinding Gear Games\\Path of Exile\\Content.ggpk",
	}

	ok, err := dlg.ShowOpen(app.mw)
	if err != nil {
		log.Println(err)
		return
	}
	if !ok {
		return
	}

	f, err := ggpk.Open(dlg.FilePath)
	if err != nil {
		walk.MsgBox(
			app.mw,
			"Failed to open archive",
			fmt.Sprintf("Failed to open %s: %s", dlg.FilePath, err),
			walk.MsgBoxIconStop,
		)
		return
	}

	app.ggpk = f
	app.ggpkFilename = dlg.FilePath

	app.treeView.SetModel(nil)
	app.treeView.SetModel(&app.treeModel)
}
