package formatter

type Formatter interface {
	Format(value float64, unit string) string
}

var (
	DurationFormatter   = NewDurationFormatter()
	BoolFormatter       = NewBoolFormatter()
	PercentFormatter    = NewPercentFormatter()
	NoneFormatter       = NewNoneFormatter()
	DataFormatter       = NewDataFormatter()
	DataRateFormatter   = NewDataRateFormatter()
	ThroughputFormatter = NewThroughputFormatter()
)

func FromUnit(u string) Formatter {
	switch u {
	case "ns", "us", "ms", "s", "m", "h", "d":
		return DurationFormatter
	case "bytes", "decbytes", "bits", "decbits", "kbytes", "decKbytes", "mbytes", "decMbytes", "gbytes", "decGbytes", "tbytes", "decTbytes", "pbytes", "decPbytes":
		return DataFormatter
	case "binBps", "Bps", "binbps", "bps", "KiBs", "Kibits", "KBs", "Kbits", "MiBs", "Mibits", "MBs", "Mbits", "GiBs", "Gibits", "GBs", "Gbits", "TiBs", "Tibits", "TBs", "Tbits", "PiBs", "Pibits", "PBs", "Pbits":
		return DataRateFormatter
	case "percent", "percentunit":
		return PercentFormatter
	case "bool", "bool_yes_no", "bool_true_false", "bool_1_0":
		return BoolFormatter
	case "cps", "ops", "reqps", "rps", "wps", "iops", "cpm", "opm", "rpm", "wpm":
		return ThroughputFormatter
	default:
		return NoneFormatter
	}
}
