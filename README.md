# WebSearch Server
Server that use the power of bing web search

# How to use
First build it:

``` go build websearch.go ```

The web server is listening to environment variable PORT, example:

```
PORT=3030 ./websearch
```

## Parameters
- q
- keyword
- advancedQuery
- tor
- header
- count

## Example
Output is JSON
``` 
http://localhost:3030/?keyword=calculus+theory&count=20 
```
will output the 20 search results (if any)
 
```
http://localhost:3030/?keyword=calculus&count=20&header=true
```
will output 20 search results and each result has header field

```
http://localhost:3030/?keyword=calculus&count=20&tor=true
```
will output 20 search results, but use tor connection when searching

```
http://localhost:3030/?keyword=best+sort+algorithm&advancedQuery=filetype:pdf&count=20
```
will output 20 search results, with filetype pdf
 
```
http://localhost:3030/?q=best+sort+algorithm+filetype:pdf&count=20
```
Use q is alias for keyword + advancedQuery. But, you will get ugly logs

# TODO
- Better code documentation
- Better readme
- Fix some bugs
- Fix typo
- Refactor code

# License
The MIT License (MIT)

Copyright (c) 2015 Ribhararnus Pracutiar

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

