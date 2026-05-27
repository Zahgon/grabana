package golang

import (
	"github.com/K-Phoen/jennifer/jen"
	"github.com/K-Phoen/sdk"
)

func (encoder *Encoder) encodeTargets(targets []sdk.Target, grabanaPackage string) []jen.Code {
	_ = "STUB: not implemented"
	return nil
}

func (encoder *Encoder) encodeTarget(target sdk.Target, grabanaPackage string) jen.Code {
	_ = "STUB: not implemented"
	// looks like a prometheus target
	return *new(jen.Code)
}

/*
	// looks like graphite
	if target.Target != "" {
		return encoder.encodeGraphiteTarget(target)
	}

	// looks like influxdb
	if target.Measurement != "" {
		return encoder.encodeInfluxDBTarget(target)
	}

	// looks like stackdriver
	if target.MetricType != "" {
		return encoder.encodeStackdriverTarget(target)
	}
*/

func (encoder *Encoder) encodePrometheusTarget(target sdk.Target, grabanaPackage string) jen.Code {
	_ = "STUB: not implemented"
	return *new(jen.Code)
}

// only emit code if the default isn't used
