file = open("input.txt", "r")
problem = [line.split(sep="\n") for line in file.read().split(sep="\n\n")]
stacks = []
stack_number = 0
blocks = ""

# Initialize stacks
for i in range(0, int(problem[0][len(problem[0])-1].strip().split(sep="   ")[-1])):
    stacks.append([])

# Initialize initial state of crates
for i in range(len(problem[0])-2, -1, -1):
    for j in range(1, len(problem[0][i]), 4):
        if problem[0][i][j] != ' ':
            stacks[stack_number].append(problem[0][i][j])
        stack_number += 1
    stack_number = 0

# Move crates
for instruction in problem[1]:
    words = instruction.split(sep=" ")
    for i in range(0, int(words[1])):
        item = stacks[int(words[3])-1].pop()
        stacks[int(words[5])-1].append(item)

# Get crates on top
for stack in stacks:
    blocks += stack.pop()

print(f'After the rearrangement procedure completes, what crate ends up on top of each stack? {blocks}')

stacks = []
blocks = ""

# Initialize stacks
for i in range(0, int(problem[0][len(problem[0])-1].strip().split(sep="   ")[-1])):
    stacks.append([])

# Initialize initial state of crates
for i in range(len(problem[0])-2, -1, -1):
    for j in range(1, len(problem[0][i]), 4):
        if problem[0][i][j] != ' ':
            stacks[stack_number].append(problem[0][i][j])
        stack_number += 1
    stack_number = 0

# Move crates
for instruction in problem[1]:
    words = instruction.split(sep=" ")
    for i in range(int(words[1]), 0, -1):
        item = stacks[int(words[3])-1].pop(-i)
        stacks[int(words[5])-1].append(item)

# Get crates on top
for stack in stacks:
    blocks += stack.pop()

print(f'After the rearrangement procedure completes, what crate ends up on top of each stack? {blocks}')
