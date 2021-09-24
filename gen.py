import random

file = open("ints.txt", "a")

i = 0
while (i < 1e6):
  file.write(str(random.randint(0, 1e9)) + "\n")
  i+=1