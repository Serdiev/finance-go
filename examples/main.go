package main

import (
	"fmt"

	"github.com/Serdiev/finance-go/chart"
	"github.com/Serdiev/finance-go/crypto"
	"github.com/Serdiev/finance-go/datetime"
	"github.com/Serdiev/finance-go/equity"
	"github.com/Serdiev/finance-go/etf"
	"github.com/Serdiev/finance-go/forex"
	"github.com/Serdiev/finance-go/future"
	"github.com/Serdiev/finance-go/index"
	"github.com/Serdiev/finance-go/mutualfund"
	"github.com/Serdiev/finance-go/options"
	"github.com/Serdiev/finance-go/quote"
)

// This file lists several usage examples of this library
// and can be used to verify behavior.
func main() {

	// Basic options example.
	// --------------------
	{
		fmt.Println("Options stradle example\n====================")
		fmt.Println()
		iter := options.GetStraddle("AAPL")

		fmt.Println(iter.Meta())

		for iter.Next() {
			fmt.Println(iter.Straddle().Strike)
		}
		if iter.Err() != nil {
			fmt.Println(iter.Err())
		}
		fmt.Println()
	}

	// Basic quote example.
	// --------------------
	{
		fmt.Println("Quote example\n====================")
		fmt.Println()
		q, err := quote.Get("GOOG")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(q)
		}
		fmt.Println()
	}

	// Basic chart example.
	// --------------------
	{
		fmt.Println("Chart example\n====================")
		fmt.Println()
		params := &chart.Params{
			Symbol:   "GOOG",
			Interval: datetime.OneHour,
		}
		iter := chart.Get(params)

		for iter.Next() {
			b := iter.Bar()
			fmt.Println(datetime.FromUnix(b.Timestamp))

		}
		if iter.Err() != nil {
			fmt.Println(iter.Err())
		}
		fmt.Println()
	}

	// Basic crypto example.
	// --------------------
	{
		fmt.Println("Crypto example\n====================")
		fmt.Println()
		q, err := crypto.Get("BTC-USD")

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(q)
		}
		fmt.Println()
	}

	// Basic equity example.
	// --------------------
	{
		fmt.Println("Equity example\n====================")
		fmt.Println()
		symbols := []string{"AAPL", "GOOG", "MSFT"}
		iter := equity.List(symbols)

		if iter.Err() != nil {
			fmt.Println(iter.Err())
			fmt.Println()
		} else {
			for iter.Next() {
				q := iter.Equity()
				fmt.Println(q)
				fmt.Println()
			}
		}
	}

	// Basic ETF example.
	// --------------------
	{
		fmt.Println("ETF example\n====================")
		fmt.Println()
		q, err := etf.Get("SPY")

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(q)
		}
		fmt.Println()
	}

	// Basic forex example.
	// --------------------
	{
		fmt.Println("Forex example\n====================")
		fmt.Println()
		q, err := forex.Get("CADUSD=X")

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(q)
		}
		fmt.Println()
	}

	// Basic future example.
	// --------------------
	{
		fmt.Println("Future example\n====================")
		fmt.Println()
		q, err := future.Get("CL=F")

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(q)
		}
		fmt.Println()
	}

	// Basic index example.
	// --------------------
	{
		fmt.Println("Index example\n====================")
		fmt.Println()
		q, err := index.Get("^DJI")

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(q)
		}
		fmt.Println()
	}

	// Basic mutual fund example.
	// --------------------
	{
		fmt.Println("Mutual fund example\n====================")
		fmt.Println()
		q, err := mutualfund.Get("FMAGX")

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(q)
		}
		fmt.Println()
	}
}
