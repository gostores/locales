package es_PY

import (
	"math"
	"strconv"
	"time"

	"github.com/govenue/locales"
	"github.com/govenue/locales/currency"
)

type es_PY struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	percentSuffix          string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
	currencyPositivePrefix string
	currencyNegativePrefix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	periodsAbbreviated     []string
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'es_PY' locale
func New() locales.Translator {
	return &es_PY{
		locale:                 "es_PY",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "Gs.", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UZS", "VEB", "VEF", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "ene.", "feb.", "mar.", "abr.", "may.", "jun.", "jul.", "ago.", "sept.", "oct.", "nov.", "dic."},
		monthsNarrow:           []string{"", "E", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
		daysAbbreviated:        []string{"dom.", "lun.", "mar.", "mié.", "jue.", "vie.", "sáb."},
		daysNarrow:             []string{"D", "L", "M", "X", "J", "V", "S"},
		daysShort:              []string{"do", "lu", "ma", "mi", "ju", "vi", "sa"},
		daysWide:               []string{"domingo", "lunes", "martes", "miércoles", "jueves", "viernes", "sábado"},
		periodsAbbreviated:     []string{"a. m.", "p. m."},
		periodsNarrow:          []string{"a. m.", "p. m."},
		periodsWide:            []string{"a. m.", "p. m."},
		erasAbbreviated:        []string{"a. C.", "d. C."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"antes de Cristo", "después de Cristo"},
		timezones:              map[string]string{"HNPMX": "hora estándar del Pacífico de México", "HECU": "hora de verano de Cuba", "PST": "hora estándar del Pacífico", "LHDT": "hora de verano de Lord Howe", "AEDT": "hora de verano de Australia oriental", "ARST": "hora de verano de Argentina", "ACST": "hora estándar de Australia central", "AKDT": "hora de verano de Alaska", "MESZ": "hora de verano de Europa central", "HENOMX": "hora de verano del noroeste de México", "HEOG": "hora de verano de Groenlandia occidental", "HKT": "hora estándar de Hong Kong", "CHADT": "hora de verano de Chatham", "CST": "hora estándar central", "BOT": "hora de Bolivia", "NZDT": "hora de verano de Nueva Zelanda", "MST": "hora estándar de las Montañas", "WITA": "hora de Indonesia central", "HNOG": "hora estándar de Groenlandia occidental", "SAST": "hora de Sudáfrica", "HNT": "hora estándar de Terranova", "HEEG": "hora de verano de Groenlandia oriental", "HAT": "hora de verano de Terranova", "HAST": "hora estándar de Hawái-Aleutiano", "EAT": "hora de África oriental", "AKST": "hora estándar de Alaska", "SRT": "hora de Surinam", "UYST": "hora de verano de Uruguay", "MDT": "hora de verano de las Montañas", "OEZ": "hora estándar de Europa oriental", "WIT": "hora de Indonesia oriental", "ACWST": "hora estándar de Australia centroccidental", "WART": "hora estándar de Argentina occidental", "ADT": "hora de verano del Atlántico", "HEPMX": "hora de verano del Pacífico de México", "BT": "hora de Bután", "UYT": "hora estándar de Uruguay", "HADT": "hora de verano de Hawái-Aleutiano", "WARST": "hora de verano de Argentina occidental", "GYT": "hora de Guyana", "ACDT": "hora de verano de Australia central", "CAT": "hora de África central", "WIB": "hora de Indonesia occidental", "AST": "hora estándar del Atlántico", "∅∅∅": "hora de verano de las Azores", "AWST": "hora estándar de Australia occidental", "MEZ": "hora estándar de Europa central", "CHAST": "hora estándar de Chatham", "HEPM": "hora de verano de San Pedro y Miquelón", "NZST": "hora estándar de Nueva Zelanda", "OESZ": "hora de verano de Europa oriental", "AEST": "hora estándar de Australia oriental", "CLST": "hora de verano de Chile", "EDT": "hora de verano oriental", "ChST": "hora estándar de Chamorro", "WEZ": "hora estándar de Europa occidental", "WESZ": "hora de verano de Europa occidental", "AWDT": "hora de verano de Australia occidental", "JDT": "hora de verano de Japón", "WAT": "hora estándar de África occidental", "WAST": "hora de verano de África occidental", "GFT": "hora de la Guayana Francesa", "EST": "hora estándar oriental", "MYT": "hora de Malasia", "VET": "hora de Venezuela", "HNNOMX": "hora estándar del noroeste de México", "ART": "hora estándar de Argentina", "HNCU": "hora estándar de Cuba", "ECT": "hora de Ecuador", "CDT": "hora de verano central", "ACWDT": "hora de verano de Australia centroccidental", "HNEG": "hora estándar de Groenlandia oriental", "HKST": "hora de verano de Hong Kong", "CLT": "hora estándar de Chile", "COT": "hora estándar de Colombia", "COST": "hora de verano de Colombia", "TMST": "hora de verano de Turkmenistán", "SGT": "hora de Singapur", "TMT": "hora estándar de Turkmenistán", "PDT": "hora de verano del Pacífico", "LHST": "hora estándar de Lord Howe", "JST": "hora estándar de Japón", "IST": "hora estándar de la India", "GMT": "hora del meridiano de Greenwich", "HNPM": "hora estándar de San Pedro y Miquelón"},
	}
}

// Locale returns the current translators string locale
func (es *es_PY) Locale() string {
	return es.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'es_PY'
func (es *es_PY) PluralsCardinal() []locales.PluralRule {
	return es.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'es_PY'
func (es *es_PY) PluralsOrdinal() []locales.PluralRule {
	return es.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'es_PY'
func (es *es_PY) PluralsRange() []locales.PluralRule {
	return es.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'es_PY'
func (es *es_PY) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'es_PY'
func (es *es_PY) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'es_PY'
func (es *es_PY) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (es *es_PY) MonthAbbreviated(month time.Month) string {
	return es.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (es *es_PY) MonthsAbbreviated() []string {
	return es.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (es *es_PY) MonthNarrow(month time.Month) string {
	return es.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (es *es_PY) MonthsNarrow() []string {
	return es.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (es *es_PY) MonthWide(month time.Month) string {
	return es.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (es *es_PY) MonthsWide() []string {
	return es.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (es *es_PY) WeekdayAbbreviated(weekday time.Weekday) string {
	return es.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (es *es_PY) WeekdaysAbbreviated() []string {
	return es.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (es *es_PY) WeekdayNarrow(weekday time.Weekday) string {
	return es.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (es *es_PY) WeekdaysNarrow() []string {
	return es.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (es *es_PY) WeekdayShort(weekday time.Weekday) string {
	return es.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (es *es_PY) WeekdaysShort() []string {
	return es.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (es *es_PY) WeekdayWide(weekday time.Weekday) string {
	return es.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (es *es_PY) WeekdaysWide() []string {
	return es.daysWide
}

// Decimal returns the decimal point of number
func (es *es_PY) Decimal() string {
	return es.decimal
}

// Group returns the group of number
func (es *es_PY) Group() string {
	return es.group
}

// Group returns the minus sign of number
func (es *es_PY) Minus() string {
	return es.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'es_PY' and handles both Whole and Real numbers based on 'v'
func (es *es_PY) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, es.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, es.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, es.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'es_PY' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (es *es_PY) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, es.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, es.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, es.percentSuffix...)

	b = append(b, es.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'es_PY'
func (es *es_PY) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := es.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, es.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, es.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(es.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, es.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, es.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, es.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'es_PY'
// in accounting notation.
func (es *es_PY) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := es.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, es.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, es.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(es.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, es.currencyNegativePrefix[j])
		}

		b = append(b, es.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(es.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, es.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, es.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'es_PY'
func (es *es_PY) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'es_PY'
func (es *es_PY) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, es.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'es_PY'
func (es *es_PY) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, es.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'es_PY'
func (es *es_PY) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, es.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, es.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'es_PY'
func (es *es_PY) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, es.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'es_PY'
func (es *es_PY) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, es.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, es.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'es_PY'
func (es *es_PY) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, es.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, es.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'es_PY'
func (es *es_PY) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, es.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, es.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()

	if btz, ok := es.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	b = append(b, []byte{0x29}...)

	return string(b)
}
