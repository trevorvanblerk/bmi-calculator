package main

import (
    "fmt"
    "math"
)

func bmi(weight float64, height float64) (float64, string) {
    bmi := math.Round(weight / (height * height) * 100) / 100
    var classification string
    if bmi < 18.5 {
        classification = "(Underweight)"
    } else if bmi < 25 {
        classification = "(Normal weight)"
    } else if bmi < 30 {
        classification = "(Overweight)"
    } else {
        classification = "(Obese)"
    }
    return bmi, classification
}

func main() {
    var weight float64
    var height float64
    
    fmt.Println("Enter your weight (in kgs): ")
    fmt.Scan(&weight)
    fmt.Println("Enter your height (in meters): ")
    fmt.Scan(&height)
    
    bmi, classification := bmi(weight, height)
    fmt.Println("Your BMI is:", bmi, classification)
}
