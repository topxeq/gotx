package main

import (
	"reflect"

	"github.com/topxeq/xiaoxian"
)

func init() {
	GotxSymbols["github.com/topxeq/xiaoxian"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"CalCosineSimilarityOfVectors": reflect.ValueOf(xiaoxian.CalCosineSimilarityOfVectors),
		"CleanChineseSentence":         reflect.ValueOf(xiaoxian.CleanChineseSentence),
		"CleanEnglish":                 reflect.ValueOf(xiaoxian.CleanEnglish),
		"DownloadPagePostOnlyBaiduNLP": reflect.ValueOf(xiaoxian.DownloadPagePostOnlyBaiduNLP),
		"EnsureValidEnglishOnly":       reflect.ValueOf(xiaoxian.EnsureValidEnglishOnly),
		"GetArticleDifficultyEnOL":     reflect.ValueOf(xiaoxian.GetArticleDifficultyEnOL),
		"GetNamedEntityEnOL":           reflect.ValueOf(xiaoxian.GetNamedEntityEnOL),
		"GetVectorCnBaiduOL":           reflect.ValueOf(xiaoxian.GetVectorCnBaiduOL),
		"LoadD2VModel":                 reflect.ValueOf(xiaoxian.LoadD2VModel),
		"LoadD2VModelWithDicts":        reflect.ValueOf(xiaoxian.LoadD2VModelWithDicts),
		"NerCnBaiduOL":                 reflect.ValueOf(xiaoxian.NerCnBaiduOL),
		"NewD2VModel":                  reflect.ValueOf(xiaoxian.NewD2VModel),
		"NewD2VModelFromDicts":         reflect.ValueOf(xiaoxian.NewD2VModelFromDicts),
		"NewD2VModelFromSegmenter":     reflect.ValueOf(xiaoxian.NewD2VModelFromSegmenter),
		"NewPosTaggerCn":               reflect.ValueOf(xiaoxian.NewPosTaggerCn),
		"NewTokenizerCn":               reflect.ValueOf(xiaoxian.NewTokenizerCn),
		"NewTreebankWordTokenizer":     reflect.ValueOf(xiaoxian.NewTreebankWordTokenizer),
		"ParseSentenceEnOL":            reflect.ValueOf(xiaoxian.ParseSentenceEnOL),
		"SentimentCnBaiduOL":           reflect.ValueOf(xiaoxian.SentimentCnBaiduOL),
		"SplitArticleCn":               reflect.ValueOf(xiaoxian.SplitArticleCn),
		"SplitArticleEn":               reflect.ValueOf(xiaoxian.SplitArticleEn),
		"SplitArticleEnOL":             reflect.ValueOf(xiaoxian.SplitArticleEnOL),
		"TagEnOL":                      reflect.ValueOf(xiaoxian.TagEnOL),
		"TokenizeCnBaiduOL":            reflect.ValueOf(xiaoxian.TokenizeCnBaiduOL),
		"TokenizeEn":                   reflect.ValueOf(xiaoxian.TokenizeEn),
		"TokenizeEnOL":                 reflect.ValueOf(xiaoxian.TokenizeEnOL),
		"TrainDoc2VecModel":            reflect.ValueOf(xiaoxian.TrainDoc2VecModel),
		"XiaoXianDefaultTid":           reflect.ValueOf(xiaoxian.XiaoXianDefaultTid),
		"XiaoXianHost":                 reflect.ValueOf(xiaoxian.XiaoXianHost),

		// type definitions
		"D2VModel":              reflect.ValueOf((*xiaoxian.D2VModel)(nil)),
		"PosTaggerCn":           reflect.ValueOf((*xiaoxian.PosTaggerCn)(nil)),
		"TokenizerCn":           reflect.ValueOf((*xiaoxian.TokenizerCn)(nil)),
		"TreebankWordTokenizer": reflect.ValueOf((*xiaoxian.TreebankWordTokenizer)(nil)),
	}
}
