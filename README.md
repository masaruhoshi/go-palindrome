Palindromes are word, phrases, numbers, or other sequence of characters 
which reads the same backward as forward ([https://en.wikipedia.org/wiki/Palindrome]).

    racecar
    Was it a cat I saw?
    Dammit I'm Mad

Observe that phrase palindromes ignore punctuation, spaces and capital letters. In other languages where characters have variations due to accent notation, those characters are also converted to their primitive form in order to be considered as a palindrome.

	DÁBALE ARROZ A LA ZORRA EL ABAD
	SOCORRAM-ME, SUBI NO ÔNIBUS EM MARROCOS

Photenic languages like Japanese require special analisis to be considered as a palindrome. It's required to use their phonetic notation to evaluate them. In order words, ideograms representation (kanji) must be converted to phonetic format (hiragana) before.

	竹藪焼けた => たけやぶやけた
	私負けましたわ => わたしまけましたわ

## Character normalization
Thanks to utf-8 representation, characteres in different languages can
be expressed in a single and common encoding, 8-bit based.
However, it also introduced new problems. For example, "é" can be 
either represented as a single character \u00e9 or the combination of
e and \u0301. Or the ⁹ and the regular digit 9. Or yet, 'K' ("\u004B") 
and 'K' (Kelvin sign "\u212A").

## Normalization
Although Go is able to handle character collation quite well using on
unnormalized strings, evaluating palindromes in different languages
require normalizing the string. To achieve we, I decided using the 
text/transform and text/unicode/norm packages.

# Dev

```
export GOPATH=$PWD
export PATH=$PATH:/usr/local/go/bin
```

# Packages

    github.com/julienschmidt/httprouter
    gopkg.in/mgo.v2
    golang.org/x/text/transform
    golang.org/x/text/unicode/norm

# Examples

| English                              | Spanish / Portuguese                    | Japanese                 |
|--------------------------------------|-----------------------------------------|--------------------------|
| racecar                              | DÁBALE ARROZ A LA ZORRA EL ABAD         | たけやぶやけた              |
| Go hang a salami, I'm a lasagna hog  | ROMA ME TEM AMOR                        | わたしまけましたわ           |
| Rats live on no evil star            | SOCORRAM-ME, SUBI NO ÔNIBUS EM MARROCOS | なるとをとるな              |
| Live on time, emit no evil           |                                         | しなもんぱんもれもんぱんもなし |
| Mr. Owl ate my metal worm            |                                         | よのなかほかほかなのよ        |
| Was it a cat I saw?                  |                                         | たしかにかした               |
| Dammit I'm Mad                       |                                         |                           |

# TODO
* Dockerfile
* Move application settings to its own struct
 * ~Mongo host and credentials~
 * ~Database~
 * Collection names
* ~Have database functions to its own package~

# Licence
This project is licensed unter Apache License 2.0. You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
