package main

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
)

const shortcutsWindowUI = ` <?xml version="1.0" encoding="UTF-8"?>
<interface>
  <object class="GtkShortcutsWindow" id="shortcuts-window">
    <property name="modal">1</property>
    <child>
      <object class="GtkShortcutsSection">
        <property name="visible">1</property>
        <property name="section-name">shortcuts</property>
        <property name="max-height">12</property>
        <child>
          <object class="GtkShortcutsGroup">
            <property name="visible">1</property>
            <property name="title">Comic Navigation</property>
            <child>
              <object class="GtkShortcutsShortcut">
                <property name="visible">1</property>
                <property name="accelerator">&lt;ctrl&gt;Left</property>
                <property name="title">Previous Comic</property>
              </object>
            </child>
            <child>
              <object class="GtkShortcutsShortcut">
                <property name="visible">1</property>
                <property name="accelerator">&lt;ctrl&gt;Right</property>
                <property name="title">Next Comic</property>
              </object>
            </child>
            <child>
              <object class="GtkShortcutsShortcut">
                <property name="visible">1</property>
                <property name="accelerator">&lt;ctrl&gt;r</property>
                <property name="title">Random Comic</property>
              </object>
            </child>
            <child>
              <object class="GtkShortcutsShortcut">
                <property name="visible">1</property>
                <property name="accelerator">&lt;ctrl&gt;f</property>
                <property name="title">Search Comics</property>
              </object>
            </child>
          </object>
        </child>
        <child>
          <object class="GtkShortcutsGroup">
            <property name="visible">1</property>
            <property name="title">Application Actions</property>
            <child>
              <object class="GtkShortcutsShortcut">
                <property name="visible">1</property>
                <property name="accelerator">&lt;ctrl&gt;n</property>
                <property name="title">Open New Window</property>
              </object>
            </child>
            <child>
              <object class="GtkShortcutsShortcut">
                <property name="accelerator">&lt;ctrl&gt;p</property>
                <property name="visible">1</property>
                <property name="title">Open Properties Window</property>
              </object>
            </child>
            <child>
              <object class="GtkShortcutsShortcut">
                <property name="visible">1</property>
                <property name="accelerator">&lt;ctrl&gt;q</property>
                <property name="title">Quit Application</property>
              </object>
            </child>
          </object>
        </child>
      </object>
    </child>
  </object>
</interface>
`

func NewShortcutsWindow() (*gtk.ShortcutsWindow, error) {
	builder, err := gtk.BuilderNew()
	if err != nil {
		return nil, err
	}

	err = builder.AddFromString(shortcutsWindowUI)
	if err != nil {
		return nil, err
	}

	obj, err := builder.GetObject("shortcuts-window")
	if err != nil {
		return nil, err
	}

	return obj.(*gtk.ShortcutsWindow), nil
}

var shortcutsWindow *gtk.ShortcutsWindow

func (app *Application) ShowShortcuts() {
	var err error
	if shortcutsWindow == nil {
		shortcutsWindow, err = NewShortcutsWindow()
		if err != nil {
			log.Print(err)
			return
		}

		// We want to keep the about dialog around in case we want to
		// show it again.
		shortcutsWindow.HideOnDelete()
		shortcutsWindow.Connect("hide", func() {
			app.application.RemoveWindow(&shortcutsWindow.Window)
		})

	}

	// Set our parent window as the active window, but avoid
	// accidentally setting ourself as the parent window.
	win := app.application.GetActiveWindow()
	if win.Native() != shortcutsWindow.Native() {
		shortcutsWindow.SetTransientFor(win)
	}

	app.application.AddWindow(&shortcutsWindow.Window)
	shortcutsWindow.Present()
}
