# termcolor

## Installation

```bash
go install github.com/yasukotelin/termcolor
```

## Usage

```bash
termcolor
```

<img src="images/termcolor.png" />

## Using with Docker

### Image build

```
$ git clone https://github.com/yasukotelin/termcolor.git
$ cd termcolor
$ docker image build -t termcolor:latest .
```

### Run

```
$ docker container run -it termcolor

root@***:/go/src/github.com/yasukotelin/termcolor# ./termcolor
```

