# Game of life

It's just a toy I did on a boring Sunday after a Code Retreat Event 
in Munich. It's just to practise my Go skills, as I found this 
language very interesting and productive.

## Installation

```
$ go get github.com/leandrosansilva/toy_gameoflife
```

## Running
```
$ $GOPATH/bin/toy_gameoflife --config description.json -i
imported_glider=imported/dir/any_life.lif
```

Where description.json looks like this:

```json
{
  "Size": {
    "Width": 130,
    "Height": 48
  },

  "GenerationDuration": "100ms",

  "RandomCells": 3,

  "Circular": true,

  "Positions": [
    [10,10],
  ]

  "Species": {
    "lwss": [
      [0,1,1,0,0],
      [1,1,1,1,0],
      [1,1,0,1,1],
      [0,0,1,1,0]
    ]
  },

  "Population": [
    {
      "Specie": "imported_glider",
      "Position": [0,0]
    },

    {
      "Specie": "lwss",
      "Position": [40,20]
    }
  ]
}
```


If you do not want to download the source code but have Docker installed, 
first write a config.json file in the current directory and run:


```
$ docker run --rm -it -v $PWD:/toy:ro golang bash -c 'go get github.com/leandrosansilva/toy_gameoflife && toy_gameoflife --config /toy/config.json'
```
