def anagrams(word, words):
    sorted_word_a = sorted(word)
    sorted_word_a = "".join(sorted_word_a)
    match = []
    for i in range(len(words)):
        sorted_word_b = sorted(words[i])
        sorted_word_b = "".join(sorted_word_b)
        if sorted_word_b == sorted_word_a:
            match.append(words[i])

    return match