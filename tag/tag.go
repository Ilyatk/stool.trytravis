package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("qwewe", len(os.Args), os.Args)
        fmt.Printf("%#v", os.Args)

return 
	if len(os.Args) != 3 {
		fmt.Println("usage test v123.56.1 2 ")
		return
	}


	i, err := strconv.Atoi(os.Args[2])
	if err != nil || i < 1 || i > 3 {
		return
	}

	rv := regexp.MustCompile("[Vv]?(\\d{1,})\\.(\\d{1,})\\.(\\d{1,})")

	m := rv.FindStringSubmatch(os.Args[1])
	if len(m) != 4 {
		if i == 1 {
			fmt.Printf("%s", "1")
		} else {
			fmt.Printf("%s", "0")
		}
		return
	}

	//fmt.Printf("%q\n", m)
	fmt.Printf("%s", m[i])

	// fmt.Printf("%q\n", rv.FindStringSubmatch("-axxxbyc-"))
	// fmt.Printf("%q\n", rv.FindStringSubmatch("v1"))
	// fmt.Printf("%q\n", rv.FindStringSubmatch("v123"))
	// fmt.Printf("%q\n", rv.FindStringSubmatch("v123.42"))
	// fmt.Printf("%q\n", rv.FindStringSubmatch("v123.42."))
	// fmt.Printf("%q\n", rv.FindStringSubmatch("v123.42.q"))
	// fmt.Printf("%q\n", rv.FindStringSubmatch("v123.42.5"))
	// fmt.Printf("%q\n", rv.FindStringSubmatch("v123.42.5qweqwe"))

	// re := regexp.MustCompile("a(x*)b(y|z)c")
	// fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-"))
	// fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))
}
