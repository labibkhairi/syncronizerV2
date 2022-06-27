package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"prima-integrasi.com/fendiya/syncronizer/internal/datastruct"
	"prima-integrasi.com/fendiya/syncronizer/internal/db"
	"prima-integrasi.com/fendiya/syncronizer/internal/dto"
)

var myEnv map[string]string
var couchDB db.ConnectionCouchbaseSDK
var oracleDB db.Connection

func init() {
	//DB Declaration
	couchDB = db.Couchbase{
		db.DbProperties{
			Hostname: os.Getenv("DB_HOSTNAME"),
			Port:     os.Getenv("DB_PORT"),
			Dbname:   os.Getenv("DB_NAME"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD")}}
	oracleDB = db.Oracle{
		db.DbProperties{
			Hostname: os.Getenv("DB_HOSTNAME1"),
			Port:     os.Getenv("DB_PORT1"),
			Dbname:   os.Getenv("DB_NAME1"),
			Username: os.Getenv("DB_USERNAME1"),
			Password: os.Getenv("DB_PASSWORD1")}}

}

func Sync(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	awb := vars["awb"]
	log.Println(awb)

	ds := fetchFromSource(awb)
	isSuccess := insertToTarget(ds)

	resp := dto.Response{
		Awb:    awb,
		Status: "Failed",
	}

	if isSuccess {
		resp.Status = "Success"
	}
	jsonResponse, _ := json.Marshal(resp)
	w.Write(jsonResponse)

}

func SyncCnoteDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	date := vars["date"]
	log.Println(date)

	ds := fetchFromCnoteDate(date)
	isSuccess := insertToCouchbase(ds)

	resp := dto.ResponseDate{
		Date:   date,
		Status: "Failed",
	}

	if isSuccess {
		resp.Status = "Success"
	}
	jsonResponse, _ := json.Marshal(resp)
	w.Write(jsonResponse)
}

func SyncCmsReturn(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	awb := vars["awb"]
	log.Println(awb)

	ds := fetchFromCmsReturn(awb)
	isSuccess := insertCmsReturn(ds)

	resp := dto.Response{
		Awb:    awb,
		Status: "Failed",
	}

	if isSuccess {
		resp.Status = "Success"
	}
	jsonResponse, _ := json.Marshal(resp)
	w.Write(jsonResponse)
}

// function get data from oracle db 101
func fetchFromCnoteDate(date string) []datastruct.T_dwh {
	log.Println("Fetch from oracle")
	//OPEN DB
	c := oracleDB.OpenConn()
	defer oracleDB.CloseConn(c)

	// select data from oracle
	//query := "SELECT * FROM jnedwh.T_DWH where cnote_date = '" + date + "'"
	query := "SELECT * FROM jnedwh.T_DWH WHERE trunc(cnote_date) = TO_DATE('" + date + "', 'DDMMYY')"
	queryResult, err := c.Query(query)

	if err != nil {
		log.Fatalf("error %v", err)
	}
	var res []datastruct.T_dwh
	for queryResult.Next() {
		var row datastruct.T_dwh

		queryResult.Scan(&row.CREATE_DATE, &row.CNOTE_NO, &row.CNOTE_DATE, &row.CNOTE_CRDATE, &row.CNOTE_BRANCH_ID, &row.BRANCH_REGION, &row.CNOTE_ORIGIN, &row.ORIGIN_NAME, &row.ORIGIN_ZONE, &row.CNOTE_CUST_NO, &row.CNOTE_CUST_TYPE, &row.CUST_NAME, &row.CUST_ADDR1, &row.CUST_ADDR2, &row.CUST_ADDR3, &row.CUST_PHONE, &row.CUST_ZIP, &row.CUST_NA, &row.MARKETPLACE_TYPE, &row.MARKETPLACE_NAME, &row.CNOTE_SHIPPER_NAME, &row.CNOTE_SHIPPER_CONTACT, &row.CNOTE_SHIPPER_ADDR1, &row.CNOTE_SHIPPER_ADDR2, &row.CNOTE_SHIPPER_ADDR3, &row.CNOTE_SHIPPER_PHONE, &row.CNOTE_SHIPPER_ZIP, &row.CNOTE_RECEIVER_NAME, &row.CNOTE_RECEIVER_CONTACT, &row.CNOTE_RECEIVER_ADDR1, &row.CNOTE_RECEIVER_ADDR2, &row.CNOTE_RECEIVER_ADDR3, &row.CNOTE_RECEIVER_PHONE, &row.CNOTE_RECEIVER_ZIP, &row.CNOTE_DESTINATION_ID, &row.BRANCH_DEST_REGION, &row.BRANCH_DESTINATION, &row.CNOTE_DESTINATION, &row.DESTINATION_NAME, &row.DESTINATION_CODE, &row.DESTINATION_ZONE, &row.CNOTE_SERVICES_CODE, &row.ROUTE_ETD_FROM, &row.ROUTE_ETD_THRU, &row.CNOTE_SHIPMENT_TYPE, &row.CNOTE_TRX_TYPE, &row.CNOTE_PAYMENT_TYPE, &row.CNOTE_QTY, &row.CNOTE_WEIGHT, &row.CNOTE_DIM, &row.CNOTE_GOODS_TYPE, &row.CNOTE_PACKING, &row.CNOTE_GOODS_DESCR, &row.CNOTE_GOODS_VALUE, &row.CNOTE_SPECIAL_INS, &row.CNOTE_INSURANCE_ID, &row.CNOTE_INSURANCE_VALUE, &row.CNOTE_AMOUNT, &row.CNOTE_ADDITIONAL_FEE, &row.CNOTE_COD, &row.COD_NO, &row.COD_GOODS_AMOUNT, &row.COD_AMOUNT, &row.CNOTE_CASHLESS, &row.JLC_NO, &row.JLC_NAME, &row.JLC_DISCOUNT, &row.HYBRID_BRANCH, &row.HYBRID_CUST_NO, &row.HYBRID_CUST_NAME, &row.HYBRID_CUST_ADDR1, &row.HYBRID_CUST_ADDR2, &row.HYBRID_CUST_ADDR3, &row.HYBRID_CUST_PHONE, &row.HYBRID_CUST_ZIP, &row.CNOTE_CANCEL, &row.CNOTE_HOLD, &row.CNOTE_USER, &row.CNOTE_USER_ZONE, &row.R_CNOTE_FREIGHT_CHARGE, &row.PUBLISH_RATE, &row.CASHREG_NO, &row.CASHREG_DATE, &row.CASHREG_USER_ID, &row.CASHREG_USER_ZONE, &row.CASHREG_CRDATE, &row.PICKUP_NO, &row.PICKUP_COURIER_ID, &row.PICKUP_COURIER_ZONE, &row.PICKUP_DATE, &row.PICKUP_CRDATE, &row.PICKUP_MERCHAN_ID, &row.PICKUP_LATITUDE, &row.PICKUP_LONGITUDE, &row.PU_FIRST_ATTTEMP_STATUS_CODE, &row.PU_FIRST_ATTTEMP_STATUS_DESC, &row.PU_FIRST_ATTTEMP_STATUS_DATE, &row.PU_LAST_ATTEMP_STATUS_CODE, &row.PU_LAST_ATTEMP_STATUS_DESC, &row.PU_LAST_ATTEMP_STATUS_DATE, &row.PU_REF_ID, &row.HO_NO, &row.HO_DATE, &row.HO_COURIER_ID, &row.HO_CDATE, &row.RECEIVING_AGENT_NO, &row.RECEIVING_AGENT_DATE, &row.RECEIVING_AGENT_BRANCH, &row.RECEIVING_AGENT_COURIER_ID, &row.RECEIVING_AGENT_USER_ID, &row.RECEIVING_AGENT_USER_ZONE, &row.RECEIVING_AGENT_CRDATE, &row.RECEIVING_OUT_NO, &row.RECEIVING_OUT_DATE, &row.RECEIVING_OUT_BRANCH, &row.RECEIVING_OUT_COURIER_ID, &row.RECEIVING_OUT_USER_ID, &row.RECEIVING_OUT_USER_ZONE, &row.RECEIVING_OUT_CRDATE, &row.MANIFEST_OUTB_NO, &row.MANIFEST_OUTB_ORIGIN, &row.MANIFEST_OUTB_DATE, &row.MANIFEST_OUTB_BAG_NO, &row.MANIFEST_OUTB_USER_ID, &row.MANIFEST_OUTB_USER_ZONE, &row.MANIFEST_OUTB_CRDATE, &row.SMU_NO, &row.SMU_SCHD_NO, &row.SMU_SCH_DATE, &row.SMU_DATE, &row.SMU_ETD, &row.SMU_ETA, &row.SMU_REMARKS, &row.SMU_REMARKS_DATE, &row.SMU_QTY, &row.SMU_WEIGHT, &row.SMU_FLAG_APPROVE, &row.SMU_FLAG_CANCEL, &row.SMU_DESTINATION, &row.MANIFEST_TRS1_NO, &row.MANIFEST_TRS1_ORIGIN, &row.MANIFEST_TRS1_DATE, &row.MANIFEST_TRS1_BAG_NO, &row.MANIFEST_TRS1_USER_ID, &row.MANIFEST_TRS1_USER_ZONE, &row.MANIFEST_TRS1_CRDATE, &row.MANIFEST_TRSL_NO, &row.MANIFEST_TRSL_ORIGIN, &row.MANIFEST_TRSL_DATE, &row.MANIFEST_TRSL_BAG_NO, &row.MANIFEST_TRSL_USER_ID, &row.MANIFEST_TRSL_USER_ZONE, &row.MANIFEST_TRSL_CRDATE, &row.MANIFEST_INB_NO, &row.MANIFEST_INB_ORIGIN, &row.MANIFEST_INB_DATE, &row.MANIFEST_INB_BAG_NO, &row.MANIFEST_INB_USER_ID, &row.MANIFEST_INB_USER_ZONE, &row.MANIFEST_INB_CRDATE, &row.MANIFEST_BAG_NO, &row.MANIFEST_BAG_DATE, &row.MANIFEST_BAG_BAG_NO, &row.MANIFEST_BAG_USER_ID, &row.MANIFEST_BAG_USER_ZONE, &row.MANIFEST_BAG_CRDATE, &row.PRA_MRSHEET_NO, &row.PRA_MRSHEET_DATE, &row.PRA_MRSHEET_BRANCH, &row.PRA_MRSHEET_ZONE, &row.PRA_MRSHEET_COURIER_ID, &row.PRA_COURIER_ZONE_CODE, &row.PRA_MRSHEET_UID, &row.PRA_USER_ZONE_CODE, &row.PRA_CREATION_DATE, &row.MTA_OUT_MANIFEST_NO, &row.MTA_OUT_MANIFEST_DATE, &row.MTA_OUT_BRANCH_ID, &row.MTA_OUT_DESTINATION, &row.MTA_OUT_MANIFEST_UID, &row.MTA_OUT_USER_ZONE_CODE, &row.MTA_OUT_ESB_TIME, &row.MTA_INB_MANIFEST_NO, &row.MTA_INB_MANIFEST_DATE, &row.MTA_INB_BRANCH_ID, &row.MTA_INB_DESTINATION, &row.MTA_INB_MANIFEST_UID, &row.MTA_INB_USER_ZONE_CODE, &row.MTA_INB_ESB_TIME, &row.MHOCNOTE_NO, &row.MHOCNOTE_DATE, &row.MHOCNOTE_BRANCH_ID, &row.MHOCNOTE_ZONE, &row.MHOCNOTE_ZONE_DEST, &row.MHOCNOTE_USER_ID, &row.MHOCNOTE_USER_ZONE_CODE, &row.DHOCNOTE_TDATE, &row.MHICNOTE_NO, &row.MHICNOTE_DATE, &row.MHICNOTE_BRANCH_ID, &row.MHICNOTE_ZONE, &row.MHICNOTE_USER_ID, &row.MHICNOTE_USER_ZONE_CODE, &row.DHICNOTE_TDATE, &row.MRSHEET1_NO, &row.MRSHEET1_DATE, &row.MRSHEET1_BRANCH, &row.MRSHEET1_COURIER_ID, &row.MRSHEET1_UID, &row.MRSHEET1_USER_ZONE_CODE, &row.MRSHEET1_CREATION_DATE, &row.MRSHEETL_NO, &row.MRSHEETL_DATE, &row.MRSHEETL_BRANCH, &row.MRSHEETL_COURIER_ID, &row.MRSHEETL_UID, &row.MRSHEETL_USER_ZONE_CODE, &row.MRSHEETL_CREATION_DATE, &row.POD1_DRSHEET_NO, &row.POD1_MRSHEET_DATE, &row.POD1_MRSHEET_BRANCH, &row.POD1_MRSHEET_COURIER_ID, &row.POD1_COURIER_ZONE_CODE, &row.POD1_DRSHEET_DATE, &row.POD1_DRSHEET_RECEIVER, &row.POD1_DRSHEET_STATUS, &row.POD1_LATITUDE, &row.POD1_LONGITUDE, &row.POD1_EPOD_URL, &row.POD1_EPOD_URL_PIC, &row.POD1_DRSHEET_UID, &row.POD1_USER_ZONE_CODE, &row.POD1_DRSHEET_UDATE, &row.PODL_DRSHEET_NO, &row.PODL_MRSHEET_DATE, &row.PODL_MRSHEET_BRANCH, &row.PODL_MRSHEET_COURIER_ID, &row.PODL_COURIER_ZONE_CODE, &row.PODL_DRSHEET_DATE, &row.PODL_DRSHEET_RECEIVER, &row.PODL_DRSHEET_STATUS, &row.PODL_LATITUDE, &row.PODL_LONGITUDE, &row.PODL_EPOD_URL, &row.PODL_EPOD_URL_PIC, &row.PODL_DRSHEET_UID, &row.PODL_USER_ZONE_CODE, &row.PODL_DRSHEET_UDATE, &row.DO_NO, &row.DO_DATE, &row.RDO_NO, &row.RDO_DATE, &row.SHIPPER_PROVIDER, &row.CNOTE_REFNO, &row.MANIFEST_OUTB_APPROVED, &row.MANIFEST_INB_APPROVED, &row.SMU_BAG_BUX, &row.SMU_TGL_MASTER_BAG, &row.SMU_USER_MASTER_BAG, &row.SMU_NO_MASTER_BAG, &row.SMU_MANIFEST_DESTINATION, &row.MANIFEST_COST_WEIGHT, &row.MANIFEST_ACT_WEIGHT, &row.DWH_PACKING_FEE, &row.DWH_SURCHARGE, &row.DWH_DISC_REV_TYPE, &row.DWH_DISCOUNT_AMT, &row.DWH_FCHARGE_AFT_DISC_AMT, &row.DWH_CUST_DISC_IC, &row.DWH_CUST_DISC_DM, &row.DWH_RT_PACKING_FEE, &row.DWH_RT_FREIGHT_CHARGE, &row.DWH_RT_SURCHARGE, &row.DWH_RT_DISC_AMT, &row.DWH_RT_FCHARGE_AFT_DISC_AMT, &row.DWH_PAYTYPE, &row.DWH_EPAY_VEND, &row.DWH_EPAY_TRXID, &row.DWH_VAT_FCHARGE_AFT_DISC, &row.DWH_VAT_RT_FCHARGE_AFT_DISC)

		log.Println("Cnote No :" + row.CNOTE_NO)
		res = append(res, row)
	}
	return res
}

//function get data from cms return db 101
func fetchFromCmsReturn(awb string) []datastruct.CMS_RETURN_REQCNOTE {
	log.Println("Fetch from Oracle")
	//OPEN DB
	c := oracleDB.OpenConn()
	defer oracleDB.CloseConn(c)

	//select data from Oracle
	query := "select * from cms_return_reqcnote where reqcnote_cnote_no = '" + awb + "'"
	queryResult, err := c.Query(query)
	if err != nil {
		log.Fatalf("error %v", err)
	}
	var res []datastruct.CMS_RETURN_REQCNOTE
	for queryResult.Next() {
		var row datastruct.CMS_RETURN_REQCNOTE

		queryResult.Scan(&row.REQCNOTE_CNOTE_NO, &row.REQCNOTE_CNOTE_NO_RT, &row.REQCNOTE_HRS_NO, &row.REQCNOTE_REQ_USER, &row.REQCNOTE_AMOUNT1, &row.REQCNOTE_AMOUNT2, &row.REQCNOTE_AMOUNT3, &row.REQCNOTE_CDATE)

		log.Println("here :" + row.REQCNOTE_CNOTE_NO)
		res = append(res, row)
	}
	return res

}

//function get data from oracle
func fetchFromSource(awb string) []datastruct.T_dwh {
	log.Println("Fetch from Oracle")
	//OPEN DB
	c := oracleDB.OpenConn()
	defer oracleDB.CloseConn(c)

	//select data from Oracle
	var res []datastruct.T_dwh
	query := "SELECT * FROM jnedwh.T_DWH where cnote_no = '" + awb + "'"
	queryResult, err := c.Query(query)

	if err != nil {
		log.Fatalf("error %v", err)
	}

	for queryResult.Next() {
		var row datastruct.T_dwh

		queryResult.Scan(&row.CREATE_DATE, &row.CNOTE_NO, &row.CNOTE_DATE, &row.CNOTE_CRDATE, &row.CNOTE_BRANCH_ID, &row.BRANCH_REGION, &row.CNOTE_ORIGIN, &row.ORIGIN_NAME, &row.ORIGIN_ZONE, &row.CNOTE_CUST_NO, &row.CNOTE_CUST_TYPE, &row.CUST_NAME, &row.CUST_ADDR1, &row.CUST_ADDR2, &row.CUST_ADDR3, &row.CUST_PHONE, &row.CUST_ZIP, &row.CUST_NA, &row.MARKETPLACE_TYPE, &row.MARKETPLACE_NAME, &row.CNOTE_SHIPPER_NAME, &row.CNOTE_SHIPPER_CONTACT, &row.CNOTE_SHIPPER_ADDR1, &row.CNOTE_SHIPPER_ADDR2, &row.CNOTE_SHIPPER_ADDR3, &row.CNOTE_SHIPPER_PHONE, &row.CNOTE_SHIPPER_ZIP, &row.CNOTE_RECEIVER_NAME, &row.CNOTE_RECEIVER_CONTACT, &row.CNOTE_RECEIVER_ADDR1, &row.CNOTE_RECEIVER_ADDR2, &row.CNOTE_RECEIVER_ADDR3, &row.CNOTE_RECEIVER_PHONE, &row.CNOTE_RECEIVER_ZIP, &row.CNOTE_DESTINATION_ID, &row.BRANCH_DEST_REGION, &row.BRANCH_DESTINATION, &row.CNOTE_DESTINATION, &row.DESTINATION_NAME, &row.DESTINATION_CODE, &row.DESTINATION_ZONE, &row.CNOTE_SERVICES_CODE, &row.ROUTE_ETD_FROM, &row.ROUTE_ETD_THRU, &row.CNOTE_SHIPMENT_TYPE, &row.CNOTE_TRX_TYPE, &row.CNOTE_PAYMENT_TYPE, &row.CNOTE_QTY, &row.CNOTE_WEIGHT, &row.CNOTE_DIM, &row.CNOTE_GOODS_TYPE, &row.CNOTE_PACKING, &row.CNOTE_GOODS_DESCR, &row.CNOTE_GOODS_VALUE, &row.CNOTE_SPECIAL_INS, &row.CNOTE_INSURANCE_ID, &row.CNOTE_INSURANCE_VALUE, &row.CNOTE_AMOUNT, &row.CNOTE_ADDITIONAL_FEE, &row.CNOTE_COD, &row.COD_NO, &row.COD_GOODS_AMOUNT, &row.COD_AMOUNT, &row.CNOTE_CASHLESS, &row.JLC_NO, &row.JLC_NAME, &row.JLC_DISCOUNT, &row.HYBRID_BRANCH, &row.HYBRID_CUST_NO, &row.HYBRID_CUST_NAME, &row.HYBRID_CUST_ADDR1, &row.HYBRID_CUST_ADDR2, &row.HYBRID_CUST_ADDR3, &row.HYBRID_CUST_PHONE, &row.HYBRID_CUST_ZIP, &row.CNOTE_CANCEL, &row.CNOTE_HOLD, &row.CNOTE_USER, &row.CNOTE_USER_ZONE, &row.R_CNOTE_FREIGHT_CHARGE, &row.PUBLISH_RATE, &row.CASHREG_NO, &row.CASHREG_DATE, &row.CASHREG_USER_ID, &row.CASHREG_USER_ZONE, &row.CASHREG_CRDATE, &row.PICKUP_NO, &row.PICKUP_COURIER_ID, &row.PICKUP_COURIER_ZONE, &row.PICKUP_DATE, &row.PICKUP_CRDATE, &row.PICKUP_MERCHAN_ID, &row.PICKUP_LATITUDE, &row.PICKUP_LONGITUDE, &row.PU_FIRST_ATTTEMP_STATUS_CODE, &row.PU_FIRST_ATTTEMP_STATUS_DESC, &row.PU_FIRST_ATTTEMP_STATUS_DATE, &row.PU_LAST_ATTEMP_STATUS_CODE, &row.PU_LAST_ATTEMP_STATUS_DESC, &row.PU_LAST_ATTEMP_STATUS_DATE, &row.PU_REF_ID, &row.HO_NO, &row.HO_DATE, &row.HO_COURIER_ID, &row.HO_CDATE, &row.RECEIVING_AGENT_NO, &row.RECEIVING_AGENT_DATE, &row.RECEIVING_AGENT_BRANCH, &row.RECEIVING_AGENT_COURIER_ID, &row.RECEIVING_AGENT_USER_ID, &row.RECEIVING_AGENT_USER_ZONE, &row.RECEIVING_AGENT_CRDATE, &row.RECEIVING_OUT_NO, &row.RECEIVING_OUT_DATE, &row.RECEIVING_OUT_BRANCH, &row.RECEIVING_OUT_COURIER_ID, &row.RECEIVING_OUT_USER_ID, &row.RECEIVING_OUT_USER_ZONE, &row.RECEIVING_OUT_CRDATE, &row.MANIFEST_OUTB_NO, &row.MANIFEST_OUTB_ORIGIN, &row.MANIFEST_OUTB_DATE, &row.MANIFEST_OUTB_BAG_NO, &row.MANIFEST_OUTB_USER_ID, &row.MANIFEST_OUTB_USER_ZONE, &row.MANIFEST_OUTB_CRDATE, &row.SMU_NO, &row.SMU_SCHD_NO, &row.SMU_SCH_DATE, &row.SMU_DATE, &row.SMU_ETD, &row.SMU_ETA, &row.SMU_REMARKS, &row.SMU_REMARKS_DATE, &row.SMU_QTY, &row.SMU_WEIGHT, &row.SMU_FLAG_APPROVE, &row.SMU_FLAG_CANCEL, &row.SMU_DESTINATION, &row.MANIFEST_TRS1_NO, &row.MANIFEST_TRS1_ORIGIN, &row.MANIFEST_TRS1_DATE, &row.MANIFEST_TRS1_BAG_NO, &row.MANIFEST_TRS1_USER_ID, &row.MANIFEST_TRS1_USER_ZONE, &row.MANIFEST_TRS1_CRDATE, &row.MANIFEST_TRSL_NO, &row.MANIFEST_TRSL_ORIGIN, &row.MANIFEST_TRSL_DATE, &row.MANIFEST_TRSL_BAG_NO, &row.MANIFEST_TRSL_USER_ID, &row.MANIFEST_TRSL_USER_ZONE, &row.MANIFEST_TRSL_CRDATE, &row.MANIFEST_INB_NO, &row.MANIFEST_INB_ORIGIN, &row.MANIFEST_INB_DATE, &row.MANIFEST_INB_BAG_NO, &row.MANIFEST_INB_USER_ID, &row.MANIFEST_INB_USER_ZONE, &row.MANIFEST_INB_CRDATE, &row.MANIFEST_BAG_NO, &row.MANIFEST_BAG_DATE, &row.MANIFEST_BAG_BAG_NO, &row.MANIFEST_BAG_USER_ID, &row.MANIFEST_BAG_USER_ZONE, &row.MANIFEST_BAG_CRDATE, &row.PRA_MRSHEET_NO, &row.PRA_MRSHEET_DATE, &row.PRA_MRSHEET_BRANCH, &row.PRA_MRSHEET_ZONE, &row.PRA_MRSHEET_COURIER_ID, &row.PRA_COURIER_ZONE_CODE, &row.PRA_MRSHEET_UID, &row.PRA_USER_ZONE_CODE, &row.PRA_CREATION_DATE, &row.MTA_OUT_MANIFEST_NO, &row.MTA_OUT_MANIFEST_DATE, &row.MTA_OUT_BRANCH_ID, &row.MTA_OUT_DESTINATION, &row.MTA_OUT_MANIFEST_UID, &row.MTA_OUT_USER_ZONE_CODE, &row.MTA_OUT_ESB_TIME, &row.MTA_INB_MANIFEST_NO, &row.MTA_INB_MANIFEST_DATE, &row.MTA_INB_BRANCH_ID, &row.MTA_INB_DESTINATION, &row.MTA_INB_MANIFEST_UID, &row.MTA_INB_USER_ZONE_CODE, &row.MTA_INB_ESB_TIME, &row.MHOCNOTE_NO, &row.MHOCNOTE_DATE, &row.MHOCNOTE_BRANCH_ID, &row.MHOCNOTE_ZONE, &row.MHOCNOTE_ZONE_DEST, &row.MHOCNOTE_USER_ID, &row.MHOCNOTE_USER_ZONE_CODE, &row.DHOCNOTE_TDATE, &row.MHICNOTE_NO, &row.MHICNOTE_DATE, &row.MHICNOTE_BRANCH_ID, &row.MHICNOTE_ZONE, &row.MHICNOTE_USER_ID, &row.MHICNOTE_USER_ZONE_CODE, &row.DHICNOTE_TDATE, &row.MRSHEET1_NO, &row.MRSHEET1_DATE, &row.MRSHEET1_BRANCH, &row.MRSHEET1_COURIER_ID, &row.MRSHEET1_UID, &row.MRSHEET1_USER_ZONE_CODE, &row.MRSHEET1_CREATION_DATE, &row.MRSHEETL_NO, &row.MRSHEETL_DATE, &row.MRSHEETL_BRANCH, &row.MRSHEETL_COURIER_ID, &row.MRSHEETL_UID, &row.MRSHEETL_USER_ZONE_CODE, &row.MRSHEETL_CREATION_DATE, &row.POD1_DRSHEET_NO, &row.POD1_MRSHEET_DATE, &row.POD1_MRSHEET_BRANCH, &row.POD1_MRSHEET_COURIER_ID, &row.POD1_COURIER_ZONE_CODE, &row.POD1_DRSHEET_DATE, &row.POD1_DRSHEET_RECEIVER, &row.POD1_DRSHEET_STATUS, &row.POD1_LATITUDE, &row.POD1_LONGITUDE, &row.POD1_EPOD_URL, &row.POD1_EPOD_URL_PIC, &row.POD1_DRSHEET_UID, &row.POD1_USER_ZONE_CODE, &row.POD1_DRSHEET_UDATE, &row.PODL_DRSHEET_NO, &row.PODL_MRSHEET_DATE, &row.PODL_MRSHEET_BRANCH, &row.PODL_MRSHEET_COURIER_ID, &row.PODL_COURIER_ZONE_CODE, &row.PODL_DRSHEET_DATE, &row.PODL_DRSHEET_RECEIVER, &row.PODL_DRSHEET_STATUS, &row.PODL_LATITUDE, &row.PODL_LONGITUDE, &row.PODL_EPOD_URL, &row.PODL_EPOD_URL_PIC, &row.PODL_DRSHEET_UID, &row.PODL_USER_ZONE_CODE, &row.PODL_DRSHEET_UDATE, &row.DO_NO, &row.DO_DATE, &row.RDO_NO, &row.RDO_DATE, &row.SHIPPER_PROVIDER, &row.CNOTE_REFNO, &row.MANIFEST_OUTB_APPROVED, &row.MANIFEST_INB_APPROVED, &row.SMU_BAG_BUX, &row.SMU_TGL_MASTER_BAG, &row.SMU_USER_MASTER_BAG, &row.SMU_NO_MASTER_BAG, &row.SMU_MANIFEST_DESTINATION, &row.MANIFEST_COST_WEIGHT, &row.MANIFEST_ACT_WEIGHT, &row.DWH_PACKING_FEE, &row.DWH_SURCHARGE, &row.DWH_DISC_REV_TYPE, &row.DWH_DISCOUNT_AMT, &row.DWH_FCHARGE_AFT_DISC_AMT, &row.DWH_CUST_DISC_IC, &row.DWH_CUST_DISC_DM, &row.DWH_RT_PACKING_FEE, &row.DWH_RT_FREIGHT_CHARGE, &row.DWH_RT_SURCHARGE, &row.DWH_RT_DISC_AMT, &row.DWH_RT_FCHARGE_AFT_DISC_AMT, &row.DWH_PAYTYPE, &row.DWH_EPAY_VEND, &row.DWH_EPAY_TRXID, &row.DWH_VAT_FCHARGE_AFT_DISC, &row.DWH_VAT_RT_FCHARGE_AFT_DISC)

		log.Println("here :" + row.CNOTE_NO)
		res = append(res, row)
	}
	// log.Printf("here : %v", res[0])
	return res

}

func checkNull() {

}

//function insert cms_return to couchbase mapr
func insertCmsReturn(datas []datastruct.CMS_RETURN_REQCNOTE) bool {
	log.Println("insert to couchbase mapr")

	//OPEN DB
	c := couchDB.OpenConn()
	//OPEN bucket, scope and collection
	log.Println("bucket name :" + os.Getenv("DB_NAME"))
	bucket := c.Bucket(os.Getenv("DB_NAME"))

	err := bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		log.Fatal(err)
	}

	col := bucket.Scope("DWH").Collection("cms_return_reqcnote")

	for _, data := range datas {
		key := data.REQCNOTE_CNOTE_NO
		_, err = col.Upsert(key, data, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("insert to Couchbase Success")
	return true
}

//function insert to couchbase mapr
func insertToCouchbase(datas []datastruct.T_dwh) bool {
	log.Println("insert to couchbase mapr")

	//OPEN DB
	c := couchDB.OpenConn()
	//OPEN bucket, scope and collection
	log.Println("bucket name :" + os.Getenv("DB_NAME"))
	bucket := c.Bucket(os.Getenv("DB_NAME"))

	err := bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		log.Fatal(err)
	}

	col := bucket.Scope("DWH").Collection("T_DWH")

	for _, data := range datas {
		key := data.CNOTE_NO
		_, err = col.Upsert(key, data, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("insert to Couchbase Success")
	return true
}

//function insert to Couchbase
func insertToTarget(datas []datastruct.T_dwh) bool {
	log.Println("insert to Couchbase")

	//OPEN DB
	c := couchDB.OpenConn()
	//OPEN bucket, scope and collection
	log.Println("bucket name :" + os.Getenv("DB_NAME"))
	bucket := c.Bucket(os.Getenv("DB_NAME"))

	err := bucket.WaitUntilReady(10*time.Second, nil)
	if err != nil {
		log.Fatal(err)
	}

	col := bucket.Scope("DWH").Collection("T_DWH")

	for _, data := range datas {
		key := data.CNOTE_NO

		// newData := datastruct.T_dwh{
		// 	CREATE_DATE:                  data.CREATE_DATE.Time,
		// 	CNOTE_NO:                     data.CNOTE_NO.String,
		// 	CNOTE_DATE:                   data.CNOTE_DATE.Time,
		// 	CNOTE_CRDATE:                 data.CNOTE_CRDATE.Time,
		// 	CNOTE_BRANCH_ID:              data.CNOTE_BRANCH_ID.String,
		// 	BRANCH_REGION:                data.BRANCH_REGION.String,
		// 	CNOTE_ORIGIN:                 data.CNOTE_ORIGIN.String,
		// 	ORIGIN_NAME:                  data.ORIGIN_NAME.String,
		// 	ORIGIN_ZONE:                  data.ORIGIN_ZONE.String,
		// 	CNOTE_CUST_NO:                data.CNOTE_CUST_NO.String,
		// 	CNOTE_CUST_TYPE:              data.CNOTE_CUST_TYPE.String,
		// 	CUST_NAME:                    data.CUST_NAME.String,
		// 	CUST_ADDR1:                   data.CUST_ADDR1.String,
		// 	CUST_ADDR2:                   data.CUST_ADDR2.String,
		// 	CUST_ADDR3:                   data.CUST_ADDR3.String,
		// 	CUST_PHONE:                   data.CUST_PHONE.String,
		// 	CUST_ZIP:                     data.CUST_ZIP.String,
		// 	CUST_NA:                      data.CUST_NA.String,
		// 	MARKETPLACE_TYPE:             data.MARKETPLACE_TYPE.String,
		// 	MARKETPLACE_NAME:             data.MARKETPLACE_NAME.String,
		// 	CNOTE_SHIPPER_NAME:           data.CNOTE_SHIPPER_NAME.String,
		// 	CNOTE_SHIPPER_CONTACT:        data.CNOTE_SHIPPER_CONTACT.String,
		// 	CNOTE_SHIPPER_ADDR1:          data.CNOTE_SHIPPER_ADDR1.String,
		// 	CNOTE_SHIPPER_ADDR2:          data.CNOTE_SHIPPER_ADDR2.String,
		// 	CNOTE_SHIPPER_ADDR3:          data.CNOTE_SHIPPER_ADDR3.String,
		// 	CNOTE_SHIPPER_PHONE:          data.CNOTE_SHIPPER_PHONE.String,
		// 	CNOTE_SHIPPER_ZIP:            data.CNOTE_SHIPPER_ZIP.String,
		// 	CNOTE_RECEIVER_NAME:          data.CNOTE_RECEIVER_NAME.String,
		// 	CNOTE_RECEIVER_CONTACT:       data.CNOTE_RECEIVER_CONTACT.String,
		// 	CNOTE_RECEIVER_ADDR1:         data.CNOTE_RECEIVER_ADDR1.String,
		// 	CNOTE_RECEIVER_ADDR2:         data.CNOTE_RECEIVER_ADDR2.String,
		// 	CNOTE_RECEIVER_ADDR3:         data.CNOTE_RECEIVER_ADDR3.String,
		// 	CNOTE_RECEIVER_PHONE:         data.CNOTE_RECEIVER_PHONE.String,
		// 	CNOTE_RECEIVER_ZIP:           data.CNOTE_RECEIVER_ZIP.String,
		// 	CNOTE_DESTINATION_ID:         data.CNOTE_DESTINATION_ID.String,
		// 	BRANCH_DEST_REGION:           data.BRANCH_DEST_REGION.String,
		// 	BRANCH_DESTINATION:           data.BRANCH_DESTINATION.String,
		// 	CNOTE_DESTINATION:            data.CNOTE_DESTINATION.String,
		// 	DESTINATION_NAME:             data.DESTINATION_NAME.String,
		// 	DESTINATION_CODE:             data.DESTINATION_CODE.String,
		// 	DESTINATION_ZONE:             data.DESTINATION_ZONE.String,
		// 	CNOTE_SERVICES_CODE:          data.CNOTE_SERVICES_CODE.String,
		// 	ROUTE_ETD_FROM:               data.ROUTE_ETD_FROM.Float64,
		// 	ROUTE_ETD_THRU:               data.ROUTE_ETD_THRU.Float64,
		// 	CNOTE_SHIPMENT_TYPE:          data.CNOTE_SHIPMENT_TYPE.String,
		// 	CNOTE_TRX_TYPE:               data.CNOTE_TRX_TYPE.String,
		// 	CNOTE_PAYMENT_TYPE:           data.CNOTE_PAYMENT_TYPE.String,
		// 	CNOTE_QTY:                    data.CNOTE_QTY.Float64,
		// 	CNOTE_WEIGHT:                 data.CNOTE_WEIGHT.Float64,
		// 	CNOTE_DIM:                    data.CNOTE_DIM.Float64,
		// 	CNOTE_GOODS_TYPE:             data.CNOTE_GOODS_TYPE.String,
		// 	CNOTE_PACKING:                data.CNOTE_PACKING.String,
		// 	CNOTE_GOODS_DESCR:            data.CNOTE_GOODS_DESCR.String,
		// 	CNOTE_GOODS_VALUE:            data.CNOTE_GOODS_VALUE.Float64,
		// 	CNOTE_SPECIAL_INS:            data.CNOTE_SPECIAL_INS.String,
		// 	CNOTE_INSURANCE_ID:           data.CNOTE_INSURANCE_ID.String,
		// 	CNOTE_INSURANCE_VALUE:        data.CNOTE_INSURANCE_VALUE.Float64,
		// 	CNOTE_AMOUNT:                 data.CNOTE_AMOUNT.Float64,
		// 	CNOTE_ADDITIONAL_FEE:         data.CNOTE_ADDITIONAL_FEE.Float64,
		// 	CNOTE_COD:                    data.CNOTE_COD.String,
		// 	COD_NO:                       data.COD_NO.String,
		// 	COD_GOODS_AMOUNT:             data.COD_GOODS_AMOUNT.Float64,
		// 	COD_AMOUNT:                   data.COD_AMOUNT.Float64,
		// 	CNOTE_CASHLESS:               data.CNOTE_CASHLESS.String,
		// 	JLC_NO:                       data.JLC_NO.String,
		// 	JLC_NAME:                     data.JLC_NAME.String,
		// 	JLC_DISCOUNT:                 data.JLC_DISCOUNT.Float64,
		// 	HYBRID_BRANCH:                data.HYBRID_BRANCH.String,
		// 	HYBRID_CUST_NO:               data.HYBRID_CUST_NO.String,
		// 	HYBRID_CUST_NAME:             data.HYBRID_CUST_NAME.String,
		// 	HYBRID_CUST_ADDR1:            data.HYBRID_CUST_ADDR1.String,
		// 	HYBRID_CUST_ADDR2:            data.HYBRID_CUST_ADDR2.String,
		// 	HYBRID_CUST_ADDR3:            data.HYBRID_CUST_ADDR3.String,
		// 	HYBRID_CUST_PHONE:            data.HYBRID_CUST_PHONE.String,
		// 	HYBRID_CUST_ZIP:              data.HYBRID_CUST_ZIP.String,
		// 	CNOTE_CANCEL:                 data.CNOTE_CANCEL.String,
		// 	CNOTE_HOLD:                   data.CNOTE_HOLD.String,
		// 	CNOTE_USER:                   data.CNOTE_USER.String,
		// 	CNOTE_USER_ZONE:              data.CNOTE_USER_ZONE.String,
		// 	R_CNOTE_FREIGHT_CHARGE:       data.R_CNOTE_FREIGHT_CHARGE.Float64,
		// 	PUBLISH_RATE:                 data.PUBLISH_RATE.Float64,
		// 	CASHREG_NO:                   data.CASHREG_NO.String,
		// 	CASHREG_DATE:                 data.CASHREG_DATE.Time,
		// 	CASHREG_USER_ID:              data.CASHREG_USER_ID.String,
		// 	CASHREG_USER_ZONE:            data.CASHREG_USER_ZONE.String,
		// 	CASHREG_CRDATE:               data.CASHREG_CRDATE.Time,
		// 	PICKUP_NO:                    data.PICKUP_NO.String,
		// 	PICKUP_COURIER_ID:            data.PICKUP_COURIER_ID.String,
		// 	PICKUP_COURIER_ZONE:          data.PICKUP_COURIER_ZONE.String,
		// 	PICKUP_DATE:                  data.PICKUP_DATE.Time,
		// 	PICKUP_CRDATE:                data.PICKUP_CRDATE.Time,
		// 	PICKUP_MERCHAN_ID:            data.PICKUP_MERCHAN_ID.String,
		// 	PICKUP_LATITUDE:              data.PICKUP_LATITUDE.String,
		// 	PICKUP_LONGITUDE:             data.PICKUP_LONGITUDE.String,
		// 	PU_FIRST_ATTTEMP_STATUS_CODE: data.PU_FIRST_ATTTEMP_STATUS_CODE.String,
		// 	PU_FIRST_ATTTEMP_STATUS_DESC: data.PU_FIRST_ATTTEMP_STATUS_DESC.String,
		// 	PU_FIRST_ATTTEMP_STATUS_DATE: data.PU_FIRST_ATTTEMP_STATUS_DATE.Time,
		// 	PU_LAST_ATTEMP_STATUS_CODE:   data.PU_LAST_ATTEMP_STATUS_CODE.String,
		// 	PU_LAST_ATTEMP_STATUS_DESC:   data.PU_LAST_ATTEMP_STATUS_DESC.String,
		// 	PU_LAST_ATTEMP_STATUS_DATE:   data.PU_LAST_ATTEMP_STATUS_DATE.Time,
		// 	PU_REF_ID:                    data.PU_REF_ID.String,
		// 	HO_NO:                        data.HO_NO.String,
		// 	HO_DATE:                      data.HO_DATE.Time,
		// 	HO_COURIER_ID:                data.HO_COURIER_ID.String,
		// 	HO_CDATE:                     data.HO_CDATE.Time,
		// 	RECEIVING_AGENT_NO:           data.RECEIVING_AGENT_NO.String,
		// 	RECEIVING_AGENT_DATE:         data.RECEIVING_AGENT_DATE.Time,
		// 	RECEIVING_AGENT_BRANCH:       data.RECEIVING_AGENT_BRANCH.String,
		// 	RECEIVING_AGENT_COURIER_ID:   data.RECEIVING_AGENT_COURIER_ID.String,
		// 	RECEIVING_AGENT_USER_ID:      data.RECEIVING_AGENT_USER_ID.String,
		// 	RECEIVING_AGENT_USER_ZONE:    data.RECEIVING_AGENT_USER_ZONE.String,
		// 	RECEIVING_AGENT_CRDATE:       data.RECEIVING_AGENT_CRDATE.Time,
		// 	RECEIVING_OUT_NO:             data.RECEIVING_OUT_NO.String,
		// 	RECEIVING_OUT_DATE:           data.RECEIVING_OUT_DATE.Time,
		// 	RECEIVING_OUT_BRANCH:         data.RECEIVING_OUT_BRANCH.String,
		// 	RECEIVING_OUT_COURIER_ID:     data.RECEIVING_OUT_COURIER_ID.String,
		// 	RECEIVING_OUT_USER_ID:        data.RECEIVING_OUT_USER_ID.String,
		// 	RECEIVING_OUT_USER_ZONE:      data.RECEIVING_OUT_USER_ZONE.String,
		// 	RECEIVING_OUT_CRDATE:         data.RECEIVING_OUT_CRDATE.Time,
		// 	MANIFEST_OUTB_NO:             data.MANIFEST_OUTB_NO.String,
		// 	MANIFEST_OUTB_ORIGIN:         data.MANIFEST_OUTB_ORIGIN.String,
		// 	MANIFEST_OUTB_DATE:           data.MANIFEST_OUTB_DATE.Time,
		// 	MANIFEST_OUTB_BAG_NO:         data.MANIFEST_OUTB_BAG_NO.String,
		// 	MANIFEST_OUTB_USER_ID:        data.MANIFEST_OUTB_USER_ID.String,
		// 	MANIFEST_OUTB_USER_ZONE:      data.MANIFEST_OUTB_USER_ZONE.String,
		// 	MANIFEST_OUTB_CRDATE:         data.MANIFEST_OUTB_CRDATE.Time,
		// 	SMU_NO:                       data.SMU_NO.String,
		// 	SMU_SCHD_NO:                  data.SMU_SCHD_NO.String,
		// 	SMU_SCH_DATE:                 data.SMU_SCH_DATE.Time,
		// 	SMU_DATE:                     data.SMU_DATE.Time,
		// 	SMU_ETD:                      data.SMU_ETD.Time,
		// 	SMU_ETA:                      data.SMU_ETA.Time,
		// 	SMU_REMARKS:                  data.SMU_REMARKS.String,
		// 	SMU_REMARKS_DATE:             data.SMU_REMARKS_DATE.Time,
		// 	SMU_QTY:                      data.SMU_QTY.Float64,
		// 	SMU_WEIGHT:                   data.SMU_WEIGHT.Float64,
		// 	SMU_FLAG_APPROVE:             data.SMU_FLAG_APPROVE.String,
		// 	SMU_FLAG_CANCEL:              data.SMU_FLAG_CANCEL.String,
		// 	SMU_DESTINATION:              data.SMU_DESTINATION.String,
		// 	MANIFEST_TRS1_NO:             data.MANIFEST_TRS1_NO.String,
		// 	MANIFEST_TRS1_ORIGIN:         data.MANIFEST_TRS1_ORIGIN.String,
		// 	MANIFEST_TRS1_DATE:           data.MANIFEST_TRS1_DATE.Time,
		// 	MANIFEST_TRS1_BAG_NO:         data.MANIFEST_TRS1_BAG_NO.String,
		// 	MANIFEST_TRS1_USER_ID:        data.MANIFEST_TRS1_USER_ID.String,
		// 	MANIFEST_TRS1_USER_ZONE:      data.MANIFEST_TRS1_USER_ZONE.String,
		// 	MANIFEST_TRS1_CRDATE:         data.MANIFEST_TRS1_CRDATE.Time,
		// 	MANIFEST_TRSL_NO:             data.MANIFEST_TRSL_NO.String,
		// 	MANIFEST_TRSL_ORIGIN:         data.MANIFEST_TRSL_ORIGIN.String,
		// 	MANIFEST_TRSL_DATE:           data.MANIFEST_TRSL_DATE.Time,
		// 	MANIFEST_TRSL_BAG_NO:         data.MANIFEST_TRSL_BAG_NO.String,
		// 	MANIFEST_TRSL_USER_ID:        data.MANIFEST_TRSL_USER_ID.String,
		// 	MANIFEST_TRSL_USER_ZONE:      data.MANIFEST_TRSL_USER_ZONE.String,
		// 	MANIFEST_TRSL_CRDATE:         data.MANIFEST_TRSL_CRDATE.Time,
		// 	MANIFEST_INB_NO:              data.MANIFEST_INB_NO.String,
		// 	MANIFEST_INB_ORIGIN:          data.MANIFEST_INB_ORIGIN.String,
		// 	MANIFEST_INB_DATE:            data.MANIFEST_INB_DATE.Time,
		// 	MANIFEST_INB_BAG_NO:          data.MANIFEST_INB_BAG_NO.String,
		// 	MANIFEST_INB_USER_ID:         data.MANIFEST_INB_USER_ID.String,
		// 	MANIFEST_INB_USER_ZONE:       data.MANIFEST_INB_USER_ZONE.String,
		// 	MANIFEST_INB_CRDATE:          data.MANIFEST_INB_CRDATE.Time,
		// 	MANIFEST_BAG_NO:              data.MANIFEST_BAG_NO.String,
		// 	MANIFEST_BAG_DATE:            data.MANIFEST_BAG_DATE.Time,
		// 	MANIFEST_BAG_BAG_NO:          data.MANIFEST_BAG_BAG_NO.String,
		// 	MANIFEST_BAG_USER_ID:         data.MANIFEST_BAG_USER_ID.String,
		// 	MANIFEST_BAG_USER_ZONE:       data.MANIFEST_BAG_USER_ZONE.String,
		// 	MANIFEST_BAG_CRDATE:          data.MANIFEST_BAG_CRDATE.Time,
		// 	PRA_MRSHEET_NO:               data.PRA_MRSHEET_NO.String,
		// 	PRA_MRSHEET_DATE:             data.PRA_MRSHEET_DATE.Time,
		// 	PRA_MRSHEET_BRANCH:           data.PRA_MRSHEET_BRANCH.String,
		// 	PRA_MRSHEET_ZONE:             data.PRA_MRSHEET_ZONE.String,
		// 	PRA_MRSHEET_COURIER_ID:       data.PRA_MRSHEET_COURIER_ID.String,
		// 	PRA_COURIER_ZONE_CODE:        data.PRA_COURIER_ZONE_CODE.String,
		// 	PRA_MRSHEET_UID:              data.PRA_MRSHEET_UID.String,
		// 	PRA_USER_ZONE_CODE:           data.PRA_USER_ZONE_CODE.String,
		// 	PRA_CREATION_DATE:            data.PRA_CREATION_DATE.Time,
		// 	MTA_OUT_MANIFEST_NO:          data.MTA_OUT_MANIFEST_NO.String,
		// 	MTA_OUT_MANIFEST_DATE:        data.MTA_OUT_MANIFEST_DATE.Time,
		// 	MTA_OUT_BRANCH_ID:            data.MTA_OUT_BRANCH_ID.String,
		// 	MTA_OUT_DESTINATION:          data.MTA_OUT_DESTINATION.String,
		// 	MTA_OUT_MANIFEST_UID:         data.MTA_OUT_MANIFEST_UID.String,
		// 	MTA_OUT_USER_ZONE_CODE:       data.MTA_OUT_USER_ZONE_CODE.String,
		// 	MTA_OUT_ESB_TIME:             data.MTA_OUT_ESB_TIME.Time,
		// 	MTA_INB_MANIFEST_NO:          data.MTA_INB_MANIFEST_NO.String,
		// 	MTA_INB_MANIFEST_DATE:        data.MTA_INB_MANIFEST_DATE.Time,
		// 	MTA_INB_BRANCH_ID:            data.MTA_INB_BRANCH_ID.String,
		// 	MTA_INB_DESTINATION:          data.MTA_INB_DESTINATION.String,
		// 	MTA_INB_MANIFEST_UID:         data.MTA_INB_MANIFEST_UID.String,
		// 	MTA_INB_USER_ZONE_CODE:       data.MTA_INB_USER_ZONE_CODE.String,
		// 	MTA_INB_ESB_TIME:             data.MTA_INB_ESB_TIME.Time,
		// 	MHOCNOTE_NO:                  data.MHOCNOTE_NO.String,
		// 	MHOCNOTE_DATE:                data.MHOCNOTE_DATE.Time,
		// 	MHOCNOTE_BRANCH_ID:           data.MHOCNOTE_BRANCH_ID.String,
		// 	MHOCNOTE_ZONE:                data.MHOCNOTE_ZONE.String,
		// 	MHOCNOTE_ZONE_DEST:           data.MHOCNOTE_ZONE_DEST.String,
		// 	MHOCNOTE_USER_ID:             data.MHOCNOTE_USER_ID.String,
		// 	MHOCNOTE_USER_ZONE_CODE:      data.MHOCNOTE_USER_ZONE_CODE.String,
		// 	DHOCNOTE_TDATE:               data.DHOCNOTE_TDATE.Time,
		// 	MHICNOTE_NO:                  data.MHICNOTE_NO.String,
		// 	MHICNOTE_DATE:                data.MHICNOTE_DATE.Time,
		// 	MHICNOTE_BRANCH_ID:           data.MHICNOTE_BRANCH_ID.String,
		// 	MHICNOTE_ZONE:                data.MHICNOTE_ZONE.String,
		// 	MHICNOTE_USER_ID:             data.MHICNOTE_USER_ID.String,
		// 	MHICNOTE_USER_ZONE_CODE:      data.MHICNOTE_USER_ZONE_CODE.String,
		// 	DHICNOTE_TDATE:               data.DHICNOTE_TDATE.Time,
		// 	MRSHEET1_NO:                  data.MRSHEET1_NO.String,
		// 	MRSHEET1_DATE:                data.MRSHEET1_DATE.Time,
		// 	MRSHEET1_BRANCH:              data.MRSHEET1_BRANCH.String,
		// 	MRSHEET1_COURIER_ID:          data.MRSHEET1_COURIER_ID.String,
		// 	MRSHEET1_UID:                 data.MRSHEET1_UID.String,
		// 	MRSHEET1_USER_ZONE_CODE:      data.MRSHEET1_USER_ZONE_CODE.String,
		// 	MRSHEET1_CREATION_DATE:       data.MRSHEET1_CREATION_DATE.Time,
		// 	MRSHEETL_NO:                  data.MRSHEETL_NO.String,
		// 	MRSHEETL_DATE:                data.MRSHEETL_DATE.Time,
		// 	MRSHEETL_BRANCH:              data.MRSHEETL_BRANCH.String,
		// 	MRSHEETL_COURIER_ID:          data.MRSHEETL_COURIER_ID.String,
		// 	MRSHEETL_UID:                 data.MRSHEETL_UID.String,
		// 	MRSHEETL_USER_ZONE_CODE:      data.MRSHEETL_USER_ZONE_CODE.String,
		// 	MRSHEETL_CREATION_DATE:       data.MRSHEETL_CREATION_DATE.Time,
		// 	POD1_DRSHEET_NO:              data.POD1_DRSHEET_NO.String,
		// 	POD1_MRSHEET_DATE:            data.POD1_MRSHEET_DATE.Time,
		// 	POD1_MRSHEET_BRANCH:          data.POD1_MRSHEET_BRANCH.String,
		// 	POD1_MRSHEET_COURIER_ID:      data.POD1_MRSHEET_COURIER_ID.String,
		// 	POD1_COURIER_ZONE_CODE:       data.POD1_COURIER_ZONE_CODE.String,
		// 	POD1_DRSHEET_DATE:            data.POD1_DRSHEET_DATE.Time,
		// 	POD1_DRSHEET_RECEIVER:        data.POD1_DRSHEET_RECEIVER.String,
		// 	POD1_DRSHEET_STATUS:          data.POD1_DRSHEET_STATUS.String,
		// 	POD1_LATITUDE:                data.POD1_LATITUDE.String,
		// 	POD1_LONGITUDE:               data.POD1_LONGITUDE.String,
		// 	POD1_EPOD_URL:                data.POD1_EPOD_URL.String,
		// 	POD1_EPOD_URL_PIC:            data.POD1_EPOD_URL_PIC.String,
		// 	POD1_DRSHEET_UID:             data.POD1_DRSHEET_UID.String,
		// 	POD1_USER_ZONE_CODE:          data.POD1_USER_ZONE_CODE.String,
		// 	POD1_DRSHEET_UDATE:           data.POD1_DRSHEET_UDATE.Time,
		// 	PODL_DRSHEET_NO:              data.PODL_DRSHEET_NO.String,
		// 	PODL_MRSHEET_DATE:            data.PODL_MRSHEET_DATE.Time,
		// 	PODL_MRSHEET_BRANCH:          data.PODL_MRSHEET_BRANCH.String,
		// 	PODL_MRSHEET_COURIER_ID:      data.PODL_MRSHEET_COURIER_ID.String,
		// 	PODL_COURIER_ZONE_CODE:       data.PODL_COURIER_ZONE_CODE.String,
		// 	PODL_DRSHEET_DATE:            data.PODL_DRSHEET_DATE.Time,
		// 	PODL_DRSHEET_RECEIVER:        data.PODL_DRSHEET_RECEIVER.String,
		// 	PODL_DRSHEET_STATUS:          data.PODL_DRSHEET_STATUS.String,
		// 	PODL_LATITUDE:                data.PODL_LATITUDE.String,
		// 	PODL_LONGITUDE:               data.PODL_LONGITUDE.String,
		// 	PODL_EPOD_URL:                data.PODL_EPOD_URL.String,
		// 	PODL_EPOD_URL_PIC:            data.PODL_EPOD_URL_PIC.String,
		// 	PODL_DRSHEET_UID:             data.PODL_DRSHEET_UID.String,
		// 	PODL_USER_ZONE_CODE:          data.PODL_USER_ZONE_CODE.String,
		// 	PODL_DRSHEET_UDATE:           data.PODL_DRSHEET_UDATE.Time,
		// 	DO_NO:                        data.DO_NO.String,
		// 	DO_DATE:                      data.DO_DATE.Time,
		// 	RDO_NO:                       data.RDO_NO.String,
		// 	RDO_DATE:                     data.RDO_DATE.Time,
		// 	SHIPPER_PROVIDER:             data.SHIPPER_PROVIDER.String,
		// 	CNOTE_REFNO:                  data.CNOTE_REFNO.String,
		// 	MANIFEST_OUTB_APPROVED:       data.MANIFEST_OUTB_APPROVED.String,
		// 	MANIFEST_INB_APPROVED:        data.MANIFEST_INB_APPROVED.String,
		// 	SMU_BAG_BUX:                  data.SMU_BAG_BUX.String,
		// 	SMU_TGL_MASTER_BAG:           data.SMU_TGL_MASTER_BAG.Time,
		// 	SMU_USER_MASTER_BAG:          data.SMU_USER_MASTER_BAG.String,
		// 	SMU_NO_MASTER_BAG:            data.SMU_NO_MASTER_BAG.String,
		// 	SMU_MANIFEST_DESTINATION:     data.SMU_MANIFEST_DESTINATION.String,
		// 	MANIFEST_COST_WEIGHT:         data.MANIFEST_COST_WEIGHT.Float64,
		// 	MANIFEST_ACT_WEIGHT:          data.MANIFEST_ACT_WEIGHT.Float64,
		// 	DWH_PACKING_FEE:              data.DWH_PACKING_FEE.Float64,
		// 	DWH_SURCHARGE:                data.DWH_SURCHARGE.Float64,
		// 	DWH_DISC_REV_TYPE:            data.DWH_DISC_REV_TYPE.String,
		// 	DWH_DISCOUNT_AMT:             data.DWH_DISCOUNT_AMT.Float64,
		// 	DWH_FCHARGE_AFT_DISC_AMT:     data.DWH_FCHARGE_AFT_DISC_AMT.Float64,
		// 	DWH_CUST_DISC_IC:             data.DWH_CUST_DISC_IC.Float64,
		// 	DWH_CUST_DISC_DM:             data.DWH_CUST_DISC_DM.Float64,
		// 	DWH_RT_PACKING_FEE:           data.DWH_RT_PACKING_FEE.Float64,
		// 	DWH_RT_FREIGHT_CHARGE:        data.DWH_RT_FREIGHT_CHARGE.Float64,
		// 	DWH_RT_SURCHARGE:             data.DWH_RT_SURCHARGE.Float64,
		// 	DWH_RT_DISC_AMT:              data.DWH_RT_DISC_AMT.Float64,
		// 	DWH_RT_FCHARGE_AFT_DISC_AMT:  data.DWH_RT_FCHARGE_AFT_DISC_AMT.Float64,
		// 	DWH_PAYTYPE:                  data.DWH_PAYTYPE.String,
		// 	DWH_EPAY_VEND:                data.DWH_EPAY_VEND.String,
		// 	DWH_EPAY_TRXID:               data.DWH_EPAY_TRXID.String,
		// 	DWH_VAT_FCHARGE_AFT_DISC:     data.DWH_VAT_FCHARGE_AFT_DISC.Float64,
		// 	DWH_VAT_RT_FCHARGE_AFT_DISC:  data.DWH_VAT_RT_FCHARGE_AFT_DISC.Float64,
		// }
		_, err = col.Upsert(key, data, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("insert to Couchbase Success")
	return true
}
