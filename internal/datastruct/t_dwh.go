package datastruct

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type T_dwh struct {
	CREATE_DATE                  NullTime    `json:"CREATE_DATE" db:"CREATE_DATE"`
	CNOTE_NO                     string      `json:"CNOTE_NO" db:"CNOTE_NO"`
	CNOTE_DATE                   NullTime    `json:"CNOTE_DATE" db:"CNOTE_DATE"`
	CNOTE_CRDATE                 NullTime    `json:"CNOTE_CRDATE" db:"CNOTE_CRDATE"`
	CNOTE_BRANCH_ID              NullString  `json:"CNOTE_BRANCH_ID" db:"CNOTE_BRANCH_ID"`
	BRANCH_REGION                NullString  `json:"BRANCH_REGION" db:"BRANCH_REGION"`
	CNOTE_ORIGIN                 NullString  `json:"CNOTE_ORIGIN" db:"CNOTE_ORIGIN"`
	ORIGIN_NAME                  NullString  `json:"ORIGIN_NAME" db:"ORIGIN_NAME"`
	ORIGIN_ZONE                  NullString  `json:"ORIGIN_ZONE" db:"ORIGIN_ZONE"`
	CNOTE_CUST_NO                NullString  `json:"CNOTE_CUST_NO" db:"CNOTE_CUST_NO"`
	CNOTE_CUST_TYPE              NullString  `json:"CNOTE_CUST_TYPE" db:"CNOTE_CUST_TYPE"`
	CUST_NAME                    NullString  `json:"CUST_NAME" db:"CUST_NAME"`
	CUST_ADDR1                   NullString  `json:"CUST_ADDR1" db:"CUST_ADDR1"`
	CUST_ADDR2                   NullString  `json:"CUST_ADDR2" db:"CUST_ADDR2"`
	CUST_ADDR3                   NullString  `json:"CUST_ADDR3" db:"CUST_ADDR3"`
	CUST_PHONE                   NullString  `json:"CUST_PHONE" db:"CUST_PHONE"`
	CUST_ZIP                     NullString  `json:"CUST_ZIP" db:"CUST_ZIP"`
	CUST_NA                      NullString  `json:"CUST_NA" db:"CUST_NA"`
	MARKETPLACE_TYPE             NullString  `json:"MARKETPLACE_TYPE" db:"MARKETPLACE_TYPE"`
	MARKETPLACE_NAME             NullString  `json:"MARKETPLACE_NAME" db:"MARKETPLACE_NAME"`
	CNOTE_SHIPPER_NAME           NullString  `json:"CNOTE_SHIPPER_NAME" db:"CNOTE_SHIPPER_NAME"`
	CNOTE_SHIPPER_CONTACT        NullString  `json:"CNOTE_SHIPPER_CONTACT" db:"CNOTE_SHIPPER_CONTACT"`
	CNOTE_SHIPPER_ADDR1          NullString  `json:"CNOTE_SHIPPER_ADDR1" db:"CNOTE_SHIPPER_ADDR1"`
	CNOTE_SHIPPER_ADDR2          NullString  `json:"CNOTE_SHIPPER_ADDR2" db:"CNOTE_SHIPPER_ADDR2"`
	CNOTE_SHIPPER_ADDR3          NullString  `json:"CNOTE_SHIPPER_ADDR3" db:"CNOTE_SHIPPER_ADDR3"`
	CNOTE_SHIPPER_PHONE          NullString  `json:"CNOTE_SHIPPER_PHONE" db:"CNOTE_SHIPPER_PHONE"`
	CNOTE_SHIPPER_ZIP            NullString  `json:"CNOTE_SHIPPER_ZIP" db:"CNOTE_SHIPPER_ZIP"`
	CNOTE_RECEIVER_NAME          NullString  `json:"CNOTE_RECEIVER_NAME" db:"CNOTE_RECEIVER_NAME"`
	CNOTE_RECEIVER_CONTACT       NullString  `json:"CNOTE_RECEIVER_CONTACT" db:"CNOTE_RECEIVER_CONTACT"`
	CNOTE_RECEIVER_ADDR1         NullString  `json:"CNOTE_RECEIVER_ADDR1" db:"CNOTE_RECEIVER_ADDR1"`
	CNOTE_RECEIVER_ADDR2         NullString  `json:"CNOTE_RECEIVER_ADDR2" db:"CNOTE_RECEIVER_ADDR2"`
	CNOTE_RECEIVER_ADDR3         NullString  `json:"CNOTE_RECEIVER_ADDR3" db:"CNOTE_RECEIVER_ADDR3"`
	CNOTE_RECEIVER_PHONE         NullString  `json:"CNOTE_RECEIVER_PHONE" db:"CNOTE_RECEIVER_PHONE"`
	CNOTE_RECEIVER_ZIP           NullString  `json:"CNOTE_RECEIVER_ZIP" db:"CNOTE_RECEIVER_ZIP"`
	CNOTE_DESTINATION_ID         NullString  `json:"CNOTE_DESTINATION_ID" db:"CNOTE_DESTINATION_ID"`
	BRANCH_DEST_REGION           NullString  `json:"BRANCH_DEST_REGION" db:"BRANCH_DEST_REGION"`
	BRANCH_DESTINATION           NullString  `json:"BRANCH_DESTINATION" db:"BRANCH_DESTINATION"`
	CNOTE_DESTINATION            NullString  `json:"CNOTE_DESTINATION" db:"CNOTE_DESTINATION"`
	DESTINATION_NAME             NullString  `json:"DESTINATION_NAME" db:"DESTINATION_NAME"`
	DESTINATION_CODE             NullString  `json:"DESTINATION_CODE" db:"DESTINATION_CODE"`
	DESTINATION_ZONE             NullString  `json:"DESTINATION_ZONE" db:"DESTINATION_ZONE"`
	CNOTE_SERVICES_CODE          NullString  `json:"CNOTE_SERVICES_CODE" db:"CNOTE_SERVICES_CODE"`
	ROUTE_ETD_FROM               NullFloat64 `json:"ROUTE_ETD_FROM" db:"ROUTE_ETD_FROM"`
	ROUTE_ETD_THRU               NullFloat64 `json:"ROUTE_ETD_THRU" db:"ROUTE_ETD_THRU"`
	CNOTE_SHIPMENT_TYPE          NullString  `json:"CNOTE_SHIPMENT_TYPE" db:"CNOTE_SHIPMENT_TYPE"`
	CNOTE_TRX_TYPE               NullString  `json:"CNOTE_TRX_TYPE" db:"CNOTE_TRX_TYPE"`
	CNOTE_PAYMENT_TYPE           NullString  `json:"CNOTE_PAYMENT_TYPE" db:"CNOTE_PAYMENT_TYPE"`
	CNOTE_QTY                    NullFloat64 `json:"CNOTE_QTY" db:"CNOTE_QTY"`
	CNOTE_WEIGHT                 NullFloat64 `json:"CNOTE_WEIGHT" db:"CNOTE_WEIGHT"`
	CNOTE_DIM                    NullFloat64 `json:"CNOTE_DIM" db:"CNOTE_DIM"`
	CNOTE_GOODS_TYPE             NullString  `json:"CNOTE_GOODS_TYPE" db:"CNOTE_GOODS_TYPE"`
	CNOTE_PACKING                NullString  `json:"CNOTE_PACKING" db:"CNOTE_PACKING"`
	CNOTE_GOODS_DESCR            NullString  `json:"CNOTE_GOODS_DESCR" db:"CNOTE_GOODS_DESCR"`
	CNOTE_GOODS_VALUE            NullFloat64 `json:"CNOTE_GOODS_VALUE" db:"CNOTE_GOODS_VALUE"`
	CNOTE_SPECIAL_INS            NullString  `json:"CNOTE_SPECIAL_INS" db:"CNOTE_SPECIAL_INS"`
	CNOTE_INSURANCE_ID           NullString  `json:"CNOTE_INSURANCE_ID" db:"CNOTE_INSURANCE_ID"`
	CNOTE_INSURANCE_VALUE        NullFloat64 `json:"CNOTE_INSURANCE_VALUE" db:"CNOTE_INSURANCE_VALUE"`
	CNOTE_AMOUNT                 NullFloat64 `json:"CNOTE_AMOUNT" db:"CNOTE_AMOUNT"`
	CNOTE_ADDITIONAL_FEE         NullFloat64 `json:"CNOTE_ADDITIONAL_FEE" db:"CNOTE_ADDITIONAL_FEE"`
	CNOTE_COD                    NullString  `json:"CNOTE_COD" db:"CNOTE_COD"`
	COD_NO                       NullString  `json:"COD_NO" db:"COD_NO"`
	COD_GOODS_AMOUNT             NullFloat64 `json:"COD_GOODS_AMOUNT" db:"COD_GOODS_AMOUNT"`
	COD_AMOUNT                   NullFloat64 `json:"COD_AMOUNT" db:"COD_AMOUNT"`
	CNOTE_CASHLESS               NullString  `json:"CNOTE_CASHLESS" db:"CNOTE_CASHLESS"`
	JLC_NO                       NullString  `json:"JLC_NO" db:"JLC_NO"`
	JLC_NAME                     NullString  `json:"JLC_NAME" db:"JLC_NAME"`
	JLC_DISCOUNT                 NullFloat64 `json:"JLC_DISCOUNT" db:"JLC_DISCOUNT"`
	HYBRID_BRANCH                NullString  `json:"HYBRID_BRANCH" db:"HYBRID_BRANCH"`
	HYBRID_CUST_NO               NullString  `json:"HYBRID_CUST_NO" db:"HYBRID_CUST_NO"`
	HYBRID_CUST_NAME             NullString  `json:"HYBRID_CUST_NAME" db:"HYBRID_CUST_NAME"`
	HYBRID_CUST_ADDR1            NullString  `json:"HYBRID_CUST_ADDR1" db:"HYBRID_CUST_ADDR1"`
	HYBRID_CUST_ADDR2            NullString  `json:"HYBRID_CUST_ADDR2" db:"HYBRID_CUST_ADDR2"`
	HYBRID_CUST_ADDR3            NullString  `json:"HYBRID_CUST_ADDR3" db:"HYBRID_CUST_ADDR3"`
	HYBRID_CUST_PHONE            NullString  `json:"HYBRID_CUST_PHONE" db:"HYBRID_CUST_PHONE"`
	HYBRID_CUST_ZIP              NullString  `json:"HYBRID_CUST_ZIP" db:"HYBRID_CUST_ZIP"`
	CNOTE_CANCEL                 NullString  `json:"CNOTE_CANCEL" db:"CNOTE_CANCEL"`
	CNOTE_HOLD                   NullString  `json:"CNOTE_HOLD" db:"CNOTE_HOLD"`
	CNOTE_USER                   NullString  `json:"CNOTE_USER" db:"CNOTE_USER"`
	CNOTE_USER_ZONE              NullString  `json:"CNOTE_USER_ZONE" db:"CNOTE_USER_ZONE"`
	R_CNOTE_FREIGHT_CHARGE       NullFloat64 `json:"R_CNOTE_FREIGHT_CHARGE" db:"R_CNOTE_FREIGHT_CHARGE"`
	PUBLISH_RATE                 NullFloat64 `json:"PUBLISH_RATE" db:"PUBLISH_RATE"`
	CASHREG_NO                   NullString  `json:"CASHREG_NO" db:"CASHREG_NO"`
	CASHREG_DATE                 NullTime    `json:"CASHREG_DATE" db:"CASHREG_DATE"`
	CASHREG_USER_ID              NullString  `json:"CASHREG_USER_ID" db:"CASHREG_USER_ID"`
	CASHREG_USER_ZONE            NullString  `json:"CASHREG_USER_ZONE" db:"CASHREG_USER_ZONE"`
	CASHREG_CRDATE               NullTime    `json:"CASHREG_CRDATE" db:"CASHREG_CRDATE"`
	PICKUP_NO                    NullString  `json:"PICKUP_NO" db:"PICKUP_NO"`
	PICKUP_COURIER_ID            NullString  `json:"PICKUP_COURIER_ID" db:"PICKUP_COURIER_ID"`
	PICKUP_COURIER_ZONE          NullString  `json:"PICKUP_COURIER_ZONE" db:"PICKUP_COURIER_ZONE"`
	PICKUP_DATE                  NullTime    `json:"PICKUP_DATE" db:"PICKUP_DATE"`
	PICKUP_CRDATE                NullTime    `json:"PICKUP_CRDATE" db:"PICKUP_CRDATE"`
	PICKUP_MERCHAN_ID            NullString  `json:"PICKUP_MERCHAN_ID" db:"PICKUP_MERCHAN_ID"`
	PICKUP_LATITUDE              NullString  `json:"PICKUP_LATITUDE" db:"PICKUP_LATITUDE"`
	PICKUP_LONGITUDE             NullString  `json:"PICKUP_LONGITUDE" db:"PICKUP_LONGITUDE"`
	PU_FIRST_ATTTEMP_STATUS_CODE NullString  `json:"PU_FIRST_ATTTEMP_STATUS_CODE" db:"PU_FIRST_ATTTEMP_STATUS_CODE"`
	PU_FIRST_ATTTEMP_STATUS_DESC NullString  `json:"PU_FIRST_ATTTEMP_STATUS_DESC" db:"PU_FIRST_ATTTEMP_STATUS_DESC"`
	PU_FIRST_ATTTEMP_STATUS_DATE NullTime    `json:"PU_FIRST_ATTTEMP_STATUS_DATE" db:"PU_FIRST_ATTTEMP_STATUS_DATE"`
	PU_LAST_ATTEMP_STATUS_CODE   NullString  `json:"PU_LAST_ATTEMP_STATUS_CODE" db:"PU_LAST_ATTEMP_STATUS_CODE"`
	PU_LAST_ATTEMP_STATUS_DESC   NullString  `json:"PU_LAST_ATTEMP_STATUS_DESC" db:"PU_LAST_ATTEMP_STATUS_DESC"`
	PU_LAST_ATTEMP_STATUS_DATE   NullTime    `json:"PU_LAST_ATTEMP_STATUS_DATE" db:"PU_LAST_ATTEMP_STATUS_DATE"`
	PU_REF_ID                    NullString  `json:"PU_REF_ID" db:"PU_REF_ID"`
	HO_NO                        NullString  `json:"HO_NO" db:"HO_NO"`
	HO_DATE                      NullTime    `json:"HO_DATE" db:"HO_DATE"`
	HO_COURIER_ID                NullString  `json:"HO_COURIER_ID" db:"HO_COURIER_ID"`
	HO_CDATE                     NullTime    `json:"HO_CDATE" db:"HO_CDATE"`
	RECEIVING_AGENT_NO           NullString  `json:"RECEIVING_AGENT_NO" db:"RECEIVING_AGENT_NO"`
	RECEIVING_AGENT_DATE         NullTime    `json:"RECEIVING_AGENT_DATE" db:"RECEIVING_AGENT_DATE"`
	RECEIVING_AGENT_BRANCH       NullString  `json:"RECEIVING_AGENT_BRANCH" db:"RECEIVING_AGENT_BRANCH"`
	RECEIVING_AGENT_COURIER_ID   NullString  `json:"RECEIVING_AGENT_COURIER_ID" db:"RECEIVING_AGENT_COURIER_ID"`
	RECEIVING_AGENT_USER_ID      NullString  `json:"RECEIVING_AGENT_USER_ID" db:"RECEIVING_AGENT_USER_ID"`
	RECEIVING_AGENT_USER_ZONE    NullString  `json:"RECEIVING_AGENT_USER_ZONE" db:"RECEIVING_AGENT_USER_ZONE"`
	RECEIVING_AGENT_CRDATE       NullTime    `json:"RECEIVING_AGENT_CRDATE" db:"RECEIVING_AGENT_CRDATE"`
	RECEIVING_OUT_NO             NullString  `json:"RECEIVING_OUT_NO" db:"RECEIVING_OUT_NO"`
	RECEIVING_OUT_DATE           NullTime    `json:"RECEIVING_OUT_DATE" db:"RECEIVING_OUT_DATE"`
	RECEIVING_OUT_BRANCH         NullString  `json:"RECEIVING_OUT_BRANCH" db:"RECEIVING_OUT_BRANCH"`
	RECEIVING_OUT_COURIER_ID     NullString  `json:"RECEIVING_OUT_COURIER_ID" db:"RECEIVING_OUT_COURIER_ID"`
	RECEIVING_OUT_USER_ID        NullString  `json:"RECEIVING_OUT_USER_ID" db:"RECEIVING_OUT_USER_ID"`
	RECEIVING_OUT_USER_ZONE      NullString  `json:"RECEIVING_OUT_USER_ZONE" db:"RECEIVING_OUT_USER_ZONE"`
	RECEIVING_OUT_CRDATE         NullTime    `json:"RECEIVING_OUT_CRDATE" db:"RECEIVING_OUT_CRDATE"`
	MANIFEST_OUTB_NO             NullString  `json:"MANIFEST_OUTB_NO" db:"MANIFEST_OUTB_NO"`
	MANIFEST_OUTB_ORIGIN         NullString  `json:"MANIFEST_OUTB_ORIGIN" db:"MANIFEST_OUTB_ORIGIN"`
	MANIFEST_OUTB_DATE           NullTime    `json:"MANIFEST_OUTB_DATE" db:"MANIFEST_OUTB_DATE"`
	MANIFEST_OUTB_BAG_NO         NullString  `json:"MANIFEST_OUTB_BAG_NO" db:"MANIFEST_OUTB_BAG_NO"`
	MANIFEST_OUTB_USER_ID        NullString  `json:"MANIFEST_OUTB_USER_ID" db:"MANIFEST_OUTB_USER_ID"`
	MANIFEST_OUTB_USER_ZONE      NullString  `json:"MANIFEST_OUTB_USER_ZONE" db:"MANIFEST_OUTB_USER_ZONE"`
	MANIFEST_OUTB_CRDATE         NullTime    `json:"MANIFEST_OUTB_CRDATE" db:"MANIFEST_OUTB_CRDATE"`
	SMU_NO                       NullString  `json:"SMU_NO" db:"SMU_NO"`
	SMU_SCHD_NO                  NullString  `json:"SMU_SCHD_NO" db:"SMU_SCHD_NO"`
	SMU_SCH_DATE                 NullTime    `json:"SMU_SCH_DATE" db:"SMU_SCH_DATE"`
	SMU_DATE                     NullTime    `json:"SMU_DATE" db:"SMU_DATE"`
	SMU_ETD                      NullTime    `json:"SMU_ETD" db:"SMU_ETD"`
	SMU_ETA                      NullTime    `json:"SMU_ETA" db:"SMU_ETA"`
	SMU_REMARKS                  NullString  `json:"SMU_REMARKS" db:"SMU_REMARKS"`
	SMU_REMARKS_DATE             NullTime    `json:"SMU_REMARKS_DATE" db:"SMU_REMARKS_DATE"`
	SMU_QTY                      NullFloat64 `json:"SMU_QTY" db:"SMU_QTY"`
	SMU_WEIGHT                   NullFloat64 `json:"SMU_WEIGHT" db:"SMU_WEIGHT"`
	SMU_FLAG_APPROVE             NullString  `json:"SMU_FLAG_APPROVE" db:"SMU_FLAG_APPROVE"`
	SMU_FLAG_CANCEL              NullString  `json:"SMU_FLAG_CANCEL" db:"SMU_FLAG_CANCEL"`
	SMU_DESTINATION              NullString  `json:"SMU_DESTINATION" db:"SMU_DESTINATION"`
	MANIFEST_TRS1_NO             NullString  `json:"MANIFEST_TRS1_NO" db:"MANIFEST_TRS1_NO"`
	MANIFEST_TRS1_ORIGIN         NullString  `json:"MANIFEST_TRS1_ORIGIN" db:"MANIFEST_TRS1_ORIGIN"`
	MANIFEST_TRS1_DATE           NullTime    `json:"MANIFEST_TRS1_DATE" db:"MANIFEST_TRS1_DATE"`
	MANIFEST_TRS1_BAG_NO         NullString  `json:"MANIFEST_TRS1_BAG_NO" db:"MANIFEST_TRS1_BAG_NO"`
	MANIFEST_TRS1_USER_ID        NullString  `json:"MANIFEST_TRS1_USER_ID" db:"MANIFEST_TRS1_USER_ID"`
	MANIFEST_TRS1_USER_ZONE      NullString  `json:"MANIFEST_TRS1_USER_ZONE" db:"MANIFEST_TRS1_USER_ZONE"`
	MANIFEST_TRS1_CRDATE         NullTime    `json:"MANIFEST_TRS1_CRDATE" db:"MANIFEST_TRS1_CRDATE"`
	MANIFEST_TRSL_NO             NullString  `json:"MANIFEST_TRSL_NO" db:"MANIFEST_TRSL_NO"`
	MANIFEST_TRSL_ORIGIN         NullString  `json:"MANIFEST_TRSL_ORIGIN" db:"MANIFEST_TRSL_ORIGIN"`
	MANIFEST_TRSL_DATE           NullTime    `json:"MANIFEST_TRSL_DATE" db:"MANIFEST_TRSL_DATE"`
	MANIFEST_TRSL_BAG_NO         NullString  `json:"MANIFEST_TRSL_BAG_NO" db:"MANIFEST_TRSL_BAG_NO"`
	MANIFEST_TRSL_USER_ID        NullString  `json:"MANIFEST_TRSL_USER_ID" db:"MANIFEST_TRSL_USER_ID"`
	MANIFEST_TRSL_USER_ZONE      NullString  `json:"MANIFEST_TRSL_USER_ZONE" db:"MANIFEST_TRSL_USER_ZONE"`
	MANIFEST_TRSL_CRDATE         NullTime    `json:"MANIFEST_TRSL_CRDATE" db:"MANIFEST_TRSL_CRDATE"`
	MANIFEST_INB_NO              NullString  `json:"MANIFEST_INB_NO" db:"MANIFEST_INB_NO"`
	MANIFEST_INB_ORIGIN          NullString  `json:"MANIFEST_INB_ORIGIN" db:"MANIFEST_INB_ORIGIN"`
	MANIFEST_INB_DATE            NullTime    `json:"MANIFEST_INB_DATE" db:"MANIFEST_INB_DATE"`
	MANIFEST_INB_BAG_NO          NullString  `json:"MANIFEST_INB_BAG_NO" db:"MANIFEST_INB_BAG_NO"`
	MANIFEST_INB_USER_ID         NullString  `json:"MANIFEST_INB_USER_ID" db:"MANIFEST_INB_USER_ID"`
	MANIFEST_INB_USER_ZONE       NullString  `json:"MANIFEST_INB_USER_ZONE" db:"MANIFEST_INB_USER_ZONE"`
	MANIFEST_INB_CRDATE          NullTime    `json:"MANIFEST_INB_CRDATE" db:"MANIFEST_INB_CRDATE"`
	MANIFEST_BAG_NO              NullString  `json:"MANIFEST_BAG_NO" db:"MANIFEST_BAG_NO"`
	MANIFEST_BAG_DATE            NullTime    `json:"MANIFEST_BAG_DATE" db:"MANIFEST_BAG_DATE"`
	MANIFEST_BAG_BAG_NO          NullString  `json:"MANIFEST_BAG_BAG_NO" db:"MANIFEST_BAG_BAG_NO"`
	MANIFEST_BAG_USER_ID         NullString  `json:"MANIFEST_BAG_USER_ID" db:"MANIFEST_BAG_USER_ID"`
	MANIFEST_BAG_USER_ZONE       NullString  `json:"MANIFEST_BAG_USER_ZONE" db:"MANIFEST_BAG_USER_ZONE"`
	MANIFEST_BAG_CRDATE          NullTime    `json:"MANIFEST_BAG_CRDATE" db:"MANIFEST_BAG_CRDATE"`
	PRA_MRSHEET_NO               NullString  `json:"PRA_MRSHEET_NO" db:"PRA_MRSHEET_NO"`
	PRA_MRSHEET_DATE             NullTime    `json:"PRA_MRSHEET_DATE" db:"PRA_MRSHEET_DATE"`
	PRA_MRSHEET_BRANCH           NullString  `json:"PRA_MRSHEET_BRANCH" db:"PRA_MRSHEET_BRANCH"`
	PRA_MRSHEET_ZONE             NullString  `json:"PRA_MRSHEET_ZONE" db:"PRA_MRSHEET_ZONE"`
	PRA_MRSHEET_COURIER_ID       NullString  `json:"PRA_MRSHEET_COURIER_ID" db:"PRA_MRSHEET_COURIER_ID"`
	PRA_COURIER_ZONE_CODE        NullString  `json:"PRA_COURIER_ZONE_CODE" db:"PRA_COURIER_ZONE_CODE"`
	PRA_MRSHEET_UID              NullString  `json:"PRA_MRSHEET_UID" db:"PRA_MRSHEET_UID"`
	PRA_USER_ZONE_CODE           NullString  `json:"PRA_USER_ZONE_CODE" db:"PRA_USER_ZONE_CODE"`
	PRA_CREATION_DATE            NullTime    `json:"PRA_CREATION_DATE" db:"PRA_CREATION_DATE"`
	MTA_OUT_MANIFEST_NO          NullString  `json:"MTA_OUT_MANIFEST_NO" db:"MTA_OUT_MANIFEST_NO"`
	MTA_OUT_MANIFEST_DATE        NullTime    `json:"MTA_OUT_MANIFEST_DATE" db:"MTA_OUT_MANIFEST_DATE"`
	MTA_OUT_BRANCH_ID            NullString  `json:"MTA_OUT_BRANCH_ID" db:"MTA_OUT_BRANCH_ID"`
	MTA_OUT_DESTINATION          NullString  `json:"MTA_OUT_DESTINATION" db:"MTA_OUT_DESTINATION"`
	MTA_OUT_MANIFEST_UID         NullString  `json:"MTA_OUT_MANIFEST_UID" db:"MTA_OUT_MANIFEST_UID"`
	MTA_OUT_USER_ZONE_CODE       NullString  `json:"MTA_OUT_USER_ZONE_CODE" db:"MTA_OUT_USER_ZONE_CODE"`
	MTA_OUT_ESB_TIME             NullTime    `json:"MTA_OUT_ESB_TIME" db:"MTA_OUT_ESB_TIME"`
	MTA_INB_MANIFEST_NO          NullString  `json:"MTA_INB_MANIFEST_NO" db:"MTA_INB_MANIFEST_NO"`
	MTA_INB_MANIFEST_DATE        NullTime    `json:"MTA_INB_MANIFEST_DATE" db:"MTA_INB_MANIFEST_DATE"`
	MTA_INB_BRANCH_ID            NullString  `json:"MTA_INB_BRANCH_ID" db:"MTA_INB_BRANCH_ID"`
	MTA_INB_DESTINATION          NullString  `json:"MTA_INB_DESTINATION" db:"MTA_INB_DESTINATION"`
	MTA_INB_MANIFEST_UID         NullString  `json:"MTA_INB_MANIFEST_UID" db:"MTA_INB_MANIFEST_UID"`
	MTA_INB_USER_ZONE_CODE       NullString  `json:"MTA_INB_USER_ZONE_CODE" db:"MTA_INB_USER_ZONE_CODE"`
	MTA_INB_ESB_TIME             NullTime    `json:"MTA_INB_ESB_TIME" db:"MTA_INB_ESB_TIME"`
	MHOCNOTE_NO                  NullString  `json:"MHOCNOTE_NO" db:"MHOCNOTE_NO"`
	MHOCNOTE_DATE                NullTime    `json:"MHOCNOTE_DATE" db:"MHOCNOTE_DATE"`
	MHOCNOTE_BRANCH_ID           NullString  `json:"MHOCNOTE_BRANCH_ID" db:"MHOCNOTE_BRANCH_ID"`
	MHOCNOTE_ZONE                NullString  `json:"MHOCNOTE_ZONE" db:"MHOCNOTE_ZONE"`
	MHOCNOTE_ZONE_DEST           NullString  `json:"MHOCNOTE_ZONE_DEST" db:"MHOCNOTE_ZONE_DEST"`
	MHOCNOTE_USER_ID             NullString  `json:"MHOCNOTE_USER_ID" db:"MHOCNOTE_USER_ID"`
	MHOCNOTE_USER_ZONE_CODE      NullString  `json:"MHOCNOTE_USER_ZONE_CODE" db:"MHOCNOTE_USER_ZONE_CODE"`
	DHOCNOTE_TDATE               NullTime    `json:"DHOCNOTE_TDATE" db:"DHOCNOTE_TDATE"`
	MHICNOTE_NO                  NullString  `json:"MHICNOTE_NO" db:"MHICNOTE_NO"`
	MHICNOTE_DATE                NullTime    `json:"MHICNOTE_DATE" db:"MHICNOTE_DATE"`
	MHICNOTE_BRANCH_ID           NullString  `json:"MHICNOTE_BRANCH_ID" db:"MHICNOTE_BRANCH_ID"`
	MHICNOTE_ZONE                NullString  `json:"MHICNOTE_ZONE" db:"MHICNOTE_ZONE"`
	MHICNOTE_USER_ID             NullString  `json:"MHICNOTE_USER_ID" db:"MHICNOTE_USER_ID"`
	MHICNOTE_USER_ZONE_CODE      NullString  `json:"MHICNOTE_USER_ZONE_CODE" db:"MHICNOTE_USER_ZONE_CODE"`
	DHICNOTE_TDATE               NullTime    `json:"DHICNOTE_TDATE" db:"DHICNOTE_TDATE"`
	MRSHEET1_NO                  NullString  `json:"MRSHEET1_NO" db:"MRSHEET1_NO"`
	MRSHEET1_DATE                NullTime    `json:"MRSHEET1_DATE" db:"MRSHEET1_DATE"`
	MRSHEET1_BRANCH              NullString  `json:"MRSHEET1_BRANCH" db:"MRSHEET1_BRANCH"`
	MRSHEET1_COURIER_ID          NullString  `json:"MRSHEET1_COURIER_ID" db:"MRSHEET1_COURIER_ID"`
	MRSHEET1_UID                 NullString  `json:"MRSHEET1_UID" db:"MRSHEET1_UID"`
	MRSHEET1_USER_ZONE_CODE      NullString  `json:"MRSHEET1_USER_ZONE_CODE" db:"MRSHEET1_USER_ZONE_CODE"`
	MRSHEET1_CREATION_DATE       NullTime    `json:"MRSHEET1_CREATION_DATE" db:"MRSHEET1_CREATION_DATE"`
	MRSHEETL_NO                  NullString  `json:"MRSHEETL_NO" db:"MRSHEETL_NO"`
	MRSHEETL_DATE                NullTime    `json:"MRSHEETL_DATE" db:"MRSHEETL_DATE"`
	MRSHEETL_BRANCH              NullString  `json:"MRSHEETL_BRANCH" db:"MRSHEETL_BRANCH"`
	MRSHEETL_COURIER_ID          NullString  `json:"MRSHEETL_COURIER_ID" db:"MRSHEETL_COURIER_ID"`
	MRSHEETL_UID                 NullString  `json:"MRSHEETL_UID" db:"MRSHEETL_UID"`
	MRSHEETL_USER_ZONE_CODE      NullString  `json:"MRSHEETL_USER_ZONE_CODE" db:"MRSHEETL_USER_ZONE_CODE"`
	MRSHEETL_CREATION_DATE       NullTime    `json:"MRSHEETL_CREATION_DATE" db:"MRSHEETL_CREATION_DATE"`
	POD1_DRSHEET_NO              NullString  `json:"POD1_DRSHEET_NO" db:"POD1_DRSHEET_NO"`
	POD1_MRSHEET_DATE            NullTime    `json:"POD1_MRSHEET_DATE" db:"POD1_MRSHEET_DATE"`
	POD1_MRSHEET_BRANCH          NullString  `json:"POD1_MRSHEET_BRANCH" db:"POD1_MRSHEET_BRANCH"`
	POD1_MRSHEET_COURIER_ID      NullString  `json:"POD1_MRSHEET_COURIER_ID" db:"POD1_MRSHEET_COURIER_ID"`
	POD1_COURIER_ZONE_CODE       NullString  `json:"POD1_COURIER_ZONE_CODE" db:"POD1_COURIER_ZONE_CODE"`
	POD1_DRSHEET_DATE            NullTime    `json:"POD1_DRSHEET_DATE" db:"POD1_DRSHEET_DATE"`
	POD1_DRSHEET_RECEIVER        NullString  `json:"POD1_DRSHEET_RECEIVER" db:"POD1_DRSHEET_RECEIVER"`
	POD1_DRSHEET_STATUS          NullString  `json:"POD1_DRSHEET_STATUS" db:"POD1_DRSHEET_STATUS"`
	POD1_LATITUDE                NullString  `json:"POD1_LATITUDE" db:"POD1_LATITUDE"`
	POD1_LONGITUDE               NullString  `json:"POD1_LONGITUDE" db:"POD1_LONGITUDE"`
	POD1_EPOD_URL                NullString  `json:"POD1_EPOD_URL" db:"POD1_EPOD_URL"`
	POD1_EPOD_URL_PIC            NullString  `json:"POD1_EPOD_URL_PIC" db:"POD1_EPOD_URL_PIC"`
	POD1_DRSHEET_UID             NullString  `json:"POD1_DRSHEET_UID" db:"POD1_DRSHEET_UID"`
	POD1_USER_ZONE_CODE          NullString  `json:"POD1_USER_ZONE_CODE" db:"POD1_USER_ZONE_CODE"`
	POD1_DRSHEET_UDATE           NullTime    `json:"POD1_DRSHEET_UDATE" db:"POD1_DRSHEET_UDATE"`
	PODL_DRSHEET_NO              NullString  `json:"PODL_DRSHEET_NO" db:"PODL_DRSHEET_NO"`
	PODL_MRSHEET_DATE            NullTime    `json:"PODL_MRSHEET_DATE" db:"PODL_MRSHEET_DATE"`
	PODL_MRSHEET_BRANCH          NullString  `json:"PODL_MRSHEET_BRANCH" db:"PODL_MRSHEET_BRANCH"`
	PODL_MRSHEET_COURIER_ID      NullString  `json:"PODL_MRSHEET_COURIER_ID" db:"PODL_MRSHEET_COURIER_ID"`
	PODL_COURIER_ZONE_CODE       NullString  `json:"PODL_COURIER_ZONE_CODE" db:"PODL_COURIER_ZONE_CODE"`
	PODL_DRSHEET_DATE            NullTime    `json:"PODL_DRSHEET_DATE" db:"PODL_DRSHEET_DATE"`
	PODL_DRSHEET_RECEIVER        NullString  `json:"PODL_DRSHEET_RECEIVER" db:"PODL_DRSHEET_RECEIVER"`
	PODL_DRSHEET_STATUS          NullString  `json:"PODL_DRSHEET_STATUS" db:"PODL_DRSHEET_STATUS"`
	PODL_LATITUDE                NullString  `json:"PODL_LATITUDE" db:"PODL_LATITUDE"`
	PODL_LONGITUDE               NullString  `json:"PODL_LONGITUDE" db:"PODL_LONGITUDE"`
	PODL_EPOD_URL                NullString  `json:"PODL_EPOD_URL" db:"PODL_EPOD_URL"`
	PODL_EPOD_URL_PIC            NullString  `json:"PODL_EPOD_URL_PIC" db:"PODL_EPOD_URL_PIC"`
	PODL_DRSHEET_UID             NullString  `json:"PODL_DRSHEET_UID" db:"PODL_DRSHEET_UID"`
	PODL_USER_ZONE_CODE          NullString  `json:"PODL_USER_ZONE_CODE" db:"PODL_USER_ZONE_CODE"`
	PODL_DRSHEET_UDATE           NullTime    `json:"PODL_DRSHEET_UDATE" db:"PODL_DRSHEET_UDATE"`
	DO_NO                        NullString  `json:"DO_NO" db:"DO_NO"`
	DO_DATE                      NullTime    `json:"DO_DATE" db:"DO_DATE"`
	RDO_NO                       NullString  `json:"RDO_NO" db:"RDO_NO"`
	RDO_DATE                     NullTime    `json:"RDO_DATE" db:"RDO_DATE"`
	SHIPPER_PROVIDER             NullString  `json:"SHIPPER_PROVIDER" db:"SHIPPER_PROVIDER"`
	CNOTE_REFNO                  NullString  `json:"CNOTE_REFNO" db:"CNOTE_REFNO"`
	MANIFEST_OUTB_APPROVED       NullString  `json:"MANIFEST_OUTB_APPROVED" db:"MANIFEST_OUTB_APPROVED"`
	MANIFEST_INB_APPROVED        NullString  `json:"MANIFEST_INB_APPROVED" db:"MANIFEST_INB_APPROVED"`
	SMU_BAG_BUX                  NullString  `json:"SMU_BAG_BUX" db:"SMU_BAG_BUX"`
	SMU_TGL_MASTER_BAG           NullTime    `json:"SMU_TGL_MASTER_BAG" db:"SMU_TGL_MASTER_BAG"`
	SMU_USER_MASTER_BAG          NullString  `json:"SMU_USER_MASTER_BAG" db:"SMU_USER_MASTER_BAG"`
	SMU_NO_MASTER_BAG            NullString  `json:"SMU_NO_MASTER_BAG" db:"SMU_NO_MASTER_BAG"`
	SMU_MANIFEST_DESTINATION     NullString  `json:"SMU_MANIFEST_DESTINATION" db:"SMU_MANIFEST_DESTINATION"`
	MANIFEST_COST_WEIGHT         NullFloat64 `json:"MANIFEST_COST_WEIGHT" db:"MANIFEST_COST_WEIGHT"`
	MANIFEST_ACT_WEIGHT          NullFloat64 `json:"MANIFEST_ACT_WEIGHT" db:"MANIFEST_ACT_WEIGHT"`
	DWH_PACKING_FEE              NullFloat64 `json:"DWH_PACKING_FEE" db:"DWH_PACKING_FEE"`
	DWH_SURCHARGE                NullFloat64 `json:"DWH_SURCHARGE" db:"DWH_SURCHARGE"`
	DWH_DISC_REV_TYPE            NullString  `json:"DWH_DISC_REV_TYPE" db:"DWH_DISC_REV_TYPE"`
	DWH_DISCOUNT_AMT             NullFloat64 `json:"DWH_DISCOUNT_AMT" db:"DWH_DISCOUNT_AMT"`
	DWH_FCHARGE_AFT_DISC_AMT     NullFloat64 `json:"DWH_FCHARGE_AFT_DISC_AMT" db:"DWH_FCHARGE_AFT_DISC_AMT"`
	DWH_CUST_DISC_IC             NullFloat64 `json:"DWH_CUST_DISC_IC" db:"DWH_CUST_DISC_IC"`
	DWH_CUST_DISC_DM             NullFloat64 `json:"DWH_CUST_DISC_DM" db:"DWH_CUST_DISC_DM"`
	DWH_RT_PACKING_FEE           NullFloat64 `json:"DWH_RT_PACKING_FEE" db:"DWH_RT_PACKING_FEE"`
	DWH_RT_FREIGHT_CHARGE        NullFloat64 `json:"DWH_RT_FREIGHT_CHARGE" db:"DWH_RT_FREIGHT_CHARGE"`
	DWH_RT_SURCHARGE             NullFloat64 `json:"DWH_RT_SURCHARGE" db:"DWH_RT_SURCHARGE"`
	DWH_RT_DISC_AMT              NullFloat64 `json:"DWH_RT_DISC_AMT" db:"DWH_RT_DISC_AMT"`
	DWH_RT_FCHARGE_AFT_DISC_AMT  NullFloat64 `json:"DWH_RT_FCHARGE_AFT_DISC_AMT" db:"DWH_RT_FCHARGE_AFT_DISC_AMT"`
	DWH_PAYTYPE                  NullString  `json:"DWH_PAYTYPE" db:"DWH_PAYTYPE"`
	DWH_EPAY_VEND                NullString  `json:"DWH_EPAY_VEND" db:"DWH_EPAY_VEND"`
	DWH_EPAY_TRXID               NullString  `json:"DWH_EPAY_TRXID" db:"DWH_EPAY_TRXID"`
	DWH_VAT_FCHARGE_AFT_DISC     NullFloat64 `json:"DWH_VAT_FCHARGE_AFT_DISC" db:"DWH_VAT_FCHARGE_AFT_DISC"`
	DWH_VAT_RT_FCHARGE_AFT_DISC  NullFloat64 `json:"DWH_VAT_RT_FCHARGE_AFT_DISC" db:"DWH_VAT_RT_FCHARGE_AFT_DISC"`
}

