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
- [ ] TODO - Use a more efficient math approach.
- [ ] TODO - Create a better cli using urfave/cli or cobra.
