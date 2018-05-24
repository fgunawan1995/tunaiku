package main

func main() {
	db := InitDb()
	defer db.Close()
	FillPrime(1000000)
	FillEvenOdd(1000000)
	FillNominal(15000000)
	ReadCsv()
}