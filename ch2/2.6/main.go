package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		// 从标准输入读取参数
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("请输入转换参数（例如：32°F、5ft、10kg）：")
		input, _ := reader.ReadString('\n')
		convert(input)
	} else {
		// 从命令行读取参数
		input := strings.Join(os.Args[1:], " ")
		convert(input)
	}
}

func convert(input string) {
	input = strings.TrimSpace(input)
	unit := input[len(input)-2:]
	valueStr := input[:len(input)-2]

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		fmt.Println("无效的数值:", valueStr)
		return
	}

	switch unit {
	case "°C":
		f := celsiusToFahrenheit(value)
		fmt.Printf("%.2f°C = %.2f°F\n", value, f)
	case "°F":
		c := fahrenheitToCelsius(value)
		fmt.Printf("%.2f°F = %.2f°C\n", value, c)
	case "ft":
		m := feetToMeters(value)
		fmt.Printf("%.2fft = %.2fm\n", value, m)
	case "m":
		ft := metersToFeet(value)
		fmt.Printf("%.2fm = %.2fft\n", value, ft)
	case "lb":
		kg := poundsToKilograms(value)
		fmt.Printf("%.2flb = %.2fkg\n", value, kg)
	case "kg":
		lb := kilogramsToPounds(value)
		fmt.Printf("%.2fkg = %.2flb\n", value, lb)
	default:
		fmt.Println("不支持的单位:", unit)
	}
}

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func feetToMeters(feet float64) float64 {
	return feet * 0.3048
}

func metersToFeet(meters float64) float64 {
	return meters / 0.3048
}

func poundsToKilograms(pounds float64) float64 {
	return pounds * 0.45359237
}

func kilogramsToPounds(kilograms float64) float64 {
	return kilograms / 0.45359237
}

// package main

// import (
// 	"2.6/tempconv"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	//fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC) // "Brrrr! -273.15°C"
// 	//fmt.Println(tempconv.CToF(tempconv.BoilingC))     // "212°F"
// 	//fmt.Println(tempconv.CToK(tempconv.FreezingC))
// 	//fmt.Println(tempconv.FToK(12))
// 	for _, arg := range os.Args[1:] {
// 		t, err := strconv.ParseFloat(arg, 64)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "cf:%v\n", err)
// 			os.Exit(1)
// 		}
// 		if t == 0.0 {
// 			var b float64
// 			fmt.Scanln(&b)
// 			t = b
// 		}
// 		f := tempconv.Fahrenheit(t)
// 		c := tempconv.Celsius(t)
// 		fmt.Printf("%s=%s,%s=%s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
// 	}
// }
