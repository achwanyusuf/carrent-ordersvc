package model

import (
	"context"
	"time"

	"github.com/achwanyusuf/carrent-lib/pkg/logger"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/grpcmodel"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/psqlmodel"
	"github.com/volatiletech/null/v8"
)

func FillUpdateCar(car *psqlmodel.Car, v *grpcmodel.UpdateCarRequest) {
	if v.CarName != nil {
		car.CarName = *v.CarName
	}

	if v.DayRate != nil {
		car.DayRate = *v.DayRate
	}

	if v.MonthRate != nil {
		car.MonthRate = *v.MonthRate
	}

	if v.Image != nil {
		car.Image = *v.Image
	}
}

func FillUpdateOrder(ctx context.Context, log logger.Logger, order *psqlmodel.Order, v *grpcmodel.UpdateOrderRequest) {
	if v.CarId != nil {
		carID := int(*v.CarId)
		order.CarID = carID
	}

	if v.OrderDate != nil {
		orderDate, err := time.Parse(time.RFC3339, *v.OrderDate)
		if err != nil {
			log.Error(ctx, err)
		}
		order.OrderDate = orderDate
	}

	if v.PickupDate != nil {
		pickupDate, err := time.Parse(time.RFC3339, *v.PickupDate)
		if err != nil {
			log.Error(ctx, err)
		}
		order.PickupDate = pickupDate
	}

	if v.DropoffDate != nil {
		dropoffDate, err := time.Parse(time.RFC3339, *v.DropoffDate)
		if err != nil {
			log.Error(ctx, err)
		}
		order.DropoffDate = dropoffDate
	}

	if v.DropoffLocation != nil {
		order.DropoffLocation = *v.DropoffLocation
	}

	if v.DropoffLat != nil {
		order.DropoffLat = *v.DropoffLat
	}

	if v.DropoffLong != nil {
		order.DropoffLong = *v.DropoffLong
	}

	if v.PickupLocation != nil {
		order.PickupLocation = *v.PickupLocation
	}

	if v.PickupLat != nil {
		order.PickupLat = *v.PickupLat
	}

	if v.PickupLong != nil {
		order.PickupLong = *v.PickupLong
	}
}

func TransformSingleCarReplyToCar(ctx context.Context, v *grpcmodel.SingleCarReply, log logger.Logger) Car {
	var (
		deletedAt null.Time
	)
	if v == nil {
		return Car{}
	}

	createdAt, err := time.Parse(time.RFC3339, v.CreatedAt)
	if err != nil {
		log.Error(ctx, err, "error parsing time")
	}

	updatedAt, err := time.Parse(time.RFC3339, v.UpdatedAt)
	if err != nil {
		log.Error(ctx, err, "error parsing time")
	}

	if v.DeletedAt != nil {
		tParse, err := time.Parse(time.RFC3339, *v.DeletedAt)
		if err != nil {
			log.Error(ctx, err, "error parsing time")
		}
		deletedAt = null.TimeFrom(tParse)
	}

	return Car{
		ID:        v.Id,
		CarName:   v.CarName,
		DayRate:   v.DayRate,
		MonthRate: v.MonthRate,
		Image:     v.Image,
		BaseInformation: BaseInformation{
			CreatedBy: v.CreatedBy,
			CreatedAt: createdAt,
			UpdatedBy: v.UpdatedBy,
			UpdatedAt: updatedAt,
			DeletedBy: null.Int64FromPtr(v.DeletedBy),
			DeletedAt: deletedAt,
		},
	}
}

func TransformCarByParamReplyToCar(ctx context.Context, v *grpcmodel.GetCarByParamReply, log logger.Logger) ([]Car, Pagination) {
	var (
		cars       []Car
		pagination Pagination
	)

	for _, val := range v.Data {
		createdAt, err := time.Parse(time.RFC3339, val.CreatedAt)
		if err != nil {
			log.Error(ctx, err, "error parsing time")
		}

		updatedAt, err := time.Parse(time.RFC3339, val.UpdatedAt)
		if err != nil {
			log.Error(ctx, err, "error parsing time")
		}

		deletedAt := null.Time{}
		if val.DeletedAt != nil {
			tParse, err := time.Parse(time.RFC3339, *val.DeletedAt)
			if err != nil {
				log.Error(ctx, err, "error parsing time")
			}
			deletedAt = null.TimeFrom(tParse)
		}
		cars = append(cars, Car{
			ID:        val.Id,
			CarName:   val.CarName,
			DayRate:   val.DayRate,
			MonthRate: val.MonthRate,
			Image:     val.Image,
			BaseInformation: BaseInformation{
				CreatedBy: val.CreatedBy,
				CreatedAt: createdAt,
				UpdatedBy: val.UpdatedBy,
				UpdatedAt: updatedAt,
				DeletedBy: null.Int64FromPtr(val.DeletedBy),
				DeletedAt: deletedAt,
			},
		})
	}

	if v.Pagination != nil {
		pagination = Pagination{
			CurrentPage:     v.Pagination.CurrentPage,
			CurrentElements: v.Pagination.CurrentElement,
			TotalPages:      v.Pagination.TotalPages,
			TotalElements:   v.Pagination.TotalElements,
			SortBy:          v.Pagination.SortBy,
		}
	}

	return cars, pagination
}

