package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"time"
)

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

var print = fmt.Println

type response1 struct {
	Page   int
	Fruits []string
}
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	//var strs = []string{"Apple", "peach", "pear", "plum"}
	//fmt.Println(Index(strs, "pear"))
	//fmt.Println(Map(strs, func(s string) string {
	//	s = s + "?"
	//	return s
	//}))
	//
	//print(s.Contains("ABC","A"))
	//print(s.Count("ABCDA","A"))
	//print(s.HasPrefix("ABC", "A"))
	//print(s.HasSuffix("ABC", "C"))

	//bolB, _ := json.Marshal(true)
	//print(string(bolB))
	//
	//intB, _ := json.Marshal(1)
	//print(string(intB))
	//
	//str := `{"page":1,"fruits":["apa","pec"]}`
	//res := response2{}
	//json.Unmarshal([]byte(str), &res)
	//print(res)
	//print(res.Fruits[0])
	//
	//enc := json.NewEncoder(os.Stdout)
	//d := map[string] int{"apple":5,"lettuce":7}
	//enc.Encode(d)

	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan(){
	//	var text = scanner.Text()
	//	if strings.Contains(text, "ELK"){
	//		print(text)
	//	}
	//}
	//if err:=scanner.Err();err!=nil{
	//	print(os.Stderr, "error", err)
	//	os.Exit(1)
	//}

	//args := os.Args
	//ragswithoutProg := os.Args[1:]
	//fmt.Println(args)
	//fmt.Println(ragswithoutProg)

	//wordPtr := flag.String("start","foo","starting")
	//fmt.Println("word:", *wordPtr)
	//flag.Parse()
	//resp, err := http.Get("http://gobyexample.com")
	//if err != nil{
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//fmt.Println("response status:", resp.Status)
	//scanner := bufio.NewScanner(resp.Body)
	//for i:=0; scanner.Scan()&& i<5;i++{
	//	fmt.Println(scanner.Text())
	//}
	//if err:=scanner.Err(); err!=nil{
	//	panic(err)
	//}
	//
	//http.HandleFunc("/hello", hello)
	//http.HandleFunc("/headers", headers)
	//http.ListenAndServe(":8989",nil)

	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("< date")
	fmt.Println(string(dateOut))
	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep \n goodbye grep "))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsout, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsout))
}
func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	defer fmt.Println("server done")
	select {
	case <-time.After(2 * time.Second):
		fmt.Fprintf(w, "hello!\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		interError := http.StatusInternalServerError
		http.Error(w, err.Error(), interError)

	}
}
func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v : %v \n", name, h)
		}
	}
}
