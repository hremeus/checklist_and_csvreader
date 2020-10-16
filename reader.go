package main

import (
    "encoding/csv"
    "strconv"
    "fmt"
    "os"
)

type Product struct {
	Name string
    TotalCost, TotalQuantity int
}

func main() {
    lines, err := ReadCsv("input.csv")
    if err != nil {
        panic(err)
    }

	products := make(map[string]*Product)

    for _, line := range lines[1:] {
    	quantity, err := strconv.Atoi(line[4])
    	if err != nil {
	        panic(err)
	    }

    	cost, err := strconv.Atoi(line[3])
    	if err != nil {
	        panic(err)
	    }
		
    	if _, ok := products[line[2]]; ok {
			products[line[2]].TotalCost += cost * quantity
			products[line[2]].TotalQuantity += quantity
		} else {
			products[line[2]] = &Product{line[2], cost * quantity, quantity}
		}
    }

    for _, product := range products {
        fmt.Println(product.Name, " ", "total quantity:", product.TotalQuantity, " ", "total cost:", product.TotalCost)
    }
}

func ReadCsv(filename string) ([][]string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return [][]string{}, err
    }
    defer f.Close()

    lines, err := csv.NewReader(f).ReadAll()
    if err != nil {
        return [][]string{}, err
    }

    return lines, nil
}