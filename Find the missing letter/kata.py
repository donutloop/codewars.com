def find_missing_letter(chars):
    j = 0
    for i in range(1, len(chars)):
        diff = ord(chars[i]) - ord(chars[j])
        if diff > 1:
            return chr(ord(chars[j]) + 1)
        j = j + 1

    return
