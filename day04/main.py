test = """Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
"""


def num_matching(card_line):
    nums = card_line.split(":")[1]
    winning_nums, my_nums = nums.split("|")
    print(winning_nums, my_nums)

    winning_nums = set(map(int, filter(bool, winning_nums.split(" "))))
    my_nums = set(map(int, filter(bool, my_nums.split(" "))))

    return sum((1 for num in my_nums if num in winning_nums))


def part1(input_: str):
    total = 0
    for line in input_.splitlines(keepends=False):
        if matching := num_matching(line):
            total += (matching - 1) ** 2


def part2(input_: str):
    lines = input_.splitlines(keepends=False)
    n = len(lines)
    counts = [1] * n
    for i, line in enumerate(lines):
        matching = num_matching(line)
        print(i + 1, "matching", matching)
        for j in range(i + 1, i + matching + 1):
            counts[j] += counts[i]

    print(counts)
    print(sum(counts))


with open("input.txt", "r") as f:
    input = f.read()

part1(input)
part2(input)
