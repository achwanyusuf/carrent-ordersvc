package svcerr

import (
	"net/http"

	"github.com/achwanyusuf/carrent-lib/pkg/errormsg"
)

const (
	CodeBadRequest = iota + 40000
	CodePSQLErrorTransaction
	CodePSQLErrorCommit
	CodePSQLErrorRollback
	CodePSQLErrorInsert
	CodePSQLErrorUpdate
	CodePSQLErrorDelete
	CodePSQLErrorGet
	CodeErrorGRPCClient
	CodeInvalidCarName
	CodeInvalidDayRate
	CodeInvalidMonthRate
	CodeInvalidImage
	CodeInvalidCarID
	CodeInvalidOrderDate
	CodeInvalidPickupDate
	CodeInvalidDropoffDate
	CodeInvalidPickupLocation
	CodeInvalidPickupLat
	CodeInvalidPickupLong
	CodeInvalidDropoffLocation
	CodeInvalidDropoffLat
	CodeInvalidDropoffLong

	CodeNotAuthorized = 401000
	CodeNotFound      = 404000
)

var (
	OrderSVCPSQLErrorTransaction = ErrMsg[CodePSQLErrorTransaction]
	OrderSVCPSQLErrorCommit      = ErrMsg[CodePSQLErrorCommit]
	OrderSVCPSQLErrorRollback    = ErrMsg[CodePSQLErrorRollback]
	OrderSVCPSQLErrorInsert      = ErrMsg[CodePSQLErrorInsert]
	OrderSVCPSQLErrorUpdate      = ErrMsg[CodePSQLErrorUpdate]
	OrderSVCPSQLErrorDelete      = ErrMsg[CodePSQLErrorDelete]
	OrderSVCPSQLErrorGet         = ErrMsg[CodePSQLErrorGet]
	OrderSVCNotAuthorized        = ErrMsg[CodeNotAuthorized]
	OrderSVCNotFound             = ErrMsg[CodeNotFound]
	OrderSVCBadRequest           = ErrMsg[CodeBadRequest]
	OrderSVCErrorGRPCClient      = ErrMsg[CodeErrorGRPCClient]

	OrderSVCCodeInvalidCarName         = ErrMsg[CodeInvalidCarName]
	OrderSVCCodeInvalidDayRate         = ErrMsg[CodeInvalidDayRate]
	OrderSVCCodeInvalidMonthRate       = ErrMsg[CodeInvalidMonthRate]
	OrderSVCCodeInvalidImage           = ErrMsg[CodeInvalidImage]
	OrderSVCCodeInvalidCarID           = ErrMsg[CodeInvalidCarID]
	OrderSVCCodeInvalidOrderDate       = ErrMsg[CodeInvalidOrderDate]
	OrderSVCCodeInvalidPickupDate      = ErrMsg[CodeInvalidPickupDate]
	OrderSVCCodeInvalidDropoffDate     = ErrMsg[CodeInvalidDropoffDate]
	OrderSVCCodeInvalidPickupLocation  = ErrMsg[CodeInvalidPickupLocation]
	OrderSVCCodeInvalidPickupLat       = ErrMsg[CodeInvalidPickupLat]
	OrderSVCCodeInvalidPickupLong      = ErrMsg[CodeInvalidPickupLong]
	OrderSVCCodeInvalidDropoffLocation = ErrMsg[CodeInvalidDropoffLocation]
	OrderSVCCodeInvalidDropoffLat      = ErrMsg[CodeInvalidDropoffLat]
	OrderSVCCodeInvalidDropoffLong     = ErrMsg[CodeInvalidDropoffLong]
)

