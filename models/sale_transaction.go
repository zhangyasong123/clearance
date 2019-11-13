package models

import (
	"context"
	"errors"
	"time"

	"clearance/clearance-adapter-for-sale-record/factory"

	"github.com/go-xorm/xorm"
)

type SaleTransaction struct {
	Id                     int64                    `json:"id"`
	TransactionId          int64                    `json:"transactionId" xorm:"index" validate:"required"`
	OrderId                int64                    `json:"orderId" xorm:"index default 0" validate:"required"`
	RefundId               int64                    `json:"refundId" xorm:"index default 0" `
	EmpId                  string                   `json:"empId" xorm:"index VARCHAR(50)"`
	StoreId                int64                    `json:"storeId" xorm:"index default 0" validate:"required"`
	ShopCode               string                   `json:"shopCode" xorm:"index VARCHAR(30) notnull" validate:"required"`
	SalesmanId             int64                    `json:"salesmanId" xorm:"index default 0" validate:"required"`
	CustomerId             int64                    `json:"customerId" xorm:"index default 0" validate:"required"`
	TransactionCreatedId   int64                    `json:"transactionCreatedId" xorm:"index default 0" validate:"required"`
	TotalListPrice         float64                  `json:"totalListPrice" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
	TotalSalePrice         float64                  `json:"totalSalePrice" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
	TotalTransactionPrice  float64                  `json:"totalTransactionPrice" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
	TotalDiscountPrice     float64                  `json:"totalDiscountPrice" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
	SaleDate               time.Time                `json:"saleDate"`
	Mileage                float64                  `json:"mileage" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
	MileagePrice           float64                  `json:"mileagePrice" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
	ObtainMileage          float64                  `json:"obtainMileage" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
	OuterOrderNo           string                   `json:"outerOrderNo" xorm:"index VARCHAR(30) notnull" validate:"required"`
	TransactionChannelType string                   `json:"transactionChannelType" xorm:"index VARCHAR(30) notnull"`
	BaseTrimCode           string                   `json:"baseTrimCode" xorm:"index VARCHAR(30)"`
	Dtls                   []SaleTransactionDtl     `json:"dtls" xorm:"-"`
	Payments               []SaleTransactionPayment `json:"payments" xorm:"-"`
	WhetherSend            bool                     `json:"whetherSend" xorm:"index default false"`
}

type SaleTransactionDtl struct {
	Id                             int64   `json:"id"`
	SaleTransactionId              int64   `json:"saleTransactionId" xorm:"index" validate:"required"`
	Quantity                       int64   `json:"quantity" xorm:"notnull" validate:"required"`
	SalePrice                      float64 `json:"salePrice" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
	ListPrice                      float64 `json:"listPrice" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
	TotalDiscountPrice             float64 `json:"totalDiscountPrice" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
	OrderItemId                    int64   `json:"orderItemId" xorm:"index notnull" validate:"required"`
	RefundItemId                   int64   `json:"refundItemId" xorm:"index default 0" `
	ProductId                      int64   `json:"productId" xorm:"index notnull" validate:"required"`
	SkuId                          int64   `json:"skuId" xorm:"index notnull" validate:"gte=0"`
	BrandCode                      string  `json:"brandCode" xorm:"index VARCHAR(30) notnull" validate:"required"`
	BrandId                        int64   `json:"brandId" xorm:"index default 0"`
	ItemCode                       string  `json:"itemCode" xorm:"index VARCHAR(60)"`
	ItemFee                        float64 `json:"itemFee" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
	TotalListPrice                 float64 `json:"totalListPrice" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
	TotalTransactionPrice          float64 `json:"totalTransactionPrice" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
	TotalDistributedCartOfferPrice float64 `json:"totalDistributedCartOfferPrice" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
	TotalDistributedItemOfferPrice float64 `json:"totalDistributedItemOfferPrice" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
	TotalDistributedPaymentPrice   float64 `json:"totalDistributedPaymentPrice" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
	TotalSalePrice                 float64 `json:"totalSalePrice" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
	DistributedCashPrice           float64 `json:"distributedCashPrice" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
	Mileage                        float64 `json:"mileage" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
	MileagePrice                   float64 `json:"mileagePrice" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
	TransactionId                  int64   `json:"transactionId" xorm:"index default 0" validate:"required"`
	TransactionDtlId               int64   `json:"transactionDtlId" xorm:"index default 0" validate:"required"`
	ObtainMileage                  float64 `json:"obtainMileage" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
}

type SaleTransactionPayment struct {
	Id                int64     `json:"id" xorm:"pk notnull autoincr"`
	SaleTransactionId int64     `json:"saleTransactionId" xorm:"index" validate:"required"`
	SeqNo             int64     `json:"seqNo" xorm:"index default 0" validate:"required"`
	PayMethod         string    `json:"payMethod"`
	PayAmt            float64   `json:"payAmt" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
	CreatedAt         time.Time `json:"CreatedBy"`
	TransactionId     int64     `json:"transactionId" xorm:"index default 0" validate:"required"`
}

