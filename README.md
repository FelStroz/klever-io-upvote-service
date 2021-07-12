<div markdown="1" align="center">

![](./front/src/assets/klever.png)

</div>

<div markdown="1" align="center">

# Klever IO Upvote Service

Service to upvote or downvote a cryptocurrency as a test to Klever IO

![Platform](https://img.shields.io/npm/l/react?label=GO&logo=go)
![](https://img.shields.io/npm/l/react?label=React&logo=react)
![Platform](https://img.shields.io/badge/Styling-Css-lightgreen.svg?style=flat)

<br>

[![Hits](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2FFelStroz%2Fklever-io-upvote-service&count_bg=%2379C83D&title_bg=%23888888&icon=go.svg&icon_color=%235355C9&title=Acessos&edge_flat=false)](https://github.com/FelStroz/klever-io-upvote-service)

</div>

- ## Instalation

---

- With Makefile

Run the command bellow to start the server
```shell
$ make server
```

And the command bellow to start the client
```shell
$ make client
```

Install all dependencies for the frontend
```shell
$ make install-front
```

Run the frontend
```shell
$ make run-front
```

<br>

If you want tests just run
```shell
$ make test
```

<br>

- Without Makefile

In the root directory run
```shell
$ go run server/server.go
```
And then
```shell
$ go run client/main.go 
```
To run the frontend you need to install all dependencies with NPM or yarn
```shell
$ cd front
$ npm i
$ npm start
```
To run the tests run the command bellow
```shell
$ cd server
$ go test -v 
```

Done.

<div markdown="1" align="end">

Made by [Felipe Strozberg](https://github.com/FelStroz)

</div>

### [Back to the Top](#klever-io-upvote-service)