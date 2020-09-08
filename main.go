package main

type app struct {
	config *AppConfig
}

func (a app) run() {
	newTokoDelivery(a.config).Create()
}

func newApp() app {
	config := newConfig()
	return app{config}
}

func main() {
	newApp().run()
}
