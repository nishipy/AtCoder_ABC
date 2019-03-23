s = int(input())
l=[s]
i=1
def func(a):
  if a%2==0:
    return(a/2)
  else:
    return(3*a+1)

while (len(l)==len(set(l))):
  i+=1
  s = func(s)
  l.append(s)
print(i)
