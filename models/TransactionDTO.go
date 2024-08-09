package models

type TransactionDTO struct {
	TRX_ID     string `json:"trx_ID"`
	PRODUCT_ID uint   `json:"product_ID"`
	USER_ID    uint   `json:"user_ID"`
	STATUS     int    `json:"status"`
	TOTAL      int    `json:"total"`
	PAY_DATE   string `json:"pay_DATE"`
	PAY_TYPE   string `json:"pay_TYPE"`
}
