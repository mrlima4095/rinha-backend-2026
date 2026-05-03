package payments

import (
	"slices"
	"time"

	"github.com/robogg133/rinha-backend-2026/pkg/mccrisk"
	"github.com/robogg133/rinha-backend-2026/pkg/vector"
)

type Payment struct {
	ID string `json:"id"`

	LastTransaction *struct {
		Timestamp     time.Time `json:"timestamp"`
		KmFromCurrent float32   `json:"km_from_current"`
	} `json:"last_transaction"`

	Transaction struct {
		Amount       float32   `json:"amount"`
		Installments uint8     `json:"installments"`
		RequestedAt  time.Time `json:"requested_at"`
	} `json:"transaction"`

	Customer struct {
		AvgAmount      float32  `json:"avg_amount"`
		TxCount24h     uint8    `json:"tx_count_24h"`
		KnownMerchants []string `json:"known_merchants"`
	} `json:"customer"`

	Merchant struct {
		ID        string  `json:"id"`
		MCC       string  `json:"mcc"`
		AVGAmount float32 `json:"avg_amount"`
	} `json:"merchant"`

	Terminal struct {
		IsOnline    bool    `json:"is_online"`
		KmFromHome  float32 `json:"km_from_home"`
		CardPresent bool    `json:"card_present"`
	} `json:"terminal"`
}

func (p *Payment) ToVector() *vector.Vector {

	v := new(vector.Vector)

	v.SetAmount(p.Transaction.Amount / vector.MaxAmount)                                      // 0
	v.SetInstallments(float32(p.Transaction.Installments) / float32(vector.MaxInstallments))  // 1
	v.SetAmountVsAVG((p.Transaction.Amount / p.Customer.AvgAmount) / vector.AmountVsAVGRatio) // 2

	v.SetHourOfDay(float32(p.Transaction.RequestedAt.Hour()) / 23)        // 3
	v.SetDayOfWeek(weekdayToNum[p.Transaction.RequestedAt.Weekday()] / 6) // 4

	p.applyLastTransaction(v) // 5 and 6

	v.SetKmFromHome(p.Terminal.KmFromHome / vector.MaxKm)                  // 7
	v.SetTxCount24h(float32(p.Customer.TxCount24h) / vector.MaxTxCount24h) // 8

	v.SetIsOnline(boolToFloat(p.Terminal.IsOnline))       // 9
	v.SetCardPresent(boolToFloat(p.Terminal.CardPresent)) // 10

	// 11
	if !slices.Contains(p.Customer.KnownMerchants, p.Merchant.ID) {
		v.SetUnknownMerchant(1)
	} else {
		v.SetUnknownMerchant(0)
	}

	v.SetMCCRisk(mccrisk.GetMCCRisk(p.Merchant.MCC)) // 12

	v.SetMerchantAVGAmount(p.Merchant.AVGAmount / vector.MaxMerchantAVGAmount) // 13

	return v
}

func boolToFloat(b bool) float32 {
	switch b {
	case true:
		return 1
	case false:
		return 0
	default:
		return 0
	}
}
func (p *Payment) applyLastTransaction(v *vector.Vector) {
	if p.LastTransaction == nil {
		v.SetMinutesSinceLastTx(-1)
		v.SetKmFromLastTx(-1)
		return
	}

	minutes := p.Transaction.RequestedAt.Sub(p.LastTransaction.Timestamp).Minutes()
	v.SetMinutesSinceLastTx(float32(minutes / vector.MaxMinutes))

	v.SetKmFromLastTx(p.LastTransaction.KmFromCurrent / vector.MaxKm)
}

var weekdayToNum [7]float32 = [7]float32{
	6,
	0,
	1,
	2,
	3,
	4,
	5,
}