//SaleTransactionAndSaleTransactionDtl
type SaleTAndSaleTDtls struct {
	SaleTransactions    []SaleTransaction    `json:"saleTransactions"`
	SaleTransactionDtls []SaleTransactionDtl `json:"saleTransactionDtls"`
}

type SaleRecordIdSuccessMapping struct {
	Id                int64     `json:"id"`
	SaleTransactionId int64     `json:"saleTransactionId" xorm:"index default 0" validate:"required"`
	SaleNo            string    `json:"saleNo" xorm:"index VARCHAR(30) notnull"`
	TransactionId     int64     `json:"transactionId" xorm:"index default 0"`
	OrderId           int64     `json:"orderId" xorm:"index default 0"`
	RefundId          int64     `json:"refundId" xorm:"index default 0"`
	OrderItemId       int64     `json:"orderItemId" xorm:"index default 0"`
	RefundItemId      int64     `json:"refundItemId" xorm:"index default 0"`
	DtlSeq            int64     `json:"dtlSeq" xorm:"index default 0"`
	CreatedAt         time.Time `json:"createdAt" xorm:"created"`
	CreatedBy         string    `json:"createdBy" xorm:"index VARCHAR(30) notnull"`
}

type SaleRecordIdFailMapping struct {
	Id                int64     `json:"id"`
	SaleTransactionId int64     `json:"saleTransactionId" xorm:"index default 0" validate:"required"`
	OrderId           int64     `json:"orderId" xorm:"index default 0"`
	RefundId          int64     `json:"refundId" xorm:"index default 0"`
	StoreId           int64     `json:"storeId" xorm:"index default 0"`
	TransactionId     int64     `json:"transactionId" xorm:"index default 0"`
	TransactionDtlId  int64     `json:"transactionDtlId" xorm:"index default 0"`
	Error             string    `json:"error" xorm:"VARCHAR(1000)"`
	Details           string    `json:"details" xorm:"VARCHAR(100)"`
	Data              string    `json:"data" xorm:"TEXT"`
	IsCreate          bool      `json:"isCreate" xorm:"index notnull default false"`
	CreatedAt         time.Time `json:"createdAt" xorm:"created"`
	CreatedBy         string    `json:"createdBy" xorm:"index VARCHAR(30)"`
}

type RequestInput struct {
	BrandCode         string `json:"brandCode" query:"brandCode"`
	ChannelType       string `json:"channelType" query:"channelType"`
	OrderId           int64  `json:"orderId" query:"orderId"`
	RefundId          int64  `json:"refundId" query:"refundId"`
	StartAt           string `json:"startAt" query:"startAt"`
	EndAt             string `json:"endAt" query:"endAt"`
	MaxResultCount    int    `json:"maxResultCount" query:"maxResultCount"`
	SkipCount         int    `json:"skipCount" query:"skipCount"`
	StoreId           int    `json:"storeId" query:"storeId"`
	TransactionId     int64  `json:"transactionId" query:"transactionId"`
	SaleTransactionId int64  `json:"saleTransactionId" query:"saleTransactionId"`
	SaleNo            string `json:"saleNo" query:"saleNo"`
}