func TransformSingleOrderReplyToOrder(ctx context.Context, v *grpcmodel.SingleOrderReply, log logger.Logger) Order {
	var (
		deletedAt   null.Time
		pickupDate  time.Time
		dropoffDate time.Time
		orderDate   time.Time
	)
	if v == nil {
		return Order{}
	}

	createdAt, err := time.Parse(time.RFC3339, v.CreatedAt)
	if err != nil {
		log.Error(ctx, err, "error parsing time")
	}

	updatedAt, err := time.Parse(time.RFC3339, v.UpdatedAt)
	if err != nil {
		log.Error(ctx, err, "error parsing time")
	}

	if v.DeletedAt != nil {
		tParse, err := time.Parse(time.RFC3339, *v.DeletedAt)
		if err != nil {
			log.Error(ctx, err, "error parsing time")
		}
		deletedAt = null.TimeFrom(tParse)
	}

	pickupDate, err = time.Parse(time.RFC3339, v.PickupDate)
	if err != nil {
		log.Error(ctx, err)
	}

	dropoffDate, err = time.Parse(time.RFC3339, v.DropoffDate)
	if err != nil {
		log.Error(ctx, err)
	}

	orderDate, err = time.Parse(time.RFC3339, v.OrderDate)
	if err != nil {
		log.Error(ctx, err)
	}

	return Order{
		ID:              v.Id,
		CarID:           v.CarId,
		OrderDate:       orderDate,
		PickupDate:      pickupDate,
		DropoffDate:     dropoffDate,
		PickupLocation:  v.PickupLocation,
		PickupLat:       v.PickupLat,
		PickupLong:      v.PickupLong,
		DropoffLocation: v.DropoffLocation,
		DropoffLat:      v.DropoffLat,
		DropoffLong:     v.DropoffLong,
		BaseInformation: BaseInformation{
			CreatedBy: v.CreatedBy,
			CreatedAt: createdAt,
			UpdatedBy: v.UpdatedBy,
			UpdatedAt: updatedAt,
			DeletedBy: null.Int64FromPtr(v.DeletedBy),
			DeletedAt: deletedAt,
		},
	}
}

func TransformOrderByParamReplyToOrder(ctx context.Context, v *grpcmodel.GetOrderByParamReply, log logger.Logger) ([]Order, Pagination) {
	var (
		orders      []Order
		pickupDate  time.Time
		dropoffDate time.Time
		orderDate   time.Time
		pagination  Pagination
	)

	for _, val := range v.Data {
		createdAt, err := time.Parse(time.RFC3339, val.CreatedAt)
		if err != nil {
			log.Error(ctx, err, "error parsing time")
		}

		updatedAt, err := time.Parse(time.RFC3339, val.UpdatedAt)
		if err != nil {
			log.Error(ctx, err, "error parsing time")
		}

		deletedAt := null.Time{}
		if val.DeletedAt != nil {
			tParse, err := time.Parse(time.RFC3339, *val.DeletedAt)
			if err != nil {
				log.Error(ctx, err, "error parsing time")
			}
			deletedAt = null.TimeFrom(tParse)
		}
		orders = append(orders, Order{
			ID:              val.Id,
			CarID:           val.CarId,
			OrderDate:       orderDate,
			PickupDate:      pickupDate,
			DropoffDate:     dropoffDate,
			PickupLocation:  val.PickupLocation,
			PickupLat:       val.PickupLat,
			PickupLong:      val.PickupLong,
			DropoffLocation: val.DropoffLocation,
			DropoffLat:      val.DropoffLat,
			DropoffLong:     val.DropoffLong,
			BaseInformation: BaseInformation{
				CreatedBy: val.CreatedBy,
				CreatedAt: createdAt,
				UpdatedBy: val.UpdatedBy,
				UpdatedAt: updatedAt,
				DeletedBy: null.Int64FromPtr(val.DeletedBy),
				DeletedAt: deletedAt,
			},
		})
	}

	if v.Pagination != nil {
		pagination = Pagination{
			CurrentPage:     v.Pagination.CurrentPage,
			CurrentElements: v.Pagination.CurrentElement,
			TotalPages:      v.Pagination.TotalPages,
			TotalElements:   v.Pagination.TotalElements,
			SortBy:          v.Pagination.SortBy,
		}
	}

	return orders, pagination
}
