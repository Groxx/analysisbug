package ignored

import (
	"testing"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestIt(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(),
		&analysis.Analyzer{
			Name: "test",
			Doc:  "test",
			Run: func(pass *analysis.Pass) (interface{}, error) {
				return nil, nil
			},
		},
		"analysis")
}
