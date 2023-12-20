file = open("input.txt", "r")
games = [games.split() for games in file.read().strip().split(sep="\n")]

rules = { 'A': {'X': 4,'Y': 8,'Z': 3},
          'B': {'X': 1,'Y': 5,'Z': 9}, 
          'C': {'X': 7,'Y': 2,'Z': 6} }

score = sum([rules[game[0]][game[1]] for game in games])
print(f'What would your total score be: {score}')

correct_rules = { 'A': {'X': 3,'Y': 4,'Z': 8},
                  'B': {'X': 1,'Y': 5,'Z': 9}, 
                  'C': {'X': 2,'Y': 6,'Z': 7} }

score = sum([correct_rules[game[0]][game[1]] for game in games])
print(f'What would your total score be: {score}')

