def insertSort(l):
    for j in range(1, len(l)):
        chave = l[j]
        i = j-1
        while(i >=0 and l[i]>chave):
            l[i+1] = l[i]
            i= i-1
        l[i+1] = chave
        
print("Printe todos os números separados por espaços")
l = [int(i) for i in input().split(" ") ]
print(l)
insertSort(l)
print(l)
