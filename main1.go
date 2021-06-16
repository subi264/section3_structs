package main

func main() {
	myslice := []string{"hi", "hello", "bye", "you"}
	updateSlice(myslice)
	fmt.println(myslice)

}

func updateSlice(s []string) {
	s[0] = "Bye"
}
