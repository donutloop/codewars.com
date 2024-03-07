def solve(arr):
    counter = {}
    for i in range(len(arr)):
        if arr[i] in counter:
            counter[arr[i]] = counter[arr[i]] + 1
        else:
            counter[arr[i]] = 0
        
    arr = sorted(arr, key=lambda num: (counter[num], -num), reverse=True)    
    
    return arr