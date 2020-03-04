// Code generated by goyacc -o parser.go parser.go.y. DO NOT EDIT.

//line parser.go.y:2
package sql

import __yyfmt__ "fmt"

//line parser.go.y:2

//line parser.go.y:5
type yySymType struct {
	yys                  int
	literal              string
	identifier           string
	signedNumber         int64
	statement            interface{}
	columnNameList       []string
	columnName           string
	columnDefList        []ColumnDef
	columnDef            ColumnDef
	indexedColumnList    []IndexedColumn
	indexedColumn        IndexedColumn
	name                 string
	withoutRowid         bool
	unique               bool
	bool                 bool
	collate              string
	sortOrder            SortOrder
	columnConstraint     columnConstraint
	columnConstraintList []columnConstraint
	tableConstraint      TableConstraint
	tableConstraintList  []TableConstraint
	triggerAction        TriggerAction
	trigger              Trigger
	triggerList          []Trigger
	where                Expression
	expr                 Expression
	exprList             []Expression
	float                float64
}

const ABORT = 57346
const ACTION = 57347
const AND = 57348
const ASC = 57349
const AUTOINCREMENT = 57350
const CASCADE = 57351
const COLLATE = 57352
const CONFLICT = 57353
const CONSTRAINT = 57354
const CREATE = 57355
const DEFAULT = 57356
const DEFERRABLE = 57357
const DEFERRED = 57358
const DELETE = 57359
const DESC = 57360
const FAIL = 57361
const FOREIGN = 57362
const FROM = 57363
const GLOB = 57364
const IGNORE = 57365
const IN = 57366
const INDEX = 57367
const INITIALLY = 57368
const IS = 57369
const KEY = 57370
const LIKE = 57371
const MATCH = 57372
const NO = 57373
const NOT = 57374
const NULL = 57375
const ON = 57376
const OR = 57377
const PRIMARY = 57378
const REFERENCES = 57379
const REGEXP = 57380
const REPLACE = 57381
const RESTRICT = 57382
const ROLLBACK = 57383
const ROWID = 57384
const SELECT = 57385
const SET = 57386
const TABLE = 57387
const UNIQUE = 57388
const UPDATE = 57389
const WHERE = 57390
const WITHOUT = 57391
const tBare = 57392
const tLiteral = 57393
const tIdentifier = 57394
const tOperator = 57395
const tSignedNumber = 57396
const tFloat = 57397

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ABORT",
	"ACTION",
	"AND",
	"ASC",
	"AUTOINCREMENT",
	"CASCADE",
	"COLLATE",
	"CONFLICT",
	"CONSTRAINT",
	"CREATE",
	"DEFAULT",
	"DEFERRABLE",
	"DEFERRED",
	"DELETE",
	"DESC",
	"FAIL",
	"FOREIGN",
	"FROM",
	"GLOB",
	"IGNORE",
	"IN",
	"INDEX",
	"INITIALLY",
	"IS",
	"KEY",
	"LIKE",
	"MATCH",
	"NO",
	"NOT",
	"NULL",
	"ON",
	"OR",
	"PRIMARY",
	"REFERENCES",
	"REGEXP",
	"REPLACE",
	"RESTRICT",
	"ROLLBACK",
	"ROWID",
	"SELECT",
	"SET",
	"TABLE",
	"UNIQUE",
	"UPDATE",
	"WHERE",
	"WITHOUT",
	"tBare",
	"tLiteral",
	"tIdentifier",
	"tOperator",
	"tSignedNumber",
	"tFloat",
	"'-'",
	"'+'",
	"','",
	"'('",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 80,
	59, 6,
	-2, 88,
	-1, 81,
	59, 7,
	-2, 89,
}

const yyPrivate = 57344

const yyLast = 203

