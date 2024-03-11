package model

import (
	"strings"
	"time"

	"github.com/achwanyusuf/carrent-lib/pkg/errormsg"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/grpcmodel"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/psqlmodel"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/svcerr"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var (
	GetSingleByParamCarKey string = "gspCar:%s"
	GetByParamCarKey       string = "gpCar:%s"
	GetByParamCarPgKey     string = "gppgCar:%s"
)

type GetCarByParam struct {
	ID           null.Int64   `schema:"id" json:"id"`
	CarName      null.String  `json:"car_name"`
	DayRate      null.Float64 `json:"day_rate"`
	DayRateGT    null.Float64 `json:"day_rate_gt"`
	DayRateGTE   null.Float64 `json:"day_rate_gte"`
	DayRateLT    null.Float64 `json:"day_rate_lt"`
	DayRateLTE   null.Float64 `json:"day_rate_lte"`
	MonthRate    null.Float64 `json:"month_rate"`
	MonthRateGT  null.Float64 `json:"month_rate_gt"`
	MonthRateGTE null.Float64 `json:"month_rate_gte"`
	MonthRateLT  null.Float64 `json:"month_rate_lt"`
	MonthRateLTE null.Float64 `json:"month_rate_lte"`
	Image        null.String  `json:"image"`
}

func (g *GetCarByParam) GetQuery() []qm.QueryMod {
	var res []qm.QueryMod
	if g.ID.Valid {
		res = append(res, qm.Where("id=?", g.ID.Int64))
	}

	if g.CarName.Valid {
		res = append(res, qm.Where("car_name=?", g.CarName.String))
	}

	if g.DayRate.Valid {
		res = append(res, qm.Where("day_rate=?", g.DayRate.Float64))
	}

	if g.DayRateGT.Valid {
		res = append(res, qm.Where("day_rate >?", g.DayRate.Float64))
	}

	if g.DayRateGTE.Valid {
		res = append(res, qm.Where("day_rate >=?", g.DayRate.Float64))
	}

	if g.DayRateLT.Valid {
		res = append(res, qm.Where("day_rate <?", g.DayRate.Float64))
	}

	if g.DayRateLTE.Valid {
		res = append(res, qm.Where("day_rate <=?", g.DayRate.Float64))
	}

	if g.MonthRate.Valid {
		res = append(res, qm.Where("month_rate=?", g.MonthRate.Float64))
	}

	if g.MonthRateGT.Valid {
		res = append(res, qm.Where("month_rate >?", g.MonthRateGT.Float64))
	}

	if g.MonthRateGTE.Valid {
		res = append(res, qm.Where("month_rate >=?", g.MonthRateGTE.Float64))
	}

	if g.MonthRateLT.Valid {
		res = append(res, qm.Where("month_rate <?", g.MonthRateLT.Float64))
	}

	if g.MonthRateLTE.Valid {
		res = append(res, qm.Where("month_rate <=?", g.MonthRateLTE.Float64))
	}

	if g.Image.Valid {
		res = append(res, qm.Where("image=?", g.Image.String))
	}

	return res
}

type GetCarsByParam struct {
	GetCarByParam
	OrderBy null.String `schema:"order_by" json:"order_by"`
	Limit   int64       `schema:"limit" json:"limit"`
	Page    int64       `schema:"page" json:"page"`
}

func (g *GetCarsByParam) FillGrpcClient() *grpcmodel.GetCarByParamRequest {
	return &grpcmodel.GetCarByParamRequest{
		Id:           g.ID.Ptr(),
		CarName:      g.CarName.Ptr(),
		DayRate:      g.DayRate.Ptr(),
		DayRateGt:    g.DayRateGT.Ptr(),
		DayRateGte:   g.DayRateGTE.Ptr(),
		DayRateLt:    g.DayRateLT.Ptr(),
		DayRateLte:   g.DayRateLTE.Ptr(),
		MonthRate:    g.MonthRate.Ptr(),
		MonthRateGt:  g.MonthRateGT.Ptr(),
		MonthRateGte: g.MonthRateGTE.Ptr(),
		MonthRateLt:  g.MonthRateLT.Ptr(),
		MonthRateLte: g.MonthRateLTE.Ptr(),
		Image:        g.Image.Ptr(),
		OrderBy:      g.OrderBy.Ptr(),
		Limit:        g.Limit,
		Page:         g.Page,
	}
}

