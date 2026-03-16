package config

type TimeFrameConfig struct {
	SeniorTF string
	TF       string
}

var TFConfig = map[string]TimeFrameConfig{
	"15m": {
		SeniorTF: "60m",
		TF:       "15m",
	},
	"5m": {
		SeniorTF: "15m",
		TF:       "5m",
	},
	"1m": {
		SeniorTF: "5m",
		TF:       "1m",
	},
	"1h": {
		SeniorTF: "4h",
		TF:       "60m",
	},
	"4h": {
		SeniorTF: "1d",
		TF:       "4h",
	},
}
