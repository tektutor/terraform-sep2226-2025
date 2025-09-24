package main

import "fmt"

func main() {

	//key,value can be of different data type
	toolsPath := map[string]string {
		"java_home": "/usr/lib/jdk11",
		"mvn_home" : "/usr/share/maven",
	}

	fmt.Println("Java home directory ", toolsPath["java_home"])

	//add a key,value pair into an existing map
	toolsPath["go_home"] = "/usr/go"

	//iterating a map and printing its values
	for key,value := range toolsPath {
          fmt.Println(key,value)
	}

	//delete a key-vaule pair from an existing map
	delete(toolsPath, "go_home")
	fmt.Println(toolsPath)
}