func (g *GetCarsByParam) GetQuery() []qm.QueryMod {
	var res []qm.QueryMod
	if g.ID.Valid {
		res = append(res, qm.Where("id=?", g.ID.Int64))
	}

	if g.CarName.Valid {
		res = append(res, qm.Where("car_name=?", g.CarName.String))
	}

	if g.DayRate.Valid {
		res = append(res, qm.Where("day_rate=?", g.DayRate.Float64))
	}

	if g.DayRateGT.Valid {
		res = append(res, qm.Where("day_rate >?", g.DayRate.Float64))
	}

	if g.DayRateGTE.Valid {
		res = append(res, qm.Where("day_rate >=?", g.DayRate.Float64))
	}

	if g.DayRateLT.Valid {
		res = append(res, qm.Where("day_rate <?", g.DayRate.Float64))
	}

	if g.DayRateLTE.Valid {
		res = append(res, qm.Where("day_rate <=?", g.DayRate.Float64))
	}

	if g.MonthRate.Valid {
		res = append(res, qm.Where("month_rate=?", g.MonthRate.Float64))
	}

	if g.MonthRateGT.Valid {
		res = append(res, qm.Where("month_rate >?", g.MonthRateGT.Float64))
	}

	if g.MonthRateGTE.Valid {
		res = append(res, qm.Where("month_rate >=?", g.MonthRateGTE.Float64))
	}

	if g.MonthRateLT.Valid {
		res = append(res, qm.Where("month_rate <?", g.MonthRateLT.Float64))
	}

	if g.MonthRateLTE.Valid {
		res = append(res, qm.Where("month_rate <=?", g.MonthRateLTE.Float64))
	}

	if g.Image.Valid {
		res = append(res, qm.Where("image=?", g.Image.String))
	}

	if g.OrderBy.Valid {
		order := strings.Split(g.OrderBy.String, ",")
		for _, o := range order {
			res = append(res, qm.OrderBy(o))
		}
	}

	return res
}

type CreateCar struct {
	CarName   string  `json:"car_name" validate:"required,alphanumspace,min=8,max=50"`
	DayRate   float64 `json:"day_rate" validate:"required,min=10000,max=1000000"`
	MonthRate float64 `json:"month_rate" validate:"required,min=250000,max=30000000"`
	Image     string  `json:"image" validate:"required,uri,min=5,max=256"`
	CreatedBy int64   `json:"-"`
}

func (v *CreateCar) Validate() error {
	if v.CarName == "" {
		return errormsg.WrapErr(svcerr.OrderSVCCodeInvalidCarName, nil, "invalid car name")
	}

	if v.DayRate == 0 {
		return errormsg.WrapErr(svcerr.OrderSVCCodeInvalidDayRate, nil, "invalid day rate")
	}

	if v.MonthRate == 0 {
		return errormsg.WrapErr(svcerr.OrderSVCCodeInvalidMonthRate, nil, "invalid month rate")
	}

	if v.Image == "" {
		return errormsg.WrapErr(svcerr.OrderSVCCodeInvalidImage, nil, "invalid image")
	}

	return nil
}

type UpdateCar struct {
	CarName   null.String  `json:"car_name" swaggertype:"string" validate:"required,alphanumspace,min=8,max=50"`
	DayRate   null.Float64 `json:"day_rate" swaggertype:"number" validate:"required,min=10000,max=1000000"`
	MonthRate null.Float64 `json:"month_rate" swaggertype:"number" validate:"required,min=250000,max=30000000"`
	Image     null.String  `json:"image" swaggertype:"string" validate:"required,uri,min=5,max=256"`
	UpdatedBy int64        `json:"-"`
}

func (v *UpdateCar) FillGrpcClient() *grpcmodel.UpdateCarRequest {
	return &grpcmodel.UpdateCarRequest{
		CarName:   v.CarName.Ptr(),
		DayRate:   v.DayRate.Ptr(),
		MonthRate: v.MonthRate.Ptr(),
		Image:     v.Image.Ptr(),
	}
}

func (v *UpdateCar) FillEntity(Car *psqlmodel.Car) {
	if v.CarName.Valid {
		Car.CarName = v.CarName.String
	}

	if v.DayRate.Valid {
		Car.DayRate = v.DayRate.Float64
	}

	if v.MonthRate.Valid {
		Car.MonthRate = v.MonthRate.Float64
	}

	if v.Image.Valid {
		Car.Image = v.Image.String
	}
}

type Car struct {
	ID        int64   `json:"id"`
	CarName   string  `json:"car_name"`
	DayRate   float64 `json:"day_rate"`
	MonthRate float64 `json:"month_rate"`
	Image     string  `json:"image"`
	BaseInformation
}

func TransformPSQLSingleCar(car *psqlmodel.Car) Car {
	deletedBy := null.Int64{}
	if car.DeletedBy.Valid {
		deletedBy = null.Int64From(int64(car.DeletedBy.Int))
	}

	deletedAt := null.Time{}
	if car.DeletedAt.Valid {
		deletedAt = null.TimeFrom(deletedAt.Time)
	}
	creationInfo := BaseInformation{
		CreatedBy: int64(car.CreatedBy),
		CreatedAt: car.CreatedAt,
		UpdatedBy: int64(car.UpdatedBy),
		UpdatedAt: car.UpdatedAt,
		DeletedBy: deletedBy,
		DeletedAt: deletedAt,
	}

	return Car{
		ID:              int64(car.ID),
		CarName:         car.CarName,
		DayRate:         car.DayRate,
		MonthRate:       car.MonthRate,
		Image:           car.Image,
		BaseInformation: creationInfo,
	}
}

