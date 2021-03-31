package translation

import (
	"sync"
	"translation-tool/api"
)

func Translation(text string) []TranslationResults {
	var wg sync.WaitGroup
	apis := api.GetApis()
	res := make([]TranslationResults, len(apis))
	for i, _ := range apis {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			res[i] = TranslationResults{
				Name:    apis[i].GetName(),
				Results: apis[i].Translation(text),
			}
		}(i)
	}
	wg.Wait()
	return res
}

type TranslationResults struct {
	Name    string
	Results string
}
