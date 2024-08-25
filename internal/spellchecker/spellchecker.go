package spellchecker

type SpellChecker interface {
	Check(text string) ([]byte, error)
}
