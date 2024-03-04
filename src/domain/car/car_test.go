package car_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

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
						DayRate:   1.2,
						MonthRate: 7.1,
						Image:     "http://link",
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
						DayRate:   1.2,
						MonthRate: 7.1,
						Image:     "http://link",
						CreatedBy: 0,
						CreatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
						UpdatedBy: 0,
						UpdatedAt: time.Date(2022, 2, 22, 2, 0, 0, 0, time.UTC),
					}
					rows := sqlMock.NewRows([]string{"id", "car_name", "day_rate", "month_rate", "image", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at"}).AddRow(data.ID, data.CarName, data.DayRate, data.MonthRate, data.Image, data.CreatedBy, data.CreatedAt, data.UpdatedBy, data.UpdatedAt, data.DeletedBy, data.DeletedAt)
					sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT \"cars\".* FROM \"cars\" WHERE (id=$1) AND (\"cars\".\"deleted_at\" is null) LIMIT 1;")).WithArgs(a.param.ID.Int64).WillReturnRows(rows).RowsWillBeClosed()
				},
			},
		}
		for idx, test := range tests {
			Convey(fmt.Sprintf("%d - [%s] : %s", idx, test.testType, test.testDesc), func() {
				test.mockFunc(test.args)
				car, err := acc.GetSingleByParam(test.args.ctx, test.args.cacheControl, test.args.param)
				fmt.Println(car.ID, car.CreatedAt, err)
				if test.testType == "N" {
					So(err.Error(), ShouldContainSubstring, test.wantErr)
				} else {
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
