//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build ignore

package rtl

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// 解析 dfm或者lfm资源文件的

type Char = byte

const (
	toEOF     = Char(0)
	toSymbol  = Char(1)
	toString  = Char(2)
	toInteger = Char(3)
	toFloat   = Char(4)
	toWString = Char(5)
)

var FilerSignature = []byte("TPF0")

type TParser struct {
	sfBuf         []byte
	fToken        Char
	fLastTokenStr string
	fPos          int
	fFloatType    Char
}

func NewParser(buff []byte) *TParser {
	p := new(TParser)
	p.sfBuf = buff
	p.fToken = 0
	return p
}

func (p *TParser) CheckToken(t Char) {
	if p.fToken != t {
		//ErrorFmt(SParWrongTokenType,[GetTokenName(T),GetTokenName(fToken)]);
	}
}

func (p *TParser) CheckTokenSymbol(s string) {
	p.CheckToken(toSymbol)
	if strings.Compare(p.fLastTokenStr, s) != 0 {
		//ErrorFmt(SParWrongTokenSymbol,[s,fLastTokenStr]);
	}
}
func (p *TParser) Error(ident string) {
	//ErrorStr(Ident)
}

func (p *TParser) ErrorFmt(ident string, args ...interface{}) {
	p.ErrorStr(fmt.Sprintf(ident, args...))
}

func (p *TParser) ErrorStr(message string) {
	//panic(p.EParserError.CreateFmt(Message+SParLocInfo,[SourceLine,fPos+fDeltaPos,SourcePos]);
}

func (p *TParser) HexToBinary(Stream *bytes.Buffer) {
	//	var outbuf : array[0..ParseBufSize-1] of byte;
	//b : byte;
	//i : integer;
	//	begin
	//	i:=0;
	//	SkipWhitespace;
	//	while IsHexNum do
	//	begin
	//	b:=(GetHexValue(fBuf[fPos]) shl 4);
	//	inc(fPos);
	//	CheckLoadBuffer;
	//	if not IsHexNum then
	//	Error(SParUnterminatedBinValue);
	//	b:=b or GetHexValue(fBuf[fPos]);
	//	inc(fPos);
	//	CheckLoadBuffer;
	//	outbuf[i]:=b;
	//	inc(i);
	//	if i>=ParseBufSize then
	//	begin
	//	Stream.WriteBuffer(outbuf[0],i);
	//	i:=0;
	//	end;
	//	SkipWhitespace;
	//	end;
	//	if i>0 then
	//	Stream.WriteBuffer(outbuf[0],i);
	//	NextToken;
}

func (p *TParser) NextToken() Char {
	//	SkipWhiteSpace;
	//	if fEofReached then
	//	HandleEof
	//	else
	//case fBuf[fPos] of
	//	'_','A'..'Z','a'..'z' : HandleAlphaNum;
	//	'$'                   : HandleHexNumber;
	//	'-'                   : HandleMinus;
	//	'0'..'9'              : HandleNumber;
	//	'''','#'              : HandleString
	//	else
	//	HandleUnknown;
	//	end;
	return p.fToken
}

func (p *TParser) SourcePos() int {
	return 0 //return p.fStream.Position // - fBufLen + fPos
}

func (p *TParser) TokenComponentIdent() string {
	if p.fToken != toSymbol {
		//p.ErrorFmt(SParExpected,[GetTokenName(toSymbol)]);
	}
	p.CheckLoadBuffer()
	for p.sfBuf[p.fPos] == '.' {
		ProcessChar
		p.fLastTokenStr = p.fLastTokenStr + p.GetAlphaNum
	}
	return p.fLastTokenStr
}

func (p *TParser) TokenFloat() float64 {
	result, err := strconv.ParseFloat(p.fLastTokenStr, 64)
	if err != nil {
		//ErrorFmt(SParInvalidFloat,[fLastTokenStr]);
	}
	return result
}

func (p *TParser) TokenInt() int64 {
	result, err := strconv.ParseInt(p.fLastTokenStr, 10, 64)
	if err != nil {
		// 这里要将 $ 替换为 0x
		result, err = strconv.ParseInt(p.fLastTokenStr, 16, 64)
	}
	return result
}

func (p *TParser) TokenString() string {
	switch p.fToken {
	case toWString:
		//Result:=string(fLastTokenWStr);
		return p.fLastTokenStr
	case toFloat:
		if p.fFloatType != 0 {
			return p.fLastTokenStr + string(p.fFloatType)
		} else {
			return p.fLastTokenStr
		}
	default:

		return p.fLastTokenStr
	}
}
