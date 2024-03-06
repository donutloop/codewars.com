def make_parts(arr, chunk_size):
    
    chunks = []
    c = 0
    currentChunk = []
    for i in range(len(arr)):
        if (c == chunk_size):
            chunks.append(currentChunk)
            currentChunk = []
            c = 0
        currentChunk.append(arr[i])
        c = c + 1
    if c > 0:
        chunks.append(currentChunk)
        
    return chunks