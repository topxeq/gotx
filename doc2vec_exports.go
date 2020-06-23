package main

import (
	"reflect"

	"github.com/topxeq/doc2vec/corpus"
	"github.com/topxeq/doc2vec/doc2vec"
	"github.com/topxeq/doc2vec/neuralnet"
)

func init() {
	GotxSymbols["github.com/topxeq/doc2vec/doc2vec"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"DBC2SBC":                    reflect.ValueOf(doc2vec.DBC2SBC),
		"EXP_TABLE_SIZE":             reflect.ValueOf(doc2vec.EXP_TABLE_SIZE),
		"GetNegativeSamplingWordIdx": reflect.ValueOf(doc2vec.GetNegativeSamplingWordIdx),
		"GetSigmoidValue":            reflect.ValueOf(doc2vec.GetSigmoidValue),
		"If":                         reflect.ValueOf(doc2vec.If),
		"MAX_EXP":                    reflect.ValueOf(doc2vec.MAX_EXP),
		"Max":                        reflect.ValueOf(doc2vec.Max),
		"Min":                        reflect.ValueOf(doc2vec.Min),
		"NEG_SAMPLING_TABLE_SIZE":    reflect.ValueOf(doc2vec.NEG_SAMPLING_TABLE_SIZE),
		"NewDoc2Vec":                 reflect.ValueOf(doc2vec.NewDoc2Vec),
		"PROGRESS_BAR_THRESHOLD":     reflect.ValueOf(doc2vec.PROGRESS_BAR_THRESHOLD),
		"QuickSort":                  reflect.ValueOf(doc2vec.QuickSort),
		"SBC2DBC":                    reflect.ValueOf(doc2vec.SBC2DBC),
		"THREAD_NUM":                 reflect.ValueOf(doc2vec.THREAD_NUM),

		// type definitions
		"IDoc2Vec":       reflect.ValueOf((*doc2vec.IDoc2Vec)(nil)),
		"SortItem":       reflect.ValueOf((*doc2vec.SortItem)(nil)),
		"TDoc2VecImpl":   reflect.ValueOf((*doc2vec.TDoc2VecImpl)(nil)),
		"TSortItemSlice": reflect.ValueOf((*doc2vec.TSortItemSlice)(nil)),

		// interface wrapper definitions
		"_IDoc2Vec": reflect.ValueOf((*_github_com_topxeq_doc2vec_doc2vec_IDoc2Vec)(nil)),
	}
}

// _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec is an interface wrapper for IDoc2Vec type
type _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec struct {
	WDoc2Docs           func(docidx int)
	WDoc2Words          func(docidx int)
	WDocSimCal          func(content1 string, content2 string) (dis float64)
	WFindNearestDoc     func(content string, iters int) string
	WGetCorpus          func() corpus.ICorpus
	WGetDim             func() int
	WGetLeaveOneOutKwds func(content string, iters int)
	WGetLikelihood4Doc  func(context string) (likelihood float64)
	WGetNeuralNet       func() neuralnet.INeuralNet
	WGetRound           func() int
	WInferDoc           func(content string, iters int) (rs []float32)
	WLoadModel          func(fname string) (err error)
	WSaveModel          func(fname string) (err error)
	WSen2Docs           func(content string, iters int)
	WSen2Words          func(content string, iters int)
	WTXDoc2Words        func(content string, iters int, limit int, outWordsA map[string]int) (rs []string, vec []float32)
	WTXWord2Docs        func(word string, limit int) []string
	WTXWord2Words       func(word string, limit int, outWordsA map[string]int) []string
	WTXWord2WordsIn     func(word string, limit int, inWordListA map[string]int) []string
	WTXWord2WordsInOut  func(word string, limit int, inWordListA map[string]int, outWordsA map[string]int) []string
	WTrain              func(fname string)
	WTrainFromString    func(strA string)
	WWord2Docs          func(word string)
	WWord2Words         func(word string)
}

func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) Doc2Docs(docidx int)  { W.WDoc2Docs(docidx) }
func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) Doc2Words(docidx int) { W.WDoc2Words(docidx) }
func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) DocSimCal(content1 string, content2 string) (dis float64) {
	return W.WDocSimCal(content1, content2)
}
func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) FindNearestDoc(content string, iters int) string {
	return W.WFindNearestDoc(content, iters)
}
func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) GetCorpus() corpus.ICorpus {
	return W.WGetCorpus()
}
func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) GetDim() int { return W.WGetDim() }
func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) GetLeaveOneOutKwds(content string, iters int) {
	W.WGetLeaveOneOutKwds(content, iters)
}
func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) GetLikelihood4Doc(context string) (likelihood float64) {
	return W.WGetLikelihood4Doc(context)
}
func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) GetNeuralNet() neuralnet.INeuralNet {
	return W.WGetNeuralNet()
}
func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) GetRound() int { return W.WGetRound() }
func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) InferDoc(content string, iters int) (rs []float32) {
	return W.WInferDoc(content, iters)
}
func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) LoadModel(fname string) (err error) {
	return W.WLoadModel(fname)
}
func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) SaveModel(fname string) (err error) {
	return W.WSaveModel(fname)
}
func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) Sen2Docs(content string, iters int) {
	W.WSen2Docs(content, iters)
}
func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) Sen2Words(content string, iters int) {
	W.WSen2Words(content, iters)
}
func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) TXDoc2Words(content string, iters int, limit int, outWordsA map[string]int) (rs []string, vec []float32) {
	return W.WTXDoc2Words(content, iters, limit, outWordsA)
}
func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) TXWord2Docs(word string, limit int) []string {
	return W.WTXWord2Docs(word, limit)
}
func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) TXWord2Words(word string, limit int, outWordsA map[string]int) []string {
	return W.WTXWord2Words(word, limit, outWordsA)
}
func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) TXWord2WordsIn(word string, limit int, inWordListA map[string]int) []string {
	return W.WTXWord2WordsIn(word, limit, inWordListA)
}
func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) TXWord2WordsInOut(word string, limit int, inWordListA map[string]int, outWordsA map[string]int) []string {
	return W.WTXWord2WordsInOut(word, limit, inWordListA, outWordsA)
}
func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) Train(fname string) { W.WTrain(fname) }
func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) TrainFromString(strA string) {
	W.WTrainFromString(strA)
}
func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) Word2Docs(word string)  { W.WWord2Docs(word) }
func (W _github_com_topxeq_doc2vec_doc2vec_IDoc2Vec) Word2Words(word string) { W.WWord2Words(word) }