type CslSaleMst struct {
	Id                          int64                `json:"id"`
	SaleTransactionId           int64                `json:"saleTransactionId" xorm:"index default 0"`
	TransactionId               int64                `json:"transactionId" xorm:"index default 0"`
	StoreId                     int64                `json:"storeId" xorm:"index default 0"`
	OrderId                     int64                `json:"orderId" xorm:"index default 0"`
	RefundId                    int64                `json:"refundId" xorm:"index default 0"`
	SaleNo                      string               `json:"saleNo" xorm:"index"`
	BrandCode                   string               `json:"brandCode"`
	ShopCode                    string               `json:"shopCode"`
	Dates                       string               `json:"dates"`
	PosNo                       string               `json:"posNo"`
	SeqNo                       int64                `json:"seqNo"`
	SaleMode                    string               `json:"saleMode"`
	CustNo                      string               `json:"custNo"`
	CustCardNo                  string               `json:"custCardNo"`
	CustMileagePolicyNo         int64                `json:"custMileagePolicyNo"`
	PrimaryCustEventNo          int64                `json:"primaryCustEventNo"`
	SecondaryCustEventNo        int64                `json:"secondaryCustEventNo"`
	DepartStoreReceiptNo        string               `json:"departStoreReceiptNo"`
	SaleQty                     int64                `json:"saleQty"`
	SaleAmt                     float64              `json:"saleAmt"`
	DiscountAmt                 float64              `json:"discountAmt"`
	ChinaFISaleAmt              float64              `json:"chinaFISaleAmt"`
	EstimateSaleAmt             float64              `json:"estimateSaleAmt"`
	SellingAmt                  float64              `json:"sellingAmt"`
	FeeAmt                      float64              `json:"feeAmt"`
	ActualSaleAmt               float64              `json:"actualSaleAmt"`
	UseMileage                  float64              `json:"useMileage"`
	ObtainMileage               float64              `json:"obtainMileage"`
	InUserID                    string               `json:"inUserID"`
	InDateTime                  time.Time            `json:"inDateTime"`
	ModiUserID                  string               `json:"modiUserID"`
	ModiDateTime                time.Time            `json:"modiDateTime"`
	SendState                   string               `json:"sendState"`
	SendFlag                    string               `json:"sendFlag"`
	SendSeqNo                   int64                `query:"sendSeqNo" json:"sendSeqNo"`
	SendDateTime                time.Time            `json:"sendDateTime"`
	DiscountAmtAsCost           float64              `json:"discountAmtAsCost"`
	CustDivisionCode            string               `json:"custDivisionCode"`
	MileageCustChangeStatusCode string               `json:"mileageCustChangeStatusCode"`
	CustGradeCode               string               `json:"custGradeCode"`
	PreSaleNo                   string               `json:"preSaleNo"`
	ActualSellingAmt            float64              `json:"actualSellingAmt"`
	EstimateSaleAmtForConsumer  float64              `json:"estimateSaleAmtForConsumer"`
	ShopEmpEstimateSaleAmt      float64              `json:"shopEmpEstimateSaleAmt"`
	ComplexShopSeqNo            string               `json:"complexShopSeqNo"`
	CustBrandCode               string               `json:"custBrandCode"`
	Freight                     float64              `json:"freight"`
	TMall_UseMileage            float64              `json:"tMall_UseMileage"`
	TMall_ObtainMileage         float64              `json:"tMall_ObtainMileage"`
	SaleOfficeCode              string               `json:"saleOfficeCode"`
	CreatedAt                   time.Time            `json:"createdAt" xorm:"created"`
	CslSaleDtl                  []CslSaleDtl         `json:"cslSaleDtl" xorm:"-"`
	CslSalePayment              []CslSalePayment     `json:"cslSalePayment" xorm:"-"`
	CslStaffSaleRecord          []CslStaffSaleRecord `json:"cslStaffSaleRecord" xorm:"-"`
}

