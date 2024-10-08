package spellchecker

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"kode-notes/internal/entity"
	"log"
	"log/slog"
	"net/http"
	"strings"
)

type YandexSpellChecker struct {
	logger *slog.Logger
}

func NewYandexSpellChecker(logger *slog.Logger) *YandexSpellChecker {
	return &YandexSpellChecker{
		logger: logger,
	}
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
	const op = "spellchecker.YandexSpellChecker.Check"
	y.logger = y.logger.With("op", op)

	var inputItems []InputItem
	//Yandex speller принимает текст с '+' вместо пробелов
	modifiedText := strings.ReplaceAll(text, " ", "+")

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: transport,
	}

	resp, err := client.Get(apiURL + modifiedText)
	if err != nil {
		y.logger.Error("cannot get response", slog.String("url", apiURL+modifiedText))
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		y.logger.Error("cannot read body", slog.String("url", apiURL+modifiedText))
		return nil, err
	}

	err = json.Unmarshal(body, &inputItems)
	if err != nil {
		y.logger.Error("cannot unmarshal body", slog.String("url", apiURL+modifiedText))
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
		y.logger.Error("cannot marshal body", slog.String("url", apiURL+modifiedText))
		log.Fatalf("Error marshalling output JSON: %v", err)
	}
	return mistakes, nil
}
