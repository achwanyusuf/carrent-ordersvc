package car

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

func (c *CarDep) insertPSQL(ctx *context.Context, data *psqlmodel.Car) error {
	tx, err := c.DB.BeginTx(*ctx, nil)
	if err != nil {
		return errormsg.WrapErr(svcerr.OrderSVCPSQLErrorTransaction, err, "error begin transaction")
	}

	err = data.Insert(*ctx, tx, boil.Infer())
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			c.Log.Warn(*ctx, errormsg.WrapErr(svcerr.OrderSVCPSQLErrorRollback, err, "error rollback"))
		}
		return errormsg.WrapErr(svcerr.OrderSVCPSQLErrorInsert, err, "error insert")
	}
	err = tx.Commit()
	if err != nil {
		return errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error commit")
	}
	return nil
}

func (c *CarDep) getSingleByParamPSQL(ctx *context.Context, param *model.GetCarByParam) (psqlmodel.Car, error) {
	var res psqlmodel.Car
	qr := param.GetQuery()
	car, err := psqlmodel.Cars(qr...).One(*ctx, c.DB)
	if err == sql.ErrNoRows {
		return res, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get cars")
	}

	if err != nil {
		return res, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get cars")
	}

	return *car, nil
}

func (c *CarDep) updatePSQL(ctx *context.Context, car *psqlmodel.Car) error {
	tx, err := c.DB.BeginTx(*ctx, nil)
	if err != nil {
		return errormsg.WrapErr(svcerr.OrderSVCPSQLErrorTransaction, err, "error begin transaction")
	}

	_, err = car.Update(*ctx, tx, boil.Infer())
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			c.Log.Warn(*ctx, errormsg.WrapErr(svcerr.OrderSVCPSQLErrorRollback, err, "error rollback"))
		}
		return errormsg.WrapErr(svcerr.OrderSVCPSQLErrorUpdate, err, "error update")
	}
	err = tx.Commit()
	if err != nil {
		return errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error commit")
	}
	return nil
}

func (c *CarDep) deletePSQL(ctx *context.Context, car *psqlmodel.Car, id int64, isHardDelete bool) error {
	tx, err := c.DB.BeginTx(*ctx, nil)
	if err != nil {
		return errormsg.WrapErr(svcerr.OrderSVCPSQLErrorTransaction, err, "error begin transaction")
	}

	_, err = car.Delete(*ctx, tx, isHardDelete)
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			c.Log.Warn(*ctx, errormsg.WrapErr(svcerr.OrderSVCPSQLErrorRollback, err, "error rollback"))
		}
		return errormsg.WrapErr(svcerr.OrderSVCPSQLErrorUpdate, err, "error delete")
	}

	if !isHardDelete {
		car.DeletedBy = null.NewInt(int(id), true)
		_, err = car.Update(*ctx, tx, boil.Infer())
		if err != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				c.Log.Warn(*ctx, errormsg.WrapErr(svcerr.OrderSVCPSQLErrorRollback, err, "error rollback"))
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

func (c *CarDep) getByParamPSQL(ctx *context.Context, param *model.GetCarsByParam) (psqlmodel.CarSlice, model.Pagination, error) {
	var totalPages int64 = 1
	if param.Limit == 0 {
		param.Limit = int64(c.Conf.DefaultPageLimit)
		if c.Conf.DefaultPageLimit == 0 {
			param.Limit = int64(model.DefaultPageLimit)
		}
	}

	if param.Page == 0 {
		param.Page = 1
	}

	qr := param.GetQuery()
	count, err := psqlmodel.Cars(qr...).Count(*ctx, c.DB)
	if err != nil {
		return psqlmodel.CarSlice{}, model.Pagination{}, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error count data")
	}
	qr = append(qr, qm.Offset(int((param.Page-1)*param.Limit)))
	qr = append(qr, qm.Limit(int(param.Limit)))
	cars, err := psqlmodel.Cars(qr...).All(*ctx, c.DB)
	if err == sql.ErrNoRows {
		return cars, model.Pagination{}, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get cars")
	}
	if err != nil {
		return cars, model.Pagination{}, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get cars")
	}
	if count > 0 {
		totalPages = (count / param.Limit) + 1
	}
	return cars, model.Pagination{
		CurrentPage:     param.Page,
		CurrentElements: int64(len(cars)),
		TotalElements:   count,
		TotalPages:      totalPages,
		SortBy:          param.OrderBy.String,
	}, nil
}
