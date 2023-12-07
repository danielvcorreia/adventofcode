import time

def parseGamesPossible(file):
  cube_colors_max = [('red', 12), \
      ('green', 13), ('blue', 14)]
  possible = True
  sum_ids = 0

  for game in file.readlines():
    game_info = game.strip().split(': ')[1].split('; ')
    
    for reveal in game_info:
      cubes_revealed = reveal.split(', ')

      for color in cubes_revealed:
        info = color.split(' ')

        if info[1] == cube_colors_max[0][0]:
          if int(info[0]) > cube_colors_max[0][1]:
            possible = False
            break

        if info[1] == cube_colors_max[1][0]:
          if int(info[0]) > cube_colors_max[1][1]:
            possible = False
            break

        if info[1] == cube_colors_max[2][0]:
          if int(info[0]) > cube_colors_max[2][1]:
            possible = False
            break
      
      if not possible:
        break

    if possible:
      sum_ids += int(game.split(': ')[0].split(' ')[1])

    possible = True

  return sum_ids

def parseGamesMinCubes(file):
  cube_colors = ['red', 'green', 'blue']
  sum_power_sets = 0
  red_min = 0
  green_min = 0
  blue_min = 0

  for game in file.readlines():
    game_info = game.strip().split(': ')[1].split('; ')
    for reveal in game_info:
      cubes_revealed = reveal.split(', ')

      for color in cubes_revealed:
        info = color.split(' ')

        if info[1] == cube_colors[0]:
          if int(info[0]) > red_min:
            red_min = int(info[0])

        if info[1] == cube_colors[1]:
          if int(info[0]) > green_min:
            green_min = int(info[0])

        if info[1] == cube_colors[2]:
          if int(info[0]) > blue_min:
            blue_min = int(info[0])

    sum_power_sets += red_min * green_min * blue_min
    red_min = 0
    green_min = 0
    blue_min = 0

  return sum_power_sets


file = open('games.puzzle', 'r')

start = time.time()
#sum_ids = parseGamesPossible(file)
sum_power_sets = parseGamesMinCubes(file)
end = time.time()

#print(f'The sum of the IDs of those games is: \
#{sum_ids}\nCalculated in: {end-start} seconds.')
print(f'The sum of the power of these sets is: \
{sum_power_sets}\nCalculated in: {end-start} seconds.')