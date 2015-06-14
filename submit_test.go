package main

import (
	"testing"
)

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