type CslSaleDtl struct {
	Id                                int64     `json:"id"`
	SaleTransactionId                 int64     `json:"saleTransactionId" xorm:"index default 0"`
	SaleTransactionDtlId              int64     `json:"saleTransactionDtlId" xorm:"index default 0"`
	TransactionId                     int64     `json:"transactionId" xorm:"index default 0"`
	OrderItemId                       int64     `json:"orderItemId" xorm:"index default 0"`
	RefundItemId                      int64     `json:"refundItemId" xorm:"index default 0"`
	TransactionDtlId                  int64     `json:"transactionDtlId" xorm:"index default 0"`
	SaleNo                            string    `json:"saleNo" xorm:"index"`
	DtSeq                             int64     `json:"dtSeq"`
	BrandCode                         string    `json:"brandCode"`
	ShopCode                          string    `json:"shopCode"`
	Dates                             string    `json:"dates"`
	PosNo                             string    `json:"posNo"`
	SeqNo                             int64     `json:"seqNo"`
	NormalSaleTypeCode                string    `json:"normalSaleTypeCode"`
	CustMileagePolicyNo               int64     `json:"custMileagePolicyNo"`
	PrimaryCustEventNo                int64     `json:"primaryCustEventNo"`
	PrimaryEventTypeCode              string    `json:"primaryEventTypeCode"`
	PrimaryEventSettleTypeCode        string    `json:"primaryEventSettleTypeCode"`
	SecondaryCustEventNo              int64     `json:"secondaryCustEventNo"`
	SecondaryEventTypeCode            string    `json:"secondaryEventTypeCode"`
	SecondaryEventSettleTypeCode      string    `json:"secondaryEventSettleTypeCode"`
	SaleEventNo                       int64     `json:"saleEventNo"`
	SaleEventTypeCode                 string    `json:"saleEventTypeCode"`
	SaleReturnReasonCode              string    `json:"saleReturnReasonCode"`
	ProdCode                          string    `json:"prodCode"`
	EANCode                           string    `json:"eANCode"`
	PriceTypeCode                     string    `json:"priceTypeCode"`
	SupGroupCode                      string    `json:"supGroupCode"`
	SaipType                          string    `json:"saipType"`
	NormalPrice                       float64   `json:"normalPrice"`
	Price                             float64   `json:"price"`
	PriceDecisionDate                 string    `json:"priceDecisionDate"`
	SaleQty                           int64     `json:"saleQty"`
	SaleAmt                           float64   `json:"saleAmt"`
	EventAutoDiscountAmt              float64   `json:"eventAutoDiscountAmt"`
	EventDecisionDiscountAmt          float64   `json:"eventDecisionDiscountAmt"`
	SaleEventSaleBaseAmt              float64   `json:"saleEventSaleBaseAmt"`
	SaleEventDiscountBaseAmt          float64   `json:"saleEventDiscountBaseAmt"`
	SaleEventNormalSaleRecognitionChk bool      `json:"saleEventNormalSaleRecognitionChk"`
	SaleEventInterShopSalePermitChk   bool      `json:"saleEventInterShopSalePermitChk"`
	SaleEventAutoDiscountAmt          float64   `json:"saleEventAutoDiscountAmt"`
	SaleEventManualDiscountAmt        float64   `json:"saleEventManualDiscountAmt"`
	SaleVentDecisionDiscountAmt       float64   `json:"saleVentDecisionDiscountAmt"`
	ChinaFISaleAmt                    float64   `json:"chinaFISaleAmt"`
	EstimateSaleAmt                   float64   `json:"estimateSaleAmt"`
	SellingAmt                        float64   `json:"sellingAmt"`
	NormalFee                         float64   `json:"normalFee"`
	SaleEventFee                      float64   `json:"saleEventFee"`
	ActualSaleAmt                     float64   `json:"actualSaleAmt"`
	UseMileage                        float64   `json:"useMileage"`
	PreSaleNo                         string    `json:"preSaleNo"`
	PreSaleDtSeq                      int64     `json:"preSaleDtSeq"`
	NormalFeeRate                     float64   `json:"normalFeeRate"`
	SaleEventFeeRate                  float64   `json:"saleEventFeeRate"`
	InUserID                          string    `json:"inUserID"`
	InDateTime                        time.Time `json:"inDateTime"`
	ModiUserID                        string    `json:"modiUserID"`
	ModiDateTime                      time.Time `json:"modiDateTime"`
	SendState                         string    `json:"sendState"`
	SendFlag                          string    `json:"sendFlag"`
	SendSeqNo                         int64     `json:"sendSeqNo"`
	SendDateTime                      time.Time `json:"sendDateTime"`
	DiscountAmt                       float64   `json:"discountAmt"`
	DiscountAmtAsCost                 float64   `json:"discountAmtAsCost"`
	UseMileageSettleType              string    `json:"useMileageSettleType"`
	EstimateSaleAmtForConsumer        float64   `json:"estimateSaleAmtForConsumer"`
	SaleEventDiscountAmtForConsumer   float64   `json:"saleEventDiscountAmtForConsumer"`
	ShopEmpEstimateSaleAmt            float64   `json:"shopEmpEstimateSaleAmt"`
	PromotionID                       int64     `json:"promotionID"`
	TMallEventID                      int64     `json:"tMallEventID"`
	TMall_ObtainMileage               float64   `json:"tMall_ObtainMileage"`
	SaleOfficeCode                    string    `json:"saleOfficeCode"`
	CreatedAt                         time.Time `json:"createdAt" xorm:"created"`
}

type CslSalePayment struct {
	Id                 int64     `json:"id"`
	SaleTransactionId  int64     `json:"saleTransactionId" xorm:"index default 0"`
	TransactionId      int64     `json:"transactionId" xorm:"index default 0"`
	SaleNo             string    `json:"saleNo" xorm:"index"`
	SeqNo              int64     `json:"seqNo"`
	PaymentCode        string    `json:"paymentCode"`
	PaymentAmt         float64   `json:"paymentAmt"`
	InUserID           string    `json:"inUserID"`
	InDateTime         time.Time `json:"inDateTime"`
	ModiUserID         string    `json:"modiUserID"`
	ModiDateTime       time.Time `json:"modiDateTime"`
	SendFlag           string    `json:"sendFlag"`
	SendDateTime       time.Time `json:"sendDateTime"`
	CreditCardFirmCode string    `json:"creditCardFirmCode"`
	CreatedAt          time.Time `json:"createdAt" xorm:"created"`
}

