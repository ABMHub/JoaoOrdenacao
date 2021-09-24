import random

file = open("ints.txt", "ab")

i = 0
while (i < 1e6 * 0.8):
  file.write((str(random.randint(0, 1e9)) + "\n").encode("ascii"))
  i+=1