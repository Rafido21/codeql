// Code generated by https://github.com/gagliardetto/codebox. DO NOT EDIT.

package main

import (
	"io"
	"log"
)

func TaintStepTest_LogNew_B0I0O0(sourceCQL interface{}) interface{} {
	fromLogger656 := sourceCQL.(*log.Logger)
	var intoWriter414 io.Writer
	intermediateCQL := log.New(intoWriter414, "", 0)
	link(fromLogger656, intermediateCQL)
	return intoWriter414
}

func TaintStepTest_LogLoggerFatal_B0I0O0(sourceCQL interface{}) interface{} {
	fromInterface518 := sourceCQL.(interface{})
	var intoLogger650 log.Logger
	intoLogger650.Fatal(fromInterface518)
	return intoLogger650
}

func TaintStepTest_LogLoggerFatalf_B0I0O0(sourceCQL interface{}) interface{} {
	fromString784 := sourceCQL.(string)
	var intoLogger957 log.Logger
	intoLogger957.Fatalf(fromString784, nil)
	return intoLogger957
}

func TaintStepTest_LogLoggerFatalf_B0I1O0(sourceCQL interface{}) interface{} {
	fromInterface520 := sourceCQL.(interface{})
	var intoLogger443 log.Logger
	intoLogger443.Fatalf("", fromInterface520)
	return intoLogger443
}

func TaintStepTest_LogLoggerFatalln_B0I0O0(sourceCQL interface{}) interface{} {
	fromInterface127 := sourceCQL.(interface{})
	var intoLogger483 log.Logger
	intoLogger483.Fatalln(fromInterface127)
	return intoLogger483
}

func TaintStepTest_LogLoggerPanic_B0I0O0(sourceCQL interface{}) interface{} {
	fromInterface989 := sourceCQL.(interface{})
	var intoLogger982 log.Logger
	intoLogger982.Panic(fromInterface989)
	return intoLogger982
}

func TaintStepTest_LogLoggerPanicf_B0I0O0(sourceCQL interface{}) interface{} {
	fromString417 := sourceCQL.(string)
	var intoLogger584 log.Logger
	intoLogger584.Panicf(fromString417, nil)
	return intoLogger584
}

func TaintStepTest_LogLoggerPanicf_B0I1O0(sourceCQL interface{}) interface{} {
	fromInterface991 := sourceCQL.(interface{})
	var intoLogger881 log.Logger
	intoLogger881.Panicf("", fromInterface991)
	return intoLogger881
}

func TaintStepTest_LogLoggerPanicln_B0I0O0(sourceCQL interface{}) interface{} {
	fromInterface186 := sourceCQL.(interface{})
	var intoLogger284 log.Logger
	intoLogger284.Panicln(fromInterface186)
	return intoLogger284
}

func TaintStepTest_LogLoggerPrint_B0I0O0(sourceCQL interface{}) interface{} {
	fromInterface908 := sourceCQL.(interface{})
	var intoLogger137 log.Logger
	intoLogger137.Print(fromInterface908)
	return intoLogger137
}

func TaintStepTest_LogLoggerPrintf_B0I0O0(sourceCQL interface{}) interface{} {
	fromString494 := sourceCQL.(string)
	var intoLogger873 log.Logger
	intoLogger873.Printf(fromString494, nil)
	return intoLogger873
}

func TaintStepTest_LogLoggerPrintf_B0I1O0(sourceCQL interface{}) interface{} {
	fromInterface599 := sourceCQL.(interface{})
	var intoLogger409 log.Logger
	intoLogger409.Printf("", fromInterface599)
	return intoLogger409
}

func TaintStepTest_LogLoggerPrintln_B0I0O0(sourceCQL interface{}) interface{} {
	fromInterface246 := sourceCQL.(interface{})
	var intoLogger898 log.Logger
	intoLogger898.Println(fromInterface246)
	return intoLogger898
}

func TaintStepTest_LogLoggerSetOutput_B0I0O0(sourceCQL interface{}) interface{} {
	fromLogger598 := sourceCQL.(log.Logger)
	var intoWriter631 io.Writer
	fromLogger598.SetOutput(intoWriter631)
	return intoWriter631
}

func TaintStepTest_LogLoggerSetPrefix_B0I0O0(sourceCQL interface{}) interface{} {
	fromString165 := sourceCQL.(string)
	var intoLogger150 log.Logger
	intoLogger150.SetPrefix(fromString165)
	return intoLogger150
}

func TaintStepTest_LogLoggerWriter_B0I0O0(sourceCQL interface{}) interface{} {
	fromLogger340 := sourceCQL.(log.Logger)
	intoWriter471 := fromLogger340.Writer()
	return intoWriter471
}

func RunAllTaints_Log() {
	{
		source := newSource(0)
		out := TaintStepTest_LogNew_B0I0O0(source)
		sink(0, out)
	}
	{
		source := newSource(1)
		out := TaintStepTest_LogLoggerFatal_B0I0O0(source)
		sink(1, out)
	}
	{
		source := newSource(2)
		out := TaintStepTest_LogLoggerFatalf_B0I0O0(source)
		sink(2, out)
	}
	{
		source := newSource(3)
		out := TaintStepTest_LogLoggerFatalf_B0I1O0(source)
		sink(3, out)
	}
	{
		source := newSource(4)
		out := TaintStepTest_LogLoggerFatalln_B0I0O0(source)
		sink(4, out)
	}
	{
		source := newSource(5)
		out := TaintStepTest_LogLoggerPanic_B0I0O0(source)
		sink(5, out)
	}
	{
		source := newSource(6)
		out := TaintStepTest_LogLoggerPanicf_B0I0O0(source)
		sink(6, out)
	}
	{
		source := newSource(7)
		out := TaintStepTest_LogLoggerPanicf_B0I1O0(source)
		sink(7, out)
	}
	{
		source := newSource(8)
		out := TaintStepTest_LogLoggerPanicln_B0I0O0(source)
		sink(8, out)
	}
	{
		source := newSource(9)
		out := TaintStepTest_LogLoggerPrint_B0I0O0(source)
		sink(9, out)
	}
	{
		source := newSource(10)
		out := TaintStepTest_LogLoggerPrintf_B0I0O0(source)
		sink(10, out)
	}
	{
		source := newSource(11)
		out := TaintStepTest_LogLoggerPrintf_B0I1O0(source)
		sink(11, out)
	}
	{
		source := newSource(12)
		out := TaintStepTest_LogLoggerPrintln_B0I0O0(source)
		sink(12, out)
	}
	{
		source := newSource(13)
		out := TaintStepTest_LogLoggerSetOutput_B0I0O0(source)
		sink(13, out)
	}
	{
		source := newSource(14)
		out := TaintStepTest_LogLoggerSetPrefix_B0I0O0(source)
		sink(14, out)
	}
	{
		source := newSource(15)
		out := TaintStepTest_LogLoggerWriter_B0I0O0(source)
		sink(15, out)
	}
}