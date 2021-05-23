def string_to_array(s):
    words = []
    word = ""
    for i in range(0, len(s)):
        if s[i] == ' ':
            words.append(word)
            word = ""
            continue
        word += str(s[i])

    words.append(word)
    return words