type CslStaffSaleRecord struct {
	Id                int64     `json:"id"`
	SaleTransactionId int64     `json:"saleTransactionId" xorm:"index default 0"`
	TransactionId     int64     `json:"transactionId" xorm:"index default 0"`
	Dates             string    `json:"dates"`
	HREmpNo           string    `json:"hREmpNo"`
	SaleNo            string    `json:"saleNo" xorm:"index"`
	BrandCode         string    `json:"brandCode"`
	ShopCode          string    `json:"shopCode"`
	InUserID          string    `json:"inUserID"`
	InDateTime        time.Time `json:"inDateTime"`
	CreatedAt         time.Time `json:"createdAt" xorm:"created"`
}

func (srsm *SaleRecordIdSuccessMapping) CheckAndSave() error {
	saleRecordIdSuccessMapping := SaleRecordIdSuccessMapping{}
	has, err := factory.GetCfsrEngine().Where("sale_no = ?", srsm.SaleNo).And("order_item_id = ?", srsm.OrderItemId).
		And("refund_item_id = ? ", srsm.RefundItemId).Get(&saleRecordIdSuccessMapping)
	if err != nil {
		return err
	}
	if !has {
		if _, err := factory.GetCfsrEngine().Insert(srsm); err != nil {
			return err
		}
	}
	return nil
}

func (srfm *SaleRecordIdFailMapping) Save() error {
	var saleRecordIdFailMapping SaleRecordIdFailMapping
	has, err := factory.GetCfsrEngine().Where("transaction_id = ?", srfm.TransactionId).And("is_create = ?", false).Get(&saleRecordIdFailMapping)
	if err != nil {
		return err
	}
	if !has {
		if _, err := factory.GetCfsrEngine().Insert(srfm); err != nil {
			return err
		}
	} else {
		if err := srfm.Update(); err != nil {
			return err
		}
	}
	return nil
}

func (SaleRecordIdSuccessMapping) GetSaleSuccessData(saleTransactionId int64, orderId int64, itemId int64) ([]SaleRecordIdSuccessMapping, error) {
	var success []SaleRecordIdSuccessMapping
	queryBuilder := func() xorm.Interface {
		q := factory.GetCfsrEngine().Where("1 = 1")
		if saleTransactionId != 0 {
			q.And("sale_transaction_id = ?", saleTransactionId)
		}
		if orderId != 0 {
			q.And("order_id = ?", orderId)
		}
		if itemId != 0 {
			q.And("order_item_id = ?", itemId)
		}
		return q
	}
	if err := queryBuilder().Find(&success); err != nil {
		return nil, err
	}
	if len(success) == 0 {
		return nil, errors.New("SaleRecordIdSuccessMapping is not exist!")
	}
	return success, nil
}

func (requestInput RequestInput) Validate() error {
	// if requestInput.BrandCode == "" {
	// 	return errors.New("BrandCode can not be null!")
	// }
	// if requestInput.ChannelType == "" {
	// 	return errors.New("ChannelType can not be null!")
	// }
	if requestInput.StartAt != "" && requestInput.EndAt != "" {
		_, err := time.Parse("2006-01-02 15:04:05", requestInput.StartAt)
		if err != nil {
			return errors.New("Please input the correct time format!")
		}
		_, err = time.Parse("2006-01-02 15:04:05", requestInput.EndAt)
		if err != nil {
			return errors.New("Please input the correct time format!")
		}
	}
	if requestInput.OrderId == 0 && requestInput.RefundId == 0 && requestInput.StartAt == "" && requestInput.EndAt == "" {
		return errors.New("In orderId and startAt must be have one condition!")
	}
	return nil
}

func (SaleRecordIdFailMapping) GetAll(ctx context.Context, requestInput RequestInput) (int64, []SaleRecordIdFailMapping, error) {
	var failDatas []SaleRecordIdFailMapping
	query := func() xorm.Interface {
		query := factory.GetCfsrEngine().Where("1 = 1").And("is_create = ?", false)
		if requestInput.StoreId != 0 {
			query.And("store_id = ?", requestInput.StoreId)
		}
		if requestInput.TransactionId != 0 {
			query.And("transaction_id = ?", requestInput.TransactionId)
		}
		return query
	}
	totalCount, err := query().Desc("id").Limit(requestInput.MaxResultCount, requestInput.SkipCount).FindAndCount(&failDatas)
	if err != nil {
		return 0, nil, err
	}
	return totalCount, failDatas, nil
}

