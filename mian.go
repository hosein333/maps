package main

import "fmt"

func main() {

    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

    _, ok := m["k1"]
    fmt.Println("ok:", ok)
    
    n := map[string]int{"foo": 1, "bar": 2, "test": 3, "ada": 4, "aaa": 5}
    fmt.Println("map:", n)

    for key, value := range n {
	    fmt.Println(key, value)
    }


    }
