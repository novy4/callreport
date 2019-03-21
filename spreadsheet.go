package main

// We only use packages from the standard library here.
import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// In `main()`, we sketch out our program flow:
//
// * Read the CSV file,
// * calculate the desired numbers, and
// * write the results to a new CSV file.
func main() {
	rows := readOrders("orders.csv")
	clients := readClients("clients.csv")
	rows = calculate(rows, clients)
	writeOrders("ordersReport.csv", rows)
}

/*
### Reading CSV files

As the next step, we need to read in the header row, and then the data rows. The result shall be a two-dimensional slice of strings, or a slice of slices of strings.
*/

// `readOrders` takes a filename and returns a two-dimensional list of spreadsheet cells.
func readOrders(name string) [][]string {

	f, err := os.Open(name)
	// Usually we would return the error to the caller and handle
	// all errors in function `main()`. However, this is just a
	// small command-line tool, and so we use `log.Fatal()`
	// instead, in order to write the error message to the
	// terminal and exit immediately.
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", name, err.Error())
	}

	// After this point, the file has been successfully opened,
	// and we want to ensure that it gets closed when no longer
	// needed, so we add a deferred call to `f.Close()`.
	defer f.Close()

	// To read in the CSV data, we create a new CSV reader that
	// reads from the input file.
	//
	// The CSV reader is aware of the CSV data format. It
	// separates the input stream into rows and columns,
	// and returns a slice of slices of strings.
	r := csv.NewReader(f)

	// We can even adjust the reader to recognize a semicolon,
	// rather than a comma, as the column separator.
	r.Comma = ';'

	// Read the whole file at once. (We don't expect large files.)
	rows, err := r.ReadAll()

	// Again, we check for any error,
	if err != nil {
		log.Fatalln("Cannot read CSV data:", err.Error())
	}

	// and finally we can return the rows.
	return rows
}

func readClients(name string) [][]string {

	f, err := os.Open(name)
	// Usually we would return the error to the caller and handle
	// all errors in function `main()`. However, this is just a
	// small command-line tool, and so we use `log.Fatal()`
	// instead, in order to write the error message to the
	// terminal and exit immediately.
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", name, err.Error())
	}

	// After this point, the file has been successfully opened,
	// and we want to ensure that it gets closed when no longer
	// needed, so we add a deferred call to `f.Close()`.
	defer f.Close()

	// To read in the CSV data, we create a new CSV reader that
	// reads from the input file.
	//
	// The CSV reader is aware of the CSV data format. It
	// separates the input stream into rows and columns,
	// and returns a slice of slices of strings.
	r := csv.NewReader(f)

	// We can even adjust the reader to recognize a semicolon,
	// rather than a comma, as the column separator.
	r.Comma = ';'

	// Read the whole file at once. (We don't expect large files.)
	clients, err := r.ReadAll()

	// Again, we check for any error,
	if err != nil {
		log.Fatalln("Cannot read CSV data:", err.Error())
	}

	// and finally we can return the rows.
	return clients
}

/*
### Process the data

Now that the data is read in, we can loop over the rows, and read from or write to each row slice as needed.

This is where we can extract the desired information: The total price for each order, the total sales volume, and the number of ball pens sold.
*/