func (saleTransaction *SaleTransaction) Delete() error {
	queryBuilder := func() xorm.Interface {
		q := factory.GetCfsrEngine().Where("1 = 1")
		if saleTransaction.TransactionId != 0 {
			q.And("sale_transaction_id = ?", saleTransaction.Id)
		}
		return q
	}
	if _, err := queryBuilder().Delete(&SaleTransaction{}); err != nil {
		return err
	}
	if _, err := queryBuilder().Delete(&SaleTransactionDtl{}); err != nil {
		return err
	}
	if _, err := queryBuilder().Delete(&SaleTransactionPayment{}); err != nil {
		return err
	}
	return nil
}

func (saleTransaction *SaleTransaction) Update() error {
	if _, err := factory.GetCfsrEngine().ID(saleTransaction.Id).AllCols().Update(saleTransaction); err != nil {
		return err
	}
	for _, saleTransactionDtl := range saleTransaction.Dtls {
		saleTransactionDtl.SaleTransactionId = saleTransaction.Id
		if _, err := factory.GetCfsrEngine().Where("order_item_id = ?", saleTransactionDtl.OrderItemId).
			And("refund_item_id = ?", saleTransactionDtl.RefundItemId).AllCols().Update(saleTransactionDtl); err != nil {
			return err
		}
	}
	for _, payment := range saleTransaction.Payments {
		payment.SaleTransactionId = saleTransaction.Id
		if _, err := factory.GetCfsrEngine().Where("seq_no = ?", payment.SeqNo).
			And("transaction_id = ?", payment.TransactionId).AllCols().Update(payment); err != nil {
			return err
		}
	}
	return nil
}

func (SaleTransaction) Get(transactionId int64) (SaleTransaction, error) {
	var saleTransactions []struct {
		SaleTransaction    SaleTransaction    `xorm:"extends"`
		SaleTransactionDtl SaleTransactionDtl `xorm:"extends"`
	}
	if err := factory.GetCfsrEngine().Table("sale_transaction").
		Join("INNER", "sale_transaction_dtl", "sale_transaction_dtl.transaction_id = sale_transaction.transaction_id").
		Where("sale_transaction.transaction_id = ? ", transactionId).Find(&saleTransactions); err != nil {
		return SaleTransaction{}, err
	}
	var saleTransaction SaleTransaction
	for i, sale := range saleTransactions {
		if i == 0 {
			saleTransaction = sale.SaleTransaction
		}
		saleTransaction.Dtls = append(saleTransaction.Dtls, sale.SaleTransactionDtl)
	}
	return saleTransaction, nil
}

func (saleRecordIdFailMapping *SaleRecordIdFailMapping) Update() error {
	if _, err := factory.GetCfsrEngine().Where("transaction_id = ?", saleRecordIdFailMapping.TransactionId).AllCols().Update(saleRecordIdFailMapping); err != nil {
		return err
	}
	return nil
}

func (SaleRecordIdSuccessMapping) GetBySaleNo(salNo string) ([]SaleRecordIdSuccessMapping, error) {
	var successes []SaleRecordIdSuccessMapping
	queryBuilder := func() xorm.Interface {
		q := factory.GetCfsrEngine().Where("1 = 1")
		if salNo != "" {
			q.And("sale_no = ?", salNo)
		}
		return q
	}
	if err := queryBuilder().Find(&successes); err != nil {
		return nil, err
	}
	return successes, nil
}

