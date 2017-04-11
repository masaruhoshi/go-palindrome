package models

import (
	"testing"

	"pkg/utils"
)

func TestCleanStringToReturnOnlyCharactersInLowercase(t *testing.T) {
	phrase := "Go hang a salami, I'm a lasagna hog"
	expected := "gohangasalamiimalasagnahog"

	actual := cleanString(phrase)
	utils.Expect(t, actual, expected)
}

func TestCleanStringToReturnCleanedJapaneseCharacters(t *testing.T) {
	phrase := "たけやぶやけた"
	expected := "たけやふやけた"

	actual := cleanString(phrase)
	utils.Expect(t, actual, expected)
}

func TestCleanStringToReturnCleanedAccentCharacters(t *testing.T) {
	phrase := "w͢͢͝h͡o͢͡ ̸͢k̵͟n̴͘ǫw̸̛s͘ ̀́w͘͢ḩ̵at ̧̕h́o̵r͏̵rors̡ ̶͡͠lį̶e͟͟ ̶͝in͢ ͏t̕h̷̡͟e ͟͟d̛a͜r̕͡k̢̨ ͡h̴e͏a̷̢̡rt́͏ ̴̷͠ò̵̶f̸ u̧͘ní̛͜c͢͏o̷͏d̸͢e̡͝?͞"
	expected := "whoknowswhathorrorslieinthedarkheartofunicode"

	actual := cleanString(phrase)
	utils.Expect(t, actual, expected)
}

func TestValidateAssertsPalindromes(t *testing.T) {
	phrases := []string{
		"Go hang a salami, I'm a lasagna hog",
		"racecar",
		"Rats live on no evil star",
		"Live on time, emit no evil",
		"Mr. Owl ate my metal worm",
		"Was it a cat I saw?",
		"Dammit I'm Mad",

		"DÁBALE ARROZ A LA ZORRA EL ABAD",
		"ROMA ME TEM AMOR",
		"SOCORRAM-ME, SUBI NO ÔNIBUS EM MARROCOS",

		"たけやぶやけた",
		"わたしまけましたわ",
		"なるとをとるな",
		"しなもんぱんもれもんぱんもなし",
		"よのなかほかほかなのよ",
		"たしかにかした",
	}

	var palindrome Palindrome
	for _, phrase := range phrases {
		palindrome = Palindrome{Phrase: phrase}
		palindrome.Validate()

		utils.Expect(t, palindrome.Valid, true)
	}
}

func TestValidateRefutePalindromes(t *testing.T) {
	phrases := []string{
		"Not a valid palindrome",
		"Just an usual phrase",
		"Il n'y a pas qqch d'intéressant ici",
		"Uma frase qualquer sem conexão",
		"一期一会",
	}

	var palindrome Palindrome
	for _, phrase := range phrases {
		palindrome = Palindrome{Phrase: phrase}
		palindrome.Validate()

		utils.Expect(t, palindrome.Valid, false)
	}
}