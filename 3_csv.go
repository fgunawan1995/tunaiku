package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "os"
	"strconv"
)

type Stock struct {
    Date string `json:"date"`
	Open float64 `json:"open"`
	High float64 `json:"high"`
	Low float64 `json:"low"`
	AdjClose float64 `json:"adjclose"`
	Volume float64 `json:"volume"`	
}

type SimplifiedStock struct {
	Date string `json:"date"`
	Value float64 `json:"value"`	
}

func ReadCsv(){
	csvFile, _ := os.Open("data.csv")
	i := 0
    reader := csv.NewReader(bufio.NewReader(csvFile))
    var Stocks []Stock
    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        }
		i++
		if i > 1 {
			open,erropen := strconv.ParseFloat(line[1],64)
			if erropen != nil {
				log.Fatal(erropen)
			}
			high,errhigh := strconv.ParseFloat(line[2],64)
			if errhigh != nil {
				log.Fatal(errhigh)
			}
			low,errlow := strconv.ParseFloat(line[3],64)
			if errlow != nil {
				log.Fatal(errlow)
			}
			adjclose,erradjclose := strconv.ParseFloat(line[5],64)
			if erradjclose != nil {
				log.Fatal(erradjclose)
			}
			volume,errvolume := strconv.ParseFloat(line[6],64)
			if errvolume != nil {
				log.Fatal(errvolume)
			}
			Stocks = append(Stocks, Stock{
				Date: line[0],
				Open:  open,
				High: high,
				Low:  low,
				AdjClose:  adjclose,
				Volume: volume,
			})
		}
    }
	var SimplifiedStocks []int
    for _, worker := range Stocks{
		SimplifiedStocks = append(SimplifiedStocks, int(worker.AdjClose*worker.Volume))
	}
	j := 0
	maxDiff := 0
	buy := 0
	sell := 0
	diff := 0
	for i := 0; i < len(SimplifiedStocks); i++{
		for j = i; j < len(SimplifiedStocks); j++{
			diff = SimplifiedStocks[j] - SimplifiedStocks[i];
			if (diff > maxDiff){
				maxDiff = diff;
				buy = i;
				sell = j;
			}
		}	
	}
	fmt.Printf("3. Buy at %s and sell at %s",Stocks[buy].Date,Stocks[sell].Date)
}