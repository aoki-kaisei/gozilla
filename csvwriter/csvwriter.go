package csvwriter


import (
	"encoding/csv"
	"os"
	"regexp"
)


func Export(array []string,name  string) {
	hostChar := regexp.MustCompile(`https://|http://|/g`)
	nameWitoutHost := hostChar.ReplaceAllString(name, "")

	specialChar := regexp.MustCompile(`/|\:|\?|\.|"|<|>|\|\*|/g`)
	n := specialChar.ReplaceAllString(nameWitoutHost, "-")

	file, err := os.Create(n+".csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	cw := csv.NewWriter(file)
	defer cw.Flush()

	for i := 0; i < len(array); i++ {
		col := []string{array[i]}
		cw.Write(col)
	}

}