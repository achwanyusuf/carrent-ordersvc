package model

import (
	"context"
	"strings"
	"time"

	"github.com/achwanyusuf/carrent-lib/pkg/errormsg"
	"github.com/achwanyusuf/carrent-lib/pkg/logger"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/grpcmodel"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/psqlmodel"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/svcerr"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var (
	GetSingleByParamOrderKey string = "gspOrder:%s"
	GetByParamOrderKey       string = "gpOrder:%s"
	GetByParamOrderPgKey     string = "gppgOrder:%s"
)

type GetOrderByParam struct {
	ID              null.Int64   `schema:"id" json:"id"`
	CarID           null.Int64   `json:"car_id"`
	OrderDate       null.Time    `json:"order_date"`
	PickupDate      null.Time    `json:"pickup_date"`
	DropoffDate     null.Time    `json:"dropoff_date"`
	PickupLocation  null.String  `json:"pickup_location"`
	PickupLat       null.Float64 `json:"pickup_lat"`
	PickupLong      null.Float64 `json:"pickup_long"`
	DropoffLocation null.String  `json:"dropoff_location"`
	DropoffLat      null.Float64 `json:"dropoff_lat"`
	DropoffLong     null.Float64 `json:"dropoff_long"`
}

func (g *GetOrderByParam) GetQuery() []qm.QueryMod {
	var res []qm.QueryMod
	if g.ID.Valid {
		res = append(res, qm.Where("id=?", g.ID.Int64))
	}

	if g.CarID.Valid {
		res = append(res, qm.Where("car_id=?", g.CarID.Int64))
	}

	if g.OrderDate.Valid {
		res = append(res, qm.Where("order_date=?", g.OrderDate.Time))
	}

	if g.PickupDate.Valid {
		res = append(res, qm.Where("pickup_date=?", g.PickupDate.Time))
	}

	if g.DropoffDate.Valid {
		res = append(res, qm.Where("dropoff_date=?", g.DropoffDate.Time))
	}

	if g.PickupLocation.Valid {
		res = append(res, qm.Where("pickup_location=?", g.PickupLocation.String))
	}

	if g.PickupLat.Valid {
		res = append(res, qm.Where("pickup_lat=?", g.PickupLat.Float64))
	}

	if g.PickupLong.Valid {
		res = append(res, qm.Where("pickup_long=?", g.PickupLong.Float64))
	}

	if g.DropoffLocation.Valid {
		res = append(res, qm.Where("dropoff_location=?", g.DropoffLocation.String))
	}

	if g.DropoffLat.Valid {
		res = append(res, qm.Where("dropoff_lat=?", g.DropoffLat.Float64))
	}

	if g.PickupLong.Valid {
		res = append(res, qm.Where("dropoff_long=?", g.PickupLong.Float64))
	}

	return res
}

type GetOrdersByParam struct {
	GetOrderByParam
	OrderBy null.String `schema:"order_by" json:"order_by"`
	Limit   int64       `schema:"limit" json:"limit"`
	Page    int64       `schema:"page" json:"page"`
}

func (g *GetOrdersByParam) FillGrpcClient() *grpcmodel.GetOrderByParamRequest {
	res := &grpcmodel.GetOrderByParamRequest{
		Id:              g.ID.Ptr(),
		CarId:           g.CarID.Ptr(),
		PickupLocation:  g.PickupLocation.Ptr(),
		PickupLat:       g.PickupLat.Ptr(),
		PickupLong:      g.PickupLong.Ptr(),
		DropoffLocation: g.DropoffLocation.Ptr(),
		DropoffLat:      g.DropoffLat.Ptr(),
		DropoffLong:     g.DropoffLong.Ptr(),
		OrderBy:         g.OrderBy.Ptr(),
		Limit:           g.Limit,
		Page:            g.Page,
	}

	if g.OrderDate.Valid {
		orderDate := g.OrderDate.Time.Format(time.RFC3339)
		res.OrderDate = &orderDate
	}

	if g.PickupDate.Valid {
		pickupDate := g.PickupDate.Time.Format(time.RFC3339)
		res.PickupDate = &pickupDate
	}

	if g.DropoffDate.Valid {
		dropoffDate := g.DropoffDate.Time.Format(time.RFC3339)
		res.DropoffDate = &dropoffDate
	}

	return res
}

