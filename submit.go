package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "submit"
	app.Version = Version
	app.Usage = ""
	app.Author = "upamune"
	app.Email = "jajkeqos@gmail.com"
	app.Action = doMain
	app.Run(os.Args)
}

//AOJ => AOJ User
type AOJ struct {
	id         string
	pass       string
	source     string
	language   string
	problemNum string
}

// NewAOJ is Constructor
func NewAOJ(id string, pass string, source string, language string, problemNum string) *AOJ {
	aoj := &AOJ{
		id:         id,
		pass:       pass,
		source:     source,
		language:   language,
		problemNum: problemNum,
	}
	return aoj
}

func (aoj AOJ) submitCode() int {
	// valuesに値を設定しておく
	values := url.Values{}
	values.Add("userID", aoj.id)
	values.Add("password", aoj.pass)
	values.Add("sourceCode", aoj.source)
	values.Add("problemNO", aoj.problemNum)
	values.Add("language", aoj.language)

	// POSTする
	res, err := http.PostForm("http://judge.u-aizu.ac.jp/onlinejudge/servlet/Submit", values)
	if err != nil {
		println("Error is happen when submitting to AOJ")
		os.Exit(1)
	}
	defer res.Body.Close()

	return res.StatusCode
}

func (aoj *AOJ) xmlParser(xml string) (string, int64) {
	// XMLからstatusを取得する
	re, _ := regexp.Compile("(<status>\n?)(.*)(\n?</status>)")
	status := re.FindString(xml)
	status = strings.Replace(status, "<status>\n", "", 1)
	status = strings.Replace(status, "\n</status>", "", 1)
	// XMLからsubmission_dateを取得する
	re, _ = regexp.Compile("<submission_date>\n?(.*)\n?</submission_date>")
	submissionDate := re.FindString(xml)
	submissionDate = strings.Replace(submissionDate, "<submission_date>\n", "", 1)
	submissionDate = strings.Replace(submissionDate, "\n</submission_date>", "", 1)

	submissionTime, _ := strconv.ParseInt(submissionDate, 10, 64)
	submissionTime /= 1000

	return status, submissionTime
}

func (aoj *AOJ) checkSubmittedCode(submitTime int64) string {

	for i := 0; i < 5; i++ {
		// GETメソッドのvalues
		values := url.Values{}
		values.Add("user_id", aoj.id)
		values.Add("problem_id", aoj.problemNum)
		values.Add("limit", "1")

		res, err := http.Get("http://judge.u-aizu.ac.jp/onlinejudge/webservice/status_log" + "?" + values.Encode())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		// println("BODY", string(body))

		// 無理やりなのでXMLをパースするようにしたい
		status, submissionTime := aoj.xmlParser(string(body))

		// 提出した時との時間差が30秒以内だったら結果を返す
		diffTime := math.Abs(float64(submitTime) - float64(submissionTime))

		if diffTime <= 30 {
			switch status {
			case "Accepted":
				return "AC"
			case "Time Limit Exceeded":
				return "TLE"
			case "Runtime Error":
				return "RE"
			case "WA: Presentation Error":
				return "PE"
			case "Wrong Answer":
				return "WA"
			case "Compile Error":
				return "CE"
			case "Memory Limit Exceeded":
				return "MLE"
			case "Partial Points":
				return "PP"
			case "Output Limit Exceeded":
				return "OLE"
			}
		} else {
			// 3秒まってからリトライ
			time.Sleep(3 * time.Second)
		}
	}
	return ""
}

func checkFileType(filename string) string {
	// 拡張子を取得する
	re, _ := regexp.Compile("[^.]+$")
	ext := re.FindString(filename)

	// C, C++11, Java, C#, D, Ruby, Python, PHP, JavaScript に対応している
	switch ext {
	case "c":
		return "C"
	case "cpp", "cc":
		return "C++11"
	case "java":
		return "JAVA"
	case "cs":
		return "C#"
	case "d":
		return "D"
	case "rb":
		return "Ruby"
	case "py":
		return "Python"
	case "php":
		return "PHP"
	case "js":
		return "JavaScript"
	default:
		{
			println("AOJ is not unsupported this Programming Language")
			os.Exit(1)
		}
	}

	return "null"
}

func setIDPass() (string, string) {
	id := os.Getenv("AOJID")
	pass := os.Getenv("AOJPASS")
	if id == "" && pass == "" {
		println("Please set yourID and password on $AOJID and $AOJPASS")
		os.Exit(1)
	} else if id == "" {
		println("Please set yourID on $AOJID")
		os.Exit(1)
	} else if pass == "" {
		println("Please set password on $AOJPASS")
		os.Exit(1)
	}

	return id, pass
}

// arrangeNum 問題番号を4桁にして返す(0詰め)
func arrangeNum(num string) string {
	num = "0000" + num

	return num[len(num)-4:]
}

func doMain(c *cli.Context) {
	// 引数の数
	narg := len(c.Args())
	if narg != 2 {
		if narg < 2 {
			// 引数が多い時
			println("Error Argument isn't enough")
		} else {
			// 引数が少ない時
			println("Error Argument is too enough")
		}
		os.Exit(1)
	}
	// 環境変数からIDとPASSを取得
	id, pass := setIDPass()
	// 引数から問題番号とファイル名を設定
	number, filename := c.Args()[0], c.Args()[1]
	number = arrangeNum(number)

	// ファイル名からプログラミング言語を設定
	language := checkFileType(filename)
	// ファイル名のファイルを読み込む
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		println("Something error is happen when opening a file")
		os.Exit(1)
	}
	aoj := NewAOJ(id, pass, string(src), language, number)
	// println("ID =>", aoj.id, "PASS => ", aoj.pass)
	// println("ProblemNum => ", aoj.problemNum)
	// println("LANGAGE => ", aoj.language)
	// println("SRC\n", aoj.source)
	resCode := aoj.submitCode()
	submittedTime := time.Now().Unix()
	if resCode != 200 {
		println("Error is happen when submitted a code to AOJ, CODE => ", resCode)
		os.Exit(1)
	}

	time.Sleep(1 * time.Second)
	judge := aoj.checkSubmittedCode(submittedTime)
	println(judge)

}
