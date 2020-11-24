def insertSort(l):
    for j in range(1, len(l)):
        chave = l[j]
        i = j-1
        while(i >=0 and l[i]>chave):
            l[i+1] = l[i]
            i= i-1
        l[i+1] = chave
        
l = [5,2,4,6,1,3]
print(l)
insertSort(l)
print(l)
