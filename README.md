# Meli Morse @danieljvx

### `Production meli-morse:` [https://danieljvx-meli-morse.herokuapp.com](https://danieljvx-meli-morse.herokuapp.com)
| Resource | Type | Path | Body |
| ------ | ------ | ------ | ------ |
| decodeBits2Morse | POST | [https://danieljvx-meli-morse.herokuapp.com/translate/decodeBits](https://danieljvx-meli-morse.herokuapp.com/translate/decodeBits) | { "text": "string" } |
| translate2Human | POST | [https://danieljvx-meli-morse.herokuapp.com/translate/2text](https://danieljvx-meli-morse.herokuapp.com/translate/2text) | { "text": "string" } |
| translate2Morse | POST | [https://danieljvx-meli-morse.herokuapp.com/translate/2morse](https://danieljvx-meli-morse.herokuapp.com/translate/2morse) | { "text": "string" } |


### `Up conntainer:`
```bash
docker-compose up --build
```

### `Run App:`
```bash
go run .
```
### `meli-morse:`
| Resource | Type | Path | Body |
| ------ | ------ | ------ | ------ |
| decodeBits2Morse | POST | [http://localhost:8000/translate/decodeBits](http://localhost:8000/translate/decodeBits) | { "text": "string" } |
| translate2Human | POST | [http://localhost:8000/translate/2text](http://localhost:8000/translate/2text) | { "text": "string" } |
| translate2Morse | POST | [http://localhost:8000/translate/2morse](http://localhost:8000/translate/2morse) | { "text": "string" } |

### `Test App:`
```bash
go test -v ./... -cover
```

## Daniel Villanueva

Email: [villanueva.danielx@gmail.com](mail://villanueva.danielx@gmail.com)

Github: [@danieljvx](https://github.com/danieljvx)
