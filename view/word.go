package view

// ShanbeiResult Shanbay word search API
// example endpoint: http://api.shanbay.com/bdc/search/?word=XXX
type ShanbeiResult struct {
	Data struct {
		Audio          string `json:"audio"`
		AudioAddresses struct {
			Uk []string `json:"uk"`
			Us []string `json:"us"`
		} `json:"audio_addresses"`
		AudioName    string `json:"audio_name"`
		CnDefinition struct {
			Defn string `json:"defn"`
			Pos  string `json:"pos"`
		} `json:"cn_definition"`
		ConentID     int    `json:"conent_id"`
		Content      string `json:"content"`
		ContentID    int    `json:"content_id"`
		ContentType  string `json:"content_type"`
		Definition   string `json:"definition"`
		EnDefinition struct {
			Defn string `json:"defn"`
			Pos  string `json:"pos"`
		} `json:"en_definition"`
		EnDefinitions struct {
			Adj []string `json:"adj"`
			V   []string `json:"v"`
		} `json:"en_definitions"`
		HasAudio       bool   `json:"has_audio"`
		ID             int    `json:"id"`
		NumSense       int    `json:"num_sense"`
		ObjectID       int    `json:"object_id"`
		Pron           string `json:"pron"`
		Pronunciation  string `json:"pronunciation"`
		Pronunciations struct {
			Uk string `json:"uk"`
			Us string `json:"us"`
		} `json:"pronunciations"`
		SenseID int    `json:"sense_id"`
		UkAudio string `json:"uk_audio"`
		URL     string `json:"url"`
		UsAudio string `json:"us_audio"`
	} `json:"data"`
	Msg        string `json:"msg"`
	StatusCode int    `json:"status_code"`
}

// CibaResult Kingsoft Powerword Search API
// example endpoint: http://dict-co.iciba.com/api/dictionary.php?type=json&w=XXX&key=XXX
type CibaResult struct {
	Exchange struct {
		WordDone  []string `json:"word_done"`
		WordEr    string   `json:"word_er"`
		WordEst   string   `json:"word_est"`
		WordIng   []string `json:"word_ing"`
		WordPast  []string `json:"word_past"`
		WordPl    []string `json:"word_pl"`
		WordThird []string `json:"word_third"`
	} `json:"exchange"`
	IsCRI   string `json:"is_CRI"`
	Symbols []struct {
		Parts []struct {
			Means []string `json:"means"`
			Part  string   `json:"part"`
		} `json:"parts"`
		PhAm     string `json:"ph_am"`
		PhAmMp3  string `json:"ph_am_mp3"`
		PhEn     string `json:"ph_en"`
		PhEnMp3  string `json:"ph_en_mp3"`
		PhOther  string `json:"ph_other"`
		PhTtsMp3 string `json:"ph_tts_mp3"`
	} `json:"symbols"`
	WordName string `json:"word_name"`
}

// PearsonAPI is http response structure by using ldoce5 as {dictionary} parameter
// example endpoint: https://api.pearson.com/v2/dictionaries/{dictionary}/entries?headword=XXX&apikey=XXX
type PearsonAPI struct {
	Count   int `json:"count"`
	Limit   int `json:"limit"`
	Offset  int `json:"offset"`
	Results []struct {
		Datasets       []string `json:"datasets"`
		Headword       string   `json:"headword"`
		ID             string   `json:"id"`
		PartOfSpeech   string   `json:"part_of_speech"`
		Pronunciations []struct {
			Audio []struct {
				Lang string `json:"lang"`
				Type string `json:"type"`
				URL  string `json:"url"`
			} `json:"audio"`
			Ipa string `json:"ipa"`
		} `json:"pronunciations"`
		Senses []struct {
			Definition   []string `json:"definition"`
			RelatedWords []string `json:"related_words"`
		} `json:"senses"`
		URL string `json:"url"`
	} `json:"results"`
	Status int    `json:"status"`
	Total  int    `json:"total"`
	URL    string `json:"url"`
}