func TransformPSQLCar(car *psqlmodel.CarSlice) []Car {
	var res []Car
	for _, v := range *car {
		deletedBy := null.Int64{}
		if v.DeletedBy.Valid {
			deletedBy = null.Int64From(int64(v.DeletedBy.Int))
		}

		deletedAt := null.Time{}
		if v.DeletedAt.Valid {
			deletedAt = null.TimeFrom(deletedAt.Time)
		}
		creationInfo := BaseInformation{
			CreatedBy: int64(v.CreatedBy),
			CreatedAt: v.CreatedAt,
			UpdatedBy: int64(v.UpdatedBy),
			UpdatedAt: v.UpdatedAt,
			DeletedBy: deletedBy,
			DeletedAt: deletedAt,
		}

		res = append(res, Car{
			ID:              int64(v.ID),
			CarName:         v.CarName,
			DayRate:         v.DayRate,
			MonthRate:       v.MonthRate,
			Image:           v.Image,
			BaseInformation: creationInfo,
		})
	}

	return res
}

func TransformSingleCarReply(v *psqlmodel.Car) *grpcmodel.SingleCarReply {
	singleCar := &grpcmodel.SingleCarReply{
		Id:        int64(v.ID),
		CarName:   v.CarName,
		DayRate:   v.DayRate,
		MonthRate: v.MonthRate,
		Image:     v.Image,
		CreatedBy: int64(v.CreatedBy),
		CreatedAt: v.CreatedAt.Format(time.RFC3339),
		UpdatedBy: int64(v.UpdatedBy),
		UpdatedAt: v.UpdatedAt.Format(time.RFC3339),
	}

	if v.DeletedBy.Valid {
		deletedBy := int64(v.DeletedBy.Int)
		singleCar.DeletedBy = &deletedBy
	}

	if v.DeletedAt.Valid {
		deletedAt := v.DeletedAt.Time.Format(time.RFC3339)
		singleCar.DeletedAt = &deletedAt
	}

	return singleCar
}

func TransformGetCarByParamRequestToCarParam(v *grpcmodel.GetCarByParamRequest) GetCarsByParam {
	return GetCarsByParam{
		GetCarByParam: GetCarByParam{
			ID:           null.Int64FromPtr(v.Id),
			CarName:      null.StringFromPtr(v.CarName),
			DayRate:      null.Float64FromPtr(v.DayRate),
			DayRateGT:    null.Float64FromPtr(v.DayRateGt),
			DayRateGTE:   null.Float64FromPtr(v.DayRateGte),
			DayRateLT:    null.Float64FromPtr(v.DayRateLt),
			DayRateLTE:   null.Float64FromPtr(v.DayRateLte),
			MonthRate:    null.Float64FromPtr(v.MonthRate),
			MonthRateGT:  null.Float64FromPtr(v.MonthRateGt),
			MonthRateGTE: null.Float64FromPtr(v.MonthRateGte),
			MonthRateLT:  null.Float64FromPtr(v.MonthRateLt),
			MonthRateLTE: null.Float64FromPtr(v.MonthRateLte),
			Image:        null.StringFromPtr(v.Image),
		},
		OrderBy: null.StringFromPtr(v.OrderBy),
		Limit:   v.Limit,
		Page:    v.Page,
	}
}

func TransformCarToGetCarByParamReply(v *psqlmodel.CarSlice, p Pagination) *grpcmodel.GetCarByParamReply {
	var (
		data      []*grpcmodel.SingleCarReply
		deletedBy *int64
		deletedAt *string
	)
	for _, val := range *v {
		deletedBy = nil
		if val.DeletedBy.Valid {
			iParse := int64(val.DeletedBy.Int)
			deletedBy = &iParse
		}

		if val.DeletedAt.Valid {
			tParse := val.DeletedAt.Time.Format(time.RFC3339)
			deletedAt = &tParse
		}
		data = append(data, &grpcmodel.SingleCarReply{
			Id:        int64(val.ID),
			CarName:   val.CarName,
			DayRate:   val.DayRate,
			MonthRate: val.MonthRate,
			Image:     val.Image,
			CreatedBy: int64(val.CreatedBy),
			CreatedAt: val.CreatedAt.Format(time.RFC3339),
			UpdatedBy: int64(val.UpdatedBy),
			UpdatedAt: val.UpdatedAt.Format(time.RFC3339),
			DeletedBy: deletedBy,
			DeletedAt: deletedAt,
		})
	}

	return &grpcmodel.GetCarByParamReply{
		Data: data,
		Pagination: &grpcmodel.Pagination{
			CurrentPage:    p.CurrentPage,
			CurrentElement: p.CurrentElements,
			TotalPages:     p.TotalPages,
			TotalElements:  p.TotalElements,
			SortBy:         p.SortBy,
		},
	}
}
