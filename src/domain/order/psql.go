package order

import (
	"context"
	"database/sql"

	"github.com/achwanyusuf/carrent-lib/pkg/errormsg"
	"github.com/achwanyusuf/carrent-ordersvc/src/model"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/psqlmodel"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/svcerr"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (o *OrderDep) insertPSQL(ctx *context.Context, data *psqlmodel.Order) error {
	tx, err := o.DB.BeginTx(*ctx, nil)
	if err != nil {
		return errormsg.WrapErr(svcerr.OrderSVCPSQLErrorTransaction, err, "error begin transaction")
	}

	err = data.Insert(*ctx, tx, boil.Infer())
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			o.Log.Warn(*ctx, errormsg.WrapErr(svcerr.OrderSVCPSQLErrorRollback, err, "error rollback"))
		}
		return errormsg.WrapErr(svcerr.OrderSVCPSQLErrorInsert, err, "error insert")
	}
	err = tx.Commit()
	if err != nil {
		return errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error commit")
	}
	return nil
}

func (o *OrderDep) getSingleByParamPSQL(ctx *context.Context, param *model.GetOrderByParam) (psqlmodel.Order, error) {
	var res psqlmodel.Order
	qr := param.GetQuery()
	order, err := psqlmodel.Orders(qr...).One(*ctx, o.DB)
	if err == sql.ErrNoRows {
		return res, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get orders")
	}

	if err != nil {
		return res, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get orders")
	}

	return *order, nil
}

func (o *OrderDep) updatePSQL(ctx *context.Context, order *psqlmodel.Order) error {
	tx, err := o.DB.BeginTx(*ctx, nil)
	if err != nil {
		return errormsg.WrapErr(svcerr.OrderSVCPSQLErrorTransaction, err, "error begin transaction")
	}

	_, err = order.Update(*ctx, tx, boil.Infer())
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			o.Log.Warn(*ctx, errormsg.WrapErr(svcerr.OrderSVCPSQLErrorRollback, err, "error rollback"))
		}
		return errormsg.WrapErr(svcerr.OrderSVCPSQLErrorUpdate, err, "error update")
	}
	err = tx.Commit()
	if err != nil {
		return errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error commit")
	}
	return nil
}

func (o *OrderDep) deletePSQL(ctx *context.Context, order *psqlmodel.Order, id int64, isHardDelete bool) error {
	tx, err := o.DB.BeginTx(*ctx, nil)
	if err != nil {
		return errormsg.WrapErr(svcerr.OrderSVCPSQLErrorTransaction, err, "error begin transaction")
	}

	_, err = order.Delete(*ctx, tx, isHardDelete)
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			o.Log.Warn(*ctx, errormsg.WrapErr(svcerr.OrderSVCPSQLErrorRollback, err, "error rollback"))
		}
		return errormsg.WrapErr(svcerr.OrderSVCPSQLErrorUpdate, err, "error delete")
	}

	if !isHardDelete {
		order.DeletedBy = null.NewInt(int(id), true)
		_, err = order.Update(*ctx, tx, boil.Infer())
		if err != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				o.Log.Warn(*ctx, errormsg.WrapErr(svcerr.OrderSVCPSQLErrorRollback, err, "error rollback"))
			}
			return errormsg.WrapErr(svcerr.OrderSVCPSQLErrorUpdate, err, "error update")
		}
	}
	err = tx.Commit()
	if err != nil {
		return errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error commit")
	}
	return nil
}

func (o *OrderDep) getByParamPSQL(ctx *context.Context, param *model.GetOrdersByParam) (psqlmodel.OrderSlice, model.Pagination, error) {
	var totalPages int64 = 1
	if param.Limit == 0 {
		param.Limit = int64(o.Conf.DefaultPageLimit)
		if o.Conf.DefaultPageLimit == 0 {
			param.Limit = int64(model.DefaultPageLimit)
		}
	}

	if param.Page == 0 {
		param.Page = 1
	}

	qr := param.GetQuery()
	count, err := psqlmodel.Orders(qr...).Count(*ctx, o.DB)
	if err != nil {
		return psqlmodel.OrderSlice{}, model.Pagination{}, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error count data")
	}
	qr = append(qr, qm.Offset(int((param.Page-1)*param.Limit)))
	qr = append(qr, qm.Limit(int(param.Limit)))
	orders, err := psqlmodel.Orders(qr...).All(*ctx, o.DB)
	if err == sql.ErrNoRows {
		return orders, model.Pagination{}, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get orders")
	}
	if err != nil {
		return orders, model.Pagination{}, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get orders")
	}
	if count > 0 {
		totalPages = (count / param.Limit) + 1
	}
	return orders, model.Pagination{
		CurrentPage:     param.Page,
		CurrentElements: int64(len(orders)),
		TotalElements:   count,
		TotalPages:      totalPages,
		SortBy:          param.OrderBy.String,
	}, nil
}
