package stepImpl

import (
	"github.com/getgauge-contrib/gauge-go/gauge"
	m "github.com/getgauge-contrib/gauge-go/models"
	cmp "github.com/praveen4g0/comparator/pkg/comparator"
)

var _ = gauge.Step("Compare API response <table>", func(tbl *m.Table) {
	for _, row := range tbl.Rows {
		cmp.Compare(string(row.Cells[0]), string(row.Cells[1]))
	}
})
