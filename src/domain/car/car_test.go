package car_test

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/achwanyusuf/carrent-lib/pkg/errormsg"
	"github.com/achwanyusuf/carrent-lib/pkg/logger"
	"github.com/achwanyusuf/carrent-ordersvc/src/domain/car"
	"github.com/achwanyusuf/carrent-ordersvc/src/model"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/psqlmodel"

	gosqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/go-redis/redismock/v9"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestGetSingleParam(t *testing.T) {
	dbSQL, sqlMock, err := gosqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	oldDB := boil.GetDB()
	defer func() {
		dbSQL.Close()
		boil.SetDB(oldDB)
	}()
	boil.SetDB(dbSQL)
	dbRedis, redisMock := redismock.NewClientMock()
	acc := car.CarDep{
		Log:   logger.New(&logger.Config{}),
		DB:    dbSQL,
		Redis: dbRedis,
		Conf: car.Conf{
			DefaultPageLimit:    10,
			RedisExpirationTime: 30 * time.Second,
		},
	}
	ctx := context.Background()
	Convey("test get single param", t, FailureHalts, func() {
		tests := []struct {
			testType string
			testDesc string
			args     struct {
				ctx          *context.Context
				cacheControl string
				param        *model.GetCarByParam
			}
			want struct {
				result psqlmodel.Car
			}
			wantErr  string
			mockFunc func(a struct {
				ctx          *context.Context
				cacheControl string
				param        *model.GetCarByParam
			})
		}{
			{
				testType: "P",
				testDesc: "test get single with cache control must revalidate",
				args: struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarByParam
				}{
					ctx:          &ctx,
					cacheControl: model.MustRevalidate,
					param: &model.GetCarByParam{
						ID: null.NewInt64(1, true),
					},
				},
				want: struct{ result psqlmodel.Car }{
					result: psqlmodel.Car{
						ID:        1,
						CarName:   "sedan",
						DayRate:   12000,
						MonthRate: 300000,
						Image:     "http://link.com",
						CreatedBy: 0,
						CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						UpdatedBy: 0,
						UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
					},
				},
				wantErr: "",
				mockFunc: func(a struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarByParam
				}) {
					data := psqlmodel.Car{
						ID:        1,
						CarName:   "sedan",
						DayRate:   12000,
						MonthRate: 300000,
						Image:     "http://link.com",
						CreatedBy: 0,
						CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						UpdatedBy: 0,
						UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
					}
					rows := sqlMock.NewRows([]string{"id", "car_name", "day_rate", "month_rate", "image", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at"}).AddRow(data.ID, data.CarName, data.DayRate, data.MonthRate, data.Image, data.CreatedBy, data.CreatedAt, data.UpdatedBy, data.UpdatedAt, data.DeletedBy, data.DeletedAt)
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT \"cars\".* FROM \"cars\" WHERE (id=$1) AND (\"cars\".\"deleted_at\" is null) LIMIT 1;")).WithArgs(a.param.ID.Int64).WillReturnRows(rows).RowsWillBeClosed()
					dataStr, _ := json.Marshal(&data)
					redisMock.ExpectDel("gspCar:{\"id\":1,\"car_name\":null,\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null}").SetVal(0)
					redisMock.ExpectSet("gspCar:{\"id\":1,\"car_name\":null,\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null}", string(dataStr), 30*time.Second).SetVal(string(dataStr))
				},
			},
			{
				testType: "N",
				testDesc: "test get single with cache control must revalidate error no rows",
				args: struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarByParam
				}{
					ctx:          &ctx,
					cacheControl: model.MustRevalidate,
					param: &model.GetCarByParam{
						ID: null.NewInt64(1, true),
					},
				},
				wantErr: "sql: no rows in result set",
				mockFunc: func(a struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarByParam
				}) {
					rows := sqlMock.NewRows([]string{"id", "car_name", "day_rate", "month_rate", "image", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at"})
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT \"cars\".* FROM \"cars\" WHERE (id=$1) AND (\"cars\".\"deleted_at\" is null) LIMIT 1;")).WithArgs(a.param.ID.Int64).WillReturnRows(rows).RowsWillBeClosed()
				},
			},
			{
				testType: "N",
				testDesc: "test get single with cache control must revalidate error",
				args: struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarByParam
				}{
					ctx:          &ctx,
					cacheControl: model.MustRevalidate,
					param: &model.GetCarByParam{
						ID: null.NewInt64(1, true),
					},
				},
				wantErr: "psqlmodel: failed to execute a one query for cars: bind failed to execute query: error get data",
				mockFunc: func(a struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarByParam
				}) {
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT \"cars\".* FROM \"cars\" WHERE (id=$1) AND (\"cars\".\"deleted_at\" is null) LIMIT 1;")).WithArgs(a.param.ID.Int64).WillReturnError(errors.New("error get data")).RowsWillBeClosed()
				},
			},
			{
				testType: "P",
				testDesc: "test get single with cache control none",
				args: struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarByParam
				}{
					ctx: &ctx,
					param: &model.GetCarByParam{
						ID: null.NewInt64(1, true),
					},
				},
				want: struct{ result psqlmodel.Car }{
					result: psqlmodel.Car{
						ID:        1,
						CarName:   "sedan",
						DayRate:   12000,
						MonthRate: 300000,
						Image:     "http://link.com",
						CreatedBy: 0,
						CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						UpdatedBy: 0,
						UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
					},
				},
				wantErr: "",
				mockFunc: func(a struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarByParam
				}) {
					data := psqlmodel.Car{
						ID:        1,
						CarName:   "sedan",
						DayRate:   12000,
						MonthRate: 300000,
						Image:     "http://link.com",
						CreatedBy: 0,
						CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						UpdatedBy: 0,
						UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
					}
					strData, _ := json.Marshal(data)
					redisMock.ExpectGet("gspCar:{\"id\":1,\"car_name\":null,\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null}").SetVal(string(strData))
				},
			},
			{
				testType: "P",
				testDesc: "test get single with cache control none redis nil",
				args: struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarByParam
				}{
					ctx: &ctx,
					param: &model.GetCarByParam{
						ID: null.NewInt64(1, true),
					},
				},
				want: struct{ result psqlmodel.Car }{
					result: psqlmodel.Car{
						ID:        1,
						CarName:   "sedan",
						DayRate:   12000,
						MonthRate: 300000,
						Image:     "http://link.com",
						CreatedBy: 0,
						CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						UpdatedBy: 0,
						UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
					},
				},
				wantErr: "",
				mockFunc: func(a struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarByParam
				}) {
					redisMock.ExpectGet("gspCar:{\"id\":1,\"car_name\":null,\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null}").RedisNil()
					data := psqlmodel.Car{
						ID:        1,
						CarName:   "sedan",
						DayRate:   12000,
						MonthRate: 300000,
						Image:     "http://link.com",
						CreatedBy: 0,
						CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						UpdatedBy: 0,
						UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
					}
					rows := sqlMock.NewRows([]string{"id", "car_name", "day_rate", "month_rate", "image", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at"}).AddRow(data.ID, data.CarName, data.DayRate, data.MonthRate, data.Image, data.CreatedBy, data.CreatedAt, data.UpdatedBy, data.UpdatedAt, data.DeletedBy, data.DeletedAt)
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT \"cars\".* FROM \"cars\" WHERE (id=$1) AND (\"cars\".\"deleted_at\" is null) LIMIT 1;")).WithArgs(a.param.ID.Int64).WillReturnRows(rows)
					dataStr, _ := json.Marshal(&data)
					redisMock.ExpectDel("gspCar:{\"id\":1,\"car_name\":null,\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null}").SetVal(0)
					redisMock.ExpectSet("gspCar:{\"id\":1,\"car_name\":null,\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null}", string(dataStr), 30*time.Second).SetVal(string(dataStr))
				},
			},
			{
				testType: "N",
				testDesc: "test get single with cache control none error delete redis",
				args: struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarByParam
				}{
					ctx: &ctx,
					param: &model.GetCarByParam{
						ID: null.NewInt64(1, true),
					},
				},
				wantErr: "error delete redis",
				mockFunc: func(a struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarByParam
				}) {
					redisMock.ExpectGet("gspCar:{\"id\":1,\"car_name\":null,\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null}").RedisNil()
					data := psqlmodel.Car{
						ID:        1,
						CarName:   "sedan",
						DayRate:   12000,
						MonthRate: 300000,
						Image:     "http://link.com",
						CreatedBy: 0,
						CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						UpdatedBy: 0,
						UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
					}
					rows := sqlMock.NewRows([]string{"id", "car_name", "day_rate", "month_rate", "image", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at"}).AddRow(data.ID, data.CarName, data.DayRate, data.MonthRate, data.Image, data.CreatedBy, data.CreatedAt, data.UpdatedBy, data.UpdatedAt, data.DeletedBy, data.DeletedAt)
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT \"cars\".* FROM \"cars\" WHERE (id=$1) AND (\"cars\".\"deleted_at\" is null) LIMIT 1;")).WithArgs(a.param.ID.Int64).WillReturnRows(rows)
					redisMock.ExpectDel("gspCar:{\"id\":1,\"car_name\":null,\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null}").SetErr(errors.New("error delete redis"))
				},
			},

			{
				testType: "N",
				testDesc: "test get single with cache control none error delete redis",
				args: struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarByParam
				}{
					ctx: &ctx,
					param: &model.GetCarByParam{
						ID: null.NewInt64(1, true),
					},
				},
				wantErr: "error set redis",
				mockFunc: func(a struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarByParam
				}) {
					redisMock.ExpectGet("gspCar:{\"id\":1,\"car_name\":null,\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null}").RedisNil()
					data := psqlmodel.Car{
						ID:        1,
						CarName:   "sedan",
						DayRate:   12000,
						MonthRate: 300000,
						Image:     "http://link.com",
						CreatedBy: 0,
						CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						UpdatedBy: 0,
						UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
					}
					rows := sqlMock.NewRows([]string{"id", "car_name", "day_rate", "month_rate", "image", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at"}).AddRow(data.ID, data.CarName, data.DayRate, data.MonthRate, data.Image, data.CreatedBy, data.CreatedAt, data.UpdatedBy, data.UpdatedAt, data.DeletedBy, data.DeletedAt)
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT \"cars\".* FROM \"cars\" WHERE (id=$1) AND (\"cars\".\"deleted_at\" is null) LIMIT 1;")).WithArgs(a.param.ID.Int64).WillReturnRows(rows)
					dataStr, _ := json.Marshal(&data)
					redisMock.ExpectDel("gspCar:{\"id\":1,\"car_name\":null,\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null}").SetVal(0)
					redisMock.ExpectSet("gspCar:{\"id\":1,\"car_name\":null,\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null}", string(dataStr), 30*time.Second).SetErr(errors.New("error set redis"))
				},
			},
			{
				testType: "N",
				testDesc: "test get single with cache control none error get redis",
				args: struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarByParam
				}{
					ctx: &ctx,
					param: &model.GetCarByParam{
						ID: null.NewInt64(1, true),
					},
				},
				wantErr: "error get redis",
				mockFunc: func(a struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarByParam
				}) {
					redisMock.ExpectGet("gspCar:{\"id\":1,\"car_name\":null,\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null}").SetErr(errors.New("error get redis"))
				},
			},
		}
		for idx, test := range tests {
			Convey(fmt.Sprintf("%d - [%s] : %s", idx, test.testType, test.testDesc), func() {
				test.mockFunc(test.args)
				car, err := acc.GetSingleByParam(test.args.ctx, test.args.cacheControl, test.args.param)
				if test.testType == "N" {
					e := errormsg.GetErrorData(err)
					So(e.DebugError.Error(), ShouldEqual, test.wantErr)
				} else {
					So(err, ShouldBeNil)
					So(car.ID, ShouldEqual, test.want.result.ID)
					So(car.CarName, ShouldEqual, test.want.result.CarName)
					So(car.DayRate, ShouldEqual, test.want.result.DayRate)
					So(car.MonthRate, ShouldEqual, test.want.result.MonthRate)
					So(car.Image, ShouldEqual, test.want.result.Image)
					So(car.CreatedBy, ShouldEqual, test.want.result.CreatedBy)
					So(car.CreatedAt, ShouldEqual, test.want.result.CreatedAt)
					So(car.UpdatedBy, ShouldEqual, test.want.result.UpdatedBy)
					So(car.UpdatedAt, ShouldEqual, test.want.result.UpdatedAt)
					So(car.DeletedBy, ShouldEqual, test.want.result.DeletedBy)
					So(car.DeletedAt, ShouldEqual, test.want.result.DeletedAt)
				}
			})
		}
	})
}

