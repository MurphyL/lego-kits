package main

import (
	"dsa/internal/data_dict"
	"os"
	"strings"
	"testing"
)

type testDataDictEntry struct {
	group, label, value, intro string
}

func (d *testDataDictEntry) Group() string {
	return d.group
}

func (d *testDataDictEntry) Label() string {
	return d.label
}

func (d *testDataDictEntry) Value() string {
	return d.value
}

func (d *testDataDictEntry) Intro() string {
	return d.intro
}

func TestDataDict(t *testing.T) {
	dd := data_dict.New(func(group, label string) []data_dict.Entry {
		ret := make([]data_dict.Entry, 0)
		envs := os.Environ()
		for _, entry := range envs {
			l, v, ok := strings.Cut(entry, "=")
			if ok {
				if label == l {
					ret = append(ret, &testDataDictEntry{
						group: "environ",
						label: l,
						value: v,
						intro: entry,
					})
				}
			}
		}
		return ret
	})
	x := dd.Value("environ", "JAVA_HOME")
	t.Log("hello", x)
}
