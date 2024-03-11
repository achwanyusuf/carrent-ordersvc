package order

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/achwanyusuf/carrent-lib/pkg/errormsg"
	"github.com/achwanyusuf/carrent-lib/pkg/logger"
	"github.com/achwanyusuf/carrent-ordersvc/src/model"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/svcerr"
	"github.com/achwanyusuf/carrent-ordersvc/src/usecase/order"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

type OrderDep struct {
	log      logger.Logger
	order    order.OrderInterface
	conf     Conf
	validate *validator.Validate
}

type Conf struct{}

type OrderInterface interface {
	Create(ctx *gin.Context)
	UpdateByID(ctx *gin.Context)
	Read(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	DeleteByID(ctx *gin.Context)
}

func New(conf Conf, log *logger.Logger, c order.OrderInterface, validate *validator.Validate) OrderInterface {
	return &OrderDep{
		conf:     conf,
		log:      *log,
		order:    c,
		validate: validate,
	}
}

// Create Order godoc
// @Summary Create Order
// @Description Create order data
// @Tags order
// @Accept json
// @Produce json
// @Security OAuth2Password
// @Param data body model.CreateOrder true "Order Data"
// @Success 200 {object} model.SingleOrderResponse
// @Success 400 {object} model.SingleOrderResponse
// @Success 500 {object} model.SingleOrderResponse
// @Router /order [post]
func (o *OrderDep) Create(ctx *gin.Context) {
	var (
		orderInput model.CreateOrder
		result     model.Order
		response   model.SingleOrderResponse
	)

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		statusCode := response.Transform(ctx, o.log, http.StatusCreated, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error read body"))
		ctx.JSON(statusCode, response)
		return
	}

	if err = json.Unmarshal(body, &orderInput); err != nil {
		statusCode := response.Transform(ctx, o.log, http.StatusCreated, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error unmarshal body"))
		ctx.JSON(statusCode, response)
		return
	}

	if err = o.validate.Struct(orderInput); err != nil {
		statusCode := response.Transform(ctx, o.log, http.StatusCreated, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error validate struct"))
		ctx.JSON(statusCode, response)
		return
	}

	orderInput.CreatedBy = ctx.Value("id").(int64)
	result, err = o.order.Create(ctx, orderInput)
	if err != nil {
		statusCode := response.Transform(ctx, o.log, http.StatusCreated, err)
		ctx.JSON(statusCode, response)
		return
	}

	response.Data = result

	statusCode := response.Transform(ctx, o.log, http.StatusCreated, nil)
	ctx.JSON(statusCode, response)
}

// Update Order Data godoc
// @Summary Update order data
// @Description Update order data
// @Tags order
// @Accept json
// @Produce json
// @Security OAuth2Password
// @Param id path string true "update by id"
// @Param data body model.UpdateOrder true "Order Data"
// @Success 200 {object} model.SingleOrderResponse
// @Success 400 {object} model.SingleOrderResponse
// @Success 500 {object} model.SingleOrderResponse
// @Router /order/{id} [put]
func (o *OrderDep) UpdateByID(ctx *gin.Context) {
	var (
		updateData model.UpdateOrder
		response   model.SingleOrderResponse
	)

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		statusCode := response.Transform(ctx, o.log, http.StatusOK, errormsg.WrapErr(errormsg.Error400, err, "error get id"))
		ctx.JSON(statusCode, response)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		statusCode := response.Transform(ctx, o.log, http.StatusOK, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error read body"))
		ctx.JSON(statusCode, response)
		return
	}

	if err = json.Unmarshal(body, &updateData); err != nil {
		statusCode := response.Transform(ctx, o.log, http.StatusOK, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error unmarshal body"))
		ctx.JSON(statusCode, response)
		return
	}
	updateData.UpdatedBy = ctx.Value("id").(int64)
	scope := ctx.Value("scope").(string)
	if scope != model.SuperAdminScope {
		id = ctx.Value("id").(int64)
	}
	if err = o.validate.Struct(updateData); err != nil {
		statusCode := response.Transform(ctx, o.log, http.StatusOK, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error validate struct"))
		ctx.JSON(statusCode, response)
		return
	}
	result, err := o.order.UpdateByID(ctx, id, updateData)
	if err != nil {
		statusCode := response.Transform(ctx, o.log, http.StatusOK, err)
		ctx.JSON(statusCode, response)
		return
	}

	response.Data = result

	statusCode := response.Transform(ctx, o.log, http.StatusOK, nil)
	ctx.JSON(statusCode, response)
}

// Get Orders Data godoc
// @Summary Get orders data
// @Description Get orders data
// @Tags order
// @Accept json
// @Produce json
// @Security OAuth2Password
// @Param id query number false "search by id"
// @Param car_id query number false "search by car id"
// @Param order_date query string false "search by order date"
// @Param pickup_date query string false "search by pickup date"
// @Param dropoff_date query string false "search by dropoff date"
// @Param pickup_location query string false "search by pickup location"
// @Param pickup_lat query string false "search by lat"
// @Param pickup_long query string false "search by long"
// @Param dropoff_location query string false "search by dropoff location"
// @Param dropoff_lat query string false "search by lat"
// @Param dropoff_long query string false "search by long"
// @Param page query int false " "
// @Param limit query int false " "
// @Param Cache-Control header string false "Request Cache Control" Enums(must-revalidate, none)
// @Success 200 {object} model.OrdersResponse
// @Success 400 {object} model.OrdersResponse
// @Success 500 {object} model.OrdersResponse
// @Router /order [get]
func (o *OrderDep) Read(ctx *gin.Context) {
	var (
		param    model.GetOrdersByParam
		response model.OrdersResponse
	)
	cacheControl := ctx.GetHeader("Cache-Control")
	var decoder = schema.NewDecoder()
	err := decoder.Decode(&param, ctx.Request.URL.Query())
	if err != nil {
		statusCode := response.Transform(ctx, o.log, http.StatusOK, err)
		ctx.JSON(statusCode, response)
		return
	}
	orders, pagination, err := o.order.GetByParam(ctx, cacheControl, param)
	if err != nil {
		statusCode := response.Transform(ctx, o.log, http.StatusOK, err)
		ctx.JSON(statusCode, response)
		return
	}

	response.Data = orders
	response.Pagination = pagination

	statusCode := response.Transform(ctx, o.log, http.StatusOK, nil)
	ctx.JSON(statusCode, response)
}

// Get Orders Data godoc
// @Summary Get orders data
// @Description Get orders data
// @Tags order
// @Accept json
// @Produce json
// @Security OAuth2Password
// @Param id path string true "get by id"
// @Param Cache-Control header string false "Request Cache Control" Enums(must-revalidate, none)
// @Success 200 {object} model.SingleOrderResponse
// @Success 400 {object} model.SingleOrderResponse
// @Success 500 {object} model.SingleOrderResponse
// @Router /order/{id} [get]
func (o *OrderDep) GetByID(ctx *gin.Context) {
	var response model.SingleOrderResponse
	cacheControl := ctx.GetHeader("Cache-Control")
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		statusCode := response.Transform(ctx, o.log, http.StatusOK, errormsg.WrapErr(errormsg.Error400, err, "error get id"))
		ctx.JSON(statusCode, response)
		return
	}
	result, err := o.order.GetByID(ctx, cacheControl, id)
	if err != nil {
		statusCode := response.Transform(ctx, o.log, http.StatusOK, err)
		ctx.JSON(statusCode, response)
		return
	}

	response.Data = result

	statusCode := response.Transform(ctx, o.log, http.StatusOK, nil)
	ctx.JSON(statusCode, response)
}

// Delete Order Data godoc
// @Summary Delete order data
// @Description Delete order data
// @Tags order
// @Accept json
// @Produce json
// @Security OAuth2Password
// @Param id path string true "delete by id"
// @Success 200 {object} model.EmptyResponse
// @Success 400 {object} model.EmptyResponse
// @Success 500 {object} model.EmptyResponse
// @Router /order/{id} [delete]
func (o *OrderDep) DeleteByID(ctx *gin.Context) {
	var (
		response model.EmptyResponse
	)
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		statusCode := response.Transform(ctx, o.log, http.StatusOK, errormsg.WrapErr(errormsg.Error400, err, "error get id"))
		ctx.JSON(statusCode, response)
		return
	}
	scope := ctx.Value("scope").(string)
	if scope != model.SuperAdminScope {
		id = ctx.Value("id").(int64)
	}
	err = o.order.DeleteByID(ctx, ctx.Value("id").(int64), id)
	if err != nil {
		statusCode := response.Transform(ctx, o.log, http.StatusOK, err)
		ctx.JSON(statusCode, response)
		return
	}

	statusCode := response.Transform(ctx, o.log, http.StatusOK, nil)
	ctx.JSON(statusCode, response)
}
