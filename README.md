# codecata.com - kata19 - Word Chains 

![codecata.com](https://imgur.com/download/CJozxMr)

This repository contains my personal [kata19 Word Chains](http://codekata.com/kata/kata19-word-chains/) implementation.

## Usage

I didn't used any dependencies, so no vendoring is needed.

```
cd cmd
go run main.go -start=cat -end=dog
```

## Initial planning 

After reading the kata, I realised that this kata were not going to be as easy as others.

I've wrote some services that used a graph database (neo4j) and this "linked" words problem did recall on me traversal search. 

I think that I would enjoy creating some "linked" word nodes and perform concurrency traversal search on them.

## Implementation

Finally I didn't created a full Go concurrent solution. Even when I'm using it for dictionary loading and to collect results.

I didn't do so as I realised that it would be so complex to read (while developing it), and It might affect the code review.

DISCLAIMER: I still can implement it.

This approach loads the words from the dictionary mentioned on the Kata comments and link them in ~2 min (i9 8 Cores Mackbook Pro).

This approach also takes few seconds to find a solution, but given the number words on the dictionary (>300k). I Ctrl-C the execution as chains keeps getting 500+ results easily.

## Posible next steps

- [ ] TODO - Move onto a full Go concurrency solution.
- [ ] TODO - Use a more efficient scoring approcach - there are chains that are more likely to produce better results. 
- [ ] TODO - Use a more efficient math approach.
- [ ] TODO - Create a better cli using urfave/cli or cobra.

## Analysis

I built the `words` package using the `-gcflags -m=2` to check for Heap scapes during the two most intensive parts:

- Words Initialization
- Traverse search

### Words Initialization

Looks like there is no Heap scapes. I was expecting them due the extensive use of pointer semantics that we use to relate the words between them:

```
/word.go:36:7: (*Word).isLinkable w does not escape
./word.go:36:27: (*Word).isLinkable word does not escape
./word.go:15:21: leaking param content: words
./word.go:15:21: 	from words[i] (dot of pointer) at ./word.go:18:47
./word.go:15:21: 	from append(w.LinkedWords, words[i]) (appended to slice) at ./word.go:18:26
./word.go:15:7: leaking param content: w
./word.go:15:7: 	from w.LinkedWords (dot of pointer) at ./word.go:18:28
./word.go:15:7: 	from *w.LinkedWords (indirection) at ./word.go:18:28
./word.go:15:7: 	from append(w.LinkedWords, words[i]) (appendee slice) at ./word.go:18:26
```
### Traverse search

There are some Heap scapes but not in the "intesive parts", just on the initialization and cleanup:

```
./traverse.go:102:5: t.Mutex escapes to heap
./traverse.go:102:5: 	from t.Mutex (passed to call[argument escapes]) at ./traverse.go:102:10
./traverse.go:94:7: leaking param: t
./traverse.go:94:7: 	from t.Mutex (dot of pointer) at ./traverse.go:102:5
./traverse.go:94:7: 	from t.Mutex (address-of) at ./traverse.go:102:5
./traverse.go:94:7: 	from t.Mutex (passed to call[argument escapes]) at ./traverse.go:102:10
./traverse.go:108:5: t.Mutex escapes to heap
./traverse.go:108:5: 	from t.Mutex (passed to call[argument escapes]) at ./traverse.go:108:12
./traverse.go:65:25: leaking param content: stepWord
./traverse.go:65:25: 	from stepWord.Term (dot of pointer) at ./traverse.go:66:32
./traverse.go:65:25: 	from append(chain, stepWord.Term) (appended to slice) at ./traverse.go:66:16
./traverse.go:65:41: leaking param content: chain
./traverse.go:65:41: 	from *chain (indirection) at ./traverse.go:65:41
./traverse.go:65:41: 	from append(chain, stepWord.Term) (appendee slice) at ./traverse.go:66:16
./traverse.go:65:41: leaking param: chain
./traverse.go:65:41: 	from t.Results <- chain (send) at ./traverse.go:69:13
./traverse.go:65:7: (*Traverse).step t does not escape
```
### Trace exploring

I used the trace tool, to check the word initialization, looks like GC and CPU utilization are fine during the process:

[Trace tool](https://i.imgur.com/IOU06rv.png)