var ErrMsg = map[int]errormsg.Message{
	CodePSQLErrorCommit: {
		Code:       CodePSQLErrorCommit,
		StatusCode: http.StatusBadRequest,
		Message:    "Terdapat kesalahan dalam pembuatan data!",
		Translation: errormsg.Translation{
			EN: "There was an error in creating the data ",
		},
	},
	CodePSQLErrorRollback: {
		Code:       CodePSQLErrorRollback,
		StatusCode: http.StatusBadRequest,
		Message:    "Terdapat kesalahan dalam pembuatan data!",
		Translation: errormsg.Translation{
			EN: "There was an error in creating the data ",
		},
	},
	CodePSQLErrorInsert: {
		Code:       CodePSQLErrorInsert,
		StatusCode: http.StatusBadRequest,
		Message:    "Terdapat kesalahan dalam pembuatan data!",
		Translation: errormsg.Translation{
			EN: "There was an error in creating the data ",
		},
	},
	CodePSQLErrorUpdate: {
		Code:       CodePSQLErrorUpdate,
		StatusCode: http.StatusBadRequest,
		Message:    "Terdapat kesalahan dalam mengubah data!",
		Translation: errormsg.Translation{
			EN: "There was an error in updating the data ",
		},
	},
	CodePSQLErrorDelete: {
		Code:       CodePSQLErrorDelete,
		StatusCode: http.StatusBadRequest,
		Message:    "Terdapat kesalahan dalam menghapus data!",
		Translation: errormsg.Translation{
			EN: "There was an error in deleting the data ",
		},
	},
	CodePSQLErrorGet: {
		Code:       CodePSQLErrorGet,
		StatusCode: http.StatusBadRequest,
		Message:    "Terdapat kesalahan dalam pengambilan data!",
		Translation: errormsg.Translation{
			EN: "There was an error in get data!",
		},
	},
	CodeNotFound: {
		Code:       CodeNotFound,
		StatusCode: http.StatusNotFound,
		Message:    "Data tidak ditemukan!",
		Translation: errormsg.Translation{
			EN: "Data not found!",
		},
	},
	CodeNotAuthorized: {
		Code:       CodeNotAuthorized,
		StatusCode: http.StatusUnauthorized,
		Message:    "Akses tidak diijinkan! Silakan login kembali!",
		Translation: errormsg.Translation{
			EN: "Access not authorized! Please login again!",
		},
	},
	CodeBadRequest: {
		Code:       CodeBadRequest,
		StatusCode: http.StatusBadRequest,
		Message:    "Kesalahan input. Silakan cek kembali masukan anda!",
		Translation: errormsg.Translation{
			EN: "Invalid input. Please validate your input!",
		},
	},
	CodePSQLErrorTransaction: {
		Code:       CodePSQLErrorTransaction,
		StatusCode: http.StatusBadRequest,
		Message:    "Terdapat kesalahan dalam pembuatan data!",
		Translation: errormsg.Translation{
			EN: "There was an error in creating the data ",
		},
	},
	CodeErrorGRPCClient: {
		Code:       CodeErrorGRPCClient,
		StatusCode: http.StatusBadRequest,
		Message:    "Terdapat kesalahan dalam proses pengiriman data!",
		Translation: errormsg.Translation{
			EN: "There was an error in data delivering",
		},
	},
	CodeInvalidCarName: {
		Code:       CodeInvalidCarName,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom nama mobil harus diisi!",
		Translation: errormsg.Translation{
			EN: "Car's name field should not be empty!",
		},
	},
	CodeInvalidDayRate: {
		Code:       CodeInvalidDayRate,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom biaya harian harus diisi!",
		Translation: errormsg.Translation{
			EN: "Day rate field should not be empty!",
		},
	},
	CodeInvalidMonthRate: {
		Code:       CodeInvalidMonthRate,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom biaya bulanan harus diisi!",
		Translation: errormsg.Translation{
			EN: "Month rate field should not be empty!",
		},
	},
	CodeInvalidImage: {
		Code:       CodeInvalidImage,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom gambar harus diisi!",
		Translation: errormsg.Translation{
			EN: "Image field should not be empty!",
		},
	},
	CodeInvalidCarID: {
		Code:       CodeInvalidCarID,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom mobil harus diisi!",
		Translation: errormsg.Translation{
			EN: "Car field should not be empty!",
		},
	},
	CodeInvalidOrderDate: {
		Code:       CodeInvalidOrderDate,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom tanggal pemesanan harus diisi!",
		Translation: errormsg.Translation{
			EN: "Order date field should not be empty!",
		},
	},
	CodeInvalidPickupDate: {
		Code:       CodeInvalidPickupDate,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom tanggal penjemputan harus diisi!",
		Translation: errormsg.Translation{
			EN: "Pickup date field should not be empty!",
		},
	},
	CodeInvalidDropoffDate: {
		Code:       CodeInvalidDropoffDate,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom tanggal pengantaran harus diisi!",
		Translation: errormsg.Translation{
			EN: "Dropoff date field should not be empty!",
		},
	},
	CodeInvalidPickupLocation: {
		Code:       CodeInvalidPickupLocation,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom lokasi penjemputan harus diisi!",
		Translation: errormsg.Translation{
			EN: "Pickup location field should not be empty!",
		},
	},
	CodeInvalidPickupLat: {
		Code:       CodeInvalidPickupLat,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom garis lintang penjemputan harus diisi!",
		Translation: errormsg.Translation{
			EN: "Pickup latitude field should not be empty!",
		},
	},
	CodeInvalidPickupLong: {
		Code:       CodeInvalidPickupLong,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom garis bujur penjemputan harus diisi!",
		Translation: errormsg.Translation{
			EN: "Pickup longitude field should not be empty!",
		},
	},
	CodeInvalidDropoffLocation: {
		Code:       CodeInvalidDropoffLocation,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom lokasi pengantaran harus diisi!",
		Translation: errormsg.Translation{
			EN: "Dropoff location should not be empty!",
		},
	},
	CodeInvalidDropoffLat: {
		Code:       CodeInvalidDropoffLat,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom garis lintang pengantaran harus diisi!",
		Translation: errormsg.Translation{
			EN: "Dropoff latitude field should not be empty!",
		},
	},
	CodeInvalidDropoffLong: {
		Code:       CodeInvalidDropoffLong,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom garis bujur pengantaran harus diisi!",
		Translation: errormsg.Translation{
			EN: "Dropoff longitude field should not be empty!",
		},
	},
}
