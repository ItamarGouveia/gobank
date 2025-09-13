package main

func main() {

	server := NewAPISerever(":3000")
	server.Run()
}
