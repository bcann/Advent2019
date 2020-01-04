package main

import (
    "fmt"

)

/*
Build a map where all objects are keys and the object they directly orbit is the value
*/
func main() {
   input := []int{2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,1,0,2,2,2,2,2,2,2,2,2,0,2,0,0,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,0,2,2,1,1,2,2,0,2,2,2,1,0,2,1,2,2,2,0,2,2,2,0,2,1,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,0,2,1,2,2,2,1,0,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,0,0,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,0,2,2,1,1,2,2,0,2,2,2,1,0,2,2,2,2,2,2,2,2,2,0,2,0,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,1,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,1,2,2,2,2,2,2,2,2,2,1,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,0,2,2,2,2,2,1,2,2,2,2,2,0,2,2,2,2,2,2,2,1,2,2,2,2,2,2,0,2,2,2,2,2,2,0,2,2,2,1,2,2,2,2,2,1,2,2,0,2,2,2,2,2,2,1,2,2,2,2,2,2,2,0,0,1,2,2,2,2,0,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,1,1,2,2,2,2,2,2,2,2,2,0,2,2,1,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,0,2,2,0,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,0,0,2,2,0,0,2,2,0,1,2,1,2,2,2,2,2,2,2,0,2,2,2,2,0,2,2,2,2,2,2,0,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,0,1,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,0,2,2,1,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,1,1,2,2,0,2,2,2,1,2,2,0,2,2,2,1,2,2,2,2,2,0,2,2,0,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,0,2,2,2,2,2,1,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,0,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,1,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,0,1,2,2,1,0,2,2,0,1,2,2,2,2,2,0,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,1,1,2,2,2,2,2,1,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,1,2,2,1,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,0,2,2,1,1,2,2,0,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,0,2,2,2,2,2,2,2,0,0,0,2,2,2,0,1,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,0,2,1,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,1,0,2,2,2,1,2,2,2,1,2,2,1,2,2,2,0,2,2,2,1,0,2,1,2,1,2,0,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,0,2,2,1,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,1,2,1,1,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,0,2,2,1,2,2,2,2,2,0,2,2,2,2,2,0,2,2,2,0,2,2,2,0,2,2,0,1,2,2,2,2,2,2,1,0,2,2,2,0,2,2,2,2,2,1,2,2,2,2,1,2,2,2,2,2,2,1,2,1,2,2,2,2,2,1,1,2,2,2,1,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,1,0,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,0,2,2,2,2,1,0,2,2,0,0,2,2,2,1,2,2,0,0,2,2,2,1,2,2,0,0,2,2,2,0,2,1,2,2,2,2,2,0,2,2,0,2,2,2,2,2,2,0,0,0,2,2,2,2,2,1,1,1,2,2,0,1,1,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,0,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,0,2,2,2,2,1,1,2,2,0,0,2,2,2,2,2,2,2,1,2,2,2,0,2,2,0,1,2,2,2,2,2,1,2,2,2,2,2,0,2,2,0,2,2,2,2,2,2,0,0,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,0,1,2,2,2,2,2,2,2,2,2,1,2,0,1,2,0,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,0,2,2,1,2,2,2,2,2,1,0,2,2,2,0,0,2,2,2,1,2,2,2,1,2,2,0,2,2,2,0,0,2,2,1,0,2,1,2,0,2,1,2,2,2,2,2,0,2,2,0,2,2,2,2,2,2,1,0,0,2,2,2,2,2,1,1,1,2,2,2,1,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,1,2,2,2,2,2,2,2,2,2,2,1,2,0,1,2,1,1,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,1,2,2,2,2,2,2,2,2,1,1,2,2,2,0,1,2,2,2,1,2,2,2,2,2,2,2,0,2,2,0,1,2,2,0,1,2,0,2,2,2,2,2,2,2,0,2,0,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,0,2,1,2,2,1,0,0,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,1,1,2,2,2,2,2,2,2,2,2,2,2,1,1,2,0,2,2,2,2,2,2,1,2,2,2,2,1,2,2,2,2,2,2,2,1,2,2,2,2,2,1,0,2,2,2,1,2,2,2,2,0,2,2,2,0,2,2,1,1,2,2,0,2,2,2,1,1,2,1,2,2,2,0,2,2,2,1,2,2,2,2,0,2,2,2,2,2,2,0,0,2,2,2,2,2,2,1,1,2,2,2,2,0,1,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,1,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,1,2,2,2,2,2,2,2,2,2,0,2,1,2,2,0,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,0,2,2,1,2,2,2,2,2,2,2,2,2,2,0,2,2,2,0,1,2,2,2,0,2,2,2,2,2,2,0,0,2,2,1,1,2,0,2,1,2,0,2,2,2,1,2,0,2,2,0,2,2,2,2,2,2,1,0,1,2,2,2,2,2,0,0,2,2,2,2,0,1,2,2,1,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,0,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,1,1,2,2,2,2,2,2,2,2,2,1,2,1,0,2,2,2,2,2,2,2,1,1,2,2,2,2,2,2,2,2,2,1,2,2,1,2,2,2,2,2,2,2,2,2,2,2,0,2,2,1,2,2,2,2,1,2,2,2,0,2,2,0,1,2,2,1,1,2,1,2,2,2,2,2,2,2,0,2,2,2,2,0,2,2,2,2,2,2,2,1,0,2,2,2,2,2,2,2,2,2,2,1,0,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,1,2,2,2,2,2,2,2,2,2,2,1,2,2,1,2,1,1,2,2,2,2,0,2,2,2,2,2,0,2,1,2,2,0,2,2,1,2,2,2,2,2,0,0,2,2,2,2,1,2,2,2,1,2,2,2,1,2,2,0,2,2,2,0,1,2,2,0,2,2,0,2,2,2,1,2,2,2,2,2,0,2,2,1,2,2,2,2,2,2,1,2,0,2,2,2,2,2,1,0,2,2,2,2,0,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,0,2,2,0,1,2,2,2,2,1,1,2,2,2,2,1,2,0,2,2,0,2,2,0,2,2,2,2,2,1,0,2,2,2,1,1,2,2,0,0,2,2,2,0,2,2,2,1,2,2,2,2,2,2,1,2,2,0,2,0,2,2,2,2,2,0,2,0,2,2,0,2,2,2,2,2,2,2,2,0,2,2,2,2,2,0,2,1,2,2,2,0,1,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,1,1,2,2,2,2,2,2,2,2,2,0,2,1,0,2,2,1,2,2,2,2,2,0,2,2,2,2,0,2,1,2,2,1,2,2,2,2,2,2,2,2,0,0,2,2,2,2,1,2,2,1,2,2,2,2,0,2,2,1,2,2,2,1,0,2,2,0,1,2,0,2,1,2,2,2,2,2,1,2,0,2,2,0,2,2,2,2,2,2,0,1,0,2,2,2,2,2,2,0,0,2,2,2,2,1,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,1,2,2,2,2,2,2,2,2,2,1,2,2,2,2,1,2,2,2,2,2,0,0,2,2,2,2,0,2,0,2,2,2,2,2,2,2,2,2,2,2,0,0,2,2,2,0,2,2,2,0,0,2,2,2,2,2,2,0,1,2,2,0,1,2,2,2,2,2,2,2,2,2,0,2,2,2,0,2,0,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,1,1,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,0,1,2,2,2,2,2,2,0,2,2,0,2,2,0,2,0,0,2,2,2,2,0,0,2,1,2,2,0,2,2,2,2,0,2,2,1,2,2,2,2,2,2,0,2,2,2,0,1,2,2,1,2,2,2,2,0,2,2,0,2,2,2,1,0,2,2,2,1,2,1,2,2,2,0,2,2,2,0,2,1,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,1,0,1,2,2,1,0,2,2,0,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,1,2,2,2,2,2,2,2,2,2,2,2,1,1,2,0,2,2,2,2,2,2,1,2,2,2,2,0,2,2,2,2,2,2,2,1,2,2,2,2,2,0,0,2,2,2,1,2,2,2,0,1,2,2,2,0,2,2,2,0,2,2,0,0,2,2,2,1,2,2,2,1,2,1,2,2,2,1,2,1,2,2,2,2,2,2,2,2,2,0,0,1,2,2,2,2,2,1,2,0,2,2,1,0,2,2,2,1,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,0,0,2,2,2,2,2,2,2,2,2,2,2,1,1,2,2,2,2,0,1,2,0,0,2,0,2,2,0,2,1,2,2,0,2,2,0,2,2,2,2,2,2,2,2,2,2,2,0,2,2,1,1,2,2,2,1,2,2,0,1,2,2,1,2,2,2,2,1,2,2,2,0,2,1,2,0,2,1,2,1,2,2,1,2,2,2,2,2,2,1,1,2,2,2,2,2,2,2,0,1,2,2,0,2,2,2,1,1,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,1,2,2,2,2,2,2,2,1,2,2,1,2,0,0,2,1,1,2,0,0,2,1,2,2,1,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,1,2,2,1,1,2,2,2,0,2,2,0,0,2,2,0,1,2,2,1,0,2,2,2,0,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,1,2,2,2,2,2,1,2,0,2,2,2,2,1,2,2,2,2,2,2,1,2,2,2,2,2,2,1,2,2,2,2,2,1,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,0,0,2,1,1,2,0,2,2,1,2,2,2,2,2,1,2,0,2,2,0,2,2,1,2,2,2,2,2,2,2,2,2,2,1,1,2,2,2,0,2,2,2,2,2,2,0,1,2,2,0,2,2,2,0,1,2,2,2,0,2,0,2,1,2,2,2,0,2,2,1,2,2,2,2,2,2,1,2,2,2,2,2,2,2,0,2,0,2,2,1,2,2,2,0,1,2,2,2,0,2,2,2,2,2,2,1,2,2,2,2,2,1,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,0,2,2,2,2,2,0,2,2,1,2,2,2,2,0,2,1,2,2,1,2,2,1,2,2,2,2,2,0,1,2,2,2,0,2,2,2,0,1,2,2,2,0,2,2,0,1,2,2,1,2,2,2,1,0,2,2,2,0,2,0,2,0,2,0,2,1,2,2,0,2,2,2,2,2,2,0,1,0,2,2,2,2,2,1,2,1,2,2,0,0,1,2,0,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,1,2,2,2,2,2,2,0,2,2,0,2,1,1,2,2,0,2,2,1,2,0,0,2,0,2,2,1,2,1,2,2,1,2,2,2,2,2,2,2,2,1,1,2,2,2,2,1,2,2,0,0,2,2,2,1,2,2,0,0,2,2,1,2,2,2,1,0,2,0,2,1,2,0,2,2,2,0,2,0,2,2,0,2,2,2,2,2,2,1,0,1,2,2,2,2,2,0,1,2,2,2,2,1,0,2,1,0,2,2,2,1,2,2,2,2,2,2,1,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,0,2,2,1,2,0,2,2,1,2,2,2,1,2,2,2,2,1,2,2,1,2,0,2,2,0,2,2,0,2,2,2,2,2,0,1,2,2,2,2,2,2,2,1,2,2,2,2,1,2,2,1,2,2,2,1,1,2,2,1,0,2,0,2,0,2,1,2,2,2,1,0,0,2,2,0,2,2,2,2,2,2,2,1,0,2,2,2,2,2,2,0,1,2,2,0,2,0,2,1,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,1,2,1,0,2,2,2,2,2,2,1,2,2,2,2,0,2,2,1,0,2,2,2,2,2,0,2,1,2,2,1,2,0,2,2,2,2,2,1,2,2,2,2,2,0,1,2,2,2,0,0,2,2,1,1,2,2,2,2,2,2,2,0,2,2,0,0,2,2,2,0,2,1,2,1,2,2,2,0,2,0,0,0,2,2,1,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,0,2,2,0,1,1,2,2,1,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,0,0,2,2,2,2,1,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,1,2,1,1,2,2,2,2,2,2,2,2,2,0,2,0,2,2,2,1,2,1,2,2,0,2,2,2,2,2,0,2,1,2,2,1,2,2,0,2,2,2,2,2,1,1,2,2,2,1,2,2,2,2,0,2,2,2,1,2,2,2,0,2,2,1,1,2,2,0,1,2,0,2,1,2,0,2,0,2,0,1,0,2,2,0,2,2,2,2,2,2,2,0,2,2,2,2,2,2,1,0,1,2,2,0,2,2,2,1,0,2,2,2,1,2,2,2,2,2,2,1,2,2,2,2,2,2,1,2,2,2,2,0,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,0,2,1,1,2,0,2,2,2,2,2,2,2,2,2,2,1,2,1,2,2,0,1,2,2,1,2,2,2,2,0,2,2,2,2,1,2,2,2,2,2,2,2,2,0,2,2,2,2,0,2,2,2,0,1,2,2,2,2,2,2,1,1,2,2,1,1,2,2,0,2,2,2,2,1,2,0,2,2,2,1,1,0,2,2,1,2,2,2,2,2,2,1,0,0,2,2,2,2,2,0,1,0,2,2,2,2,0,2,2,1,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,1,0,2,2,2,2,1,2,2,2,1,2,2,2,2,2,2,2,2,2,1,2,2,0,2,1,2,2,1,2,2,2,2,1,2,2,2,2,2,1,2,0,2,2,2,2,2,1,2,2,1,2,2,1,2,2,2,2,2,2,2,0,2,2,2,2,2,2,1,2,2,2,2,2,2,2,1,2,2,2,2,1,2,2,0,0,2,2,0,1,2,2,1,2,2,1,2,0,2,1,2,2,2,0,0,2,2,2,0,2,2,2,2,2,2,1,0,2,2,2,2,2,2,0,1,1,2,2,2,1,2,2,1,0,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,1,2,2,2,2,0,2,2,2,0,2,2,2,2,2,2,2,2,2,1,2,2,0,2,1,1,2,1,2,2,2,2,1,2,2,1,2,2,0,2,2,1,2,2,0,2,1,2,2,2,2,2,2,2,1,2,2,0,2,2,2,2,2,2,2,2,2,1,2,2,2,1,1,1,2,1,0,2,2,2,1,2,2,2,0,2,2,2,0,2,2,2,0,2,0,2,0,2,0,2,2,2,1,0,2,2,2,2,2,2,2,2,2,2,2,1,0,2,2,2,2,2,0,2,0,2,2,0,1,1,2,0,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,1,2,2,0,2,0,2,2,0,2,2,2,2,1,2,2,2,2,0,2,2,1,2,2,0,0,2,1,1,2,2,2,2,1,2,0,2,2,0,2,2,1,2,2,2,2,2,1,0,2,2,2,0,1,0,2,1,2,2,2,2,0,2,2,0,1,2,2,0,0,2,2,2,0,2,1,2,1,2,1,2,1,2,0,0,2,2,2,1,2,2,2,2,2,2,2,0,1,2,2,2,2,2,1,0,2,2,2,0,2,1,2,0,0,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,1,0,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,1,2,1,2,1,2,2,1,2,2,0,2,0,0,2,0,0,2,0,0,2,0,2,2,2,2,2,2,2,0,2,2,0,2,2,0,2,2,2,2,2,1,1,2,2,2,2,1,0,2,0,1,2,2,2,2,2,2,1,0,2,2,0,2,2,2,0,0,2,0,2,0,2,0,2,2,2,0,0,0,2,2,0,2,2,2,2,2,2,0,1,2,2,2,2,2,2,2,2,2,2,0,1,2,0,2,1,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,0,2,1,0,2,2,2,2,2,2,1,2,2,1,2,1,0,2,2,1,2,0,1,2,2,2,2,2,2,2,0,2,1,2,2,2,2,0,1,2,2,2,2,2,1,1,2,2,2,0,1,1,2,0,1,2,2,2,0,2,2,0,2,2,2,1,0,2,2,0,0,2,2,2,1,2,0,2,2,2,1,0,0,2,2,0,2,2,2,2,2,2,1,0,1,2,2,2,2,2,1,1,1,2,1,0,1,1,2,1,0,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,0,2,2,2,1,2,2,2,2,2,2,2,2,2,1,2,2,0,2,2,0,2,1,2,2,2,2,0,2,2,2,2,0,2,2,0,1,2,0,2,2,2,0,2,0,2,2,0,2,0,2,2,1,2,0,0,2,2,2,2,2,1,1,2,2,2,2,1,2,2,1,2,2,2,2,1,2,2,1,2,2,2,0,0,2,2,2,0,2,1,2,2,2,1,2,2,2,0,1,2,2,2,0,2,2,2,2,2,2,0,2,1,2,2,2,2,2,1,0,2,2,0,0,1,2,2,2,1,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,1,0,2,2,2,2,1,2,2,2,0,2,2,2,2,2,2,2,2,2,1,2,2,2,2,0,2,2,2,2,2,2,2,1,2,2,1,2,2,2,2,2,1,2,1,0,2,1,0,2,0,2,2,0,2,2,2,2,0,2,1,0,2,2,2,2,2,2,2,2,2,2,1,0,1,2,0,0,2,2,2,0,2,2,2,1,2,2,0,2,2,2,2,0,2,2,2,0,2,0,2,2,2,2,1,1,2,2,2,2,2,2,2,2,2,1,1,1,2,2,2,2,2,0,0,2,2,1,2,2,1,2,0,2,2,2,2,1,2,2,2,2,2,2,1,2,2,2,2,2,0,2,2,2,2,2,1,2,2,2,1,2,2,2,2,2,2,2,2,2,1,2,2,0,2,1,1,2,2,2,0,2,2,2,2,2,0,2,0,1,2,1,0,2,2,1,2,0,2,2,1,2,2,2,2,1,2,2,1,2,0,0,2,2,2,2,2,2,1,2,2,2,2,1,0,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,0,2,1,2,2,2,0,2,2,2,0,0,2,2,2,0,2,2,2,2,2,2,1,1,2,2,2,2,2,2,2,1,2,2,0,2,2,0,2,2,2,2,2,2,1,2,2,2,2,2,2,1,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,0,2,2,1,2,2,2,2,0,2,0,2,2,1,2,2,1,2,1,1,2,2,2,2,1,2,2,1,1,2,1,2,2,2,2,2,2,2,0,2,1,1,2,2,2,2,2,2,0,2,2,2,2,2,0,2,1,1,2,2,2,0,2,2,2,1,2,2,0,1,2,2,2,0,2,1,2,1,2,1,2,1,2,1,0,0,2,1,1,2,2,2,2,0,2,2,2,1,2,2,2,2,2,2,1,1,2,1,1,1,0,2,0,0,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,1,1,2,2,2,2,0,2,2,2,0,2,2,2,2,2,2,2,2,2,1,2,2,2,2,1,0,2,2,2,1,2,2,1,2,2,0,2,1,2,2,1,1,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,0,2,2,2,2,2,2,2,1,2,2,2,0,2,2,2,2,0,2,2,2,1,2,2,0,2,2,2,0,1,2,2,0,2,2,0,2,0,2,2,2,0,2,1,0,1,2,0,1,2,2,2,2,2,2,1,0,0,2,2,2,2,2,1,1,0,2,2,0,1,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,0,2,1,2,2,2,1,2,2,1,2,2,2,2,1,1,2,2,1,2,0,2,2,1,1,2,0,2,2,1,2,2,2,2,2,2,0,0,2,2,2,2,2,0,2,2,2,2,1,1,2,2,0,2,2,2,2,0,2,2,1,1,0,2,1,1,2,2,1,2,2,2,2,1,2,2,2,2,2,2,0,1,2,0,2,2,1,2,2,1,2,1,1,0,2,2,2,2,2,0,1,1,2,2,1,1,0,2,1,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,0,1,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,1,2,2,1,0,0,1,2,0,0,0,2,2,0,2,2,2,2,2,0,2,1,0,2,2,0,2,1,2,0,1,2,2,2,2,1,2,2,1,2,0,2,2,2,2,2,2,0,0,2,2,2,2,0,1,2,0,0,2,2,2,1,2,2,0,2,1,1,1,2,2,0,0,2,2,1,2,0,2,1,2,2,2,2,1,0,2,1,1,2,1,2,2,0,2,2,1,0,2,2,2,2,2,2,0,0,2,2,1,2,2,2,1,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,0,2,2,2,1,2,2,2,2,2,2,2,2,2,1,2,1,1,1,0,2,2,1,2,0,2,2,0,2,2,0,2,0,0,2,1,0,2,0,2,2,1,1,0,0,2,2,2,2,2,2,2,0,2,2,2,2,2,2,0,2,0,0,2,2,2,2,1,0,2,0,0,2,2,2,0,2,2,1,1,1,1,0,0,2,1,1,1,2,2,2,2,2,1,2,0,2,2,2,2,2,0,0,2,2,2,2,0,2,2,1,2,2,2,2,2,2,2,0,0,2,1,1,2,1,2,2,0,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,0,1,2,2,2,2,0,2,2,2,1,2,2,2,2,2,2,2,2,2,1,2,0,1,0,2,1,2,0,2,0,2,2,1,2,2,2,2,1,0,2,1,1,2,0,0,2,2,2,0,2,2,2,0,2,1,2,2,0,2,0,0,2,2,2,1,2,2,2,2,2,2,0,2,1,2,1,0,2,2,2,2,2,2,1,2,1,1,2,1,2,0,1,0,2,2,2,1,2,1,2,1,2,2,1,2,2,1,0,2,1,2,2,2,2,0,0,1,2,2,2,2,2,0,1,1,2,0,1,1,0,2,1,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,0,0,2,0,2,0,2,2,2,2,2,1,2,1,0,2,2,0,2,1,1,2,0,2,2,2,2,2,0,2,0,2,2,2,2,1,1,2,2,2,1,2,1,0,2,2,2,1,0,2,2,0,0,2,2,2,1,2,2,0,1,2,2,2,2,2,1,0,2,2,1,2,1,2,2,2,0,2,1,0,2,2,1,0,2,0,2,2,0,2,0,1,2,2,2,2,2,2,0,0,1,2,0,0,0,1,2,0,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,1,0,1,0,0,2,1,0,1,2,2,0,2,2,1,2,2,1,2,2,1,2,0,2,2,1,1,0,0,2,2,0,2,2,2,2,0,2,2,1,2,2,2,0,2,0,1,2,2,2,1,0,2,2,1,2,2,2,2,0,2,2,0,2,0,2,0,1,2,1,0,2,2,1,2,1,2,2,2,1,2,2,1,2,2,2,2,2,0,2,2,2,2,2,1,2,2,2,2,2,2,0,0,0,2,0,1,1,1,2,1,2,2,2,2,0,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,1,0,0,0,1,2,1,2,2,2,2,1,2,2,1,2,1,1,2,2,1,2,2,1,2,2,2,2,0,2,2,1,2,1,2,2,0,2,1,2,2,2,2,2,2,0,0,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,0,1,0,2,0,0,2,2,0,0,2,2,2,1,2,2,2,2,2,1,2,1,2,2,0,2,2,2,2,2,2,0,1,1,2,2,2,2,2,1,0,2,2,1,2,1,0,2,2,0,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,0,2,2,2,1,2,2,2,2,2,0,2,2,2,2,2,1,2,1,2,2,2,2,0,2,2,2,1,2,2,1,2,0,0,2,2,0,2,2,2,2,2,1,0,0,2,2,1,2,0,2,2,2,2,1,0,2,0,2,2,2,0,1,2,2,2,2,1,2,2,2,1,2,2,2,0,2,2,2,2,2,1,0,2,2,1,1,0,2,1,2,2,2,0,2,2,2,2,0,0,2,2,1,2,1,2,2,1,2,2,2,1,2,2,2,2,2,0,0,0,2,2,2,1,2,2,1,2,2,2,2,0,2,2,2,2,2,2,1,2,2,2,2,2,1,1,2,2,2,2,0,2,2,2,0,2,2,2,2,2,1,2,2,2,1,2,0,2,0,1,0,2,0,1,2,2,2,1,2,2,2,2,1,0,2,2,1,2,0,2,2,1,2,1,0,2,2,0,2,0,2,2,2,2,2,2,2,0,2,0,2,0,2,2,2,2,0,1,0,2,2,1,2,2,2,0,2,2,1,0,1,0,2,2,2,1,1,0,2,1,2,0,2,0,2,2,2,0,2,2,2,2,1,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,1,1,2,0,2,2,0,2,1,1,2,2,2,0,2,2,2,2,2,2,1,2,2,2,2,2,1,0,2,2,2,2,1,2,2,2,0,2,2,2,2,2,0,2,2,2,2,2,2,0,2,0,0,2,1,2,0,2,2,1,2,2,2,2,2,0,2,1,1,2,0,0,2,1,2,2,1,2,2,0,2,0,2,2,1,2,1,1,2,1,2,2,2,0,1,2,2,2,2,2,0,2,2,1,2,2,2,1,2,2,1,0,1,0,2,0,2,2,1,1,2,2,2,1,2,0,2,2,2,2,1,0,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,1,2,2,1,2,2,2,1,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,0,2,2,2,1,2,0,1,2,0,0,2,2,2,2,2,2,1,2,2,1,2,0,0,2,2,1,2,1,1,2,1,2,2,1,2,2,2,2,0,2,2,0,2,2,1,2,2,2,1,2,2,2,2,2,2,2,1,0,2,0,2,2,2,2,1,2,2,2,2,0,1,2,2,2,0,0,0,2,2,2,1,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,0,2,1,2,2,2,2,2,2,2,0,2,2,2,0,1,2,2,2,0,2,2,2,2,2,2,2,2,0,2,2,1,2,2,2,2,2,1,1,2,2,2,2,1,2,2,2,1,2,2,2,2,2,1,2,2,2,1,2,2,2,2,1,0,2,1,0,2,2,2,0,2,2,2,2,2,1,2,2,1,2,2,0,2,0,2,1,2,2,2,2,2,2,2,2,0,2,2,1,2,2,0,2,2,2,1,2,2,2,1,1,1,2,0,1,2,2,2,0,2,2,1,1,2,2,1,0,2,0,0,2,0,1,2,2,2,1,2,2,2,0,2,0,2,2,0,2,1,2,2,0,2,2,1,0,2,2,2,2,2,2,1,2,2,1,2,1,2,2,1,2,2,2,2,1,2,2,2,2,2,2,0,2,2,2,2,2,0,0,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,1,1,2,2,2,2,1,0,0,2,2,1,2,2,2,2,2,2,2,2,1,2,2,1,2,2,0,2,1,2,2,2,2,1,2,2,0,2,2,1,2,1,2,1,2,0,2,2,2,2,0,1,1,2,2,0,2,2,2,2,2,2,1,1,1,0,0,0,2,0,2,2,1,2,2,2,2,0,2,2,2,1,2,2,2,0,2,2,2,2,2,1,2,1,0,2,2,2,2,2,2,1,2,2,0,1,2,0,0,2,0,0,2,1,2,2,2,2,2,1,2,2,2,2,2,2,2,2,0,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,1,2,2,2,1,2,1,0,0,0,1,2,1,0,2,2,2,1,2,2,2,2,2,1,2,0,0,2,0,0,2,0,2,2,1,2,2,1,2,1,2,2,0,2,0,1,2,1,1,1,2,0,0,2,2,2,2,2,2,2,0,1,2,2,2,1,2,2,1,2,0,1,0,0,2,1,0,0,1,0,2,1,2,0,2,1,2,1,0,1,2,2,2,2,2,2,2,1,2,2,0,2,2,2,2,2,2,0,1,1,0,2,2,1,2,2,1,1,2,1,2,1,2,2,2,0,2,2,1,2,2,2,2,2,1,2,2,2,2,2,1,2,2,2,1,2,2,2,2,2,2,2,2,2,0,2,0,1,1,1,0,2,2,1,0,2,2,0,2,2,1,2,2,1,2,2,1,2,2,0,2,2,0,0,2,2,2,2,2,1,2,2,2,2,2,2,2,0,1,0,2,1,0,2,2,2,2,0,2,2,2,1,2,0,2,1,2,2,0,1,1,1,1,0,2,0,2,1,2,1,2,0,2,0,2,0,2,1,0,2,2,1,1,2,1,0,2,1,2,1,2,2,2,2,2,2,2,0,0,1,0,2,0,2,2,2,1,2,2,1,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,0,2,2,2,2,0,2,2,2,2,2,2,2,2,2,0,2,2,2,1,2,0,0,1,0,0,2,2,1,2,2,2,2,2,2,0,2,1,2,2,0,2,2,1,0,2,1,1,2,2,2,2,1,2,2,2,2,0,2,0,2,2,0,1,0,2,0,2,2,2,2,0,2,2,2,1,2,2,2,2,1,1,2,1,0,0,1,1,1,2,0,2,0,0,1,2,1,2,1,2,2,2,2,1,1,2,2,2,2,1,0,2,2,2,1,0,0,2,2,2,2,2,0,2,2,1,2,0,2,2,2,2,1,2,2,2,0,2,2,2,1,2,2,1,2,2,2,2,2,1,1,2,2,2,2,0,2,2,2,0,2,2,2,2,2,2,2,2,2,1,2,0,1,2,2,2,2,2,1,2,2,2,1,2,2,0,2,1,0,2,0,2,2,1,0,2,1,1,1,1,2,2,2,2,0,2,2,1,2,2,0,2,1,0,2,2,0,0,2,2,2,1,1,2,2,1,1,2,2,2,1,2,2,0,0,0,1,0,2,2,2,0,1,2,0,2,1,2,2,2,0,2,1,0,2,2,2,0,2,2,0,2,0,2,1,0,1,2,2,2,2,2,0,0,0,2,2,2,0,1,2,1,0,2,1,2,0,2,2,2,1,2,2,2,2,2,2,2,2,0,0,2,2,2,2,0,2,2,2,1,2,2,2,2,2,0,2,2,2,2,2,1,2,2,0,0,2,0,0,2,2,2,0,2,2,2,2,1,2,2,1,0,2,1,0,2,2,1,2,1,2,2,2,2,2,2,2,0,2,2,2,2,2,1,0,2,2,2,2,2,2,2,1,0,2,1,0,2,2,2,1,2,2,1,1,0,1,0,2,2,1,1,1,2,2,2,1,2,2,2,0,2,2,0,0,2,0,2,2,2,2,2,0,2,0,2,0,2,2,2,2,2,0,2,1,0,2,2,1,2,2,0,0,2,1,2,0,2,2,2,1,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,1,2,2,1,0,2,0,1,2,2,2,0,2,2,0,2,2,1,2,1,0,2,1,0,2,2,0,0,0,2,2,1,2,0,2,2,0,2,0,2,2,0,1,1,2,2,2,2,2,2,0,1,2,2,0,0,2,0,2,1,2,2,1,0,1,2,1,2,2,1,0,0,0,0,2,2,2,0,2,1,2,2,0,0,2,0,1,2,0,0,2,1,2,1,1,2,2,2,2,2,2,0,0,1,1,0,1,0,0,2,2,1,2,2,2,1,2,2,2,1,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,1,2,2,2,1,2,1,0,2,2,1,2,0,2,2,2,2,0,2,2,0,2,2,0,1,0,2,2,0,0,2,1,2,0,2,2,2,1,2,2,2,2,0,2,1,1,2,2,1,2,2,1,2,2,2,2,2,1,1,2,1,0,2,2,2,1,0,2,2,1,2,0,1,0,2,0,2,1,1,1,2,0,2,1,2,0,2,2,1,0,2,0,1,2,0,0,2,1,2,1,1,1,2,2,2,2,2,0,1,2,1,2,0,2,2,2,2,0,2,0,2,2,2,2,2,1,2,2,0,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,1,2,2,2,2,2,1,2,2,2,2,2,2,1,1,1,1,2,2,2,2,2,2,2,2,2,1,2,2,1,0,0,0,2,1,1,2,2,2,0,2,2,2,2,2,1,2,2,0,2,0,2,2,1,2,2,2,2,2,2,2,2,1,2,1,2,1,1,2,0,2,1,0,2,2,1,1,1,2,2,2,2,2,2,0,2,2,1,2,1,2,0,2,0,0,2,2,2,0,2,0,2,1,1,2,0,2,1,2,2,2,2,2,2,0,0,1,0,0,1,0,2,1,1,2,0,2,2,2,2,2,0,2,2,2,2,2,2,2,2,0,1,2,2,2,2,1,2,2,2,2,2,2,2,2,2,0,2,2,2,0,2,0,2,2,0,0,2,1,1,0,2,2,1,2,2,0,2,1,2,1,1,2,2,1,2,2,0,2,1,1,2,2,1,2,2,2,2,2,2,2,1,2,1,2,0,2,0,0,2,2,2,2,1,2,2,1,0,2,0,2,0,2,2,1,0,1,0,2,0,2,0,2,1,0,2,2,0,2,2,2,0,2,0,2,1,2,2,2,2,0,0,1,2,2,2,1,2,2,2,2,2,2,0,2,1,1,2,2,0,1,2,0,2,2,2,2,0,2,2,2,1,2,2,0,2,2,2,2,2,2,1,2,2,2,2,0,2,1,2,1,2,2,2,2,2,0,2,2,2,2,2,0,1,0,0,1,2,2,1,2,0,2,0,2,2,0,2,1,1,2,1,1,2,2,0,2,1,0,0,2,2,2,0,2,2,2,2,0,2,0,0,2,2,0,0,2,1,2,2,2,2,2,2,2,2,1,0,2,0,2,2,1,2,2,0,2,0,2,1,2,1,2,2,2,2,2,2,2,2,2,0,2,2,1,1,2,1,1,2,2,1,0,2,2,0,1,0,2,2,2,2,2,1,1,2,2,0,2,0,0,2,2,0,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,0,2,2,2,0,2,2,1,0,1,1,2,0,1,1,2,2,0,2,2,2,2,0,0,2,2,0,2,0,1,2,0,0,1,2,2,2,1,2,2,2,2,2,2,0,2,2,1,2,1,2,2,2,2,2,2,1,2,0,2,1,2,2,1,2,0,2,2,0,1,2,1,1,1,2,1,1,0,0,1,2,0,2,0,2,0,2,2,2,1,2,2,2,2,1,1,1,1,2,2,2,0,2,2,2,2,2,1,2,2,0,0,0,1,2,2,1,1,2,0,2,1,2,2,2,1,2,2,2,2,2,2,2,2,0,1,2,2,2,2,1,2,1,2,0,2,2,2,2,2,0,2,2,2,1,2,1,2,2,2,0,2,1,0,0,0,2,1,2,2,2,2,2,0,1,0,0,2,2,2,2,2,2,1,1,2,1,2,2,1,2,2,1,2,2,2,2,0,0,1,2,1,0,2,2,2,2,1,1,2,2,2,2,2,2,0,1,2,2,1,2,2,0,1,2,1,1,0,1,0,2,1,2,2,2,0,2,0,2,0,2,0,2,2,2,0,0,0,2,1,2,1,2,2,2,2,2,0,0,0,1,1,2,1,1,2,1,0,2,0,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,1,2,2,2,2,1,2,2,2,0,2,2,2,2,2,0,2,2,2,0,2,1,0,2,1,2,2,1,0,0,2,1,1,2,2,2,2,1,2,2,0,1,2,0,1,2,2,1,0,2,2,1,1,2,0,2,2,2,2,1,1,2,1,1,1,2,2,0,2,2,2,1,0,2,2,1,1,2,2,2,1,2,2,0,2,1,1,2,1,2,0,1,1,0,2,2,2,2,0,2,0,2,0,2,0,2,1,0,2,2,2,0,0,2,2,2,2,2,2,2,2,2,0,0,1,1,2,0,1,0,2,2,1,2,0,2,1,1,2,2,0,2,2,0,2,2,2,2,2,1,2,2,2,2,2,2,2,1,2,1,2,2,2,2,2,0,2,2,2,2,2,2,1,1,1,2,2,2,0,0,2,0,0,2,2,0,2,0,0,2,2,0,2,2,2,2,1,2,2,2,2,0,2,2,1,2,2,1,0,2,0,2,2,1,0,2,1,2,2,2,2,1,2,2,2,2,0,2,0,2,2,0,2,1,0,0,1,1,2,2,1,1,2,2,2,2,1,2,0,2,1,2,0,2,1,2,0,0,2,1,0,0,1,2,0,2,0,2,2,2,2,2,1,2,1,2,0,2,0,2,2,1,2,2,1,2,1,2,2,2,0,2,2,0,2,2,2,2,2,2,1,2,2,2,2,0,2,2,2,0,2,2,2,2,2,0,2,2,2,2,2,2,1,0,0,1,2,1,1,2,0,0,0,2,2,2,2,1,0,2,2,0,0,1,0,2,2,0,0,0,2,1,2,2,1,2,2,0,1,1,1,2,2,1,2,2,1,1,2,1,2,1,0,2,2,1,0,2,0,2,1,2,2,0,2,0,1,0,2,2,2,1,0,0,0,2,2,2,0,2,2,2,2,2,2,2,1,2,2,0,2,2,1,2,1,0,0,0,2,2,2,2,2,1,1,0,2,0,0,1,2,0,1,2,1,2,0,2,2,2,0,2,2,0,2,2,2,2,2,2,2,2,2,2,2,1,2,1,2,2,2,2,2,2,2,0,2,2,2,1,2,1,0,2,0,0,2,1,2,2,2,0,0,2,2,0,2,1,2,2,2,0,2,0,2,2,1,1,0,2,2,2,0,2,2,2,2,0,1,0,1,2,0,0,2,2,2,2,2,1,2,0,0,0,2,0,0,2,0,2,0,1,2,0,1,1,2,1,0,2,1,1,0,2,1,2,0,2,0,2,0,2,0,1,1,2,2,1,2,1,1,1,1,2,1,1,2,0,2,2,2,2,2,2,1,1,1,1,1,1,2,1,2,2,2,2,1,2,2,2,1,2,2,2,2,0,2,2,2,1,1,2,2,2,2,0,2,1,2,1,2,0,2,2,0,2,2,2,2,1,2,2,0,0,2,0,2,2,2,1,1,2,2,2,2,1,2,0,0,0,1,0,2,2,0,2,0,0,2,1,2,0,0,2,1,2,2,1,2,1,2,2,0,0,2,2,1,2,2,2,2,2,2,2,2,1,1,2,1,2,1,2,2,2,2,0,1,0,1,2,1,1,1,0,0,2,0,2,1,2,0,2,1,1,1,2,2,1,2,2,2,1,1,2,0,2,2,2,2,2,2,2,1,1,1,0,1,1,1,2,2,1,0,2,0,2,0,1,2,2,2,2,2,0,2,1,2,2,2,1,2,2,2,2,2,0,2,1,2,2,0,1,2,2,0,2,2,2,2,0,2,1,2,1,0,0,2,1,2,1,2,0,2,2,2,0,2,2,0,1,0,1,2,0,0,2,2,2,0,1,2,0,2,2,0,2,2,2,2,1,0,2,0,0,0,2,0,1,2,0,2,1,2,2,2,0,1,2,0,2,1,0,2,0,0,0,0,2,2,2,0,1,0,2,2,2,2,2,2,2,1,2,2,0,2,2,1,2,2,0,1,0,2,2,0,2,1,2,2,2,2,2,1,1,1,2,2,2,2,0,2,1,0,2,2,2,0,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,0,0,0,2,2,0,2,2,2,2,2,2,1,2,2,0,0,2,0,0,1,0,2,0,2,2,1,2,2,2,0,2,0,1,0,0,2,1,0,0,2,2,0,0,2,1,2,2,0,0,0,0,2,2,1,2,2,1,1,2,0,2,2,1,0,2,2,1,2,2,2,0,2,2,1,0,1,0,1,0,2,1,2,1,2,2,2,1,2,2,2,0,2,0,1,1,2,0,1,2,2,0,1,1,2,0,0,2,0,2,2,2,2,2,2,1,1,0,2,2,0,2,1,2,2,0,2,2,2,2,2,1,2,2,0,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,1,2,2,0,2,2,2,2,2,2,0,1,0,2,0,2,0,0,1,0,0,0,2,2,2,2,2,2,2,1,2,0,1,0,2,0,0,1,2,2,0,2,2,0,2,2,0,1,0,0,2,0,1,0,2,2,1,2,2,2,1,0,1,2,0,2,2,1,2,1,2,2,0,2,2,0,0,1,2,0,0,2,2,1,2,0,2,1,2,2,2,0,1,0,2,0,1,2,2,1,1,0,2,0,1,1,1,2,2,2,2,2,2,2,0,0,0,2,1,2,0,0,2,1,2,0,0,2,2,2,2,2,0,2,2,2,2,2,0,1,2,2,2,2,0,2,2,2,1,2,0,2,2,1,0,2,2,2,1,2,0,1,2,2,1,2,2,1,2,1,1,1,2,2,1,2,0,2,0,0,0,2,1,0,2,1,2,1,2,2,0,1,2,1,2,2,1,0,0,2,2,0,0,0,2,0,2,2,2,2,2,0,1,2,2,1,2,2,2,1,2,2,0,1,2,1,1,2,2,0,2,1,2,2,2,0,2,1,2,1,2,1,2,0,2,0,1,2,0,1,2,2,2,0,2,1,0,2,2,2,2,0,2,1,2,2,2,0,1,2,2,2,2,0,2,2,1,2,2,0,2,2,0,2,1,2,0,2,0,1,2,2,2,2,2,2,1,2,2,0,1,2,2,2,0,2,2,2,0,2,0,1,2,2,0,2,0,2,1,2,2,1,2,2,1,2,1,2,0,2,0,1,0,2,2,0,1,2,2,2,2,2,2,2,2,1,1,1,1,2,2,0,2,0,0,1,0,2,2,2,2,2,0,2,1,0,2,1,2,0,0,2,2,1,2,1,2,2,2,1,0,0,1,1,2,2,2,1,2,1,2,0,2,2,2,0,1,2,2,1,0,1,2,2,0,2,2,2,2,2,2,1,2,2,0,0,2,2,2,2,2,2,2,1,2,1,2,2,2,2,2,2,2,2,1,2,1,2,2,1,2,2,2,2,0,2,1,2,0,1,1,2,2,0,1,2,2,2,0,2,0,1,1,1,0,2,0,1,0,1,1,1,2,2,0,2,0,2,1,0,0,1,2,0,2,0,1,0,0,2,0,0,2,2,2,0,0,0,0,2,2,2,1,1,0,0,1,2,1,2,0,1,0,2,2,2,2,0,2,0,1,2,1,1,2,1,1,2,2,2,1,0,2,1,2,1,2,1,2,2,2,0,2,2,2,2,1,2,0,2,1,1,2,0,0,0,0,2,2,2,1,0,0,2,1,1,2,0,2,2,1,2,2,1,2,0,1,2,2,0,2,2,2,2,1,2,1,2,2,1,2,2,2,2,2,2,0,2,1,2,1,2,2,0,1,2,2,2,0,2,1,1,0,0,2,2,0,0,2,1,0,0,2,2,2,2,0,2,1,0,0,0,1,0,2,2,1,1,1,2,2,0,2,2,2,0,1,2,1,0,2,2,1,1,1,0,2,2,0,2,0,1,2,2,2,0,2,2,2,0,0,2,2,0,1,2,0,1,2,0,1,0,0,0,2,1,2,1,2,1,2,1,0,0,2,0,1,2,2,1,0,1,2,2,1,1,0,2,2,0,2,2,2,1,2,2,2,2,2,2,2,0,2,0,2,1,1,2,2,1,2,2,2,2,0,2,2,2,1,2,2,2,2,2,1,2,2,2,0,0,2,2,2,0,0,2,2,2,2,2,1,0,2,0,1,2,0,2,1,1,2,0,2,2,0,2,0,1,0,0,2,0,2,0,2,1,0,2,2,2,1,1,2,0,2,0,1,1,1,2,2,1,2,2,2,1,0,2,2,2,0,1,1,2,0,0,2,0,2,1,0,2,0,2,2,0,0,1,2,2,1,1,2,2,2,2,2,2,2,1,2,2,2,1,2,1,0,2,0,0,2,1,2,2,2,2,2,2,2,1,0,0,1,2,2,1,0,2,1,2,0,0,2,0,2,2,1,2,2,2,2,2,1,2,2,2,2,2,0,0,2,2,2,2,2,2,1,2,1,2,1,2,2,1,0,2,2,2,2,2,2,1,2,0,0,2,1,2,2,0,0,2,2,2,1,2,2,1,1,1,0,1,0,2,2,0,2,2,2,2,1,1,2,1,2,2,1,0,1,0,2,1,2,0,0,2,1,2,1,2,1,2,2,2,2,1,2,1,2,1,0,2,1,2,1,0,1,1,2,2,0,1,0,0,0,2,2,1,2,0,2,2,1,1,2,1,2,2,0,2,0,1,2,0,1,2,0,2,2,1,2,1,0,2,0,0,0,0,1,2,0,1,2,2,2,1,0,2,2,0,2,2,1,2,2,2,0,0,0,1,2,2,2,2,1,2,2,2,2,1,0,2,2,0,2,2,2,2,0,2,1,0,1,1,1,2,0,0,1,2,1,1,2,2,2,2,1,2,1,2,1,0,2,1,2,1,1,0,0,2,2,0,2,1,2,0,0,2,2,1,2,2,2,2,0,0,1,2,1,2,0,1,1,2,1,0,2,0,2,1,1,2,2,1,0,2,0,0,2,2,0,1,2,2,1,2,2,0,2,2,2,2,0,1,2,1,1,2,2,0,1,0,2,2,1,0,2,2,2,2,2,2,2,1,0,1,0,1,2,2,0,0,2,0,2,1,0,2,2,2,2,2,1,2,2,2,2,1,0,1,2,2,2,2,1,2,1,2,1,0,2,2,2,0,1,2,2,2,2,2,0,0,0,1,0,2,1,1,2,0,1,0,2,2,1,2,1,1,2,0,0,0,0,0,2,1,0,2,0,2,1,0,2,0,2,2,0,2,0,0,2,0,1,0,2,1,2,2,0,2,0,0,1,2,1,0,2,0,2,1,0,2,0,2,2,0,1,2,2,1,1,1,1,1,2,0,2,1,2,1,2,0,1,2,2,1,2,2,2,0,1,1,2,0,2,0,1,2,2,2,1,2,1,1,1,0,1,0,0,2,0,2,2,1,2,2,0,2,2,1,2,2,2,2,1,2,0,1,1,2,2,2,2,2,2,2,2,2,0,2,0,2,2,2,1,2,2,2,2,2,0,1,1,0,2,2,2,0,1,1,0,0,2,2,0,2,2,2,1,0,2,1,2,2,2,1,1,1,2,2,2,2,2,2,2,1,0,2,0,2,2,0,1,0,1,0,2,2,1,2,0,2,0,2,0,2,2,2,2,0,0,2,1,2,0,1,2,0,2,2,0,2,1,2,0,0,2,2,2,0,2,2,1,0,2,1,0,2,0,1,1,2,2,2,0,0,0,2,2,2,2,2,2,0,2,1,0,2,2,2,0,2,2,2,2,0,2,2,2,2,2,2,0,2,1,2,0,0,0,2,2,2,2,2,0,2,1,2,0,2,0,2,2,0,1,2,0,2,0,2,0,0,0,0,1,2,0,2,1,2,0,2,2,2,1,2,1,1,0,0,2,0,2,2,2,2,2,2,0,2,2,1,2,1,2,1,1,2,1,2,2,0,0,0,0,0,0,2,2,2,1,0,0,2,1,1,2,0,2,1,0,2,2,1,1,0,2,2,2,2,1,2,1,1,2,1,2,0,2,1,2,2,0,1,2,0,2,2,1,2,0,1,2,1,2,1,0,2,2,2,2,2,1,1,1,2,0,0,1,2,0,1,2,1,2,1,0,2,0,0,2,2,1,2,0,2,2,1,0,1,2,2,2,2,1,2,2,2,2,0,1,2,2,2,0,2,0,2,1,2,1,1,0,2,2,2,1,2,2,0,0,2,2,2,1,2,1,0,1,2,2,1,1,1,2,2,1,2,0,2,2,0,2,1,2,0,0,2,0,1,1,2,1,0,0,0,1,2,0,2,2,1,2,2,2,0,2,2,2,0,1,2,1,0,1,1,1,0,2,2,1,0,1,1,2,1,2,1,2,2,2,0,1,2,2,1,2,2,0,1,1,1,2,1,0,0,0,2,2,0,2,2,0,0,1,2,2,2,0,2,2,1,2,1,2,0,1,2,1,1,2,2,2,2,2,2,2,1,0,2,2,2,2,2,1,2,1,2,2,1,0,2,2,1,1,2,2,2,0,2,2,0,1,1,2,2,1,0,1,2,0,1,2,2,2,2,0,2,1,0,1,1,0,2,2,0,2,0,0,2,0,0,2,0,2,0,2,2,0,0,0,2,0,2,1,0,0,2,0,2,2,1,2,2,2,1,2,1,2,1,1,2,0,1,0,0,2,0,2,0,1,1,0,0,2,2,2,2,2,2,2,1,0,2,2,0,0,2,2,2,1,0,2,1,2,2,2,2,2,1,2,1,1,0,1,1,0,0,2,2,1,0,2,1,2,2,2,2,0,0,2,2,2,2,2,2,1,1,2,1,1,2,2,2,0,2,0,2,0,2,1,2,2,0,0,2,2,2,0,2,0,2,2,2,0,2,2,0,1,0,0,2,2,2,2,2,2,2,0,1,1,0,0,0,2,1,0,2,0,2,0,0,2,0,2,1,0,0,2,0,2,1,1,2,0,1,0,2,2,2,0,2,1,2,2,1,2,0,2,1,2,2,1,1,0,1,0,1,2,0,1,1,2,1,2,2,2,1,2,2,2,1,2,2,2,0,1,2,2,1,2,1,2,2,0,0,2,2,2,2,0,2,1,2,2,0,1,0,2,2,0,2,2,1,2,2,1,2,0,2,2,2,0,2,0,2,2,0,0,1,0,2,2,2,0,2,2,2,0,2,1,2,2,1,1,2,0,2,0,2,0,0,1,2,0,2,2,1,1,2,1,2,2,2,0,2,0,1,2,2,2,1,2,0,2,2,2,1,0,2,2,1,2,2,2,0,1,1,0,2,1,2,2,0,2,2,2,2,0,2,1,2,2,2,0,2,2,2,2,2,1,2,1,2,2,2,0,2,2,0,2,1,0,0,2,0,2,0,2,0,2,2,0,1,2,1,0,2,0,0,0,1,2,0,1,2,2,2,2,2,1,0,0,2,0,0,0,0,0,2,1,1,1,1,0,1,1,2,2,2,2,2,0,2,0,2,1,0,1,2,0,2,2,2,2,2,1,2,0,2,0,2,2,2,2,2,1,2,2,2,0,0,0,1,2,2,2,0,2,0,0,2,2,2,0,2,2,1,0,0,1,0,0,1,2,0,0,2,2,2,1,2,2,0,2,1,1,0,2,0,1,0,2,0,2,0,2,0,0,2,0,1,1,2,1,1,2,1,2,2,0,0,1,1,1,2,1,0,2,2,1,0,1,0,2,2,2,2,2,0,2,0,2,1,2,1,2,2,0,2,0,1,2,0,2,1,0,2,2,1,1,0,1,2,1,2,2,1,1,2,0,0,2,2,1,0,0,2,2,1,2,2,1,2,1,2,0,1,0,2,2,2,2,2,1,2,2,2,1,1,1,2,2,1,2,2,2,2,0,2,2,0,1,2,1,2,1,2,1,1,0,0,2,2,0,2,0,1,2,0,0,0,2,0,2,2,2,0,0,2,2,0,2,2,2,0,2,1,0,1,2,1,2,1,0,0,1,2,2,2,0,1,0,2,0,1,2,0,2,2,2,0,0,2,2,2,1,2,2,1,2,0,1,2,1,1,2,1,2,2,2,2,2,2,2,1,1,2,2,0,1,2,2,2,2,0,2,2,2,0,2,1,2,0,1,0,0,2,0,2,1,2,2,2,2,1,0,2,1,0,2,2,0,2,0,2,0,2,2,1,2,2,2,2,2,2,0,2,0,0,0,2,2,2,1,2,0,2,2,2,2,0,2,0,1,2,1,2,0,1,2,0,2,2,1,2,1,1,0,0,0,0,0,1,2,2,0,2,0,2,1,2,2,0,2,0,1,2,2,0,0,0,0,1,0,2,2,2,0,2,2,2,0,2,2,0,2,2,2,1,2,2,1,2,2,1,0,1,2,1,1,2,1,0,2,1,2,0,2,1,2,2,0,1,2,2,0,2,1,0,1,1,2,1,1,1,0,2,2,2,2,2,0,1,2,0,2,0,1,2,2,2,0,1,1,0,2,2,1,2,2,2,2,2,1,2,1,1,0,2,1,2,2,1,1,2,0,2,2,1,0,2,2,1,2,2,1,2,2,2,0,2,0,1,1,2,0,0,2,0,2,0,2,2,0,2,1,0,1,2,1,2,0,1,2,1,1,0,1,2,1,0,2,2,2,2,1,0,0,2,1,2,1,0,0,1,2,2,0,2,0,2,0,2,0,0,2,2,2,2,2,2,2,2,2,2,2,0,2,2,0,0,0,0,2,0,2,0,2,1,2,0,2,1,2,1,1,2,1,1,0,0,2,0,1,2,1,2,2,0,1,2,0,2,1,1,1,0,0,2,1,1,2,2,1,2,2,2,2,0,2,2,0,2,1,2,2,2,2,0,2,2,2,1,0,2,2,2,0,1,0,2,2,2,0,2,1,2,0,2,0,0,2,1,2,1,0,0,0,1,0,1,2,2,2,2,1,1,2,2,1,0,0,1,2,0,2,2,1,2,2,2,2,2,2,0,2,1,1,0,1,2,1,1,1,2,0,2,1,2,1,0,0,2,1,1,2,2,2,1,1,2,1,1,0,1,0,2,2,0,0,2,2,0,1,1,2,0,2,1,2,0,2,2,2,1,2,2,2,2,2,1,2,2,1,1,1,2,2,0,2,2,1,2,0,0,1,0,0,2,1,2,1,2,2,0,0,2,2,2,2,2,1,2,1,2,1,2,0,1,0,2,2,2,1,2,0,2,1,2,1,2,2,2,1,2,0,2,1,2,2,2,1,1,2,1,0,2,1,1,0,0,2,2,1,2,2,1,0,0,0,1,0,0,2,2,2,2,2,2,1,0,2,0,2,2,2,2,0,0,0,2,0,0,2,0,2,1,0,2,2,1,0,2,1,1,2,2,2,2,0,2,0,0,1,1,1,0,2,0,0,0,1,1,2,0,2,1,2,2,2,0,1,2,2,0,2,2,2,0,2,0,2,1,2,2,1,2,2,2,1,0,1,1,0,1,0,2,2,2,0,0,0,2,0,2,2,2,2,0,2,2,0,2,0,2,2,1,0,0,1,2,2,2,2,2,1,2,1,2,0,2,2,1,1,2,1,2,0,2,0,2,0,0,2,0,2,1,1,1,1,0,2,2,1,2,2,2,0,2,2,2,0,2,2,2,1,0,1,2,2,1,2,1,2,0,0,2,0,1,2,1,2,0,0,0,1,0,0,2,1,2,0,2,1,1,2,1,2,2,0,1,1,0,1,2,0,2,2,0,1,0,2,1,1,1,2,0,2,0,2,0,1,0,2,2,2,2,2,2,2,1,2,2,2,2,0,2,2,2,0,1,1,1,0,2,1,2,0,2,2,2,1,1,1,1,0,2,2,2,0,1,0,2,1,2,2,2,0,1,2,2,2,0,2,2,0,2,1,0,1,2,2,0,1,2,2,2,0,2,1,2,0,0,2,2,0,2,1,2,0,0,2,2,1,2,1,1,0,0,1,0,2,1,2,2,0,0,1,2,2,0,2,1,2,1,2,0,2,2,2,2,0,0,2,0,1,1,0,2,2,2,1,2,1,0,2,2,2,1,2,0,2,0,0,1,0,2,2,1,0,2,1,0,2,1,2,1,2,0,2,2,2,0,2,1,2,2,0,2,2,1,2,0,2,2,0,2,2,1,1,2,1,0,2,0,0,1,2,2,2,1,2,1,1,2,2,2,1,1,0,0,2,2,0,2,0,1,1,2,0,0,2,1,0,2,0,2,2,2,0,2,2,0,2,2,2,2,0,2,1,1,2,2,1,0,2,0,1,1,0,0,2,2,1,2,0,0,2,0,2,0,0,2,2,1,1,2,2,2,2,2,2,2,2,2,0,1,1,1,0,1,2,1,0,1,0,1,2,2,1,0,1,2,0,1,2,0,2,1,2,2,2,0,0,0,0,0,2,1,0,2,2,2,0,0,2,0,2,2,2,2,1,0,2,0,2,2,2,1,2,2,2,2,2,0,0,2,1,1,0,0,1,1,1,0,0,2,0,2,0,1,2,2,0,2,1,2,2,1,1,0,1,0,1,2,1,1,1,0,0,1,2,0,0,0,2,2,0,2,0,2,2,0,2,2,1,2,1,2,2,2,2,0,1,1,2,0,1,0,2,0,2,2,0,2,0,2,2,0,0,0,1,1,2,2,2,1,0,2,0,1,2,1,2,2,2,0,0,0,1,0,2,2,2,0,2,0,1,2,1,0,1,2,1,1,2,2,2,0,1,1,2,2,0,0,1,0,2,1,0,1,1,1,2,1,2,2,2,2,2,2,1,0,2,2,2,2,0,2,1,2,2,0,1,1,0,2,1,2,2,2,0,1,2,1,1,1,2,2,0,0,1,0,0,1,1,2,2,0,2,0,2,2,0,2,0,1,2,2,1,1,2,0,0,0,0,2,1,1,1,2,2,2,1,1,1,2,2,2,0,2,1,0,0,2,2,1,0,0,0,2,2,2,1,2,2,1,0,2,2,2,0,2,2,0,0,1,0,2,1,1,2,2,2,0,2,0,2,2,0,1,2,1,1,0,0,0,2,2,2,0,2,2,1,2,2,0,2,2,2,0,2,0,2,1,2,2,2,2,1,1,1,0,2,1,2,2,2,2,2,0,2,1,2,2,0,0,2,0,1,1,2,1,2,1,2,2,2,0,1,0,0,2,1,0,2,1,2,2,0,1,2,0,0,2,1,2,2,0,1,2,2,2,1,2,0,0,1,2,1,0,0,2,1,2,1,1,0,1,1,1,1,0,2,0,1,0,1,0,0,0,2,2,2,2,0,0,1,2,2,0,0,1,2,1,1,0,0,1,0,1,0,2,0,1,1,2,1,0,2,2,0,0,0,2,1,1,1,0,0,2,0,0,2,1,1,0,1,0,1,2,1,1,2,2,0,2,1,0,0,2,2,2,1,1,2,1,0,1,0,2,1,0,0,0,0,1,1,1,1,1,2,2,0,0,2,2,1,0,0,1,0,0,2,2,1,0,2,1,2,1,0,0,2,2,2,1,2,0,1,1,0,1,0,0,1,1,0,0,0,2,2,1,2,1,1,1,0,2}
  //input := []int{0,2,2,2,1,1,2,2,2,2,1,2,0,0,0,0}
  width := 25
  height := 6
  area := width * height
  layers := len(input) / area
  img := [150]int{}
  // Paint first layer (initialize img array w/ non zero vals)
  for i := 0; i < area; i++ {
    img[i] = input[i]
  }
  // Paint the rest
  for layer := 1; layer <= layers; layer++ {
    for i := area * (layer - 1); i < area * (layer); i++ {
      normalized_i := i - (area * (layer - 1))
      if img[normalized_i] == 2 { // if transparent, replace that junk
        img[normalized_i] = input[i]
      }
    }
  }
  for h := 0; h < height; h++ {
    for w := 0; w < width; w++ {
      index := w + (width * h)
      if img[index] != 1 {
        fmt.Printf(" ")
      } else {
        fmt.Printf("%d", img[index])
      }
    }
    fmt.Printf("\n")
  }
  // Solution to Part 1 of puzzle 8
  // max_zero_count := area
  // max_zero_layer := 0
  // max_layer_one_count := 0
  // max_layer_two_count := 0
  // for layer := 1; layer <= layers; layer++ {
  //   zero_count := 0
  //   one_count := 0
  //   two_count := 0
  //   for i := area * (layer - 1); i < area * (layer); i++ {
  //     if input[i] == 0 {
  //       zero_count++
  //     } else if input[i] == 1 {
  //       one_count++
  //     } else if input[i] == 2 {
  //       two_count++
  //     }
  //   }
  //   fmt.Printf("Layer: %d Zeroz: %d Onez: %d Twoz: %d\n", layer, zero_count, one_count, two_count)
  //   if zero_count < max_zero_count {
  //     max_zero_count = zero_count
  //     max_zero_layer = layer
  //     max_layer_one_count = one_count
  //     max_layer_two_count = two_count
  //   }
  // }
  // solution := max_layer_one_count * max_layer_two_count
  // fmt.Println(solution)
}