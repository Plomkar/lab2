const readline = require('readline');

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

function minWindow(s, t) {
    if (s.length === 0 || t.length === 0) return "Не найдено";

    const need = {};
    for (let ch of t) {
        need[ch] = (need[ch] || 0) + 1;
    }

    let required = t.length;
    let left = 0;
    let minLen = Infinity;
    let start = 0;

    for (let right = 0; right < s.length; right++) {
        const rc = s[right];
        if (need[rc] !== undefined) {
            if (need[rc] > 0) required--;
            need[rc]--;
        }

        while (required === 0) {
            if (right - left + 1 < minLen) {
                minLen = right - left + 1;
                start = left;
            }

            const lc = s[left];
            if (need[lc] !== undefined) {
                need[lc]++;
                if (need[lc] > 0) required++;
            }
            left++;
        }
    }

    return minLen === Infinity ? "Не найдено" : s.substr(start, minLen);
}

rl.question("Введите строку S: ", (S) => {
    rl.question("Введите строку T: ", (T) => {
        const result = minWindow(S, T);
        console.log(`Результат: ${result}`);
        rl.close();
    });
});