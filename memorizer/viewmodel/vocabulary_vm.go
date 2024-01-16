package viewmodel

type VM_vocabulary struct {
	VocabularyID  int    `json:"vocabulary_id"`
	Vocabulary    string `json:"vocabulary"`
	Detail        string `json:"detail"`
	ReferenceFrom string `json:"reference_from"`
	UsageType     string `json:"usage_type"`
	Antonyms      string `json:"antonyms"`
	Synomyms      string `json:"synonyms"`
	PartsofSpeech string `json:"parts_of_speech"`
}
