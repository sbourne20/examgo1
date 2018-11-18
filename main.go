package main

func main() {
	a := App{}
	// You need to set your Username and Password here
	a.Initialize("root", "@Bourne20", "bkd")

	a.Run(":8080")
}
