Game of life

It's just a toy I did on a boring Sunday after a Code Retreat Event 
in Munich. It's just to practise my Go skills, as I found this 
language very interesting and productive.

If you do not want to download the source code but have Docker installed, 
first write a config.json file in the current directory with the content:

```json
{
  "Size": {
    "Width": 130,
    "Height": 48
  },

  "GenerationDuration": "100ms",

  "RandomCells" : 10,

  "Circular": true,

  "Positions": [
    [10,10],
    [10,11],
    [10,12],

    [20,12],
    [20,12],
    [21,12],
    [22,12],
    [23,12],
    [24,12],
    [26,12],
    [27,12],
    [27,12],
    [28,12],
    [28,13],
    [28,14],
    [28,15]
  ]
}
```

And then execute the command:

```
$ docker run --rm -it -v $PWD:/toy:ro golang bash -c 'go get github.com/leandrosansilva/toy_gameoflife && toy_gameoflife --config /toy/config.json'
```
