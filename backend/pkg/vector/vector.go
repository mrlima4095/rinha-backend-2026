package vector

const (
	IndexAmount uint8 = iota
	IndexInstallments
	IndexAmountVsAvg
	IndexHourOfDay
	IndexDayOfWeek
	IndexMinutesSinceLastTx
	IndexKMFromLastTx
	IndexKMFromHome
	IndexTxCount24h
	IndexIsOnline
	IndexCardPresent
	IndexUnknownMerchant
	IndexMCCRisk
	IndexMerchantAVGAmount
)

const (
	MaxAmount            float32 = 1000
	MaxInstallments      uint8   = 12
	AmountVsAVGRatio     float32 = 10
	MaxMinutes           float64 = 1440
	MaxKm                float32 = 1000
	MaxTxCount24h        float32 = 20
	MaxMerchantAVGAmount float32 = 10000
)

type Vector [14]float32

// amount [0]
func (v *Vector) SetAmount(n float32) { v[0] = limit(n) }
func (v *Vector) Amount() float32     { return v[0] }

// installments [1]
func (v *Vector) SetInstallments(n float32) { v[1] = limit(n) }
func (v *Vector) Installments() float32     { return v[1] }

// amount_vs_avg [2]
func (v *Vector) SetAmountVsAVG(n float32) { v[2] = limit(n) }
func (v *Vector) AmountVsAVG() float32     { return v[2] }

// hour_of_day [3]
func (v *Vector) SetHourOfDay(n float32) { v[3] = limit(n) }
func (v *Vector) HourOfDay() float32     { return v[3] }

// day_of_week [4]
func (v *Vector) SetDayOfWeek(n float32) { v[4] = limit(n) }
func (v *Vector) DayOfWeek() float32     { return v[4] }

// minutes_since_last_tx [5]
func (v *Vector) SetMinutesSinceLastTx(n float32) { v[5] = limitIgnoreLess(n) }
func (v *Vector) MinutesSinceLastTx() float32     { return v[5] }

// km_from_last_tx [6]
func (v *Vector) SetKmFromLastTx(n float32) { v[6] = limitIgnoreLess(n) }
func (v *Vector) KmFromLastTx() float32     { return v[6] }

// km_from_home [7]
func (v *Vector) SetKmFromHome(n float32) { v[7] = limit(n) }
func (v *Vector) KmFromHome() float32     { return v[7] }

// tx_count_24h [8]
func (v *Vector) SetTxCount24h(n float32) { v[8] = limit(n) }
func (v *Vector) TxCount24h() float32     { return v[8] }

// is_online [9]
func (v *Vector) SetIsOnline(n float32) { v[9] = limit(n) }
func (v *Vector) IsOnline() float32     { return v[9] }

// card_present [10]
func (v *Vector) SetCardPresent(n float32) { v[10] = limit(n) }
func (v *Vector) CardPresent() float32     { return v[10] }

// unknown_merchant [11]
func (v *Vector) SetUnknownMerchant(n float32) { v[11] = limit(n) }
func (v *Vector) UnknownMerchant() float32     { return v[11] }

// mcc_risk [12]
func (v *Vector) SetMCCRisk(n float32) { v[12] = limit(n) }
func (v *Vector) MCCRisk() float32     { return v[12] }

// merchant_avg_amount [13]
func (v *Vector) SetMerchantAVGAmount(n float32) { v[13] = limit(n) }
func (v *Vector) MerchantAVGAmount() float32     { return v[13] }

func (v *Vector) ToSlice() []float32 { return v[:] }

func limit(n float32) float32 {
	if n > 1 {
		return 1
	}
	if n < 0 {
		return 0
	}
	return n
}

func limitIgnoreLess(n float32) float32 {
	if n > 1 {
		return 1
	}
	return n
}
