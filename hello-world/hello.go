package hello_world

import "fmt"

const helloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return helloPrefix + name
}

func main()  {
	fmt.Println(Hello("world"))
}


