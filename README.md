# GoPal [![Build Status](https://travis-ci.org/masaruhoshi/go-palindrome.svg?branch=master)](https://travis-ci.org/masaruhoshi/go-palindrome) [![Coverage Status](https://coveralls.io/repos/github/masaruhoshi/go-palindrome/badge.svg?branch=master)](https://coveralls.io/github/masaruhoshi/go-palindrome?branch=master)

`GoPal` is a simple set of REST apis built on [http://golang.org](Go) used to validate 
sets of words, numbers or phrases as palindromes.

## Palindromes

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

### Character normalization
Thanks to utf-8 representation, characteres in different languages can
be expressed in a single and common encoding, 8-bit based.
However, it also introduced new problems. For example, "é" can be 
either represented as a single character \u00e9 or the combination of
e and \u0301. Or the ⁹ and the regular digit 9. Or yet, 'K' ("\u004B") 
and 'K' (Kelvin sign "\u212A").

### Normalization
Although Go is able to handle character collation quite well using on
unnormalized strings, evaluating palindromes in different languages
require normalizing the string. To achieve we, I decided using the 
text/transform and text/unicode/norm packages.

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

## Requirements

- `go` 1.4+
- `MongoDB`
- `curl` or any REST client (I recommend [https://advancedrestclient.com/] Advanced REST Client)

## JSON Schema

    {
        "$schema": "http://json-schema.org/schema#",
        "title": "Palindrome",
        "type": "object",
        "required": ["ID", "phrase", "valid"],
        "properties": {
            "ID": {
                "type": "string",
                "description": "Unique identified"
            },
            "phrase": {
                "type": "string",
                "description": "Name of the product"
            },
            "valid": {
                "type": "boolean",
                "description": "Wether it's a valid palindrome or not"
            }
        }
    }

The JSON Schema above can be used to test the validity of the JSON code below:

    {
        "ID": "58eedfb5b7fc13821176df2c",
        "phrase": "Live on time, emit no evil",
        "valid": true
    }

## Endpoints

1. `GET /palindrome/`

List all palindromes entered

*Usage:* 

    curl http://localhost:8080/palindrome

*Result:*

    [
        {
            "ID": "58eedfb5b7fc13821176df2c",
            "phrase": "racecar",
            "valid": true
        }
    ]

*Alternative responses*:
* `null`: In case there're no records

2. `POST /palindrome/`

Add a new palindrome phrase to be validated.

*Usage:*

    curl -H "Content-Type: application/json" \
         -X POST -d '{"phrase": "racecar"}' \
         -i http://localhost:8080/palindrome

*Result:*

    HTTP/1.1 201 Created
    Date: Thu, 13 Apr 2017 02:30:47 GMT
    Content-Length: 0
    Content-Type: text/plain; charset=utf-8

*Alternative responses:*
* `HTTP/1.1 400 Bad Request`: A malformed JSON object was provided
* `HTTP/1.1 208 Already Reported`: When the palindrome was already provided
* `HTTP/1.1 500 Internal Server Error`: The database server must be down

3. `GET /palindrome/:id`

Display details about a specific palindrome

*Usage:*

    curl -i http://localhost:8080/palindrome/58eee2d7b7fc13821176df2d

*Result:*

    {
        "ID": "58eee2d7b7fc13821176df2d",
        "phrase": "Live on time, emit no evil",
        "valid": true
    }

*Alternative responses:*
1. `HTTP/1.1 412 Precondition Failed`: The ID provided isn't a valid hex value
2. `HTTP/1.1 404 Not Found`: There's not palindrome for the ID specified
3. `HTTP/1.1 500 Internal Server Error`: The database must be down.

4. `DELETE /palindrome/:id`

Deletes a given palindrome from its `id`

*Usage:*

    curl -X DELETE -i http://localhost:8080/palindrome/58eee2d7b7fc13821176df2d

*Result:*

    HTTP/1.1 202 Accepted
    Date: Thu, 13 Apr 2017 02:35:25 GMT
    Content-Length: 0
    Content-Type: text/plain; charset=utf-8

*Alternative responses:*
1. `HTTP/1.1 412 Precondition Failed`: The ID provided isn't a valid hex value
2. `HTTP/1.1 404 Not Found`: There's not palindrome for the ID specified
3. `HTTP/1.1 500 Internal Server Error`: The database must be down.

## Development

I strongly recommend using a docker container for development

```
export GOPATH=$PWD
export PATH=$PATH:/usr/local/go/bin
```

## Dependencies

* [https://github.com/julienschmidt/httprouter](HTTPRouter): lightweight router in Go
* [https://github.com/go-mgo/mgo](MGO): Mongo driver for Go
* [https://golang.org/x/text/transform](Transform) and [https://golang.org/x/text/unicode/norm](Norm): Normalization unicode characters



# TODO
* ~Dockerfile~
* Move application settings to its own struct
* ~Mongo host and credentials~
  * ~Database~
  * Collection names
* ~Have database functions to its own package~

# Licence
This project is licensed unter Apache License 2.0. You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
