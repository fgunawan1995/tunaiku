package main

import (
    "strconv"
	"math"
	"fmt"
)
//SieveOfEratosthenes
func FillPrime(value int) {
	id := 1
	var PrimeNumber PrimeNumber
	db := InitDb()
	defer db.Close()
	f := make([]bool, value)
    for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
        if f[i] == false {
            for j := i * i; j < value; j += i {
                f[j] = true
            }
        }
    }
    for i := 2; i < value; i++ {
        if f[i] == false {
            fmt.Printf("%v ", i)
			PrimeNumber.Id = id
			PrimeNumber.Numbers = i
			PrimeNumber.NameNumber = strconv.Itoa(i)
			db.Create(&PrimeNumber)
			id++
        }
    }
    fmt.Println("")
}

func FillEvenOdd(value int) {
	id := 1
	total := 0
	var EvenOddNumber EvenOddNumber
	db := InitDb()
	defer db.Close()
    for i := 2; i < value; i += 2 {
		fmt.Printf("%v ", i)
		EvenOddNumber.Id = id
		EvenOddNumber.NumberEvens = i
		EvenOddNumber.NumberOdds = i-1
		total = i+i-1
		EvenOddNumber.Total = total
		EvenOddNumber.NameTotal = strconv.Itoa(total)
		db.Create(&EvenOddNumber)
		id++
    }
}

func FillNominal(value float64) {
	id := 1
	nominal := 2000000.0
	add := 0.0
	total := 0.0
	var NominalAmount NominalAmount
	db := InitDb()
	defer db.Close()
    for {
		NominalAmount.Id = id
		NominalAmount.Nominal = nominal
		add = 0.03 * nominal
		NominalAmount.AdditionalNumber = int(add)
		total = nominal + add
		NominalAmount.Total = total
		nominal = total
		db.Create(&NominalAmount)
		if total > value{
			return
		}
		id++
    }
}