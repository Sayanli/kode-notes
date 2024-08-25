package spellchecker

import (
	"encoding/json"
	"io/ioutil"
	"kode-notes/internal/entity"
	"log"
	"net/http"
	"strings"
)

type YandexSpellChecker struct {
}

type InputItem struct {
	Code int      `json:"code"`
	Pos  int      `json:"pos"`
	Row  int      `json:"row"`
	Col  int      `json:"col"`
	Len  int      `json:"len"`
	Word string   `json:"word"`
	S    []string `json:"s"`
}

const apiURL = "https://speller.yandex.net/services/spellservice.json/checkText?text="

func (y *YandexSpellChecker) Check(text string) ([]byte, error) {
	var inputItems []InputItem
	//Yandex speller принимает текст с '+' вместо пробелов
	modifiedText := strings.ReplaceAll(text, " ", "+")

	resp, err := http.Get(apiURL + modifiedText)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &inputItems)
	if err != nil {
		log.Fatalf("Error unmarshalling input JSON: %v", err)
	}
	var outputItems []entity.Mistakes
	for _, item := range inputItems {
		outputItems = append(outputItems, entity.Mistakes{
			OriginalWord: item.Word,
			CorrectWord:  item.S,
		})
	}

	mistakes, err := json.Marshal(outputItems)
	if err != nil {
		log.Fatalf("Error marshalling output JSON: %v", err)
	}
	return mistakes, nil
}

func NewYandexSpellChecker() *YandexSpellChecker {
	return &YandexSpellChecker{}
}
