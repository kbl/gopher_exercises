package main

import (
	"book/ch02/conv"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var numbers []string
	if len(os.Args) > 1 {
		numbers = os.Args[1:]
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()
			if text == "exit" {
				break
			}
			t, _ := strconv.ParseFloat(text, 64)
			convert(t)
			numbers = append(numbers, scanner.Text())
		}
	}
	for _, a := range numbers {
		t, err := strconv.ParseFloat(a, 64)
		if err != nil {
			log.Fatal(err)
		}
		convert(t)
	}
}

func convert(t float64) {
	f := conv.Farenheit(t)
	k := conv.Kelvin(t)
	c := conv.Celsius(t)
	fmt.Printf("%s = %s = %s\n", f, f.ToC(), f.ToK())
	fmt.Printf("%s = %s = %s\n", c, c.ToK(), c.ToF())
	fmt.Printf("%s = %s = %s\n", k, k.ToF(), k.ToC())

	kg := conv.Kilogram(t)
	lbs := conv.Pound(t)
	fmt.Printf("%s = %s\n", kg, kg.ToP())
	fmt.Printf("%s = %s\n", lbs, lbs.ToK())

	meter := conv.Meter(t)
	feet := conv.Feet(t)
	fmt.Printf("%s = %s\n", meter, meter.ToF())
	fmt.Printf("%s = %s\n", feet, feet.ToM())
}