func (g *GetOrdersByParam) GetQuery() []qm.QueryMod {
	var res []qm.QueryMod
	if g.ID.Valid {
		res = append(res, qm.Where("id=?", g.ID.Int64))
	}

	if g.CarID.Valid {
		res = append(res, qm.Where("car_id=?", g.CarID.Int64))
	}

	if g.OrderDate.Valid {
		res = append(res, qm.Where("order_date=?", g.OrderDate.Time))
	}

	if g.PickupDate.Valid {
		res = append(res, qm.Where("pickup_date=?", g.PickupDate.Time))
	}

	if g.DropoffDate.Valid {
		res = append(res, qm.Where("dropoff_date=?", g.DropoffDate.Time))
	}

	if g.PickupLocation.Valid {
		res = append(res, qm.Where("pickup_location=?", g.PickupLocation.String))
	}

	if g.PickupLat.Valid {
		res = append(res, qm.Where("pickup_lat=?", g.PickupLat.Float64))
	}

	if g.PickupLong.Valid {
		res = append(res, qm.Where("pickup_long=?", g.PickupLong.Float64))
	}

	if g.DropoffLocation.Valid {
		res = append(res, qm.Where("dropoff_location=?", g.DropoffLocation.String))
	}

	if g.DropoffLat.Valid {
		res = append(res, qm.Where("dropoff_lat=?", g.DropoffLat.Float64))
	}

	if g.PickupLong.Valid {
		res = append(res, qm.Where("dropoff_long=?", g.PickupLong.Float64))
	}

	if g.OrderBy.Valid {
		order := strings.Split(g.OrderBy.String, ",")
		for _, o := range order {
			res = append(res, qm.OrderBy(o))
		}
	}

	return res
}

type CreateOrder struct {
	CarID           int64     `json:"car_id"`
	OrderDate       time.Time `json:"order_date"`
	PickupDate      time.Time `json:"pickup_date"`
	DropoffDate     time.Time `json:"dropoff_date"`
	PickupLocation  string    `json:"pickup_location"`
	PickupLat       float64   `json:"pickup_lat"`
	PickupLong      float64   `json:"pickup_long"`
	DropoffLocation string    `json:"dropoff_location"`
	DropoffLat      float64   `json:"dropoff_lat"`
	DropoffLong     float64   `json:"dropoff_long"`
	CreatedBy       int64     `json:"-"`
}

func (v *CreateOrder) Validate() error {
	if v.CarID == 0 {
		return errormsg.WrapErr(svcerr.OrderSVCCodeInvalidCarID, nil, "invalid car id")
	}

	if v.OrderDate.IsZero() {
		return errormsg.WrapErr(svcerr.OrderSVCCodeInvalidOrderDate, nil, "invalid order date")
	}

	if v.PickupDate.IsZero() {
		return errormsg.WrapErr(svcerr.OrderSVCCodeInvalidPickupDate, nil, "invalid pickup date")
	}

	if v.DropoffDate.IsZero() {
		return errormsg.WrapErr(svcerr.OrderSVCCodeInvalidDropoffDate, nil, "invalid dropoff date")
	}

	if v.PickupLocation == "" {
		return errormsg.WrapErr(svcerr.OrderSVCCodeInvalidPickupLocation, nil, "invalid pickup location")
	}

	if v.PickupLat == 0 {
		return errormsg.WrapErr(svcerr.OrderSVCCodeInvalidPickupLat, nil, "invalid pickup lat")
	}

	if v.PickupLong == 0 {
		return errormsg.WrapErr(svcerr.OrderSVCCodeInvalidPickupLong, nil, "invalid pickup long")
	}

	if v.DropoffLocation == "" {
		return errormsg.WrapErr(svcerr.OrderSVCCodeInvalidDropoffLocation, nil, "invalid dropoff location")
	}

	if v.DropoffLat == 0 {
		return errormsg.WrapErr(svcerr.OrderSVCCodeInvalidDropoffLat, nil, "invalid dropoff lat")
	}

	if v.DropoffLong == 0 {
		return errormsg.WrapErr(svcerr.OrderSVCCodeInvalidDropoffLong, nil, "invalid dropoff long")
	}

	return nil
}

type UpdateOrder struct {
	CarID           null.Int64   `json:"car_id"`
	OrderDate       null.Time    `json:"order_date"`
	PickupDate      null.Time    `json:"pickup_date"`
	DropoffDate     null.Time    `json:"dropoff_date"`
	PickupLocation  null.String  `json:"pickup_location"`
	PickupLat       null.Float64 `json:"pickup_lat"`
	PickupLong      null.Float64 `json:"pickup_long"`
	DropoffLocation null.String  `json:"dropoff_location"`
	DropoffLat      null.Float64 `json:"dropoff_lat"`
	DropoffLong     null.Float64 `json:"dropoff_long"`
	UpdatedBy       int64        `json:"-"`
}

