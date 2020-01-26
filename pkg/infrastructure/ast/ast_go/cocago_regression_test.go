package ast_go

func getRegressionFile(name string) string {
	return "testdata/regression/" + name
}

// todo: fix package issues with input
//func Test_Regression1(t *testing.T) {
//	t.Parallel()
//	g := NewGomegaWithT(t)
//
//	filePath := getRegressionFile("coll_stack")
//	results := testParser.ProcessFile(filePath + ".code")
//	g.Expect(cocatest.JSONFileBytesEqual(results, filePath+".json")).To(Equal(true))
//}
