package constants

// TBS Config List

var (
	DuplicatedAssertionLimitLength = 5
	ASSERTION_LIST = []string{
		"assert",
		"should",
		"check",    // ArchUnit,
		"maynotbe", // ArchUnit,
		"is",       // RestAssured,
		"spec",     // RestAssured,
		"verify",   // Mockito,
	}
)

var TechStopWords = []string{
	"get",
	"create",
	"update",
	"delete",
	"save",
	"post",

	"add",
	"remove",
	"insert",
	"select",

	"exist",
	"find",
	"new",
	"parse",

	"set",
	"get",
	"first",
	"last",

	"type",
	"key",
	"value",
	"equal",
	"greater",
	"greater",

	"all",
	"by",
	"id",
	"is",
	"of",
	"not",
	"with",
	"main",

	"status",
	"count",
	"equals",
	"start",
	"config",
	"sort",
	"handle",
	"handler",
	"internal",
	"cache",
	"request",
	"process",

	"parameter",
	"method",
	"class",
	"default",
	"object",
	"annotation",

	"read",
	"write",

	"bean",
	"message",
	"factory",
	"error",
	"errors",
	"exception",
	"null",
	"string",
	"init",
	"data",
	"hash",
	"convert",
	"size",
	"build",
	"return",
}

