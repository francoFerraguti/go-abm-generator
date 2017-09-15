package templates

func MainGo() string {
	return `package main

&&IMPORTS&&

func main() {
	dbhandler.ConnectToDatabase()
	router.ConfigureRouter()
	router.CreateRouter()
	router.RunRouter()
}
`	
}