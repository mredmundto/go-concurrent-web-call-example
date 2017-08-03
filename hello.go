package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/xml"
    "time"
    "runtime"
)

func main() {

    runtime.GOMAXPROCS(4)

    fmt.Println("concurrent golang web service call")

    start := time.Now()

    stockSymbols := []string{
        "goog",
        "msft",
        "aapl",
        "bbry",
        "hpq",
        "vz",
        "t",
        "tmus",
        "s",
    }

    numComplete := 0

    for _, symbol := range stockSymbols {

        go func(symbol string){

            resp, _ := http.Get("http://dev.markitondemand.com/MODApis/Api/v2/Quote?symbol=" + symbol)
            defer resp.Body.Close()
            body, _ := ioutil.ReadAll(resp.Body)

            quote := new(QuoteResponse)
            xml.Unmarshal(body, &quote)

            fmt.Printf("%s: %.2f\n", quote.Name, quote.LastPrice)
            numComplete++
        }(symbol)

    }

    for numComplete < len(stockSymbols){
        time.Sleep(10* time.Millisecond)
    }
    elapsed := time.Since(start)

    fmt.Printf("Execution time: %s", elapsed)
}

type QuoteResponse struct {
    Status string
    Name string
    LastPrice float32
}