var yyAct = [...]int{

	76, 163, 9, 128, 91, 74, 10, 71, 77, 130,
	150, 64, 78, 137, 18, 156, 72, 10, 21, 137,
	23, 138, 75, 26, 98, 29, 134, 32, 32, 34,
	26, 133, 98, 132, 127, 98, 58, 99, 55, 80,
	79, 81, 111, 65, 85, 83, 84, 62, 82, 109,
	97, 102, 96, 105, 104, 103, 63, 70, 125, 68,
	69, 87, 38, 65, 39, 66, 67, 65, 85, 83,
	84, 50, 51, 48, 102, 94, 95, 104, 103, 65,
	117, 66, 67, 22, 17, 11, 37, 12, 106, 161,
	27, 41, 94, 95, 90, 110, 107, 108, 11, 142,
	12, 68, 69, 35, 54, 118, 114, 160, 120, 121,
	122, 124, 10, 119, 129, 115, 27, 126, 6, 162,
	52, 16, 136, 131, 11, 170, 12, 165, 28, 89,
	53, 144, 59, 46, 13, 15, 40, 47, 10, 139,
	141, 59, 56, 148, 169, 61, 145, 25, 5, 167,
	146, 10, 88, 129, 153, 45, 44, 86, 166, 42,
	60, 31, 164, 155, 168, 19, 147, 92, 143, 43,
	49, 8, 158, 152, 37, 140, 101, 36, 93, 113,
	171, 135, 154, 151, 123, 73, 116, 157, 20, 33,
	159, 30, 112, 100, 57, 14, 24, 7, 149, 4,
	3, 2, 1,
}
var yyPact = [...]int{

	105, -1000, -1000, -1000, -1000, 35, 89, 63, -1000, -1000,
	-1000, -1000, -1000, 35, 140, -1000, 35, 35, 24, 35,
	-1000, -1000, 48, 94, -33, -1000, 35, 35, 35, 74,
	4, 123, 14, 123, 12, 84, -1000, 35, 162, -13,
	123, -1000, 132, -1000, -1000, 112, 35, 9, 25, 123,
	-11, -1000, 129, 2, 124, -1000, 84, -1000, 52, -1000,
	160, -1000, -1000, -1000, -1000, -1000, 25, 25, -1000, -1000,
	-8, -23, -1000, 166, 21, -1000, -6, -1000, -1000, -1000,
	-1000, -1000, -11, 13, 13, -1000, -10, -11, -17, -1000,
	-1000, 171, -1000, -1000, -1000, -1000, -1000, 25, -11, 32,
	160, 51, -11, -11, -11, -11, -2, -1000, -1000, -11,
	-26, 35, -1000, -1000, -51, -1000, -1000, -11, -1000, -1000,
	21, 21, 21, -27, 21, -1000, -34, 88, -39, -1000,
	-1000, 21, -1000, -11, -1000, -1000, 164, 35, 62, 21,
	127, -1000, 35, -1000, -1000, -1000, -1000, -1000, -49, 158,
	35, 137, -1000, -45, -1000, 156, -1000, 73, -1000, -1000,
	72, 118, 118, -1000, 111, -1000, -1000, 175, -1000, -1000,
	-1000, -1000,
}
var yyPgo = [...]int{

	0, 202, 201, 200, 199, 0, 11, 8, 12, 2,
	171, 3, 198, 197, 196, 147, 7, 16, 161, 103,
	195, 194, 193, 4, 192, 91, 136, 72, 191, 1,
	190, 187, 186, 185, 5, 184, 183, 182, 181,
}
var yyR1 = [...]int{

	0, 1, 1, 1, 6, 6, 5, 5, 7, 7,
	7, 8, 8, 8, 9, 11, 11, 12, 10, 13,
	13, 25, 25, 25, 25, 25, 25, 25, 26, 26,
	26, 27, 27, 27, 19, 19, 28, 28, 28, 24,
	24, 14, 14, 15, 15, 18, 18, 18, 18, 22,
	22, 23, 23, 23, 21, 21, 20, 20, 38, 38,
	38, 38, 38, 38, 16, 16, 33, 17, 29, 29,
	29, 29, 29, 30, 30, 31, 31, 36, 36, 37,
	37, 32, 32, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 35, 35, 35, 2, 3, 4,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	2, 1, 2, 2, 1, 1, 3, 3, 1, 1,
	3, 4, 1, 1, 2, 2, 2, 2, 0, 1,
	2, 5, 5, 11, 0, 2, 0, 3, 4, 0,
	1, 1, 3, 3, 3, 0, 1, 4, 6, 0,
	2, 0, 1, 1, 0, 2, 0, 1, 0, 3,
	3, 3, 3, 3, 1, 3, 1, 3, 2, 2,
	1, 1, 2, 3, 3, 0, 2, 0, 1, 0,
	2, 0, 2, 1, 4, 1, 1, 1, 1, 1,
	3, 3, 3, 3, 0, 1, 3, 4, 8, 10,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, 43, 13, -13, -10, -9,
	-5, 50, 52, 45, -20, 46, 58, 21, -5, 25,
	-10, -5, 59, -5, -14, -15, -5, 42, 34, 58,
	-28, -18, -5, -18, -5, -19, -15, 12, 58, 60,
	-26, -25, 36, 46, 33, 32, 10, 14, 59, -26,
	59, -27, 36, 46, 20, -5, -19, -21, 49, -25,
	28, 33, -5, -7, -6, 54, 56, 57, 50, 51,
	-7, -16, -17, -33, -34, 33, -5, -7, -8, 51,
	50, 52, 59, 56, 57, 55, 28, 59, 28, -27,
	42, -23, 7, 18, -7, -7, 60, 58, 58, 60,
	-22, 10, 53, 57, 56, 59, -34, -8, -8, 59,
	-16, 59, -24, 8, -7, -17, -32, 48, -23, -6,
	-34, -34, -34, -35, -34, 60, -16, 60, -11, -9,
	60, -34, 60, 58, 60, -38, 34, 58, 60, -34,
	11, -9, 37, 41, 4, 19, 23, 39, -5, -12,
	59, -36, 15, -11, -37, 26, 60, -31, 16, -30,
	34, 17, 47, -29, 44, 9, 40, 31, -29, 33,
	14, 5,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 0, 56, 0, 19, 18,
	14, 6, 7, 0, 0, 57, 0, 0, 0, 0,
	20, 97, 0, 0, 36, 41, 45, 45, 0, 34,
	0, 28, 46, 28, 0, 0, 42, 0, 34, 54,
	43, 29, 0, 22, 23, 0, 0, 0, 0, 44,
	0, 37, 0, 0, 0, 35, 0, 98, 0, 30,
	51, 24, 25, 26, 27, 8, 0, 0, 4, 5,
	0, 0, 64, 49, 66, 83, 0, 85, 86, 87,
	-2, -2, 0, 0, 0, 11, 0, 0, 0, 38,
	55, 39, 52, 53, 9, 10, 47, 0, 0, 81,
	51, 0, 0, 0, 0, 94, 0, 12, 13, 0,
	0, 0, 21, 40, 0, 65, 99, 0, 67, 50,
	90, 91, 92, 0, 95, 93, 0, 58, 0, 15,
	48, 82, 84, 0, 31, 32, 0, 0, 0, 96,
	0, 16, 0, 59, 60, 61, 62, 63, 0, 77,
	0, 79, 78, 0, 75, 0, 17, 33, 80, 76,
	0, 0, 0, 73, 0, 70, 71, 0, 74, 68,
	69, 72,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	59, 60, 3, 57, 58, 56,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:128
		{
			yyVAL.literal = yyDollar[1].identifier
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:131
		{
			yyVAL.literal = yyDollar[1].identifier
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:136
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:139
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:144
		{
			yyVAL.signedNumber = yyDollar[1].signedNumber
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:147
		{
			yyVAL.signedNumber = -yyDollar[2].signedNumber
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:150
		{
			yyVAL.signedNumber = yyDollar[2].signedNumber
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:155
		{
			yyVAL.float = yyDollar[1].float
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:158
		{
			yyVAL.float = -yyDollar[2].float
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:161
		{
			yyVAL.float = yyDollar[2].float
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:166
		{
			yyVAL.columnName = yyDollar[1].identifier
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:171
		{
			yyVAL.columnNameList = []string{yyDollar[1].columnName}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:174
		{
			yyVAL.columnNameList = append(yyDollar[1].columnNameList, yyDollar[3].columnName)
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:179
		{
			yyVAL.columnNameList = yyDollar[2].columnNameList
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:184
		{
			yyVAL.columnName = yyDollar[1].columnName
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:189
		{
			yyVAL.columnNameList = []string{yyDollar[1].columnName}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:192
		{
			yyVAL.columnNameList = append(yyDollar[1].columnNameList, yyDollar[3].columnName)
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:198
		{
			yyVAL.columnConstraint = ccPrimaryKey{yyDollar[3].sortOrder, yyDollar[4].bool}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:201
		{
			yyVAL.columnConstraint = ccUnique(true)
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:204
		{
			yyVAL.columnConstraint = ccNull(true)
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:207
		{
			yyVAL.columnConstraint = ccNull(false)
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:210
		{
			yyVAL.columnConstraint = ccCollate(yyDollar[2].identifier)
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:213
		{
			yyVAL.columnConstraint = ccDefault(yyDollar[2].signedNumber)
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:216
		{
			yyVAL.columnConstraint = ccDefault(yyDollar[2].literal)
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:221
		{
			yyVAL.columnConstraintList = nil
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:224
		{
			yyVAL.columnConstraintList = []columnConstraint{yyDollar[1].columnConstraint}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:227
		{
			yyVAL.columnConstraintList = append(yyDollar[1].columnConstraintList, yyDollar[2].columnConstraint)
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:232
		{
			yyVAL.tableConstraint = TablePrimaryKey{yyDollar[4].indexedColumnList}
		}
	case 32:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:235
		{
			yyVAL.tableConstraint = TableUnique{
				IndexedColumns: yyDollar[3].indexedColumnList,
			}
		}
	case 33:
		yyDollar = yyS[yypt-11 : yypt+1]
//line parser.go.y:240
		{
			yyVAL.tableConstraint = TableForeignKey{
				Columns:           yyDollar[4].columnNameList,
				ForeignTable:      yyDollar[7].identifier,
				ForeignColumns:    yyDollar[8].columnNameList,
				Deferrable:        yyDollar[9].bool,
				InitiallyDeferred: yyDollar[10].bool,
				Triggers:          yyDollar[11].triggerList,
			}
		}
	case 34:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:252
		{
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:253
		{
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:257
		{
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:258
		{
			yyVAL.tableConstraintList = []TableConstraint{yyDollar[3].tableConstraint}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:261
		{
			yyVAL.tableConstraintList = append(yyDollar[1].tableConstraintList, yyDollar[4].tableConstraint)
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:267
		{
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:268
		{
			yyVAL.bool = true
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:273
		{
			yyVAL.columnDefList = []ColumnDef{yyDollar[1].columnDef}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:276
		{
			yyVAL.columnDefList = append(yyDollar[1].columnDefList, yyDollar[3].columnDef)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:281
		{
			yyVAL.columnDef = makeColumnDef(yyDollar[1].identifier, yyDollar[2].name, yyDollar[3].columnConstraintList)
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:284
		{
			yyVAL.columnDef = makeColumnDef("ROWID", yyDollar[2].name, yyDollar[3].columnConstraintList)
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:289
		{
			yyVAL.name = ""
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:292
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 47:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:295
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 48:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:298
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 49:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:303
		{
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:304
		{
			yyVAL.collate = yyDollar[2].literal
		}
	case 51:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:309
		{
			yyVAL.sortOrder = Asc
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:312
		{
			yyVAL.sortOrder = Asc
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:315
		{
			yyVAL.sortOrder = Desc
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:320
		{
			yyVAL.withoutRowid = false
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:323
		{
			yyVAL.withoutRowid = true
		}
	case 56:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:328
		{
			yyVAL.unique = false
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:331
		{
			yyVAL.unique = true
		}
	case 58:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:336
		{
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:337
		{
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:339
		{
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:341
		{
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:343
		{
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:345
		{
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:349
		{
			yyVAL.indexedColumnList = []IndexedColumn{yyDollar[1].indexedColumn}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:352
		{
			yyVAL.indexedColumnList = append(yyDollar[1].indexedColumnList, yyDollar[3].indexedColumn)
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:357
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:362
		{
			yyVAL.indexedColumn = newIndexColumn(yyDollar[1].expr, yyDollar[2].collate, yyDollar[3].sortOrder)
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:367
		{
			yyVAL.triggerAction = ActionSetNull
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:370
		{
			yyVAL.triggerAction = ActionSetDefault
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:373
		{
			yyVAL.triggerAction = ActionCascade
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:376
		{
			yyVAL.triggerAction = ActionRestrict
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:379
		{
			yyVAL.triggerAction = ActionNoAction
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:384
		{
			yyVAL.trigger = TriggerOnDelete(yyDollar[3].triggerAction)
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:387
		{
			yyVAL.trigger = TriggerOnUpdate(yyDollar[3].triggerAction)
		}
	case 75:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:392
		{
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:393
		{
			yyVAL.triggerList = append(yyDollar[1].triggerList, yyDollar[2].trigger)
		}
	case 77:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:398
		{
			yyVAL.bool = false
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:401
		{
			yyVAL.bool = true
		}
	case 79:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:406
		{
			yyVAL.bool = false
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:409
		{
			yyVAL.bool = true
		}
	case 81:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:414
		{
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:415
		{
			yyVAL.where = yyDollar[2].expr
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:420
		{
			yyVAL.expr = nil
		}
	case 84:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:423
		{
			yyVAL.expr = ExFunction{yyDollar[1].identifier, yyDollar[3].exprList}
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:426
		{
			yyVAL.expr = yyDollar[1].signedNumber
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:429
		{
			yyVAL.expr = yyDollar[1].float
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:432
		{
			yyVAL.expr = yyDollar[1].identifier
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:435
		{
			yyVAL.expr = ExColumn(yyDollar[1].identifier)
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:438
		{
			yyVAL.expr = ExColumn(yyDollar[1].identifier)
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:441
		{
			yyVAL.expr = ExBinaryOp{yyDollar[2].identifier, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:444
		{
			yyVAL.expr = ExBinaryOp{"+", yyDollar[1].expr, yyDollar[3].expr}
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:447
		{
			yyVAL.expr = ExBinaryOp{"-", yyDollar[1].expr, yyDollar[3].expr}
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:450
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 94:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:455
		{
			yyVAL.exprList = nil
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:458
		{
			yyVAL.exprList = []Expression{yyDollar[1].expr}
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:461
		{
			yyVAL.exprList = append(yyDollar[1].exprList, yyDollar[3].expr)
		}
	case 97:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:466
		{
			yylex.(*lexer).result = SelectStmt{Columns: yyDollar[2].columnNameList, Table: yyDollar[4].identifier}
		}
	case 98:
		yyDollar = yyS[yypt-8 : yypt+1]
//line parser.go.y:471
		{
			yylex.(*lexer).result = CreateTableStmt{
				Table:        yyDollar[3].identifier,
				Columns:      yyDollar[5].columnDefList,
				Constraints:  yyDollar[6].tableConstraintList,
				WithoutRowid: yyDollar[8].withoutRowid,
			}
		}
	case 99:
		yyDollar = yyS[yypt-10 : yypt+1]
//line parser.go.y:481
		{
			yylex.(*lexer).result = CreateIndexStmt{
				Index:          yyDollar[4].identifier,
				Table:          yyDollar[6].identifier,
				Unique:         yyDollar[2].unique,
				IndexedColumns: yyDollar[8].indexedColumnList,
				Where:          yyDollar[10].where,
			}
		}
	}
	goto yystack /* stack new state and value */
}
