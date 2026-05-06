def minWindow(s: str, t: str) -> str:
    if not s or not t:
        return "Не найдено"

    need = {}
    for ch in t:
        need[ch] = need.get(ch, 0) + 1

    required = len(t)
    left = 0
    min_len = float('inf')
    start = 0

    for right in range(len(s)):
        rc = s[right]
        if rc in need:
            if need[rc] > 0:
                required -= 1
            need[rc] -= 1

        while required == 0:
            if right - left + 1 < min_len:
                min_len = right - left + 1
                start = left

            lc = s[left]
            if lc in need:
                need[lc] += 1
                if need[lc] > 0:
                    required += 1
            left += 1

    return s[start:start + min_len] if min_len != float('inf') else "Не найдено"

# Ввод от пользователя
S = input("Введите строку S: ")
T = input("Введите строку T: ")

result = minWindow(S, T)
print(f"Результат: {result}")