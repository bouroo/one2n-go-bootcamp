# A game of Pig exercise
Write a command-line program that simulates the game of Pig.
Pig is a two-player game played with a 6-sided die.
The game has the following rules:
Each turn, a player repeatedly rolls a die until either a 1 is rolled or the player decides to “hold”:
- If the player rolls a 1, they score nothing, and it becomes the next player’s turn.
- If the player rolls any other number, it is added to their turn total, and their turn continues. The player can then decide to continue to roll again or “hold”.
- If a player chooses to “hold”, their current turn total is added to their score, and it becomes the next player’s turn.

The first player to score 100 or more points wins.
For example, the first player, Donald, begins a turn with a roll of 5. Donald could hold and score 5 points but chooses to roll again. Donald rolls a 2 and could hold with a turn total of 7 points but decides to roll again. Donald rolls a 1 and must end his turn without scoring.
The next player, Alexis, rolls the sequence 4-5-3-5-6, after which she chooses to hold and adds her turn total of 23 points to her score.
The game continues, Donald taking the turn. In this turn, Donald rolls 3-2-4 and decides to hold. His total score so far is 9, and the turn passes to Alexis.
The player reaching a score of 100 or more total points wins.
Write a program in Go that simulates this two-player game. There are no human players, only computers. Each computer player plays with some fixed strategy. An example strategy is “hold after accumulating a score of at least 10 for each turn”. In this strategy, the player will keep on rolling until they accumulate a score of at least 10 in that round. If they roll a 1 before getting a turn total of 10, they score 0 points, and their turn passes to the next player.
Some example rolls for this strategy could be: - 3-4-2-4 (turn total is 13, which is greater than 10) and then the player holds and passes the turn to the next player. - 3-5-1 (turn total is 0, since a 1 is rolled) and then the player must pass the turn to the next player. - 6-4 (turn total is 10), hence player holds and passes the turn to the next player.
Similarly, another computer player may have a strategy to “hold after accumulating a score of at least 25 for each turn”. This player is “greedy” because they try to accumulate more points in a single round. This also means they may roll 1, and not get any points for that round. But, well, that’s their strategy.
The program should run a simulation where 10 such games are played between two players, where each player uses some fixed strategy. The program should output the number of wins and losses of each player.