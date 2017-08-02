package models

// Word the word searching result, stored in db
type Word struct {
	ID                string `json:"id" mgo:"id"`                                   // 数据项ID, 可用于
	CreateTime        int64  `json:"create_time" mgo:"create_time"`                 // 创建时间, unix timestamp
	UpdateTime        int64  `json:"update_time" mgo:"update_time"`                 // 更新时间, unix timestamp
	Name              string `json:"name" mgo:"name"`                               // 词义
	ParaphraseCN      string `json:"paraphrase_cn" mgo:"paraphrase_cn"`             // 中文释义,包含词性
	ParaphraseEN      string `json:"paraphrase_en" mgo:"paraphrase_en"`             // 英文释义,包含词性
	PronounceSymbolUK string `json:"pronounce_symbol_uk" mgo:"pronounce_symbol_uk"` // 英音音标
	PronounceSymbolUS string `json:"pronounce_symbol_us" mgo:"pronounce_symbol_us"` // 美音音标
	AudioUK           string `json:"audio_uk" mgo:"audio_uk"`                       // 英音发音
	AudioUS           string `json:"audio_us" mgo:"audio_us"`                       // 美音发音
	ExampleEN         string `json:"example_en" mgo:"example_en"`                   // 例句英文
	ExampleCN         string `json:"example_cn" mgo:"example_cn"`                   // 例句中文
}