type NullTime struct {
	sql.NullTime
}
type NullString struct {
	sql.NullString
}
type NullFloat64 struct {
	sql.NullFloat64
}

func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nf.Float64)
}

func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(time.RFC3339))
	return []byte(val), nil
}

type CMS_RETURN_REQCNOTE struct {
	REQCNOTE_CNOTE_NO    string
	REQCNOTE_CNOTE_NO_RT string
	REQCNOTE_HRS_NO      string
	REQCNOTE_REQ_USER    string
	REQCNOTE_AMOUNT1     string
	REQCNOTE_AMOUNT2     string
	REQCNOTE_AMOUNT3     string
	REQCNOTE_CDATE       string
}

type Dwh_and_cms struct {
	CNOTE_NO                   string      `json:"CNOTE_NO" db:"CNOTE_NO"`
	CNOTE_DATE                 NullTime    `json:"CNOTE_DATE" db:"CNOTE_DATE"`
	CNOTE_BRANCH_ID            NullString  `json:"CNOTE_BRANCH_ID" db:"CNOTE_BRANCH_ID"`
	BRANCH_REGION              NullString  `json:"BRANCH_REGION" db:"BRANCH_REGION"`
	CNOTE_ORIGIN               NullString  `json:"CNOTE_ORIGIN" db:"CNOTE_ORIGIN"`
	ORIGIN_NAME                NullString  `json:"ORIGIN_NAME" db:"ORIGIN_NAME"`
	ORIGIN_ZONE                NullString  `json:"ORIGIN_ZONE" db:"ORIGIN_ZONE"`
	CNOTE_CUST_NO              NullString  `json:"CNOTE_CUST_NO" db:"CNOTE_CUST_NO"`
	CUST_NAME                  NullString  `json:"CUST_NAME" db:"CUST_NAME"`
	MARKETPLACE_NAME           NullString  `json:"MARKETPLACE_NAME" db:"MARKETPLACE_NAME"`
	CNOTE_SHIPPER_NAME         NullString  `json:"CNOTE_SHIPPER_NAME" db:"CNOTE_SHIPPER_NAME"`
	CNOTE_RECEIVER_NAME        NullString  `json:"CNOTE_RECEIVER_NAME" db:"CNOTE_RECEIVER_NAME"`
	CNOTE_DESTINATION_ID       NullString  `json:"CNOTE_DESTINATION_ID" db:"CNOTE_DESTINATION_ID"`
	BRANCH_DEST_REGION         NullString  `json:"BRANCH_DEST_REGION" db:"BRANCH_DEST_REGION"`
	BRANCH_DESTINATION         NullString  `json:"BRANCH_DESTINATION" db:"BRANCH_DESTINATION"`
	CNOTE_DESTINATION          NullString  `json:"CNOTE_DESTINATION" db:"CNOTE_DESTINATION"`
	DESTINATION_NAME           NullString  `json:"DESTINATION_NAME" db:"DESTINATION_NAME"`
	DESTINATION_CODE           NullString  `json:"DESTINATION_CODE" db:"DESTINATION_CODE"`
	DESTINATION_ZONE           NullString  `json:"DESTINATION_ZONE" db:"DESTINATION_ZONE"`
	CNOTE_SERVICES_CODE        NullString  `json:"CNOTE_SERVICES_CODE" db:"CNOTE_SERVICES_CODE"`
	ROUTE_ETD_FROM             NullFloat64 `json:"ROUTE_ETD_FROM" db:"ROUTE_ETD_FROM"`
	ROUTE_ETD_THRU             NullFloat64 `json:"ROUTE_ETD_THRU" db:"ROUTE_ETD_THRU"`
	CNOTE_SHIPMENT_TYPE        NullString  `json:"CNOTE_SHIPMENT_TYPE" db:"CNOTE_SHIPMENT_TYPE"`
	CNOTE_GOODS_DESCR          NullString  `json:"CNOTE_GOODS_DESCR" db:"CNOTE_GOODS_DESCR"`
	CNOTE_COD                  NullString  `json:"CNOTE_COD" db:"CNOTE_COD"`
	COD_NO                     NullString  `json:"COD_NO" db:"COD_NO"`
	CNOTE_CASHLESS             NullString  `json:"CNOTE_CASHLESS" db:"CNOTE_CASHLESS"`
	JLC_NO                     NullString  `json:"JLC_NO" db:"JLC_NO"`
	HYBRID_BRANCH              NullString  `json:"HYBRID_BRANCH" db:"HYBRID_BRANCH"`
	HYBRID_CUST_NO             NullString  `json:"HYBRID_CUST_NO" db:"HYBRID_CUST_NO"`
	CNOTE_USER                 NullString  `json:"CNOTE_USER" db:"CNOTE_USER"`
	CNOTE_USER_ZONE            NullString  `json:"CNOTE_USER_ZONE" db:"CNOTE_USER_ZONE"`
	PICKUP_NO                  NullString  `json:"PICKUP_NO" db:"PICKUP_NO"`
	PICKUP_DATE                NullTime    `json:"PICKUP_DATE" db:"PICKUP_DATE"`
	PICKUP_COURIER_ID          NullString  `json:"PICKUP_COURIER_ID" db:"PICKUP_COURIER_ID"`
	PICKUP_COURIER_ZONE        NullString  `json:"PICKUP_COURIER_ZONE" db:"PICKUP_COURIER_ZONE"`
	PICKUP_MERCHAN_ID          NullString  `json:"PICKUP_MERCHAN_ID" db:"PICKUP_MERCHAN_ID"`
	PICKUP_LATITUDE            NullString  `json:"PICKUP_LATITUDE" db:"PICKUP_LATITUDE"`
	PICKUP_LONGITUDE           NullString  `json:"PICKUP_LONGITUDE" db:"PICKUP_LONGITUDE"`
	RECEIVING_AGENT_NO         NullString  `json:"RECEIVING_AGENT_NO" db:"RECEIVING_AGENT_NO"`
	RECEIVING_AGENT_CRDATE     NullTime    `json:"RECEIVING_AGENT_CRDATE" db:"RECEIVING_AGENT_CRDATE"`
	RECEIVING_OUT_CRDATE       NullTime    `json:"RECEIVING_OUT_CRDATE" db:"RECEIVING_OUT_CRDATE"`
	RECEIVING_AGENT_BRANCH     NullString  `json:"RECEIVING_AGENT_BRANCH" db:"RECEIVING_AGENT_BRANCH"`
	RECEIVING_AGENT_COURIER_ID NullString  `json:"RECEIVING_AGENT_COURIER_ID" db:"RECEIVING_AGENT_COURIER_ID"`
	RECEIVING_AGENT_USER_ID    NullString  `json:"RECEIVING_AGENT_USER_ID" db:"RECEIVING_AGENT_USER_ID"`
	MANIFEST_OUTB_USER_ZONE    NullString  `json:"MANIFEST_OUTB_USER_ZONE" db:"MANIFEST_OUTB_USER_ZONE"`
	RECEIVING_OUT_NO           NullString  `json:"RECEIVING_OUT_NO" db:"RECEIVING_OUT_NO"`
	RECEIVING_OUT_BRANCH       NullString  `json:"RECEIVING_OUT_BRANCH" db:"RECEIVING_OUT_BRANCH"`
	RECEIVING_OUT_COURIER_ID   NullString  `json:"RECEIVING_OUT_COURIER_ID" db:"RECEIVING_OUT_COURIER_ID"`
	RECEIVING_OUT_USER_ID      NullString  `json:"RECEIVING_OUT_USER_ID" db:"RECEIVING_OUT_USER_ID"`
	RECEIVING_OUT_USER_ZONE    NullString  `json:"RECEIVING_OUT_USER_ZONE" db:"RECEIVING_OUT_USER_ZONE"`
}
