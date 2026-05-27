package golang

import (
	"github.com/K-Phoen/jennifer/jen"
	"github.com/K-Phoen/sdk"
)

func (encoder *Encoder) convertLogs(panel sdk.Panel) jen.Code {
	_ = "STUB: not implemented"
	return *new(jen.Code)
}

func (encoder *Encoder) encodeLogsVizualizationSettings(panel sdk.Panel) []jen.Code {
	_ = "STUB: not implemented"
	return nil
}

func (encoder *Encoder) encodeLogsDedupStrategy(sdkDedupStrategy string) jen.Code {
	_ = "STUB: not implemented"
	return *new(jen.Code)
}

func (encoder *Encoder) encodeLogsSortOrder(sdkSortOrder string) jen.Code {
	_ = "STUB: not implemented"
	return *new(jen.Code)
}

func (encoder *Encoder) encodeLogsTarget(target sdk.Target) jen.Code {
	_ = "STUB: not implemented"
	return *new(jen.Code)
}

func logsQual(name string) *jen.Statement { _ = "STUB: not implemented"; return nil }