// `calculate` takes a spreadsheet, extracts and calculates the desired information, and returns the result as a new spreadsheet.
func calculate(rows [][]string, clients [][]string) [][]string {

	sum := 0
	var rows_new [][]string
	tsm := 0
	gee := 0
	divat := 0
	mdtrans := 0
	lestris := 0
	legalteh := 0
	tehnogrant := 0
	sacramento := 0

	// To process the data, we loop over the rows, and read from
	// or write to each row slice as needed.
	for i := range rows {

		// The first row is the header row. Here, we only want to
		// add a new header for the column that holds the total prices.
		if i == 0 {
			//rows[0] = append(rows[0], "Total")
			continue
		}

		// From the next row onwards, we calculate the total
		// price, sum up all prices, and count the number of ball
		// pens being ordered.

		// This is fairly straightforward, as we know the indexes
		// of the item name, the unit price, and the quantity.
		// The only difficulty is that all columns are string
		// values but we need the price and quantity values as
		// numeric values.

		// We know that column 2 contains the item name.
		item := rows[i][0]

		// Another obstacle we are facing here is that the prices are floating-point values but for financial calculations, we want to use precise integer calculation only. Luckily, the [`strings`](https://golang.org/pkg/strings) and [`strconv`](https://golang.org/pkg/strconv/) packages have got us covered.

		// Column 3 contains the price. Remove the decimal point using `strings.Replace()`, and
		// turn the value into an integer (representing the value in cents) using `strconv.Atoi`.
		price, err := strconv.Atoi(strings.Replace(rows[i][10], ".", "", -1))
		if err != nil {
			log.Fatalf("Cannot retrieve price of %s: %s\n", item, err)
		}

		// Column 4 contains the ordered quantity. Again, we convert the value into an integer.
		qty, err := strconv.Atoi(strings.Replace(rows[i][10], ".", "", -1))
		if err != nil {
			log.Fatalf("Cannot retrieve call price of %s: %s\n", item, err)
		}

		// Calculate the total and append it to the current row.
		total := price

		// We use a helper function to turn the total value (an integer) back into a floating-point value with two decimals, represented as a string (see below).

		//rows[i] = append(rows[i], intToFloatString(total))

		// Update the total sum
		sum += total
		/*
			for x := range clients {
				if item == temp {
					fmt.Println(clients[x][1])
					client_total += qty
				} else {
					//temp := rows[i][0]
				}

			}
		*/
		//nb += qty
		if item == "022207210" || item == "022207216" || item == "022207217" || item == "022207240" || item == "022207250" || item == "022207283" || item == "022207284" || item == "022207289" {
			tsm += qty
		}
		if item == "022207201" || item == "022207207" {
			gee += qty
		}
		if item == "022207204" {
			lestris += qty
		}
		if item == "022207224" || item == "022207226" {
			legalteh += qty
		}
		if item == "022207244" || item == "022207247" {
			tehnogrant += qty
		}
		if item == "022207249" {
			sacramento += qty
		}
		if item == "022207266" {
			divat += qty
		}
		if item == "022207277" || item == "022207279" {
			mdtrans += qty
		}
	}
	// Here we append two new rows. The first one shows the total sum, and
	// the second one shows the number of ball pens ordered.
	rows_new = append(rows_new, []string{"TSM", ":", intToFloatString(tsm - tsm/6)})
	rows_new = append(rows_new, []string{"GEE", ":", intToFloatString(gee - gee/6)})
	rows_new = append(rows_new, []string{"LESTRIS", ":", intToFloatString(lestris - lestris/6)})
	rows_new = append(rows_new, []string{"LEGAL TEHNOLOGIES", ":", intToFloatString(legalteh - legalteh/6)})
	rows_new = append(rows_new, []string{"TEHNOGRANT", ":", intToFloatString(tehnogrant - tehnogrant/6)})
	rows_new = append(rows_new, []string{"SACRAMENTO", ":", intToFloatString(sacramento - sacramento/6)})
	rows_new = append(rows_new, []string{"DIVAT", ":", intToFloatString(divat - divat/6)})
	rows_new = append(rows_new, []string{"MDTRANS", ":", intToFloatString(mdtrans - mdtrans/6)})
	rows_new = append(rows_new, []string{"", "", "", "Sum", "", intToFloatString(sum - sum/6)})
	//for row := range rows {
	//	rows = append(rows, []string{"", "", "", "client", intToFloatString(nb)})
	//	fmt.Println(item, nb)
	//}

	// Return the new spreadsheet.
	return rows_new
}

// `intToFloatString` takes an integer `n` and calculates the floating point value representing `n/100` as a string.
func intToFloatString(n int) string {
	intgr := n / 100
	frac := n - intgr*100
	return fmt.Sprintf("%d.%d", intgr, frac)
}

