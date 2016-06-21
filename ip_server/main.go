package main

import (
  "fmt"
  "net/http"
  "encoding/csv"
  "io/ioutil"
  "io"
	"strings"
)

func parse_ip_data(file_path string)  {
  in, err := ioutil.ReadFile(file_path)
  if err != nil {
    fmt.Println(err)
  }

  defer in.Close()
  r := csv.NewReader(strings.NewReader(string(in)))
  r.Comma = '\t'

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
    fmt.Println(record[0], record[1], record[2], record[3], record[4])
	}
}

func main() {
  parse_ip_data("cidr_optim.txt")
  mux := http.NewServeMux()
	mux.HandleFunc("/", IpResultHundler)
	http.ListenAndServe(":8000", mux)
}

func IpResultHundler(w http.ResponseWriter, r *http.Request)  {
  switch r.Method {
		case "GET":
      get_par := r.URL.Query()
      ip, ok := get_par["ip"]
      fmt.Println(ip, ok)
      if ok {
        fmt.Fprintf(
          w,
          "<h1>Country - %s</h1><br><h1>City - %s</h1><br><h1>Provider - %s</h1><br><h1>Timezone - %s</h1><br>",
          "Russia",
          "MOSCOW",
          "HzHZ",
          "HZ",
        )
      } else {
        w.Write([]byte("Don't get api in get params"))
      }

		default:
			w.Write([]byte("Method not supported"))
	}
}