func (SaleTransaction) GetSaleTransactions(ctx context.Context, transactionId, orderId, RefundId int64, shopCode string, maxResultCount, skipCount int) (int64, []SaleTransaction, error) {

	queryBuilder := func() xorm.Interface {
		q := factory.GetCfsrEngine().Where("1=1")
		if transactionId > 0 {
			q.And("transaction_id =?", transactionId)
		}
		if orderId > 0 {
			q.And("order_id =?", orderId)
		}
		if RefundId > 0 {
			q.And("refund_id =?", RefundId)
		}
		if shopCode != "" {
			q.And("shop_code =?", shopCode)
		}
		return q
	}
	query := queryBuilder()

	if maxResultCount > 0 {
		query.Limit(maxResultCount, skipCount)
	}

	query.Desc("transaction_id")

	var saleTransactions []SaleTransaction
	totalCount, err := query.FindAndCount(&saleTransactions)
	if err != nil {
		return 0, nil, err
	}

	var saleTransactionIds []int64
	for _, t := range saleTransactions {
		saleTransactionIds = append(saleTransactionIds, t.Id)
	}

	saleTransactionDtls, err := SaleTransaction{}.GetSaleTransactionDtls(ctx, saleTransactionIds)
	if err != nil {
		return 0, nil, err
	}

	saleTransactionPayments, err := SaleTransactionPayment{}.GetSaleTransactionPayments(ctx, saleTransactionIds)
	if err != nil {
		return 0, nil, err
	}

	for i, saleTransaction := range saleTransactions {
		for _, saleTransactionDtl := range saleTransactionDtls {
			if saleTransaction.TransactionId == saleTransactionDtl.TransactionId && saleTransaction.Id == saleTransactionDtl.SaleTransactionId {
				saleTransactions[i].Dtls = append(saleTransactions[i].Dtls, saleTransactionDtl)
			}
		}
		for _, saleTransactionPayment := range saleTransactionPayments {
			if saleTransactionPayment.SaleTransactionId == saleTransaction.Id {
				saleTransactions[i].Payments = append(saleTransactions[i].Payments, saleTransactionPayment)
			}
		}
	}

	return totalCount, saleTransactions, nil
}

func (SaleTransaction) GetSaleTransactionDtls(ctx context.Context, saleTransactionIds []int64) ([]SaleTransactionDtl, error) {
	queryBuilder := func() xorm.Interface {
		q := factory.GetCfsrEngine().Where("1=1")
		if len(saleTransactionIds) > 0 {
			q.In("sale_transaction_id", saleTransactionIds)
		}
		return q
	}
	query := queryBuilder()
	var saleTransactionDtls []SaleTransactionDtl
	if err := query.Desc("sale_transaction_id").Find(&saleTransactionDtls); err != nil {
		return nil, err
	}

	return saleTransactionDtls, nil
}

func (SaleTransactionPayment) GetSaleTransactionPayments(ctx context.Context, saleTransactionIds []int64) ([]SaleTransactionPayment, error) {
	queryBuilder := func() xorm.Interface {
		q := factory.GetCfsrEngine().Where("1=1")
		if len(saleTransactionIds) > 0 {
			q.In("sale_transaction_id", saleTransactionIds)
		}
		return q
	}
	query := queryBuilder()
	var saleTransactionPayments []SaleTransactionPayment
	if err := query.Desc("sale_transaction_id").Find(&saleTransactionPayments); err != nil {
		return nil, err
	}
	return saleTransactionPayments, nil
}

func (CslSaleMst) GetCslSaleBySaleTransactions(ctx context.Context, requestInput RequestInput) (int64, []CslSaleMst, error) {
	queryBuilder := func() xorm.Interface {
		q := factory.GetCfsrEngine().Where("1=1")
		if requestInput.TransactionId > 0 {
			q.And("transaction_id =?", requestInput.TransactionId)
		}
		if requestInput.OrderId > 0 {
			q.And("order_id =?", requestInput.OrderId)
		}
		if requestInput.RefundId > 0 {
			q.And("refund_id =?", requestInput.RefundId)
		}
		if requestInput.SaleTransactionId > 0 {
			q.And("sale_transaction_id =?", requestInput.SaleTransactionId)
		}
		if requestInput.SaleNo != "" {
			q.And("sale_no =?", requestInput.SaleNo)
		}
		return q
	}
	query := queryBuilder()

	if requestInput.MaxResultCount > 0 {
		query.Limit(requestInput.MaxResultCount, requestInput.SkipCount)
	}

	query.Desc("sale_transaction_id")

	var cslSaleMsts []CslSaleMst
	totalCount, err := query.FindAndCount(&cslSaleMsts)
	if err != nil {
		return 0, nil, err
	}
	if len(cslSaleMsts) == 0 {
		return 0, nil, nil
	}

	var ids []interface{}
	for _, cslSaleMst := range cslSaleMsts {
		ids = append(ids, cslSaleMst.SaleTransactionId)
	}
	cslSaleDtls, err := CslSaleDtl{}.GetCslDtlBySaleTransactions(ctx, ids)
	if err != nil {
		return 0, nil, err
	}
	cslSalePayments, err := CslSalePayment{}.GetCslSalePaymentBySaleTransactions(ctx, ids)
	if err != nil {
		return 0, nil, err
	}
	cslStaffSaleRecords, err := CslStaffSaleRecord{}.GetCslStaffSaleRecordBySaleTransactions(ctx, ids)
	if err != nil {
		return 0, nil, err
	}
	for i, cslSaleMst := range cslSaleMsts {
		for _, cslSaleDtl := range cslSaleDtls {
			if cslSaleMst.SaleTransactionId == cslSaleDtl.SaleTransactionId {
				cslSaleMsts[i].CslSaleDtl = append(cslSaleMsts[i].CslSaleDtl, cslSaleDtl)
			}
		}
		for _, cslSalePayment := range cslSalePayments {
			if cslSaleMst.SaleTransactionId == cslSalePayment.SaleTransactionId {
				cslSaleMsts[i].CslSalePayment = append(cslSaleMsts[i].CslSalePayment, cslSalePayment)
			}
		}
		for _, cslStaffSaleRecord := range cslStaffSaleRecords {
			if cslSaleMst.SaleTransactionId == cslStaffSaleRecord.SaleTransactionId {
				cslSaleMsts[i].CslStaffSaleRecord = append(cslSaleMsts[i].CslStaffSaleRecord, cslStaffSaleRecord)
			}
		}
	}

	return totalCount, cslSaleMsts, nil
}

