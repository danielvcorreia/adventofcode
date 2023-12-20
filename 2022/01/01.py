file = open("input.txt", "r")
bags = [elf.split() for elf in file.read().split(sep="\n\n")]
calories = [sum([eval(item) for item in bag]) for bag in bags]
calories.sort()
print(f'How many total Calories is that Elf carrying: {calories[-1]}')
print(f'How many Calories are those Elves carrying in total: {sum(calories[-3:])}')
