def marker(code_length):   
    has_repeated_chars = True
    count = code_length

    with open('input.txt', 'r') as file:
        code = file.read(code_length)

        while has_repeated_chars:
            has_repeated_chars = len(set(code)) != len(code)

            if not has_repeated_chars:
                file.close()
                return count

            else:
                code = code[1:code_length] + file.read(1)
                count += 1
            
            if not code[code_length-1]:
                file.close()
                return

print(f'How many characters need to be processed before the first start-of-packet marker is detected? {marker(4)}')

print(f'How many characters need to be processed before the first start-of-message marker is detected? {marker(14)}')
