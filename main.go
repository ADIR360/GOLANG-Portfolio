package main

import (
    "github.com/rivo/tview"
)

func main() {
    app := tview.NewApplication()

    // Menu list
    menu := tview.NewList().
        AddItem("Home", "Welcome message", 'h', nil).
        AddItem("About", "About me section", 'a', nil).
        AddItem("Projects", "List of projects", 'p', nil).
        AddItem("Contact", "Contact information", 'c', nil).
        AddItem("Resume", "My resume", 'r', nil).
        AddItem("Quit", "Press to exit", 'q', func() {
            app.Stop()
        })

    menu.SetBorder(true).SetTitle("Menu").SetTitleAlign(tview.AlignLeft)

    // Detail section (right side)
    details := tview.NewTextView().
        SetDynamicColors(true).
        SetText("Select an item from the menu")

    details.SetBorder(true).SetTitle("Details").SetTitleAlign(tview.AlignLeft)

    // Handle menu item selection
    menu.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
        switch mainText {
        case "Home":
            details.SetText("Welcome to my portfolio!\nUse the menu on the left to navigate.")
        case "About":
            details.SetText("[yellow]About Me:\n\n[white]I am a Software Developer with expertise in Go, Python, and web technologies.")
        case "Projects":
            details.SetText("[yellow]Projects:\n\n[white]1. Vulnerability Tracker\n2. Network Analysis Tool\n3. Ping Pong Game")
        case "Contact":
            details.SetText("[yellow]Contact Info:\n\n[white]Email: example@domain.com\nGitHub: github.com/your-username")
        case "Resume":
            details.SetText("[yellow]Resume:\n\n[white]Bachelor of Technology in Computer Science\nExperience in software development and cybersecurity.")
        }
    })

    // Layout: place menu on the left and details on the right
    flex := tview.NewFlex().
        AddItem(menu, 0, 1, true).
        AddItem(details, 0, 2, false)

    if err := app.SetRoot(flex, true).Run(); err != nil {
        panic(err)
    }
}
