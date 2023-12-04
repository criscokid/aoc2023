import os
import shutil

day = input("Enter day: ")

if os.path.isdir(f"cmd/{day}"):
    print(f"{day} directory already exists.")
    quit()

os.makedirs(f"cmd/{day}/part1")

os.mknod(f"cmd/{day}/part1/input.txt")
os.mknod(f"cmd/{day}/part1/small_input.txt")

shutil.copyfile("main.go.start", f"cmd/{day}/part1/main.go")
