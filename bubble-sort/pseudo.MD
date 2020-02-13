# Buble Sort
## Receive an array
## Sort from higher to less this array
## Return sorted array

1.- Ask if *arr* length is zero. If is true return *sortedArr* or *arr*, else continue.
2.- Save i equal to one.
3.- Ask if *arr* length is equal to i, if true skip sixth step, else continue.
4.- Take i and (i + 1) element, check if first is greater than second. If true then interchange position. Increment i once.
5.- Jump to third step.
6.- Unshift last element to *sortedArr* and remove from *arr*.
7.- Jump to first step with new *arr*.

```
[3 2 1] -> [2 3 1] -> [2 1 3] -> [2 1 3*] ---- [3]
[2 1] -> [1 2] -> [1 2*]                  ---- [2 3]
[1] -> [1*]                               ---- [1 2 3]

[8 2 7 9] -> i = 1 -> [2 8 7 9] -> i = 2 -> [2 7 8 9] -> i = 3 -> [2 7 8 9] -> i = 4 => [9]
[2 7 8]   -> i = 1 -> [2 7 8]   -> i = 2 -> [2 7 8]   -> i = 3 => [8 9]
[2 7]     -> i = 1 -> [2 7]     -> i = 2 => [7 8 9]
[2]       -> i = 1 => [2 7 8 9]
[] => [2 7 8 9]
```