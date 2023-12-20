file = open("input.txt", "r")
sections = [sections.split(sep=",") for sections in file.read().strip().split(sep="\n")]
n_pairs = 0

for pair in sections:
    elf1 = pair[0].split(sep="-")
    elf2 = pair[1].split(sep="-")

    if int(elf1[0]) <= int(elf2[0]) and int(elf1[1]) >= int(elf2[1]) or \
       int(elf1[0]) >= int(elf2[0]) and int(elf1[1]) <= int(elf2[1]):
        n_pairs += 1
    
print(f'In how many assignment pairs does one range fully contain the other? {n_pairs}')

n_pairs = 0
for pair in sections:
    elf1 = pair[0].split(sep="-")
    elf2 = pair[1].split(sep="-")

    if not ( int(elf1[1]) < int(elf2[0]) or int(elf2[1]) < int(elf1[0]) ):
        n_pairs += 1
    
print(f'In how many assignment pairs do the ranges overlap? {n_pairs}')