package orm

import (
	"fmt"
	"strings"
)

// generate limit sql.
func getLimitSQL(mi *modelInfo, offset int, limit int) (limits string) {
	if limit == 0 {
		limit = DefaultRowsLimit
	}
	if limit < 0 {
		// no limit
		if offset > 0 {
			maxLimit := int64(DefaultRowsLimit)
			if maxLimit == 0 {
				limits = fmt.Sprintf("OFFSET %d", offset)
			} else {
				limits = fmt.Sprintf("LIMIT %d OFFSET %d", maxLimit, offset)
			}
		}
	} else if offset <= 0 {
		limits = fmt.Sprintf("LIMIT %d", limit)
	} else {
		limits = fmt.Sprintf("LIMIT %d OFFSET %d", limit, offset)
	}
	return
}

// generate order sql.
func getOrderSQL(orders []string) (orderSQL string) {
	if len(orders) == 0 {
		return
	}

	orderSqls := make([]string, 0, len(orders))
	for _, order := range orders {
		asc := "ASC"
		if order[0] == '-' {
			asc = "DESC"
			order = order[1:]
		}

		orderSqls = append(orderSqls, fmt.Sprintf("%s %s", order, asc))
	}

	orderSQL = fmt.Sprintf("ORDER BY %s ", strings.Join(orderSqls, ", "))
	return
}
