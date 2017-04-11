/*
Copyright 2017 Masaru Hoshi.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.

You may obtain a copy of the License at
     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Palindromes are word, phrases, numbers, or other sequence of characters 
which reads the same backward as forward [1].

Examples: 
	racecar
	Was it a cat I saw?
	Dammit I'm Mad

Observe that phrase palindromes ignore punctuation, spaces and capital
letters. In other languages where characters have variations due to 
accent notation, those characters are also converted to their primitive
form in order to be considered as a palindrome.

Examples:
	DÁBALE ARROZ A LA ZORRA EL ABAD
	SOCORRAM-ME, SUBI NO ÔNIBUS EM MARROCOS

Photenic languages like Japanese require special analisis to be 
considered as a palindrome. It's required to use their phonetic
notation to evaluate them. In order words, ideograms representation
(kanji) must be converted to phonetic format (hiragana) before.

Examples:
	竹藪焼けた => たけやぶやけた
	私負けましたわ => わたしまけましたわ

= Character normalization
Thanks to utf-8 representation, characteres in different languages can
be expressed in a single and common encoding, 8-bit based.
However, it also introduced new problems. For example, "é" can be 
either represented as a single character \u00e9 or the combination of
e and \u0301. Or the ⁹ and the regular digit 9. Or yet, 'K' ("\u004B") 
and 'K' (Kelvin sign "\u212A").

== Normalization
Although Go is able to handle character collation quite well using on
unnormalized strings, evaluating palindromes in different languages
require normalizing the string. To achieve we, I decided using the 
text/transform and text/unicode/norm packages.

= References
[1] https://en.wikipedia.org/wiki/Palindrome
*/

package models

import (
	"errors"
	"strings"
	"regexp"
	"unicode"
	"unicode/utf8"

	// Third party packages
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"gopkg.in/mgo.v2/bson"
)

type Palindrome struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Phrase	string	`json:"phrase"`
	Valid 	bool	`json:"valid"`
}

/*
Validates if the instance is a valid palindrome.

The function requires a method receiver passing a reference
to the Palindrome object as we want to enforce the predictability
of the behavior.
*/
func (p *Palindrome) Validate() error {
	word := p.Phrase
	if len(word) == 0 {
		return errors.New("Invalid length")
	}

	// Clean string before starting validation
	word = cleanString(word)

	// Compare runes 1st to last position up to middle position
	for len(word) > 0 {
	    first, sizeOfFirst := utf8.DecodeRuneInString(word)
	    if sizeOfFirst == len(word) {
	        break 
	    }
	    last, sizeOfLast := utf8.DecodeLastRuneInString(word)
	    if first != last {
	        return nil
	    }
	    word = word[sizeOfFirst : len(word)-sizeOfLast]
	}

	p.Valid = true
	return nil
}


/*
Clean up the candidate phrase before validation.

For a more detailed explanation about how the utf-8 normalization
works, take a look here: https://blog.golang.org/normalization.
*/
func cleanString(s string) string {
	f := func(r rune) bool {
    	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
	}

	// cleans up any punctuation and spaces before the normalization
	re := regexp.MustCompile("[[:punct:]]|[[:space:]]")
	s = re.ReplaceAllString(strings.ToLower(s), "")

 	t := transform.Chain(norm.NFD, transform.RemoveFunc(f), norm.NFC)
    result, _, _ := transform.String(t, s)

    return result
}
