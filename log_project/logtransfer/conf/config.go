package conf

type Logtransfer struct { //LogTransfer 全局配置
	KafkaCfg `ini:"kafka"`
	ESCfg    `ini:"es"`
}

type KafkaCfg struct { //Kafka...
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type ESCfg struct { //ESCfg
	Address  string `ini:"address"`
	ChanSize int    `ini:"size"`
	Worker   int    `ini:"worker"`
}
