
# To Do

- Init new game X
- Generate field X
- Display field in current state X
- Player interaction X
  - pick square X
  - flag X
- Restart/Quit
- Re-place mine on first click
- Leaderboard
- Win con X
- loss con X

## spaces (in order of precedence)

- Flag
- Covered
- Uncovered (Blank/number)
- Mine

## spaces code

### no mine

- Covered - X
- Uncovered (blank) - U
- Covered (numbered) - 1/2/3/4 etc.
- Uncovered (numbered) - U1/U2/U3/U4 etc.
- Flag - f

### mine

- Covered - !
- Uncovered - E
- Flag - F

## spaces Display

- X/!: #
- U: " "
- 1/2/3/4 - 1/2/3/4
- F/f - ⚑ U+2691
- E - ¤

## Board

Y

|

9 # # # # # # # # #
8 # # # # # # # # #
7 # # # # # # # # #
6 # # # # # # # # #
5 # # # # # # # # #
4 # # # # # # # # #
3 # # # # # # # # #
2 # # # # # # # # #
1 # # # # # # # # #
  1 2 3 4 5 6 7 8 9 - X

## bugs to fix

- F column interacts poorly with -f trim X
- player can choose co-ordinate outside of board range, exception caused as a result X
- Too many U characters added to numbered co-ords due to
  white space uncover script X
