package utils

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Pagination struct {
	Page_number *int    `form:"page_number" json:"page_number" validate:"gte=1"`
	Page_size   *int    `form:"page_size" 	json:"page_size"   validate:"gte=1"`
	Sort_column *string `form:"sort_column" json:"-"`
	Sort_order  *string `form:"order"       json:"-"			validate:"asc | ASC | desc | DESC"`
	Pages       *int64  `form:"-" json:"pages"`
	Count       *int64  `form:"-" json:"count"`
}

type PaginationMeta struct {
}

var pageSize int = 5
var pageNumber int = 1

var sortColumn string = "created_at"
var sortOrder = "ASC"

func GetPagination(
	c *gin.Context,
	query *gorm.DB,
) (pq *gorm.DB, pm *Pagination, err error) {
	pagination := Pagination{}

	// pageSize, ok := c.GetQuery("page_size")

	if err := c.ShouldBindQuery(&pagination); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return nil, nil, err
	} else {
		validate := validator.New()

		if err := validate.Struct(&pagination); err != nil {
			c.Error(err).SetType(gin.ErrorTypeBind)
			return nil, nil, err
		}
	}

	if pagination.Sort_column == nil {
		pagination.Sort_column = &sortColumn
	}

	if pagination.Sort_order == nil {
		pagination.Sort_order = &sortOrder
	}

	if pagination.Page_size == nil {
		pagination.Page_size = &pageSize
	}

	if pagination.Page_number == nil {
		pagination.Page_number = &pageNumber
	}

	log.Debug().Msg(*pagination.Sort_column)
	log.Debug().Msg(fmt.Sprintf("%v", *pagination.Page_size))
	log.Debug().Msg(*pagination.Sort_order)
	var count int64

	if err := query.Count(&count).Error; err != nil {
		c.Error(err).SetType(ErrorTypeDB)
		return nil, nil, err
	}

	order := fmt.Sprintf(
		"%v %v",
		*pagination.Sort_column,
		*pagination.Sort_order,
	)

	pagination.Count = &count
	pages := (*pagination.Count / int64(*pagination.Page_size)) + 1
	offset := (*pagination.Page_number - 1) * (*pagination.Page_size)

	pagination.Pages = &pages

	return query.Limit(*pagination.Page_size).
			Offset(offset).
			Order(order),
		&pagination, nil
}

func takeSliceArg(arg interface{}) (out []interface{}, ok bool) {
	slice, success := takeArg(arg, reflect.Slice)
	if !success {
		ok = false
		return
	}
	c := slice.Len()
	out = make([]interface{}, c)
	for i := 0; i < c; i++ {
		modelValue := reflect.ValueOf(slice.Index(i))
		typ := modelValue.Type()
		obj := reflect.New(typ).Interface()
		out[i] = obj
	}
	return out, true
}

func takeArg(
	arg interface{},
	kind reflect.Kind,
) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}
