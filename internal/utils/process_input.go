package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func ProcessCsvFile(
	filePath string,
	process func(row []string) error,
) {
    in, err := os.Open("./input/" + filePath)
    if err != nil {
        log.Fatal(err)
    }
    defer in.Close()

    reader := csv.NewReader(in)
    counter := 1
    for {
        row, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("\n[info] now processing row %d\n", counter)
        process(row)
        fmt.Printf("[info] ended processing row %d", counter)
        counter++
    }
}
