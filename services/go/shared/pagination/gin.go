package pagination

import (
	"github.com/gin-gonic/gin"
	"github.com/lefinal/meh"
	"github.com/lefinal/nulls"
	"strconv"
)

// Query param names for usage in ParamsFromRequest.
var (
	LimitQueryParam    = "limit"
	OffsetQueryParam   = "offset"
	OrderByQueryParam  = "order_by"
	OrderDirQueryParam = "order_dir"
)

// ParamsFromRequest extracts Params from the given gin.Context using query
// parameters. In case of invalid format or other problems, an meh.ErrBadInput
// is returned.
func ParamsFromRequest(c *gin.Context) (Params, error) {
	var params Params
	// Extract limit.
	limitStr := c.Query(LimitQueryParam)
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return Params{}, meh.NewBadInputErrFromErr(err, "parse limit", meh.Details{"was": limitStr})
		}
		params.Limit = nulls.NewInt(limit)
	}
	// Extract offset.
	offsetStr := c.Query(OffsetQueryParam)
	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return Params{}, meh.NewBadInputErrFromErr(err, "parse offset", meh.Details{"was": offsetStr})
		}
		params.Offset = offset
	}
	// Extract ordering.
	orderBy := c.Query(OrderByQueryParam)
	if orderBy != "" {
		params.OrderBy = nulls.NewString(orderBy)
	}
	params.OrderDirection = c.Query(OrderDirQueryParam)
	return params, nil
}
