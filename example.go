package main

import(
	"github.com/nicholasf/example/boundaries"
)

type ExampleApp interface {
	Houses() boundaries.Houses
	Users() boundaries.Users
}

type exampleApp struct{
	houses boundaries.Houses
	users boundaries.Users
}

func ExampleApp() ExampleApp {
	pim := factories.NewPIM()
	merchandising := factories.NewMerchandising()
	pim.SetMerchandising(merchandising)
	merchandising.SetPIM(pim)

	return &foxComm{
		merchandising: merchandising,
		pim: pim,
	}
}

func (f *exampleApp) Houses() boundaries.Houses {
	return f.houses
}

func (f *exampleApp) Users() boundaries.Users {
	return f.users
}