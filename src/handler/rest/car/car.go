package car

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/achwanyusuf/carrent-lib/pkg/errormsg"
	"github.com/achwanyusuf/carrent-lib/pkg/logger"
	"github.com/achwanyusuf/carrent-ordersvc/src/model"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/svcerr"
	"github.com/achwanyusuf/carrent-ordersvc/src/usecase/car"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
)

type CarDep struct {
	log  logger.Logger
	car  car.CarInterface
	conf Conf
}

type Conf struct{}

type CarInterface interface {
	Create(ctx *gin.Context)
	UpdateByID(ctx *gin.Context)
	Read(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	DeleteByID(ctx *gin.Context)
}

func New(conf Conf, log *logger.Logger, c car.CarInterface) CarInterface {
	return &CarDep{
		conf: conf,
		log:  *log,
		car:  c,
	}
}

// Create Car godoc
// @Summary Create Car
// @Description Create car data
// @Tags car
// @Accept json
// @Produce json
// @Security OAuth2Password
// @Param data body model.CreateCar true "Car Data"
// @Success 200 {object} model.SingleCarResponse
// @Success 400 {object} model.SingleCarResponse
// @Success 500 {object} model.SingleCarResponse
// @Router /car [post]
func (c *CarDep) Create(ctx *gin.Context) {
	var (
		carInput model.CreateCar
		result   model.Car
		response model.SingleCarResponse
	)

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		statusCode := response.Transform(ctx, c.log, http.StatusCreated, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error read body"))
		ctx.JSON(statusCode, response)
		return
	}

	if err = json.Unmarshal(body, &carInput); err != nil {
		statusCode := response.Transform(ctx, c.log, http.StatusCreated, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error unmarshal body"))
		ctx.JSON(statusCode, response)
		return
	}

	carInput.CreatedBy = ctx.Value("id").(int64)
	if err = carInput.Validate(); err != nil {
		statusCode := response.Transform(ctx, c.log, http.StatusCreated, err)
		ctx.JSON(statusCode, response)
		return
	}

	result, err = c.car.Create(ctx, carInput)
	if err != nil {
		statusCode := response.Transform(ctx, c.log, http.StatusCreated, err)
		ctx.JSON(statusCode, response)
		return
	}

	response.Data = result

	statusCode := response.Transform(ctx, c.log, http.StatusCreated, nil)
	ctx.JSON(statusCode, response)
}

// Update Car Data godoc
// @Summary Update car data
// @Description Update car data
// @Tags car
// @Accept json
// @Produce json
// @Security OAuth2Password
// @Param id path string true "update by id"
// @Param data body model.UpdateCar true "Car Data"
// @Success 200 {object} model.SingleCarResponse
// @Success 400 {object} model.SingleCarResponse
// @Success 500 {object} model.SingleCarResponse
// @Router /car/{id} [put]
func (c *CarDep) UpdateByID(ctx *gin.Context) {
	var (
		updateData model.UpdateCar
		response   model.SingleCarResponse
	)

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		statusCode := response.Transform(ctx, c.log, http.StatusOK, errormsg.WrapErr(errormsg.Error400, err, "error get id"))
		ctx.JSON(statusCode, response)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		statusCode := response.Transform(ctx, c.log, http.StatusOK, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error read body"))
		ctx.JSON(statusCode, response)
		return
	}

	if err = json.Unmarshal(body, &updateData); err != nil {
		statusCode := response.Transform(ctx, c.log, http.StatusOK, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error unmarshal body"))
		ctx.JSON(statusCode, response)
		return
	}
	updateData.UpdatedBy = ctx.Value("id").(int64)
	scope := ctx.Value("scope").(string)
	if scope != model.SuperAdminScope {
		id = ctx.Value("id").(int64)
	}
	if err = updateData.Validate(); err != nil {
		statusCode := response.Transform(ctx, c.log, http.StatusOK, err)
		ctx.JSON(statusCode, response)
		return
	}
	result, err := c.car.UpdateByID(ctx, id, updateData)
	if err != nil {
		statusCode := response.Transform(ctx, c.log, http.StatusOK, err)
		ctx.JSON(statusCode, response)
		return
	}

	response.Data = result

	statusCode := response.Transform(ctx, c.log, http.StatusOK, nil)
	ctx.JSON(statusCode, response)
}

// Get Cars Data godoc
// @Summary Get cars data
// @Description Get cars data
// @Tags car
// @Accept json
// @Produce json
// @Security OAuth2Password
// @Param id query number false "search by id"
// @Param car_name query string false "search by car name"
// @Param day_rate query number false "search by day rate"
// @Param day_rate_gt query number false "search by day rate greater than"
// @Param day_rate_gte query number false "search by day rate greater than equal"
// @Param day_rate_lt query number false "search by day rate less than"
// @Param day_rate_lte query number false "search by day rate less than equal"
// @Param month_rate query number false "search by month rate"
// @Param month_rate_gt query number false "search by month rate greater than"
// @Param month_rate_gte query number false "search by month rate greater than equal"
// @Param month_rate_lt query number false "search by month rate less than"
// @Param month_rate_lte query number false "search by month rate less than equal"
// @Param image query string false "search by image"
// @Param page query int false " "
// @Param limit query int false " "
// @Param Cache-Control header string false "Request Cache Control" Enums(must-revalidate, none)
// @Success 200 {object} model.CarsResponse
// @Success 400 {object} model.CarsResponse
// @Success 500 {object} model.CarsResponse
// @Router /car [get]
func (c *CarDep) Read(ctx *gin.Context) {
	var (
		param    model.GetCarsByParam
		response model.CarsResponse
	)
	cacheControl := ctx.GetHeader("Cache-Control")
	var decoder = schema.NewDecoder()
	err := decoder.Decode(&param, ctx.Request.URL.Query())
	if err != nil {
		statusCode := response.Transform(ctx, c.log, http.StatusOK, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error unmarshal query"))
		ctx.JSON(statusCode, response)
		return
	}
	cars, pagination, err := c.car.GetByParam(ctx, cacheControl, param)
	if err != nil {
		statusCode := response.Transform(ctx, c.log, http.StatusOK, err)
		ctx.JSON(statusCode, response)
		return
	}

	response.Data = cars
	response.Pagination = pagination

	statusCode := response.Transform(ctx, c.log, http.StatusOK, nil)
	ctx.JSON(statusCode, response)
}

// Get Cars Data godoc
// @Summary Get cars data
// @Description Get cars data
// @Tags car
// @Accept json
// @Produce json
// @Security OAuth2Password
// @Param id path string true "get by id"
// @Param Cache-Control header string false "Request Cache Control" Enums(must-revalidate, none)
// @Success 200 {object} model.SingleCarResponse
// @Success 400 {object} model.SingleCarResponse
// @Success 500 {object} model.SingleCarResponse
// @Router /car/{id} [get]
func (c *CarDep) GetByID(ctx *gin.Context) {
	var response model.SingleCarResponse
	cacheControl := ctx.GetHeader("Cache-Control")
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		statusCode := response.Transform(ctx, c.log, http.StatusOK, errormsg.WrapErr(errormsg.Error400, err, "error get id"))
		ctx.JSON(statusCode, response)
		return
	}
	result, err := c.car.GetByID(ctx, cacheControl, id)
	if err != nil {
		statusCode := response.Transform(ctx, c.log, http.StatusOK, err)
		ctx.JSON(statusCode, response)
		return
	}

	response.Data = result

	statusCode := response.Transform(ctx, c.log, http.StatusOK, nil)
	ctx.JSON(statusCode, response)
}

// Delete Car Data godoc
// @Summary Delete car data
// @Description Delete car data
// @Tags car
// @Accept json
// @Produce json
// @Security OAuth2Password
// @Param id path string true "delete by id"
// @Success 200 {object} model.EmptyResponse
// @Success 400 {object} model.EmptyResponse
// @Success 500 {object} model.EmptyResponse
// @Router /car/{id} [delete]
func (c *CarDep) DeleteByID(ctx *gin.Context) {
	var (
		response model.EmptyResponse
	)
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		statusCode := response.Transform(ctx, c.log, http.StatusOK, errormsg.WrapErr(errormsg.Error400, err, "error get id"))
		ctx.JSON(statusCode, response)
		return
	}
	scope := ctx.Value("scope").(string)
	if scope != model.SuperAdminScope {
		id = ctx.Value("id").(int64)
	}
	err = c.car.DeleteByID(ctx, ctx.Value("id").(int64), id)
	if err != nil {
		statusCode := response.Transform(ctx, c.log, http.StatusOK, err)
		ctx.JSON(statusCode, response)
		return
	}

	statusCode := response.Transform(ctx, c.log, http.StatusOK, nil)
	ctx.JSON(statusCode, response)
}
