def item_priority(item):
    if item.isupper():
        return ord(item)-38
    else:
        return ord(item)-96

file = open("input.txt", "r")
rucksacks = [rucksack for rucksack in file.read().split(sep="\n")]
score = 0

for rucksack in rucksacks:
    half = int(len(rucksack)/2)

    for i in range(0, half):
        if rucksack[i] in rucksack[half:]:
            score += item_priority(rucksack[i])
            rucksack = rucksack[:half] + \
                rucksack[half:].replace(rucksack[i], "")

print(f'What is the sum of the priorities? {score}')

score = 0
for i in range(0, len(rucksacks), 3):
    for item in rucksacks[i]:
        if item in rucksacks[i+1] and item in rucksacks[i+2]:
            score += item_priority(item)
            break

print(f'What is the sum of the priorities of those item types? {score}')