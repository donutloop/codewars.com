def get_length_of_missing_array(array_of_arrays):

    if len(array_of_arrays) == 0:
        return 0

    for array in array_of_arrays:
        if array == None or len(array) == 0:
            return 0

    def key(elem):
        return len(elem)

    array_of_arrays.sort(key=key)

    for index, item in enumerate(array_of_arrays, start=1):
        diff = len(array_of_arrays[index]) - len(array_of_arrays[index-1])
        if diff > 1:
            return len(array_of_arrays[index]) - 1