func (v *UpdateOrder) FillGrpcClient() *grpcmodel.UpdateOrderRequest {
	res := &grpcmodel.UpdateOrderRequest{
		CarId:           v.CarID.Ptr(),
		PickupLocation:  v.PickupLocation.Ptr(),
		PickupLat:       v.PickupLat.Ptr(),
		PickupLong:      v.PickupLong.Ptr(),
		DropoffLocation: v.DropoffLocation.Ptr(),
		DropoffLat:      v.DropoffLat.Ptr(),
		DropoffLong:     v.DropoffLong.Ptr(),
	}

	if v.OrderDate.Valid {
		orderDate := v.OrderDate.Time.Format(time.RFC3339)
		res.OrderDate = &orderDate
	}

	if v.PickupDate.Valid {
		pickupDate := v.PickupDate.Time.Format(time.RFC3339)
		res.PickupDate = &pickupDate
	}

	if v.DropoffDate.Valid {
		dropoffDate := v.DropoffDate.Time.Format(time.RFC3339)
		res.DropoffDate = &dropoffDate
	}
	return res
}

func (v *UpdateOrder) FillEntity(Order *psqlmodel.Order) {
	if v.CarID.Valid {
		Order.CarID = int(v.CarID.Int64)
	}

	if v.OrderDate.Valid {
		Order.OrderDate = v.OrderDate.Time
	}

	if v.PickupDate.Valid {
		Order.PickupDate = v.PickupDate.Time
	}

	if v.DropoffDate.Valid {
		Order.DropoffDate = v.DropoffDate.Time
	}

	if v.PickupLocation.Valid {
		Order.PickupLocation = v.PickupLocation.String
	}

	if v.PickupLat.Valid {
		Order.PickupLat = v.PickupLat.Float64
	}

	if v.PickupLong.Valid {
		Order.PickupLong = v.PickupLong.Float64
	}

	if v.DropoffLocation.Valid {
		Order.DropoffLocation = v.DropoffLocation.String
	}

	if v.DropoffLat.Valid {
		Order.DropoffLat = v.DropoffLat.Float64
	}

	if v.DropoffLong.Valid {
		Order.DropoffLong = v.DropoffLong.Float64
	}
}

type Order struct {
	ID              int64     `json:"id"`
	CarID           int64     `json:"car_id"`
	OrderDate       time.Time `json:"order_date"`
	PickupDate      time.Time `json:"pickup_date"`
	DropoffDate     time.Time `json:"dropoff_date"`
	PickupLocation  string    `json:"pickup_location"`
	PickupLat       float64   `json:"pickup_lat"`
	PickupLong      float64   `json:"pickup_long"`
	DropoffLocation string    `json:"dropoff_location"`
	DropoffLat      float64   `json:"dropoff_lat"`
	DropoffLong     float64   `json:"dropoff_long"`
	BaseInformation
}

func TransformPSQLSingleOrder(order *psqlmodel.Order) Order {
	deletedBy := null.Int64{}
	if order.DeletedBy.Valid {
		deletedBy = null.Int64From(int64(order.DeletedBy.Int))
	}

	deletedAt := null.Time{}
	if order.DeletedAt.Valid {
		deletedAt = null.TimeFrom(deletedAt.Time)
	}

	creationInfo := BaseInformation{
		CreatedBy: int64(order.CreatedBy),
		CreatedAt: order.CreatedAt,
		UpdatedBy: int64(order.UpdatedBy),
		UpdatedAt: order.UpdatedAt,
		DeletedBy: deletedBy,
		DeletedAt: deletedAt,
	}

	return Order{
		ID:              int64(order.ID),
		CarID:           int64(order.CarID),
		OrderDate:       order.OrderDate,
		PickupDate:      order.PickupDate,
		DropoffDate:     order.DropoffDate,
		PickupLocation:  order.PickupLocation,
		PickupLat:       order.PickupLat,
		PickupLong:      order.PickupLong,
		DropoffLocation: order.DropoffLocation,
		DropoffLat:      order.DropoffLat,
		DropoffLong:     order.DropoffLong,
		BaseInformation: creationInfo,
	}
}

func TransformPSQLOrder(order *psqlmodel.OrderSlice) []Order {
	var res []Order
	for _, v := range *order {
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

		res = append(res, Order{
			ID:              int64(v.ID),
			CarID:           int64(v.CarID),
			OrderDate:       v.OrderDate,
			PickupDate:      v.PickupDate,
			DropoffDate:     v.DropoffDate,
			PickupLocation:  v.PickupLocation,
			PickupLat:       v.PickupLat,
			PickupLong:      v.PickupLong,
			DropoffLocation: v.DropoffLocation,
			DropoffLat:      v.DropoffLat,
			DropoffLong:     v.DropoffLong,
			BaseInformation: creationInfo,
		})
	}

	return res
}

