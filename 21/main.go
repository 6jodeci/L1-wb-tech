// Реализовать паттерн «адаптер» на любом примере.
package main

// Target предоставляет интерфейс, с которым должна работать система
type Target interface {
	Request() string
}

// Adaptee реализует адаптируемую систему.
type Adaptee struct{}

// NewAdapter - конструктор адаптера.
func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

// SpecificRequest -это реализация конкретного запроса
func (a *Adaptee) SpecificRequest() string {
	return "Request"
}

// Адаптер реализует интерфейс Target и является адаптером.
type Adapter struct {
	*Adaptee
}

// Request - это адаптивный метод.
func (a *Adapter) Request() string {
	return a.SpecificRequest()
}