func TestGetByParam(t *testing.T) {
	dbSQL, sqlMock, err := gosqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	oldDB := boil.GetDB()
	defer func() {
		dbSQL.Close()
		boil.SetDB(oldDB)
	}()
	boil.SetDB(dbSQL)
	dbRedis, redisMock := redismock.NewClientMock()
	acc := car.CarDep{
		Log:   logger.New(&logger.Config{}),
		DB:    dbSQL,
		Redis: dbRedis,
		Conf: car.Conf{
			DefaultPageLimit:    10,
			RedisExpirationTime: 30 * time.Second,
		},
	}
	ctx := context.Background()
	Convey("test get by param", t, FailureHalts, func() {
		tests := []struct {
			testType string
			testDesc string
			args     struct {
				ctx          *context.Context
				cacheControl string
				param        *model.GetCarsByParam
			}
			want struct {
				result     []psqlmodel.Car
				pagination model.Pagination
			}
			wantErr  string
			mockFunc func(a struct {
				ctx          *context.Context
				cacheControl string
				param        *model.GetCarsByParam
			})
		}{
			{
				testType: "P",
				testDesc: "test get with cache control must revalidate",
				args: struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}{
					ctx:          &ctx,
					cacheControl: model.MustRevalidate,
					param: &model.GetCarsByParam{
						GetCarByParam: model.GetCarByParam{
							CarName: null.StringFrom("Toyota"),
						},
					},
				},
				want: struct {
					result     []psqlmodel.Car
					pagination model.Pagination
				}{
					result: []psqlmodel.Car{
						{
							ID:        1,
							CarName:   "Toyota Calya",
							DayRate:   12000,
							MonthRate: 300000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        2,
							CarName:   "Toyota Rush",
							DayRate:   400000,
							MonthRate: 1200000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        3,
							CarName:   "Toyota Supra",
							DayRate:   1000000,
							MonthRate: 30000000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
					},
					pagination: model.Pagination{
						CurrentPage:     1,
						CurrentElements: 3,
						TotalElements:   3,
						TotalPages:      1,
						SortBy:          "",
					},
				},
				wantErr: "",
				mockFunc: func(a struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}) {
					data := []psqlmodel.Car{
						{
							ID:        1,
							CarName:   "Toyota Calya",
							DayRate:   12000,
							MonthRate: 300000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        2,
							CarName:   "Toyota Rush",
							DayRate:   400000,
							MonthRate: 1200000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        3,
							CarName:   "Toyota Supra",
							DayRate:   1000000,
							MonthRate: 30000000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
					}
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM \"cars\" WHERE (car_name like $1) AND (\"cars\".\"deleted_at\" is null);")).WithArgs("%" + a.param.CarName.String + "%").WillReturnRows(sqlMock.NewRows([]string{"COUNT(*)"}).AddRow(3)).RowsWillBeClosed()
					rows := sqlMock.NewRows([]string{"id", "car_name", "day_rate", "month_rate", "image", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at"})
					for _, d := range data {
						rows.AddRow(d.ID, d.CarName, d.DayRate, d.MonthRate, d.Image, d.CreatedBy, d.CreatedAt, d.UpdatedBy, d.UpdatedAt, d.DeletedBy, d.DeletedAt)
					}
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT \"cars\".* FROM \"cars\" WHERE (car_name like $1) AND (\"cars\".\"deleted_at\" is null) LIMIT 10;")).WithArgs("%" + a.param.CarName.String + "%").WillReturnRows(rows).RowsWillBeClosed()
					dataStr, _ := json.Marshal(&data)
					pagination := model.Pagination{
						CurrentPage:     1,
						CurrentElements: int64(len(data)),
						TotalElements:   3,
						TotalPages:      1,
						SortBy:          a.param.OrderBy.String,
					}
					paginationStr, _ := json.Marshal(&pagination)
					key := "gpCar:{\"id\":null,\"car_name\":\"Toyota\",\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null,\"order_by\":null,\"limit\":0,\"page\":0}"
					redisMock.ExpectDel(key).SetVal(0)
					redisMock.ExpectSet(key, string(dataStr), 30*time.Second).SetVal(string(dataStr))
					keypg := "gppgCar:{\"id\":null,\"car_name\":\"Toyota\",\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null,\"order_by\":null,\"limit\":0,\"page\":0}"
					redisMock.ExpectDel(keypg).SetVal(0)
					redisMock.ExpectSet(keypg, string(paginationStr), 30*time.Second).SetVal(string(paginationStr))
				},
			},
			{
				testType: "P",
				testDesc: "test get with cache control must revalidate no rows",
				args: struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}{
					ctx:          &ctx,
					cacheControl: model.MustRevalidate,
					param: &model.GetCarsByParam{
						GetCarByParam: model.GetCarByParam{
							CarName: null.StringFrom("Toyota"),
						},
					},
				},
				want: struct {
					result     []psqlmodel.Car
					pagination model.Pagination
				}{
					result: []psqlmodel.Car{},
					pagination: model.Pagination{
						CurrentPage:     1,
						CurrentElements: 0,
						TotalElements:   0,
						TotalPages:      1,
						SortBy:          "",
					},
				},
				wantErr: "",
				mockFunc: func(a struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}) {
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM \"cars\" WHERE (car_name like $1) AND (\"cars\".\"deleted_at\" is null);")).WithArgs("%" + a.param.CarName.String + "%").WillReturnRows(sqlMock.NewRows([]string{"COUNT(*)"}).AddRow(0)).RowsWillBeClosed()
					rows := sqlMock.NewRows([]string{"id", "car_name", "day_rate", "month_rate", "image", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at"})
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT \"cars\".* FROM \"cars\" WHERE (car_name like $1) AND (\"cars\".\"deleted_at\" is null) LIMIT 10;")).WithArgs("%" + a.param.CarName.String + "%").WillReturnRows(rows).RowsWillBeClosed()
					pagination := model.Pagination{
						CurrentPage:     1,
						CurrentElements: 0,
						TotalElements:   0,
						TotalPages:      1,
						SortBy:          a.param.OrderBy.String,
					}
					paginationStr, _ := json.Marshal(&pagination)
					key := "gpCar:{\"id\":null,\"car_name\":\"Toyota\",\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null,\"order_by\":null,\"limit\":0,\"page\":0}"
					redisMock.ExpectDel(key).SetVal(0)
					redisMock.ExpectSet(key, "null", 30*time.Second).SetVal("null")
					keypg := "gppgCar:{\"id\":null,\"car_name\":\"Toyota\",\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null,\"order_by\":null,\"limit\":0,\"page\":0}"
					redisMock.ExpectDel(keypg).SetVal(0)
					redisMock.ExpectSet(keypg, string(paginationStr), 30*time.Second).SetVal(string(paginationStr))
				},
			},
			{
				testType: "N",
				testDesc: "test get with cache control must revalidate error count",
				args: struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}{
					ctx:          &ctx,
					cacheControl: model.MustRevalidate,
					param: &model.GetCarsByParam{
						GetCarByParam: model.GetCarByParam{
							CarName: null.StringFrom("Toyota"),
						},
					},
				},
				wantErr: "psqlmodel: failed to count cars rows: error count",
				mockFunc: func(a struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}) {
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM \"cars\" WHERE (car_name like $1) AND (\"cars\".\"deleted_at\" is null);")).WithArgs("%" + a.param.CarName.String + "%").WillReturnError(errors.New("error count"))
				},
			},
			{
				testType: "N",
				testDesc: "test get with cache control must revalidate error get",
				args: struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}{
					ctx:          &ctx,
					cacheControl: model.MustRevalidate,
					param: &model.GetCarsByParam{
						GetCarByParam: model.GetCarByParam{
							CarName: null.StringFrom("Toyota"),
						},
					},
				},
				wantErr: "psqlmodel: failed to assign all query results to Car slice: bind failed to execute query: error get data",
				mockFunc: func(a struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}) {
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM \"cars\" WHERE (car_name like $1) AND (\"cars\".\"deleted_at\" is null);")).WithArgs("%" + a.param.CarName.String + "%").WillReturnRows(sqlMock.NewRows([]string{"COUNT(*)"}).AddRow(0)).RowsWillBeClosed()
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT \"cars\".* FROM \"cars\" WHERE (car_name like $1) AND (\"cars\".\"deleted_at\" is null) LIMIT 10;")).WithArgs("%" + a.param.CarName.String + "%").WillReturnError(errors.New("error get data"))
				},
			},
			{
				testType: "N",
				testDesc: "test get with cache control must revalidate error delete redis data",
				args: struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}{
					ctx:          &ctx,
					cacheControl: model.MustRevalidate,
					param: &model.GetCarsByParam{
						GetCarByParam: model.GetCarByParam{
							CarName: null.StringFrom("Toyota"),
						},
					},
				},
				wantErr: "error delete redis in data",
				mockFunc: func(a struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}) {
					data := []psqlmodel.Car{
						{
							ID:        1,
							CarName:   "Toyota Calya",
							DayRate:   12000,
							MonthRate: 300000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        2,
							CarName:   "Toyota Rush",
							DayRate:   400000,
							MonthRate: 1200000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        3,
							CarName:   "Toyota Supra",
							DayRate:   1000000,
							MonthRate: 30000000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
					}
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM \"cars\" WHERE (car_name like $1) AND (\"cars\".\"deleted_at\" is null);")).WithArgs("%" + a.param.CarName.String + "%").WillReturnRows(sqlMock.NewRows([]string{"COUNT(*)"}).AddRow(3)).RowsWillBeClosed()
					rows := sqlMock.NewRows([]string{"id", "car_name", "day_rate", "month_rate", "image", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at"})
					for _, d := range data {
						rows.AddRow(d.ID, d.CarName, d.DayRate, d.MonthRate, d.Image, d.CreatedBy, d.CreatedAt, d.UpdatedBy, d.UpdatedAt, d.DeletedBy, d.DeletedAt)
					}
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT \"cars\".* FROM \"cars\" WHERE (car_name like $1) AND (\"cars\".\"deleted_at\" is null) LIMIT 10;")).WithArgs("%" + a.param.CarName.String + "%").WillReturnRows(rows).RowsWillBeClosed()
					key := "gpCar:{\"id\":null,\"car_name\":\"Toyota\",\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null,\"order_by\":null,\"limit\":0,\"page\":0}"
					redisMock.ExpectDel(key).SetErr(errors.New("error delete redis in data"))
				},
			},
			{
				testType: "N",
				testDesc: "test get with cache control must revalidate error set redis data",
				args: struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}{
					ctx:          &ctx,
					cacheControl: model.MustRevalidate,
					param: &model.GetCarsByParam{
						GetCarByParam: model.GetCarByParam{
							CarName: null.StringFrom("Toyota"),
						},
					},
				},
				wantErr: "error set redis in data",
				mockFunc: func(a struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}) {
					data := []psqlmodel.Car{
						{
							ID:        1,
							CarName:   "Toyota Calya",
							DayRate:   12000,
							MonthRate: 300000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        2,
							CarName:   "Toyota Rush",
							DayRate:   400000,
							MonthRate: 1200000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        3,
							CarName:   "Toyota Supra",
							DayRate:   1000000,
							MonthRate: 30000000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
					}
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM \"cars\" WHERE (car_name like $1) AND (\"cars\".\"deleted_at\" is null);")).WithArgs("%" + a.param.CarName.String + "%").WillReturnRows(sqlMock.NewRows([]string{"COUNT(*)"}).AddRow(3)).RowsWillBeClosed()
					rows := sqlMock.NewRows([]string{"id", "car_name", "day_rate", "month_rate", "image", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at"})
					for _, d := range data {
						rows.AddRow(d.ID, d.CarName, d.DayRate, d.MonthRate, d.Image, d.CreatedBy, d.CreatedAt, d.UpdatedBy, d.UpdatedAt, d.DeletedBy, d.DeletedAt)
					}
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT \"cars\".* FROM \"cars\" WHERE (car_name like $1) AND (\"cars\".\"deleted_at\" is null) LIMIT 10;")).WithArgs("%" + a.param.CarName.String + "%").WillReturnRows(rows).RowsWillBeClosed()
					dataStr, _ := json.Marshal(&data)
					key := "gpCar:{\"id\":null,\"car_name\":\"Toyota\",\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null,\"order_by\":null,\"limit\":0,\"page\":0}"
					redisMock.ExpectDel(key).SetVal(0)
					redisMock.ExpectSet(key, string(dataStr), 30*time.Second).SetErr(errors.New("error set redis in data"))
				},
			},
			{
				testType: "N",
				testDesc: "test get with cache control must revalidate error delete redis pagination",
				args: struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}{
					ctx:          &ctx,
					cacheControl: model.MustRevalidate,
					param: &model.GetCarsByParam{
						GetCarByParam: model.GetCarByParam{
							CarName: null.StringFrom("Toyota"),
						},
					},
				},
				wantErr: "error delete redis in pagination",
				mockFunc: func(a struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}) {
					data := []psqlmodel.Car{
						{
							ID:        1,
							CarName:   "Toyota Calya",
							DayRate:   12000,
							MonthRate: 300000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        2,
							CarName:   "Toyota Rush",
							DayRate:   400000,
							MonthRate: 1200000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        3,
							CarName:   "Toyota Supra",
							DayRate:   1000000,
							MonthRate: 30000000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
					}
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM \"cars\" WHERE (car_name like $1) AND (\"cars\".\"deleted_at\" is null);")).WithArgs("%" + a.param.CarName.String + "%").WillReturnRows(sqlMock.NewRows([]string{"COUNT(*)"}).AddRow(3)).RowsWillBeClosed()
					rows := sqlMock.NewRows([]string{"id", "car_name", "day_rate", "month_rate", "image", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at"})
					for _, d := range data {
						rows.AddRow(d.ID, d.CarName, d.DayRate, d.MonthRate, d.Image, d.CreatedBy, d.CreatedAt, d.UpdatedBy, d.UpdatedAt, d.DeletedBy, d.DeletedAt)
					}
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT \"cars\".* FROM \"cars\" WHERE (car_name like $1) AND (\"cars\".\"deleted_at\" is null) LIMIT 10;")).WithArgs("%" + a.param.CarName.String + "%").WillReturnRows(rows).RowsWillBeClosed()
					dataStr, _ := json.Marshal(&data)
					key := "gpCar:{\"id\":null,\"car_name\":\"Toyota\",\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null,\"order_by\":null,\"limit\":0,\"page\":0}"
					redisMock.ExpectDel(key).SetVal(0)
					redisMock.ExpectSet(key, string(dataStr), 30*time.Second).SetVal(string(dataStr))
					keypg := "gppgCar:{\"id\":null,\"car_name\":\"Toyota\",\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null,\"order_by\":null,\"limit\":0,\"page\":0}"
					redisMock.ExpectDel(keypg).SetErr(errors.New("error delete redis in pagination"))
				},
			},
			{
				testType: "N",
				testDesc: "test get with cache control must revalidate error set redis pagination",
				args: struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}{
					ctx:          &ctx,
					cacheControl: model.MustRevalidate,
					param: &model.GetCarsByParam{
						GetCarByParam: model.GetCarByParam{
							CarName: null.StringFrom("Toyota"),
						},
					},
				},
				wantErr: "error set redis in pagination",
				mockFunc: func(a struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}) {
					data := []psqlmodel.Car{
						{
							ID:        1,
							CarName:   "Toyota Calya",
							DayRate:   12000,
							MonthRate: 300000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        2,
							CarName:   "Toyota Rush",
							DayRate:   400000,
							MonthRate: 1200000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        3,
							CarName:   "Toyota Supra",
							DayRate:   1000000,
							MonthRate: 30000000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
					}
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM \"cars\" WHERE (car_name like $1) AND (\"cars\".\"deleted_at\" is null);")).WithArgs("%" + a.param.CarName.String + "%").WillReturnRows(sqlMock.NewRows([]string{"COUNT(*)"}).AddRow(3)).RowsWillBeClosed()
					rows := sqlMock.NewRows([]string{"id", "car_name", "day_rate", "month_rate", "image", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at"})
					for _, d := range data {
						rows.AddRow(d.ID, d.CarName, d.DayRate, d.MonthRate, d.Image, d.CreatedBy, d.CreatedAt, d.UpdatedBy, d.UpdatedAt, d.DeletedBy, d.DeletedAt)
					}
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT \"cars\".* FROM \"cars\" WHERE (car_name like $1) AND (\"cars\".\"deleted_at\" is null) LIMIT 10;")).WithArgs("%" + a.param.CarName.String + "%").WillReturnRows(rows).RowsWillBeClosed()
					dataStr, _ := json.Marshal(&data)
					pagination := model.Pagination{
						CurrentPage:     1,
						CurrentElements: int64(len(data)),
						TotalElements:   3,
						TotalPages:      1,
						SortBy:          a.param.OrderBy.String,
					}
					paginationStr, _ := json.Marshal(&pagination)
					key := "gpCar:{\"id\":null,\"car_name\":\"Toyota\",\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null,\"order_by\":null,\"limit\":0,\"page\":0}"
					redisMock.ExpectDel(key).SetVal(0)
					redisMock.ExpectSet(key, string(dataStr), 30*time.Second).SetVal(string(dataStr))
					keypg := "gppgCar:{\"id\":null,\"car_name\":\"Toyota\",\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null,\"order_by\":null,\"limit\":0,\"page\":0}"
					redisMock.ExpectDel(keypg).SetVal(0)
					redisMock.ExpectSet(keypg, string(paginationStr), 30*time.Second).SetErr(errors.New("error set redis in pagination"))
				},
			},
			{
				testType: "P",
				testDesc: "test get with cache control not must revalidate",
				args: struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}{
					ctx: &ctx,
					param: &model.GetCarsByParam{
						GetCarByParam: model.GetCarByParam{
							CarName: null.StringFrom("Toyota"),
						},
					},
				},
				want: struct {
					result     []psqlmodel.Car
					pagination model.Pagination
				}{
					result: []psqlmodel.Car{
						{
							ID:        1,
							CarName:   "Toyota Calya",
							DayRate:   12000,
							MonthRate: 300000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        2,
							CarName:   "Toyota Rush",
							DayRate:   400000,
							MonthRate: 1200000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        3,
							CarName:   "Toyota Supra",
							DayRate:   1000000,
							MonthRate: 30000000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
					},
					pagination: model.Pagination{
						CurrentPage:     1,
						CurrentElements: 3,
						TotalElements:   3,
						TotalPages:      1,
						SortBy:          "",
					},
				},
				wantErr: "",
				mockFunc: func(a struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}) {
					data := []psqlmodel.Car{
						{
							ID:        1,
							CarName:   "Toyota Calya",
							DayRate:   12000,
							MonthRate: 300000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        2,
							CarName:   "Toyota Rush",
							DayRate:   400000,
							MonthRate: 1200000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        3,
							CarName:   "Toyota Supra",
							DayRate:   1000000,
							MonthRate: 30000000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
					}
					dataStr, _ := json.Marshal(&data)
					key := "gpCar:{\"id\":null,\"car_name\":\"Toyota\",\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null,\"order_by\":null,\"limit\":0,\"page\":0}"
					redisMock.ExpectGet(key).SetVal(string(dataStr))
					keypg := "gppgCar:{\"id\":null,\"car_name\":\"Toyota\",\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null,\"order_by\":null,\"limit\":0,\"page\":0}"
					pagination := model.Pagination{
						CurrentPage:     1,
						CurrentElements: int64(len(data)),
						TotalElements:   3,
						TotalPages:      1,
						SortBy:          a.param.OrderBy.String,
					}
					paginationStr, _ := json.Marshal(&pagination)
					redisMock.ExpectGet(keypg).SetVal(string(paginationStr))
				},
			},
			{
				testType: "P",
				testDesc: "test get with cache control not must revalidate redis nil",
				args: struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}{
					ctx: &ctx,
					param: &model.GetCarsByParam{
						GetCarByParam: model.GetCarByParam{
							CarName: null.StringFrom("Toyota"),
						},
					},
				},
				want: struct {
					result     []psqlmodel.Car
					pagination model.Pagination
				}{
					result: []psqlmodel.Car{
						{
							ID:        1,
							CarName:   "Toyota Calya",
							DayRate:   12000,
							MonthRate: 300000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        2,
							CarName:   "Toyota Rush",
							DayRate:   400000,
							MonthRate: 1200000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        3,
							CarName:   "Toyota Supra",
							DayRate:   1000000,
							MonthRate: 30000000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
					},
					pagination: model.Pagination{
						CurrentPage:     1,
						CurrentElements: 3,
						TotalElements:   3,
						TotalPages:      1,
						SortBy:          "",
					},
				},
				wantErr: "",
				mockFunc: func(a struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}) {
					data := []psqlmodel.Car{
						{
							ID:        1,
							CarName:   "Toyota Calya",
							DayRate:   12000,
							MonthRate: 300000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        2,
							CarName:   "Toyota Rush",
							DayRate:   400000,
							MonthRate: 1200000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        3,
							CarName:   "Toyota Supra",
							DayRate:   1000000,
							MonthRate: 30000000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
					}
					key := "gpCar:{\"id\":null,\"car_name\":\"Toyota\",\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null,\"order_by\":null,\"limit\":0,\"page\":0}"
					redisMock.ExpectGet(key).RedisNil()
					keypg := "gppgCar:{\"id\":null,\"car_name\":\"Toyota\",\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null,\"order_by\":null,\"limit\":0,\"page\":0}"
					pagination := model.Pagination{
						CurrentPage:     1,
						CurrentElements: int64(len(data)),
						TotalElements:   3,
						TotalPages:      1,
						SortBy:          a.param.OrderBy.String,
					}
					paginationStr, _ := json.Marshal(&pagination)
					redisMock.ExpectGet(keypg).SetVal(string(paginationStr))
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM \"cars\" WHERE (car_name like $1) AND (\"cars\".\"deleted_at\" is null);")).WithArgs("%" + a.param.CarName.String + "%").WillReturnRows(sqlMock.NewRows([]string{"COUNT(*)"}).AddRow(3)).RowsWillBeClosed()
					rows := sqlMock.NewRows([]string{"id", "car_name", "day_rate", "month_rate", "image", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at"})
					for _, d := range data {
						rows.AddRow(d.ID, d.CarName, d.DayRate, d.MonthRate, d.Image, d.CreatedBy, d.CreatedAt, d.UpdatedBy, d.UpdatedAt, d.DeletedBy, d.DeletedAt)
					}
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT \"cars\".* FROM \"cars\" WHERE (car_name like $1) AND (\"cars\".\"deleted_at\" is null) LIMIT 10;")).WithArgs("%" + a.param.CarName.String + "%").WillReturnRows(rows).RowsWillBeClosed()
					dataStr, _ := json.Marshal(&data)
					redisMock.ExpectDel(key).SetVal(0)
					redisMock.ExpectSet(key, string(dataStr), 30*time.Second).SetVal(string(dataStr))
					redisMock.ExpectDel(keypg).SetVal(0)
					redisMock.ExpectSet(keypg, string(paginationStr), 30*time.Second).SetVal(string(paginationStr))
				},
			},
			{
				testType: "P",
				testDesc: "test get with cache control not must revalidate redis nil",
				args: struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}{
					ctx: &ctx,
					param: &model.GetCarsByParam{
						GetCarByParam: model.GetCarByParam{
							CarName: null.StringFrom("Toyota"),
						},
					},
				},
				want: struct {
					result     []psqlmodel.Car
					pagination model.Pagination
				}{
					result: []psqlmodel.Car{
						{
							ID:        1,
							CarName:   "Toyota Calya",
							DayRate:   12000,
							MonthRate: 300000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        2,
							CarName:   "Toyota Rush",
							DayRate:   400000,
							MonthRate: 1200000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        3,
							CarName:   "Toyota Supra",
							DayRate:   1000000,
							MonthRate: 30000000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
					},
					pagination: model.Pagination{
						CurrentPage:     1,
						CurrentElements: 3,
						TotalElements:   3,
						TotalPages:      1,
						SortBy:          "",
					},
				},
				wantErr: "",
				mockFunc: func(a struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}) {
					data := []psqlmodel.Car{
						{
							ID:        1,
							CarName:   "Toyota Calya",
							DayRate:   12000,
							MonthRate: 300000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        2,
							CarName:   "Toyota Rush",
							DayRate:   400000,
							MonthRate: 1200000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        3,
							CarName:   "Toyota Supra",
							DayRate:   1000000,
							MonthRate: 30000000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
					}
					key := "gpCar:{\"id\":null,\"car_name\":\"Toyota\",\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null,\"order_by\":null,\"limit\":0,\"page\":0}"
					dataStr, _ := json.Marshal(&data)
					redisMock.ExpectGet(key).SetVal(string(dataStr))
					keypg := "gppgCar:{\"id\":null,\"car_name\":\"Toyota\",\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null,\"order_by\":null,\"limit\":0,\"page\":0}"
					pagination := model.Pagination{
						CurrentPage:     1,
						CurrentElements: int64(len(data)),
						TotalElements:   3,
						TotalPages:      1,
						SortBy:          a.param.OrderBy.String,
					}
					paginationStr, _ := json.Marshal(&pagination)
					redisMock.ExpectGet(keypg).RedisNil()
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM \"cars\" WHERE (car_name like $1) AND (\"cars\".\"deleted_at\" is null);")).WithArgs("%" + a.param.CarName.String + "%").WillReturnRows(sqlMock.NewRows([]string{"COUNT(*)"}).AddRow(3)).RowsWillBeClosed()
					rows := sqlMock.NewRows([]string{"id", "car_name", "day_rate", "month_rate", "image", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at"})
					for _, d := range data {
						rows.AddRow(d.ID, d.CarName, d.DayRate, d.MonthRate, d.Image, d.CreatedBy, d.CreatedAt, d.UpdatedBy, d.UpdatedAt, d.DeletedBy, d.DeletedAt)
					}
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT \"cars\".* FROM \"cars\" WHERE (car_name like $1) AND (\"cars\".\"deleted_at\" is null) LIMIT 10;")).WithArgs("%" + a.param.CarName.String + "%").WillReturnRows(rows).RowsWillBeClosed()
					redisMock.ExpectDel(key).SetVal(0)
					redisMock.ExpectSet(key, string(dataStr), 30*time.Second).SetVal(string(dataStr))
					redisMock.ExpectDel(keypg).SetVal(0)
					redisMock.ExpectSet(keypg, string(paginationStr), 30*time.Second).SetVal(string(paginationStr))
				},
			},
			{
				testType: "N",
				testDesc: "test get with cache control not must revalidate redis error",
				args: struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}{
					ctx: &ctx,
					param: &model.GetCarsByParam{
						GetCarByParam: model.GetCarByParam{
							CarName: null.StringFrom("Toyota"),
						},
					},
				},
				wantErr: "redis error",
				mockFunc: func(a struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}) {
					data := []psqlmodel.Car{
						{
							ID:        1,
							CarName:   "Toyota Calya",
							DayRate:   12000,
							MonthRate: 300000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        2,
							CarName:   "Toyota Rush",
							DayRate:   400000,
							MonthRate: 1200000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        3,
							CarName:   "Toyota Supra",
							DayRate:   1000000,
							MonthRate: 30000000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
					}
					key := "gpCar:{\"id\":null,\"car_name\":\"Toyota\",\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null,\"order_by\":null,\"limit\":0,\"page\":0}"
					redisMock.ExpectGet(key).SetErr(errors.New("redis error"))
					keypg := "gppgCar:{\"id\":null,\"car_name\":\"Toyota\",\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null,\"order_by\":null,\"limit\":0,\"page\":0}"
					pagination := model.Pagination{
						CurrentPage:     1,
						CurrentElements: int64(len(data)),
						TotalElements:   3,
						TotalPages:      1,
						SortBy:          a.param.OrderBy.String,
					}
					paginationStr, _ := json.Marshal(&pagination)
					redisMock.ExpectGet(keypg).SetVal(string(paginationStr))
				},
			},
			{
				testType: "N",
				testDesc: "test get with cache control not must revalidate redis error",
				args: struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}{
					ctx: &ctx,
					param: &model.GetCarsByParam{
						GetCarByParam: model.GetCarByParam{
							CarName: null.StringFrom("Toyota"),
						},
					},
				},
				wantErr: "redis error",
				mockFunc: func(a struct {
					ctx          *context.Context
					cacheControl string
					param        *model.GetCarsByParam
				}) {
					data := []psqlmodel.Car{
						{
							ID:        1,
							CarName:   "Toyota Calya",
							DayRate:   12000,
							MonthRate: 300000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        2,
							CarName:   "Toyota Rush",
							DayRate:   400000,
							MonthRate: 1200000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
						{
							ID:        3,
							CarName:   "Toyota Supra",
							DayRate:   1000000,
							MonthRate: 30000000,
							Image:     "http://link.com",
							CreatedBy: 0,
							CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
							UpdatedBy: 0,
							UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						},
					}
					dataStr, _ := json.Marshal(&data)
					key := "gpCar:{\"id\":null,\"car_name\":\"Toyota\",\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null,\"order_by\":null,\"limit\":0,\"page\":0}"
					redisMock.ExpectGet(key).SetVal(string(dataStr))
					keypg := "gppgCar:{\"id\":null,\"car_name\":\"Toyota\",\"day_rate\":null,\"day_rate_gt\":null,\"day_rate_gte\":null,\"day_rate_lt\":null,\"day_rate_lte\":null,\"month_rate\":null,\"month_rate_gt\":null,\"month_rate_gte\":null,\"month_rate_lt\":null,\"month_rate_lte\":null,\"image\":null,\"order_by\":null,\"limit\":0,\"page\":0}"
					redisMock.ExpectGet(keypg).SetErr(errors.New("redis error"))
				},
			},
		}
		for idx, test := range tests {
			Convey(fmt.Sprintf("%d - [%s] : %s", idx, test.testType, test.testDesc), func() {
				test.mockFunc(test.args)
				car, pagination, err := acc.GetByParam(test.args.ctx, test.args.cacheControl, test.args.param)
				if test.testType == "N" {
					e := errormsg.GetErrorData(err)
					So(e.DebugError.Error(), ShouldEqual, test.wantErr)
				} else {
					So(err, ShouldBeNil)
					for idx, v := range test.want.result {
						So(car[idx].ID, ShouldEqual, v.ID)
						So(car[idx].CarName, ShouldEqual, v.CarName)
						So(car[idx].DayRate, ShouldEqual, v.DayRate)
						So(car[idx].MonthRate, ShouldEqual, v.MonthRate)
						So(car[idx].Image, ShouldEqual, v.Image)
						So(car[idx].CreatedBy, ShouldEqual, v.CreatedBy)
						So(car[idx].CreatedAt, ShouldEqual, v.CreatedAt)
						So(car[idx].UpdatedBy, ShouldEqual, v.UpdatedBy)
						So(car[idx].UpdatedAt, ShouldEqual, v.UpdatedAt)
						So(car[idx].DeletedBy, ShouldEqual, v.DeletedBy)
						So(car[idx].DeletedAt, ShouldEqual, v.DeletedAt)
					}
					So(pagination.CurrentElements, ShouldEqual, test.want.pagination.CurrentElements)
					So(pagination.CurrentPage, ShouldEqual, test.want.pagination.CurrentPage)
					So(pagination.SortBy, ShouldEqual, test.want.pagination.SortBy)
					So(pagination.TotalElements, ShouldEqual, test.want.pagination.TotalElements)
					So(pagination.TotalPages, ShouldEqual, test.want.pagination.TotalPages)
				}
			})
		}
	})
}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestInsert(t *testing.T) {
	dbSQL, sqlMock, err := gosqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	oldDB := boil.GetDB()
	defer func() {
		dbSQL.Close()
		boil.SetDB(oldDB)
	}()
	boil.SetDB(dbSQL)
	dbRedis, _ := redismock.NewClientMock()
	acc := car.CarDep{
		Log:   logger.New(&logger.Config{}),
		DB:    dbSQL,
		Redis: dbRedis,
		Conf: car.Conf{
			DefaultPageLimit:    10,
			RedisExpirationTime: 30 * time.Second,
		},
	}
	ctx := context.Background()
	Convey("test insert car", t, FailureHalts, func() {
		tests := []struct {
			testType string
			testDesc string
			args     struct {
				ctx   *context.Context
				param *psqlmodel.Car
			}
			want struct {
				result psqlmodel.Car
			}
			wantErr  string
			mockFunc func(a struct {
				ctx   *context.Context
				param *psqlmodel.Car
			})
		}{
			{
				testType: "P",
				testDesc: "test create car",
				args: struct {
					ctx   *context.Context
					param *psqlmodel.Car
				}{
					ctx: &ctx,
					param: &psqlmodel.Car{
						CarName:   "Sedan",
						DayRate:   10000,
						MonthRate: 300000,
						Image:     "http://image.com",
						CreatedBy: 1,
					},
				},
				want: struct{ result psqlmodel.Car }{
					result: psqlmodel.Car{
						ID:        1,
						CarName:   "sedan",
						DayRate:   12000,
						MonthRate: 300000,
						Image:     "http://link.com",
						CreatedBy: 0,
						CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						UpdatedBy: 0,
						UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
					},
				},
				wantErr: "",
				mockFunc: func(a struct {
					ctx   *context.Context
					param *psqlmodel.Car
				}) {
					sqlMock.ExpectBegin()
					rows := sqlMock.NewRows([]string{"id", "updated_by", "deleted_by", "deleted_at"}).AddRow(1, 1, nil, nil)
					sqlMock.ExpectQuery(regexp.QuoteMeta("INSERT INTO \"cars\" (\"car_name\",\"day_rate\",\"month_rate\",\"image\",\"created_by\",\"created_at\",\"updated_at\") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING \"id\",\"updated_by\",\"deleted_by\",\"deleted_at\"")).WithArgs(a.param.CarName, a.param.DayRate, a.param.MonthRate, a.param.Image, a.param.CreatedBy, AnyTime{}, AnyTime{}).WillReturnRows(rows)
					sqlMock.ExpectCommit()
				},
			},
			{
				testType: "N",
				testDesc: "test create car error commit",
				args: struct {
					ctx   *context.Context
					param *psqlmodel.Car
				}{
					ctx: &ctx,
					param: &psqlmodel.Car{
						CarName:   "Sedan",
						DayRate:   10000,
						MonthRate: 300000,
						Image:     "http://image.com",
						CreatedBy: 1,
					},
				},
				wantErr: "commit error",
				mockFunc: func(a struct {
					ctx   *context.Context
					param *psqlmodel.Car
				}) {
					sqlMock.ExpectBegin()
					rows := sqlMock.NewRows([]string{"id", "updated_by", "deleted_by", "deleted_at"}).AddRow(1, 1, nil, nil)
					sqlMock.ExpectQuery(regexp.QuoteMeta("INSERT INTO \"cars\" (\"car_name\",\"day_rate\",\"month_rate\",\"image\",\"created_by\",\"created_at\",\"updated_at\") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING \"id\",\"updated_by\",\"deleted_by\",\"deleted_at\"")).WithArgs(a.param.CarName, a.param.DayRate, a.param.MonthRate, a.param.Image, a.param.CreatedBy, AnyTime{}, AnyTime{}).WillReturnRows(rows)
					sqlMock.ExpectCommit().WillReturnError(errors.New("commit error"))
				},
			},
			{
				testType: "N",
				testDesc: "test create car error",
				args: struct {
					ctx   *context.Context
					param *psqlmodel.Car
				}{
					ctx: &ctx,
					param: &psqlmodel.Car{
						CarName:   "Sedan",
						DayRate:   10000,
						MonthRate: 300000,
						Image:     "http://image.com",
						CreatedBy: 1,
					},
				},
				wantErr: "psqlmodel: unable to insert into cars: error insert data",
				mockFunc: func(a struct {
					ctx   *context.Context
					param *psqlmodel.Car
				}) {
					sqlMock.ExpectBegin()
					sqlMock.ExpectQuery(regexp.QuoteMeta("INSERT INTO \"cars\" (\"car_name\",\"day_rate\",\"month_rate\",\"image\",\"created_by\",\"created_at\",\"updated_at\") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING \"id\",\"updated_by\",\"deleted_by\",\"deleted_at\"")).WithArgs(a.param.CarName, a.param.DayRate, a.param.MonthRate, a.param.Image, a.param.CreatedBy, AnyTime{}, AnyTime{}).WillReturnError(errors.New("error insert data"))
					sqlMock.ExpectRollback()
				},
			},
			{
				testType: "N",
				testDesc: "test create car error transaction",
				args: struct {
					ctx   *context.Context
					param *psqlmodel.Car
				}{
					ctx: &ctx,
					param: &psqlmodel.Car{
						CarName:   "Sedan",
						DayRate:   10000,
						MonthRate: 300000,
						Image:     "http://image.com",
						CreatedBy: 1,
					},
				},
				wantErr: "error transaction begin",
				mockFunc: func(a struct {
					ctx   *context.Context
					param *psqlmodel.Car
				}) {
					sqlMock.ExpectBegin().WillReturnError(errors.New("error transaction begin"))
				},
			},
		}
		for idx, test := range tests {
			Convey(fmt.Sprintf("%d - [%s] : %s", idx, test.testType, test.testDesc), func() {
				test.mockFunc(test.args)
				err := acc.Insert(test.args.ctx, test.args.param)
				if test.testType == "N" {
					e := errormsg.GetErrorData(err)
					So(e.DebugError.Error(), ShouldEqual, test.wantErr)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestUpdate(t *testing.T) {
	dbSQL, sqlMock, err := gosqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	oldDB := boil.GetDB()
	defer func() {
		dbSQL.Close()
		boil.SetDB(oldDB)
	}()
	boil.SetDB(dbSQL)
	dbRedis, _ := redismock.NewClientMock()
	acc := car.CarDep{
		Log:   logger.New(&logger.Config{}),
		DB:    dbSQL,
		Redis: dbRedis,
		Conf: car.Conf{
			DefaultPageLimit:    10,
			RedisExpirationTime: 30 * time.Second,
		},
	}
	ctx := context.Background()
	Convey("test update car", t, FailureHalts, func() {
		tests := []struct {
			testType string
			testDesc string
			args     struct {
				ctx   *context.Context
				param *psqlmodel.Car
			}
			want struct {
				result psqlmodel.Car
			}
			wantErr  string
			mockFunc func(a struct {
				ctx   *context.Context
				param *psqlmodel.Car
			})
		}{
			{
				testType: "P",
				testDesc: "test update car",
				args: struct {
					ctx   *context.Context
					param *psqlmodel.Car
				}{
					ctx: &ctx,
					param: &psqlmodel.Car{
						ID:        1,
						CarName:   "Sedan",
						DayRate:   10000,
						MonthRate: 300000,
						Image:     "http://image.com",
						CreatedBy: 1,
						UpdatedBy: 1,
					},
				},
				want: struct{ result psqlmodel.Car }{
					result: psqlmodel.Car{
						ID:        1,
						CarName:   "sedan",
						DayRate:   12000,
						MonthRate: 300000,
						Image:     "http://link.com",
						CreatedBy: 0,
						CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						UpdatedBy: 0,
						UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
					},
				},
				wantErr: "",
				mockFunc: func(a struct {
					ctx   *context.Context
					param *psqlmodel.Car
				}) {
					sqlMock.ExpectBegin()
					sqlMock.ExpectExec(regexp.QuoteMeta("UPDATE \"cars\" SET \"car_name\"=$1,\"day_rate\"=$2,\"month_rate\"=$3,\"image\"=$4,\"created_by\"=$5,\"updated_by\"=$6,\"updated_at\"=$7,\"deleted_by\"=$8,\"deleted_at\"=$9 WHERE \"id\"=$10")).WithArgs(a.param.CarName, a.param.DayRate, a.param.MonthRate, a.param.Image, a.param.CreatedBy, a.param.UpdatedBy, AnyTime{}, nil, nil, 1).WillReturnResult(gosqlmock.NewResult(1, 1))
					sqlMock.ExpectCommit()
				},
			},
			{
				testType: "N",
				testDesc: "test update car error commit",
				args: struct {
					ctx   *context.Context
					param *psqlmodel.Car
				}{
					ctx: &ctx,
					param: &psqlmodel.Car{
						ID:        1,
						CarName:   "Sedan",
						DayRate:   10000,
						MonthRate: 300000,
						Image:     "http://image.com",
						CreatedBy: 1,
						UpdatedBy: 1,
					},
				},
				wantErr: "commit error",
				mockFunc: func(a struct {
					ctx   *context.Context
					param *psqlmodel.Car
				}) {
					sqlMock.ExpectBegin()
					sqlMock.ExpectExec(regexp.QuoteMeta("UPDATE \"cars\" SET \"car_name\"=$1,\"day_rate\"=$2,\"month_rate\"=$3,\"image\"=$4,\"created_by\"=$5,\"updated_by\"=$6,\"updated_at\"=$7,\"deleted_by\"=$8,\"deleted_at\"=$9 WHERE \"id\"=$10")).WithArgs(a.param.CarName, a.param.DayRate, a.param.MonthRate, a.param.Image, a.param.CreatedBy, a.param.UpdatedBy, AnyTime{}, nil, nil, 1).WillReturnResult(gosqlmock.NewResult(1, 1))
					sqlMock.ExpectCommit().WillReturnError(errors.New("commit error"))
				},
			},
			{
				testType: "N",
				testDesc: "test update car error",
				args: struct {
					ctx   *context.Context
					param *psqlmodel.Car
				}{
					ctx: &ctx,
					param: &psqlmodel.Car{
						ID:        1,
						CarName:   "Sedan",
						DayRate:   10000,
						MonthRate: 300000,
						Image:     "http://image.com",
						CreatedBy: 1,
						UpdatedBy: 1,
					},
				},
				wantErr: "psqlmodel: unable to update cars row: error update data",
				mockFunc: func(a struct {
					ctx   *context.Context
					param *psqlmodel.Car
				}) {
					sqlMock.ExpectBegin()
					sqlMock.ExpectExec(regexp.QuoteMeta("UPDATE \"cars\" SET \"car_name\"=$1,\"day_rate\"=$2,\"month_rate\"=$3,\"image\"=$4,\"created_by\"=$5,\"updated_by\"=$6,\"updated_at\"=$7,\"deleted_by\"=$8,\"deleted_at\"=$9 WHERE \"id\"=$10")).WithArgs(a.param.CarName, a.param.DayRate, a.param.MonthRate, a.param.Image, a.param.CreatedBy, a.param.UpdatedBy, AnyTime{}, nil, nil, 1).WillReturnError(errors.New("error update data"))
					sqlMock.ExpectRollback()
				},
			},
			{
				testType: "N",
				testDesc: "test update car error transaction",
				args: struct {
					ctx   *context.Context
					param *psqlmodel.Car
				}{
					ctx: &ctx,
					param: &psqlmodel.Car{
						ID:        1,
						CarName:   "Sedan",
						DayRate:   10000,
						MonthRate: 300000,
						Image:     "http://image.com",
						CreatedBy: 1,
						UpdatedBy: 1,
					},
				},
				wantErr: "error transaction begin",
				mockFunc: func(a struct {
					ctx   *context.Context
					param *psqlmodel.Car
				}) {
					sqlMock.ExpectBegin().WillReturnError(errors.New("error transaction begin"))
				},
			},
		}
		for idx, test := range tests {
			Convey(fmt.Sprintf("%d - [%s] : %s", idx, test.testType, test.testDesc), func() {
				test.mockFunc(test.args)
				err := acc.Update(test.args.ctx, test.args.param)
				if test.testType == "N" {
					e := errormsg.GetErrorData(err)
					So(e.DebugError.Error(), ShouldEqual, test.wantErr)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestDelete(t *testing.T) {
	dbSQL, sqlMock, err := gosqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	oldDB := boil.GetDB()
	defer func() {
		dbSQL.Close()
		boil.SetDB(oldDB)
	}()
	boil.SetDB(dbSQL)
	dbRedis, _ := redismock.NewClientMock()
	acc := car.CarDep{
		Log:   logger.New(&logger.Config{}),
		DB:    dbSQL,
		Redis: dbRedis,
		Conf: car.Conf{
			DefaultPageLimit:    10,
			RedisExpirationTime: 30 * time.Second,
		},
	}
	ctx := context.Background()
	Convey("test delete car", t, FailureHalts, func() {
		tests := []struct {
			testType string
			testDesc string
			args     struct {
				ctx          *context.Context
				param        *psqlmodel.Car
				id           int64
				isHardDelete bool
			}
			want struct {
				result psqlmodel.Car
			}
			wantErr  string
			mockFunc func(a struct {
				ctx          *context.Context
				param        *psqlmodel.Car
				id           int64
				isHardDelete bool
			})
		}{
			{
				testType: "P",
				testDesc: "test delete car",
				args: struct {
					ctx          *context.Context
					param        *psqlmodel.Car
					id           int64
					isHardDelete bool
				}{
					ctx: &ctx,
					param: &psqlmodel.Car{
						ID:        1,
						CarName:   "Sedan",
						DayRate:   10000,
						MonthRate: 300000,
						Image:     "http://image.com",
						CreatedBy: 1,
						UpdatedBy: 1,
					},
					id:           1,
					isHardDelete: false,
				},
				want: struct{ result psqlmodel.Car }{
					result: psqlmodel.Car{
						ID:        1,
						CarName:   "sedan",
						DayRate:   12000,
						MonthRate: 300000,
						Image:     "http://link.com",
						CreatedBy: 0,
						CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						UpdatedBy: 0,
						UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
					},
				},
				wantErr: "",
				mockFunc: func(a struct {
					ctx          *context.Context
					param        *psqlmodel.Car
					id           int64
					isHardDelete bool
				}) {
					sqlMock.ExpectBegin()
					sqlMock.ExpectExec(regexp.QuoteMeta("UPDATE \"cars\" SET \"car_name\"=$1,\"day_rate\"=$2,\"month_rate\"=$3,\"image\"=$4,\"created_by\"=$5,\"updated_by\"=$6,\"updated_at\"=$7,\"deleted_by\"=$8,\"deleted_at\"=$9 WHERE \"id\"=$10")).WithArgs(a.param.CarName, a.param.DayRate, a.param.MonthRate, a.param.Image, a.param.CreatedBy, a.param.UpdatedBy, AnyTime{}, a.id, AnyTime{}, 1).WillReturnResult(gosqlmock.NewResult(1, 1))
					sqlMock.ExpectCommit()
				},
			},
			{
				testType: "N",
				testDesc: "test delete car error commit",
				args: struct {
					ctx          *context.Context
					param        *psqlmodel.Car
					id           int64
					isHardDelete bool
				}{
					ctx: &ctx,
					param: &psqlmodel.Car{
						ID:        1,
						CarName:   "Sedan",
						DayRate:   10000,
						MonthRate: 300000,
						Image:     "http://image.com",
						CreatedBy: 1,
						UpdatedBy: 1,
					},
				},
				wantErr: "commit error",
				mockFunc: func(a struct {
					ctx          *context.Context
					param        *psqlmodel.Car
					id           int64
					isHardDelete bool
				}) {
					sqlMock.ExpectBegin()
					sqlMock.ExpectExec(regexp.QuoteMeta("UPDATE \"cars\" SET \"car_name\"=$1,\"day_rate\"=$2,\"month_rate\"=$3,\"image\"=$4,\"created_by\"=$5,\"updated_by\"=$6,\"updated_at\"=$7,\"deleted_by\"=$8,\"deleted_at\"=$9 WHERE \"id\"=$10")).WithArgs(a.param.CarName, a.param.DayRate, a.param.MonthRate, a.param.Image, a.param.CreatedBy, a.param.UpdatedBy, AnyTime{}, a.id, AnyTime{}, 1).WillReturnResult(gosqlmock.NewResult(1, 1))
					sqlMock.ExpectCommit().WillReturnError(errors.New("commit error"))
				},
			},
			{
				testType: "N",
				testDesc: "test delete car error",
				args: struct {
					ctx          *context.Context
					param        *psqlmodel.Car
					id           int64
					isHardDelete bool
				}{
					ctx: &ctx,
					param: &psqlmodel.Car{
						ID:        1,
						CarName:   "Sedan",
						DayRate:   10000,
						MonthRate: 300000,
						Image:     "http://image.com",
						CreatedBy: 1,
						UpdatedBy: 1,
					},
				},
				wantErr: "psqlmodel: unable to update cars row: error delete data",
				mockFunc: func(a struct {
					ctx          *context.Context
					param        *psqlmodel.Car
					id           int64
					isHardDelete bool
				}) {
					sqlMock.ExpectBegin()
					sqlMock.ExpectExec(regexp.QuoteMeta("UPDATE \"cars\" SET \"car_name\"=$1,\"day_rate\"=$2,\"month_rate\"=$3,\"image\"=$4,\"created_by\"=$5,\"updated_by\"=$6,\"updated_at\"=$7,\"deleted_by\"=$8,\"deleted_at\"=$9 WHERE \"id\"=$10")).WithArgs(a.param.CarName, a.param.DayRate, a.param.MonthRate, a.param.Image, a.param.CreatedBy, a.param.UpdatedBy, AnyTime{}, a.id, AnyTime{}, 1).WillReturnError(errors.New("error delete data"))
					sqlMock.ExpectRollback()
				},
			},
			{
				testType: "N",
				testDesc: "test delete car error transaction",
				args: struct {
					ctx          *context.Context
					param        *psqlmodel.Car
					id           int64
					isHardDelete bool
				}{
					ctx: &ctx,
					param: &psqlmodel.Car{
						ID:        1,
						CarName:   "Sedan",
						DayRate:   10000,
						MonthRate: 300000,
						Image:     "http://image.com",
						CreatedBy: 1,
						UpdatedBy: 1,
					},
				},
				wantErr: "error transaction begin",
				mockFunc: func(a struct {
					ctx          *context.Context
					param        *psqlmodel.Car
					id           int64
					isHardDelete bool
				}) {
					sqlMock.ExpectBegin().WillReturnError(errors.New("error transaction begin"))
				},
			},
			{
				testType: "P",
				testDesc: "test delete car hard delete",
				args: struct {
					ctx          *context.Context
					param        *psqlmodel.Car
					id           int64
					isHardDelete bool
				}{
					ctx: &ctx,
					param: &psqlmodel.Car{
						ID:        1,
						CarName:   "Sedan",
						DayRate:   10000,
						MonthRate: 300000,
						Image:     "http://image.com",
						CreatedBy: 1,
						UpdatedBy: 1,
					},
					id:           1,
					isHardDelete: true,
				},
				want: struct{ result psqlmodel.Car }{
					result: psqlmodel.Car{
						ID:        1,
						CarName:   "sedan",
						DayRate:   12000,
						MonthRate: 300000,
						Image:     "http://link.com",
						CreatedBy: 0,
						CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						UpdatedBy: 0,
						UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
					},
				},
				wantErr: "",
				mockFunc: func(a struct {
					ctx          *context.Context
					param        *psqlmodel.Car
					id           int64
					isHardDelete bool
				}) {
					sqlMock.ExpectBegin()
					sqlMock.ExpectExec(regexp.QuoteMeta("DELETE FROM \"cars\" WHERE \"id\"=$1")).WithArgs(a.param.ID).WillReturnResult(gosqlmock.NewResult(1, 1))
					sqlMock.ExpectCommit()
				},
			},
			{
				testType: "N",
				testDesc: "test delete car hard delete error begin tx",
				args: struct {
					ctx          *context.Context
					param        *psqlmodel.Car
					id           int64
					isHardDelete bool
				}{
					ctx: &ctx,
					param: &psqlmodel.Car{
						ID:        1,
						CarName:   "Sedan",
						DayRate:   10000,
						MonthRate: 300000,
						Image:     "http://image.com",
						CreatedBy: 1,
						UpdatedBy: 1,
					},
					id:           1,
					isHardDelete: true,
				},
				wantErr: "error begin tx",
				mockFunc: func(a struct {
					ctx          *context.Context
					param        *psqlmodel.Car
					id           int64
					isHardDelete bool
				}) {
					sqlMock.ExpectBegin().WillReturnError(errors.New("error begin tx"))
				},
			},
			{
				testType: "N",
				testDesc: "test delete car hard delete error delete",
				args: struct {
					ctx          *context.Context
					param        *psqlmodel.Car
					id           int64
					isHardDelete bool
				}{
					ctx: &ctx,
					param: &psqlmodel.Car{
						ID:        1,
						CarName:   "Sedan",
						DayRate:   10000,
						MonthRate: 300000,
						Image:     "http://image.com",
						CreatedBy: 1,
						UpdatedBy: 1,
					},
					id:           1,
					isHardDelete: true,
				},
				want: struct{ result psqlmodel.Car }{
					result: psqlmodel.Car{
						ID:        1,
						CarName:   "sedan",
						DayRate:   12000,
						MonthRate: 300000,
						Image:     "http://link.com",
						CreatedBy: 0,
						CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						UpdatedBy: 0,
						UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
					},
				},
				wantErr: "psqlmodel: unable to delete from cars: error delete data",
				mockFunc: func(a struct {
					ctx          *context.Context
					param        *psqlmodel.Car
					id           int64
					isHardDelete bool
				}) {
					sqlMock.ExpectBegin()
					sqlMock.ExpectExec(regexp.QuoteMeta("DELETE FROM \"cars\" WHERE \"id\"=$1")).WithArgs(a.param.ID).WillReturnError(errors.New("error delete data"))
				},
			},
			{
				testType: "N",
				testDesc: "test delete car hard delete",
				args: struct {
					ctx          *context.Context
					param        *psqlmodel.Car
					id           int64
					isHardDelete bool
				}{
					ctx: &ctx,
					param: &psqlmodel.Car{
						ID:        1,
						CarName:   "Sedan",
						DayRate:   10000,
						MonthRate: 300000,
						Image:     "http://image.com",
						CreatedBy: 1,
						UpdatedBy: 1,
					},
					id:           1,
					isHardDelete: true,
				},
				wantErr: "error commit",
				mockFunc: func(a struct {
					ctx          *context.Context
					param        *psqlmodel.Car
					id           int64
					isHardDelete bool
				}) {
					sqlMock.ExpectBegin()
					sqlMock.ExpectExec(regexp.QuoteMeta("DELETE FROM \"cars\" WHERE \"id\"=$1")).WithArgs(a.param.ID).WillReturnResult(gosqlmock.NewResult(1, 1))
					sqlMock.ExpectCommit().WillReturnError(errors.New("error commit"))
				},
			},
		}
		for idx, test := range tests {
			Convey(fmt.Sprintf("%d - [%s] : %s", idx, test.testType, test.testDesc), func() {
				test.mockFunc(test.args)
				err := acc.Delete(test.args.ctx, test.args.param, test.args.id, test.args.isHardDelete)
				if test.testType == "N" {
					e := errormsg.GetErrorData(err)
					So(e.DebugError.Error(), ShouldEqual, test.wantErr)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}
