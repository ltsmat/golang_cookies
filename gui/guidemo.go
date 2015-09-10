/// +build ignore

package main

import (
	"github.com/andlabs/ui"
	"log"
)

func main() {
	// This runs the code that displays our GUI.
	// All code that interfaces with package ui (except event handlers) must be run from within a ui.Do() call.
	go ui.Do(GuiMain)

	err := ui.Go()
	if err != nil {
		log.Print(err)
	}
}

/*******



********/
func GuiMain() {
	//320*240
	//480*320
	//640*480
	//800*600
	//1024*768
	WinXSize := 640
	WinYSize := 480

	// All windows must have a control inside.
	// ui.Space() creates a control that is just a blank space for us to use.
	
	//disp
	TxtMain    := ui.NewTextbox()
	GrpDisp    := ui.NewGroup("Disp", TxtMain)
	GrpDisp.SetMargined(true)
	
	//ctrl
	LblSet     := ui.NewLabel("Set:")
	FldSet     := ui.NewTextField()
	BtnSet     := ui.NewButton("Set")
	FldSet.SetText("255.255.255.255")
	
	GrdCtrl    := ui.NewGrid()
	GrdCtrl.Add(LblSet,nil,ui.East, false, ui.Fill, false, ui.Fill, 1, 1)	
	GrdCtrl.Add(FldSet,nil,ui.East, false, ui.Fill, false, ui.Fill, 1, 1)	
	GrdCtrl.Add(BtnSet,nil,ui.East, false, ui.Fill, false, ui.Fill, 1, 1)
	GrdCtrl.SetPadded(true)
	
	GrpCtrl    := ui.NewGroup("Ctrl", GrdCtrl)
	GrpCtrl.SetMargined(true)
	
	//cfg
//	TabConf    := ui.NewTab()
//	HCTTab1    := ui.NewVerticalStack(ui.NewCheckbox("funny"),
//					ui.NewSpinbox(0,20))
//	TabConf.Append("Tab 1", HCTTab1)
//	TabConf.Append("Tab 2", ui.NewTextbox())
	GrpConf    := ui.NewGroup("Conf", ui.NewTextbox())
	GrpConf.SetMargined(true)
	
	BtnAction  := ui.NewButton("Run")
	
	
	GrdMain    := ui.NewGrid()
	GrdMain.Add(GrpDisp,nil,ui.East, true, ui.Fill, true, ui.Fill, 1, 1)	
	GrdMain.Add(GrpCtrl,GrpDisp,ui.East, false, ui.Fill, true, ui.Fill, 1, 1)	
	GrdMain.Add(GrpConf,GrpDisp,ui.South, true, ui.Fill, false, ui.Fill, 1, 1)	
	GrdMain.Add(BtnAction,nil,ui.East, false, ui.Fill, false, ui.Fill, 1, 1)	
	
	GrdMain.SetPadded(true)
	TxtMain.SetText("i'm full")
	
	// Then we create a window.
	w := ui.NewWindow("Demo", WinXSize, WinYSize, GrdMain)

	// We tell package ui to destroy our window and shut down cleanly when the user closes the window by clicking the X button in the titlebar.
	w.OnClosing(func() bool {
		// This informs package ui to shut down cleanly when it can.
		ui.Stop()
		// And this informs package ui that we want to hide AND destroy the window.
		return true
	})

	// And finally, we need to show the window.
	w.Show()
}
