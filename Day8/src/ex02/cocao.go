package main

import (
	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len([]string{}), []string{})

	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("School 21")
	window.SetGeometry2(300, 200, 300, 200)
	window.Show()

	widgets.QApplication_Exec()
}