func TransformSingleOrderReply(v *psqlmodel.Order) *grpcmodel.SingleOrderReply {
	singleOrder := &grpcmodel.SingleOrderReply{
		Id:              int64(v.ID),
		CarId:           int64(v.CarID),
		OrderDate:       v.OrderDate.Format(time.RFC3339),
		PickupDate:      v.PickupDate.Format(time.RFC3339),
		DropoffDate:     v.DropoffDate.Format(time.RFC3339),
		PickupLocation:  v.PickupLocation,
		PickupLat:       v.PickupLat,
		PickupLong:      v.PickupLong,
		DropoffLocation: v.DropoffLocation,
		DropoffLat:      v.DropoffLat,
		DropoffLong:     v.DropoffLong,
		CreatedBy:       int64(v.CreatedBy),
		CreatedAt:       v.CreatedAt.Format(time.RFC3339),
		UpdatedBy:       int64(v.UpdatedBy),
		UpdatedAt:       v.UpdatedAt.Format(time.RFC3339),
	}

	if v.DeletedBy.Valid {
		deletedBy := int64(v.DeletedBy.Int)
		singleOrder.DeletedBy = &deletedBy
	}

	if v.DeletedAt.Valid {
		deletedAt := v.DeletedAt.Time.Format(time.RFC3339)
		singleOrder.DeletedAt = &deletedAt
	}

	return singleOrder
}

func TransformGetOrderByParamRequestToOrderParam(ctx context.Context, v *grpcmodel.GetOrderByParamRequest, log logger.Logger) GetOrdersByParam {
	var (
		pickupDate  null.Time
		dropoffDate null.Time
		orderDate   null.Time
	)
	if v.PickupDate != nil {
		pDate, err := time.Parse(time.RFC3339, *v.PickupDate)
		if err != nil {
			log.Error(ctx, err)
		}
		pickupDate = null.TimeFrom(pDate)
	}
	if v.DropoffDate != nil {
		dDate, err := time.Parse(time.RFC3339, *v.DropoffDate)
		if err != nil {
			log.Error(ctx, err)
		}
		dropoffDate = null.TimeFrom(dDate)
	}
	if v.OrderDate != nil {
		oDate, err := time.Parse(time.RFC3339, *v.OrderDate)
		if err != nil {
			log.Error(ctx, err)
		}
		orderDate = null.TimeFrom(oDate)
	}
	return GetOrdersByParam{
		GetOrderByParam: GetOrderByParam{
			ID:              null.Int64FromPtr(v.Id),
			OrderDate:       orderDate,
			PickupDate:      pickupDate,
			DropoffDate:     dropoffDate,
			CarID:           null.Int64FromPtr(v.CarId),
			PickupLocation:  null.StringFromPtr(v.PickupLocation),
			PickupLat:       null.Float64FromPtr(v.PickupLat),
			PickupLong:      null.Float64FromPtr(v.PickupLong),
			DropoffLocation: null.StringFromPtr(v.DropoffLocation),
			DropoffLat:      null.Float64FromPtr(v.DropoffLat),
			DropoffLong:     null.Float64FromPtr(v.DropoffLong),
		},
		OrderBy: null.StringFromPtr(v.OrderBy),
		Limit:   v.Limit,
		Page:    v.Page,
	}
}

func TransformOrderToGetOrderByParamReply(v *psqlmodel.OrderSlice, p Pagination) *grpcmodel.GetOrderByParamReply {
	var (
		data      []*grpcmodel.SingleOrderReply
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
		data = append(data, &grpcmodel.SingleOrderReply{
			Id:              int64(val.ID),
			CarId:           int64(val.CarID),
			OrderDate:       val.OrderDate.Format(time.RFC3339),
			PickupDate:      val.PickupDate.Format(time.RFC3339),
			DropoffDate:     val.DropoffDate.Format(time.RFC3339),
			PickupLocation:  val.PickupLocation,
			PickupLat:       val.PickupLat,
			PickupLong:      val.PickupLong,
			DropoffLocation: val.DropoffLocation,
			DropoffLat:      val.DropoffLat,
			DropoffLong:     val.DropoffLong,
			CreatedBy:       int64(val.CreatedBy),
			CreatedAt:       val.CreatedAt.Format(time.RFC3339),
			UpdatedBy:       int64(val.UpdatedBy),
			UpdatedAt:       val.UpdatedAt.Format(time.RFC3339),
			DeletedBy:       deletedBy,
			DeletedAt:       deletedAt,
		})
	}

	return &grpcmodel.GetOrderByParamReply{
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
