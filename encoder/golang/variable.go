package golang

import (
	"github.com/K-Phoen/jennifer/jen"
	"github.com/K-Phoen/sdk"
)

func (encoder *Encoder) encodeVariables(variables []sdk.TemplateVar) []jen.Code {
	_ = "STUB: not implemented"
	return nil
}

func (encoder *Encoder) encodeVariable(variable sdk.TemplateVar) jen.Code {
	_ = "STUB: not implemented"
	return *

	/*
		case "interval":
			encoder.encodeIntervalVar(variable)
		case "custom":
			encoder.encodeCustomVar(variable)
		case "const":
			encoder.encodeConstVar(variable)

	*/new(jen.Code)
}

/*
	case "textbox":
		encoder.encodeTextVar(variable)
*/

func (encoder *Encoder) encodeQueryVar(variable sdk.TemplateVar) jen.Code {
	_ = "STUB: not implemented"
	return *new(jen.Code)
}

// TODO: eventually we should stop using legacy stuff... :|

func (encoder *Encoder) encodeDatasourceVar(variable sdk.TemplateVar) jen.Code {
	_ = "STUB: not implemented"
	return *new(jen.Code)
}
