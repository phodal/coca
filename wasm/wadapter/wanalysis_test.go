package wadapter

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestWAnalysis_Analysis(t *testing.T) {
	g := NewGomegaWithT(t)

	var code = `package com.phodal.coca.analysis.identifier.model;

public class DataClass {
    private String date;

    public String getDate() {
        return date;
    }
}

public class Hello {

}`

	results := new(WAnalysis).Analysis(code)

	g.Expect(len(results)).To(Equal(2))
	g.Expect(results[0].Class).To(Equal("DataClass"))
	g.Expect(results[1].Class).To(Equal("Hello"))
}