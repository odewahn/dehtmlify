Converts all HTMLBook files into Markdown (i.e., all files with a `.html` extension) using [pandoc](http://johnmacfarlane.net/pandoc/demo/example9/pandocs-markdown.html).  It will strip out all `section` tags, and then run `pandoc` to convert the "stripped" html friles into markdown.  It's hacky, but it opens the door to using [ipymd](https://github.com/rossant/ipymd) with Jupyter.


## Requirements

Install [pandoc](http://johnmacfarlane.net/pandoc/installing.html)

## To Build

```
go build -o htmlbook2md *.go
```


## To Run

* Change into the directory where you have from HTMLBook files
* Run `htmlbook2md` to convert all the things

It would be a big improvement to have this thing also only run in a non-master branch in git, but that's just too much to ask for this crappy old hack.
