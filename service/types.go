package service

type RandomJoke struct {
	Joke string `json:"joke"`
}

type DecodedRandomJoke struct {
	Jokes []RandomJoke `json:"jokes"`
}

/*
./bin/joke
-----
200
[{"joke": "What do you put on a lonely grilled cheese sandwich? Provolone, but only if you have it\u2019s parmesan."}]%
*/
