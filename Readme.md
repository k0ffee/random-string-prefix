Sometimes one needs to prepend a random string to names or the like,
which must not start with a digit, but we want digits in general.

Random strings are created by constructing an alphabet and sampling
single entities from it (sampling with put-back).

Possible number of permutations for 20 different letters and three
different digits, but no digits in position one:

```
unix% dc
23 5^p
6436343
20*p
128726860
```
