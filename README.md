# Go-Solve-Wordle

Golang module that processes a Wordle-style result against a dictionary of
five-letter words to quickly identify the solution.

# Modes

## Demonstration

Syntax:

```
go build
./go-solve-wordle -mode demo
```

Randomly selects a word from a list of all Wordle solutions, then uses every
five-letter word from the dictionary to determine the best guesses to make until
finding the solution.

## Solution Finder - FUTURE

Syntax:

```

```

Provides the best first-guess based on the entire dictionary of words, then
processes a given result to reduce the possible options, and again find the best
next guess, until a solution is found.

## Game Mode

Syntax:

```
go build
./go-solve-wordle -mode play
```

Randomly selects a word from a list of all Wordle solutions then allows the user
to provide a guess and gives the result.

# Results Syntax

Possible results for each character:

- `-`: This character is not found anywhere in the answer.
- `+`: This character is not found in the current position, but is found in the
       answer.
- `*`: All other result characters indicate the correct character in an answer.

### Examples

```
answer: never
guess:  cares
result: --+e-
```

# Logic

This program attempts to find a guess that uses the most common characters in
every position. It uses a simple count logic to assign an overall weight to every
available word. Each letter in a word is given a value equal to that letter's
frequency in the same position in every other word in the list. The total weight
is the sum of the five character weights.

## Example

```
Available Words
- cares
- score
- stole
- chore
- stamp

Character Weights:
- Position 1:
  - c: 2
  - s: 3
- Position 2:
  - a: 1
  - c: 1
  - h: 1
  - t: 2
- Position 3:
  - a: 1
  - o: 3
  - r: 1
- Position 4:
  - e: 1
  - l: 1
  - m: 1
  - r: 2
- Position 5:
  - e: 3
  - p: 1
  - s: 1

Word Weights:
- cares: 2+1+1+1+1 = 6
- score: 3+1+3+2+3 = 12
- stole: 3+2+3+1+3 = 12
- chore: 2+1+3+2+3 = 11
- stamp: 3+2+1+1+1 = 8
```

In this example, both `score` and `stole` have a weight of `12`. In the event
of a tie, the word that was weighted first remains the chosen guess.

## Future Improvements

### Prefer most common character:

if `e` in position 5 is the most common letter in all positions, the program
could prefer to find a next-guess that uses that specific character+position
combination, before weighing additional characters.

### Consider duplicate characters:

If the highest scoring word uses the same character in multiple positions, the
program may reduce the weight of this word in attempt to eliminate more
characters in a single guess.

### Fix Result for Duplicate Characters:

Right now, the result logic doesn't acknowledge the unique cases for multiple
characters.

Doesn't identify single instance of a character when the guess contains two.
```
  word:   cares
  guess:  stash
  result: +-++-
```

Doesn't identify when a single instance of a character has already been used.
```
  word:   cares
  guess:  stars
  result: +-++s
```

### Better Code

I made this as a project to learn Go. A lot of it was written from official docs
and random forum/blog posts on similar problems I was trying to solve. I hope to
come back to this project occasionally when I want to try a new way to solve
this type of problem or to make structural improvements as I learn Go more.