func (CslSaleDtl) GetCslDtlBySaleTransactions(ctx context.Context, ids []interface{}) ([]CslSaleDtl, error) {
	var cslSaleDtls []CslSaleDtl
	if err := factory.GetCfsrEngine().Where("1=1").In("sale_transaction_id", ids...).Find(&cslSaleDtls); err != nil {
		return nil, err
	}
	return cslSaleDtls, nil
}

func (CslSalePayment) GetCslSalePaymentBySaleTransactions(ctx context.Context, ids []interface{}) ([]CslSalePayment, error) {
	var cslSalePayments []CslSalePayment
	if err := factory.GetCfsrEngine().Where("1=1").In("sale_transaction_id", ids...).Find(&cslSalePayments); err != nil {
		return nil, err
	}
	return cslSalePayments, nil
}

func (CslStaffSaleRecord) GetCslStaffSaleRecordBySaleTransactions(ctx context.Context, ids []interface{}) ([]CslStaffSaleRecord, error) {
	var cslStaffSaleRecords []CslStaffSaleRecord
	if err := factory.GetCfsrEngine().Where("1=1").In("sale_transaction_id", ids...).Find(&cslStaffSaleRecords); err != nil {
		return nil, err
	}
	return cslStaffSaleRecords, nil
}

func (cslSaleMst *CslSaleMst) Save() error {
	if _, err := factory.GetCfsrEngine().Insert(cslSaleMst); err != nil {
		return err
	}
	return nil
}

func (CslSaleMst) GetAll(requestInput RequestInput) ([]CslSaleMst, error) {
	var cslSaleMsts []CslSaleMst
	queryBuilder := func() xorm.Interface {
		q := factory.GetCfsrEngine().Where("1 = 1")
		if requestInput.TransactionId != 0 {
			q.And("transaction_id = ?", requestInput.TransactionId)
		}
		return q
	}
	if requestInput.MaxResultCount > 0 {
		queryBuilder().Limit(requestInput.MaxResultCount, requestInput.SkipCount)
	}
	if err := queryBuilder().Find(&cslSaleMsts); err != nil {
		return nil, err
	}
	if len(cslSaleMsts) == 0 {
		return nil, nil
	}
	return cslSaleMsts, nil
}

func (CslSaleMst) Delete(requestInput RequestInput) error {
	queryBuilder := func() xorm.Interface {
		q := factory.GetCfsrEngine().Where("1 = 1")
		if requestInput.TransactionId != 0 {
			q.And("transaction_id = ?", requestInput.TransactionId)
		}
		return q
	}
	if _, err := queryBuilder().Delete(&CslSaleDtl{}); err != nil {
		return err
	}
	if _, err := queryBuilder().Delete(&CslSaleMst{}); err != nil {
		return err
	}
	if _, err := queryBuilder().Delete(&CslSalePayment{}); err != nil {
		return err
	}
	if _, err := queryBuilder().Delete(&CslStaffSaleRecord{}); err != nil {
		return err
	}
	return nil
}

func (cslSaleDtl *CslSaleDtl) Save() error {
	if _, err := factory.GetCfsrEngine().Insert(cslSaleDtl); err != nil {
		return err
	}
	return nil
}

func (cslSalePayment *CslSalePayment) Save() error {
	if _, err := factory.GetCfsrEngine().Insert(cslSalePayment); err != nil {
		return err
	}
	return nil
}

func (cslStaffSaleRecord *CslStaffSaleRecord) Save() error {
	if _, err := factory.GetCfsrEngine().Insert(cslStaffSaleRecord); err != nil {
		return err
	}
	return nil
}
