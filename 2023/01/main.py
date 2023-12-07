import time

file = open('calibrations.puzzle', 'r')
sum_calibrations = 0
f_search = True
l_search = True
f_index = 0
l_index = -2
spelled_digits = ['one', 'two', 'three', 'four', \
   'five', 'six', 'seven', 'eight', 'nine']

f_digit_mem = ''
l_digit_mem = ''

start = time.time()

for line in file.readlines():
  while f_search or l_search:
    if not line[f_index].isalpha() and f_search:
      f_digit = line[f_index]
      f_search = False
      f_digit_mem = ''
    else:
      f_digit_mem += line[f_index]
      if len(f_digit_mem) > 2 and f_search:
        for i in range(0, len(spelled_digits)):
          if spelled_digits[i] in f_digit_mem:
            f_digit = str(i+1)
            f_search = False
    if not line[l_index].isalpha() and l_search:
      l_digit = line[l_index]
      l_search = False
    else:
      l_digit_mem = line[l_index] + l_digit_mem
      if len(l_digit_mem) > 2 and l_search:
        for i in range(0, len(spelled_digits)):
          if spelled_digits[i] in l_digit_mem:
            l_digit = str(i+1)
            l_search = False
    
    f_index += 1
    l_index -= 1

  sum_calibrations += int(f_digit + l_digit)
  f_search = True
  l_search = True
  f_index = 0
  l_index = -2
  f_digit_mem = ''
  l_digit_mem = ''

end = time.time()

print(f'The sum of all the calibrations is: \
{sum_calibrations}\nCalculated in: {end-start} seconds.')