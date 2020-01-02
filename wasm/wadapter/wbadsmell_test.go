package wadapter

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestWBadSmell_Analysis(t *testing.T) {
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

	results := new(WBadSmell).Analysis(code)

	g.Expect(results.ClassBS.OverrideSize).To(Equal(0))
	g.Expect(results.ClassBS.PublicVarSize).To(Equal(0))
}