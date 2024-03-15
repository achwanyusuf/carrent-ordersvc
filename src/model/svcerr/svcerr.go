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
	CodeInvalidCarNameMin
	CodeInvalidCarNameMax
	CodeInvalidCarNameAlphanum
	CodeInvalidDayRateMin
	CodeInvalidDayRateMax
	CodeInvalidMonthRateMin
	CodeInvalidMonthRateMax
	CodeInvalidImageMin
	CodeInvalidImageMax
	CodeInvalidImageIsURL
	CodeInvalidCarID
	CodeInvalidOrderDate
	CodeInvalidPickupDate
	CodeInvalidDropoffDate
	CodeInvalidPickupLocationMin
	CodeInvalidPickupLocationMax
	CodeInvalidPickupLat
	CodeInvalidPickupLong
	CodeInvalidDropoffLocationMin
	CodeInvalidDropoffLocationMax
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

	OrderSVCCodeInvalidCarNameMin         = ErrMsg[CodeInvalidCarNameMin]
	OrderSVCCodeInvalidCarNameMax         = ErrMsg[CodeInvalidCarNameMax]
	OrderSVCCodeInvalidCarNameAlphanum    = ErrMsg[CodeInvalidCarNameAlphanum]
	OrderSVCCodeInvalidDayRateMin         = ErrMsg[CodeInvalidDayRateMin]
	OrderSVCCodeInvalidDayRateMax         = ErrMsg[CodeInvalidDayRateMax]
	OrderSVCCodeInvalidMonthRateMin       = ErrMsg[CodeInvalidMonthRateMin]
	OrderSVCCodeInvalidMonthRateMax       = ErrMsg[CodeInvalidMonthRateMax]
	OrderSVCCodeInvalidImageMin           = ErrMsg[CodeInvalidImageMin]
	OrderSVCCodeInvalidImageMax           = ErrMsg[CodeInvalidImageMax]
	OrderSVCCodeInvalidImageIsURL         = ErrMsg[CodeInvalidImageIsURL]
	OrderSVCCodeInvalidCarID              = ErrMsg[CodeInvalidCarID]
	OrderSVCCodeInvalidOrderDate          = ErrMsg[CodeInvalidOrderDate]
	OrderSVCCodeInvalidPickupDate         = ErrMsg[CodeInvalidPickupDate]
	OrderSVCCodeInvalidDropoffDate        = ErrMsg[CodeInvalidDropoffDate]
	OrderSVCCodeInvalidPickupLocationMin  = ErrMsg[CodeInvalidPickupLocationMin]
	OrderSVCCodeInvalidPickupLocationMax  = ErrMsg[CodeInvalidPickupLocationMax]
	OrderSVCCodeInvalidPickupLat          = ErrMsg[CodeInvalidPickupLat]
	OrderSVCCodeInvalidPickupLong         = ErrMsg[CodeInvalidPickupLong]
	OrderSVCCodeInvalidDropoffLocationMin = ErrMsg[CodeInvalidDropoffLocationMin]
	OrderSVCCodeInvalidDropoffLocationMax = ErrMsg[CodeInvalidDropoffLocationMax]
	OrderSVCCodeInvalidDropoffLat         = ErrMsg[CodeInvalidDropoffLat]
	OrderSVCCodeInvalidDropoffLong        = ErrMsg[CodeInvalidDropoffLong]
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
	CodeInvalidCarNameMin: {
		Code:       CodeInvalidCarNameMin,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom nama mobil minimum 8 karakter!",
		Translation: errormsg.Translation{
			EN: "Car's name field minimum 8 character!",
		},
	},
	CodeInvalidCarNameMax: {
		Code:       CodeInvalidCarNameMax,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom nama mobil maksimum 50 karakter!",
		Translation: errormsg.Translation{
			EN: "Car's name field maximum 50 character!",
		},
	},
	CodeInvalidCarNameAlphanum: {
		Code:       CodeInvalidCarNameAlphanum,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom nama mobil hanya boleh huruf dan angka!",
		Translation: errormsg.Translation{
			EN: "Car's name field should be character or number!",
		},
	},
	CodeInvalidDayRateMin: {
		Code:       CodeInvalidDayRateMin,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom biaya harian minimal 10000!",
		Translation: errormsg.Translation{
			EN: "Day rate's field minimum 10000!",
		},
	},
	CodeInvalidDayRateMax: {
		Code:       CodeInvalidDayRateMax,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom biaya harian minimal 1000000!",
		Translation: errormsg.Translation{
			EN: "Day rate's field maximum 1000000!",
		},
	},
	CodeInvalidMonthRateMin: {
		Code:       CodeInvalidMonthRateMin,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom biaya bulanan minimal 250000!",
		Translation: errormsg.Translation{
			EN: "Month rate's field maximum 250000!",
		},
	},
	CodeInvalidMonthRateMax: {
		Code:       CodeInvalidMonthRateMax,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom biaya bulanan maksimal 30000000!",
		Translation: errormsg.Translation{
			EN: "Month rate's field maximum 30000000!",
		},
	},
	CodeInvalidImageMin: {
		Code:       CodeInvalidImageMin,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom gambar minimum 10 karakter!",
		Translation: errormsg.Translation{
			EN: "Image field minimum 10 character!",
		},
	},
	CodeInvalidImageMax: {
		Code:       CodeInvalidImageMax,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom gambar maksimum 255 karakter!",
		Translation: errormsg.Translation{
			EN: "Image field maximum 255 character!",
		},
	},
	CodeInvalidImageIsURL: {
		Code:       CodeInvalidImageIsURL,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom gambar harus berupa url!",
		Translation: errormsg.Translation{
			EN: "Image field should be url!",
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
	CodeInvalidPickupLocationMin: {
		Code:       CodeInvalidPickupLocationMin,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom lokasi penjemputan minimal 5 karakter!",
		Translation: errormsg.Translation{
			EN: "Pickup location field minimum 5 character!",
		},
	},
	CodeInvalidPickupLocationMax: {
		Code:       CodeInvalidPickupLocationMax,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom lokasi penjemputan maksimal 50 karakter!",
		Translation: errormsg.Translation{
			EN: "Pickup location field maximum 50 character",
		},
	},
	CodeInvalidPickupLat: {
		Code:       CodeInvalidPickupLat,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom garis lintang penjemputan harus dalam jangkauan -90 sampai 90!",
		Translation: errormsg.Translation{
			EN: "Pickup latitude field should in range -90 to 90!",
		},
	},
	CodeInvalidPickupLong: {
		Code:       CodeInvalidPickupLong,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom garis bujur penjemputan harus dalam jangkauan -180 sampai 180!",
		Translation: errormsg.Translation{
			EN: "Pickup longitude field should in range -180 to 180!",
		},
	},
	CodeInvalidDropoffLocationMin: {
		Code:       CodeInvalidDropoffLocationMin,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom lokasi pengantaran minimal 5 karakter!",
		Translation: errormsg.Translation{
			EN: "Dropoff location minimum 5 character!",
		},
	},
	CodeInvalidDropoffLocationMax: {
		Code:       CodeInvalidDropoffLocationMax,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom lokasi pengantaran maksimal 50 karakter!",
		Translation: errormsg.Translation{
			EN: "Dropoff location maximum 50 character!",
		},
	},
	CodeInvalidDropoffLat: {
		Code:       CodeInvalidDropoffLat,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom garis lintang dalam jangkauan -90 sampai 90!",
		Translation: errormsg.Translation{
			EN: "Dropoff latitude field should in range -90 to 90!",
		},
	},
	CodeInvalidDropoffLong: {
		Code:       CodeInvalidDropoffLong,
		StatusCode: http.StatusBadRequest,
		Message:    "Kolom garis bujur pengantaran dalam jangkauan -180 sampai 180!",
		Translation: errormsg.Translation{
			EN: "Dropoff longitude field should in range -180 to 180!",
		},
	},
}