/*
### Write the new CSV data

Finally, we write the result to a new file, using `os.Create()` and a CSV writer that knows how to turn the slice of slices of strings back into a proper CSV file.

Note that we do not set the separator to semicolon here, as we  want to create a standard CSV format this time.
*/

// `writeOrders` takes a filename and a spreadsheet and writes the spreadsheet as CSV to the file.
func writeOrders(name string, rows_new [][]string) {

	f, err := os.Create(name)
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", name, err.Error())
	}

	// We are going to write to a file, so any errors are relevant and
	// need to be logged. Hence the anonymous func instead of a one-liner.
	defer func() {
		e := f.Close()
		if e != nil {
			log.Fatalf("Cannot close '%s': %s\n", name, e.Error())
		}
	}()

	w := csv.NewWriter(f)
	err = w.WriteAll(rows_new)
}

/*
When running this code, the output file should look like this:

```
Date,Order ID,Order Item,Unit Price,Quantity,Total
2017-11-17,1,Ball Pen,1.99,50,99.50
2017-11-17,2,Notebook,12.99,10,129.90
2017-11-17,3,Binder,4.99,25,124.75
2017-11-18,4,Pencil,0.99,100,99.0
2017-11-18,5,Sketch Block,2.99,40,119.60
2017-11-19,6,Ball Pen,1.99,30,59.70
2017-11-19,7,Sketch Block,2.99,20,59.80
2017-11-19,8,Ball Pen,1.99,60,119.40
,,,Sum,,811.65
,,,Ball Pens,140,
```

And we can open it in our spreadsheet app, or in a CSV viewer, to get a nicely formatted table.

Date       | Order ID | Order Item   | Unit Price | Quantity | **Total**
-----------|:--------:|--------------|-----------:|---------:|-------:
2017-11-17 | 1        | Ball Pen     | 1.99       | 50       | **99.50**
2017-11-17 | 2        | Notebook     | 12.99      | 10       | **129.90**
2017-11-17 | 3        | Binder       | 4.99       | 25       | **124.75**
2017-11-18 | 4        | Pencil       | 0.99       | 100      | **99.0**
2017-11-18 | 5        | Sketch Block | 2.99       | 40       | **119.60**
2017-11-19 | 6        | Ball Pen     | 1.99       | 30       | **59.70**
2017-11-19 | 7        | Sketch Block | 2.99       | 20       | **59.80**
2017-11-19 | 8        | Ball Pen     | 1.99       | 60       | **119.40**
           |          | **Sum**      |            |          | **811.65**
           |          | **Ball Pens**|            | **140**  |

Here we can see our new Totals column, as well as the two new rows that show the overall sum and the number of ball pens ordered.


## How to get and run the code

Step 1: `go get` the code. Note the `-d` flag that prevents auto-installing
the binary into `$GOPATH/bin`.

This time, also note the `/...` postfix that downloads all files, not only those imported by the main package.

    go get -d github.com/appliedgo/spreadsheet/...

Step 2: `cd` to the source code directory.

    cd $GOPATH/src/github.com/appliedgo/spreadsheet

Step 3. Run the binary.

    go run spreadsheet.go

You should then find a file named `ordersReport.csv` in the current directory. Verify that it contains the expected result.


### Q&A: Why CSV?

I use CSV here, rather than the file formats used by Excel or Open/Libre Office or Numbers, in order to stay as flexible and vendor-independent as possible. If you specifically want to work with Excel sheets, a [quick search on GitHub](https://github.com/search?o=desc&q=excel+language%3Ago&s=stars&type=Repositories&utf8=%E2%9C%93) should return a couple of useful third-party libraries. I have not used any of them yet, so I can neither share any experience nor recommend a particular one.

## Links
[Wikipedia: Comma-separated values](https://en.wikipedia.org/wiki/Comma-separated_values) - Details about the CSV format.

**Happy coding!**

*/
