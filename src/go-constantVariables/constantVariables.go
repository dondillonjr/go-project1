package main

import (
    "fmt"
)

func main() {
    const myConst int = 43
    const b string = "foo"
    const d bool = true
    const c float32 = 3.14
    var   e int = 12

    fmt.Printf("%v, %T\n", myConst, myConst)
    fmt.Printf("%v\n", b)
    fmt.Printf("%v\n", c)
    fmt.Printf("%v\n", d)
    fmt.Printf("%v, %T\n", myConst +e , myConst +e)
}