package main

import "fmt"
import "github.com/gopherjs/gopherjs/js"
import nanomorph "github.com/sorribas/gopherjs-nanomorph"

type MyApp struct {
	*nanomorph.BaseApp
	Number int
}

func (app *MyApp) Render() string {
	return fmt.Sprintf(
		`<div id="app">Hey 
			<span>%d</span>
			<button onclick="app.IncNumber(2);">hey</button>
		</div>`,
		app.Number)
}

func (app *MyApp) IncNumber(n int) {
	app.Number += n
	app.ReRender()
}

func main() {
	app := MyApp{&nanomorph.BaseApp{}, 0}
	nanomorph.MountApp(&app)
	js.Global.Set("app", js.MakeWrapper(&app))
}
