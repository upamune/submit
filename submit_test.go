package main

import "testing"

// 拡張子がただしく取得できているか
func TestTrimExt(t *testing.T) {
	filename := "test.c"

	if ext := checkFileType(filename); ext != "C" {
		t.Error("ただしくCの拡張子が取得できていない")
	}
	filename = "test.cpp"
	if ext := checkFileType(filename); ext != "C++11" {
		t.Error("ただしくC++の拡張子が取得できていない")
	}
	filename = "test.cc"
	if ext := checkFileType(filename); ext != "C++11" {
		t.Error("ただしくC++の拡張子が取得できていない")
	}
	filename = "test.java"
	if ext := checkFileType(filename); ext != "JAVA" {
		t.Error("ただしくJavaの拡張子が取得できていない")
	}
	filename = "test.php"
	if ext := checkFileType(filename); ext != "PHP" {
		t.Error("ただしくPHPの拡張子が取得できていない")
	}
	filename = "test.py"
	if ext := checkFileType(filename); ext != "Python" {
		t.Error("ただしくPythonの拡張子が取得できていない")
	}
	filename = "test.rb"
	if ext := checkFileType(filename); ext != "Ruby" {
		t.Error("ただしくRubyの拡張子が取得できていない")
	}
	filename = "test.d"
	if ext := checkFileType(filename); ext != "D" {
		t.Error("ただしくDの拡張子が取得できていない")
	}
	filename = "test.cs"
	if ext := checkFileType(filename); ext != "C#" {
		t.Error("ただしくC#の拡張子が取得できていない")
	}
	filename = "test.js"
	if ext := checkFileType(filename); ext != "JavaScript" {
		t.Error("ただしくJavaScriptの拡張子が取得できていない")
	}
}

// プログラムを正常に提出できているかどうか
func TestSubmitCode(t *testing.T) {
	id, pass := setIDPass()
	aoj := NewAOJ(id, pass, "SRC", "C", "0001")
	if statusCode := aoj.submitCode(); statusCode != 200 {
		t.Error("プログラムを正常に提出できていない")
	}
}

func TestAOJStruct(t *testing.T) {
	id, pass := setIDPass()
	aoj := NewAOJ(id, pass, "SRC", "C", "0001")

	// 各フィールドにちゃんと値が入っているか
	if aoj.id != id {
		t.Error("IDがAOJ構造体に入っていない")
	}
	if aoj.pass != pass {
		t.Error("PASSがAOJ構造体に入っていない")
	}
	if aoj.source != "SRC" {
		t.Error("SOURCEがAOJ構造体に入っていない")
	}
	if aoj.language != "C" {
		t.Error("LANGUAGEがAOJ構造体に入っていない")
	}
	if aoj.problemNum != "0001" {
		t.Error("PROBLEMNUMがAOJ構造体に入っていない")
	}
}

// XMLがパースできているかテストする
func TestXMLParse(t *testing.T) {
	// テスト用のXML
	xml := `
<?xml version="1.0"?>
<status_list>
<status>
<run_id>
753727
</run_id>
<user_id>
solver
</user_id>
<problem_id>
0500
</problem_id>
<submission_date>
1377444460525
</submission_date>
<submission_date_str>
Mon Aug 26 00:27:40
</submission_date_str>
<status>
Accepted
</status>
<language>
C++
</language>
<cputime>
0
</cputime>
<memory>
1160
</memory>
<code_size>
344
</code_size>
</status>
</status_list>`
	id, pass := setIDPass()
	aoj := NewAOJ(id, pass, "SRC", "C", "0001")
	status, date := aoj.xmlParser(xml)
	if status != "Accepted" {
		t.Error("AcceptedなXMLのSTATUSをパースできていない")
	}
	if date != 1377444460 {
		t.Error("AcceptedなXMLのSUBMISSION_DATEをパースできていない")
	}
}
