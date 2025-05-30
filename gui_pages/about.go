package gui_pages

import (
	"fmt"
	"qemu-gui/helper"
	"qemu-gui/vars"
	"regexp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func About_Page(myApp fyne.App) {

	// about window
	aboutWindow := myApp.NewWindow("About")
	aboutWindow.Resize(fyne.NewSize(400, 300))

	// right area
	aboutRight := container.NewVBox(
		widget.NewLabel("Click Left Button To Use"),
	)

	// left button
	aboutButton := widget.NewButtonWithIcon("", theme.InfoIcon(), func() {
		aboutRight.RemoveAll()
		aboutRight.Add(
			widget.NewRichTextFromMarkdown(`# About QEMU GUI

QEMU GUI is a graphical user interface for the QEMU emulator.

It is written in Go and uses the Fyne toolkit for GUI development.
`),
		)
	})
	aboutLeft := container.NewVBox(
		aboutButton,
		widget.NewButtonWithIcon("", theme.ComputerIcon(), func() {
			aboutRight.RemoveAll()
			aboutRight.Add(widget.NewLabelWithStyle(
				"QEMU Arch Support Check",
				fyne.TextAlignCenter,
				fyne.TextStyle{Bold: true},
			))

			// check qemu executable
			go func() {
				for _, qemu_arch := range vars.QEMU_SUPPORTED_ARCH {
					status, output := helper.ExcutableCommand(vars.QEMU_ARCH[qemu_arch] + " --version")
					if status {

						// find version
						re := regexp.MustCompile(`QEMU emulator version (\d+\.\d+\.\d+)`)
						matches := re.FindStringSubmatch(output)

						// if version not found
						if len(matches) < 2 {
							aboutRight.Add(widget.NewLabel(
								fmt.Sprintf("%s is not installed or version not found", qemu_arch),
							))
							return
						}

						// show version
						aboutRight.Add(widget.NewLabel(
							fmt.Sprintf("%s version: %s", qemu_arch, matches[1]),
						))

					} else {
						aboutRight.Add(widget.NewLabel(
							fmt.Sprintf("%s is not installed or not found", qemu_arch),
						))
					}
				}
			}()

		}),
	)
	aboutButton.OnTapped()

	// show window
	aboutWindow.SetContent(container.NewHBox(
		aboutLeft,
		container.NewVScroll(aboutRight),
	))
	aboutWindow.Show()
}
