#include <iostream>
#include <string>
#include <climits>
#include <vector>

using namespace std;

string minWindow(string s, string t) {
    if (s.empty() || t.empty()) return "Не найдено";

    vector<int> need(128, 0);
    for (char c : t) need[c]++;

    int left = 0, right = 0;
    int required = t.length();
    int minLen = INT_MAX;
    int start = 0;

    while (right < s.length()) {
        char rc = s[right];
        if (need[rc] > 0) required--;
        need[rc]--;
        right++;

        while (required == 0) {
            if (right - left < minLen) {
                minLen = right - left;
                start = left;
            }

            char lc = s[left];
            need[lc]++;
            if (need[lc] > 0) required++;
            left++;
        }
    }

    return minLen == INT_MAX ? "Не найдено" : s.substr(start, minLen);
}

int main() {
    string S, T;
    
    cout << "Введите строку S: ";
    cin >> S;
    cout << "Введите строку T: ";
    cin >> T;
    
    string result = minWindow(S, T);
    cout << "Результат: " << result << endl;
    
    return 0;
}