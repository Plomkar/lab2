#include <iostream>
using namespace std;

int main() {
    int N;
    cin >> N;

    int prev, curr;
    cin >> prev;
    int count = 0;

    for (int i = 1; i < N; i++) {
        cin >> curr;
        if (curr == prev) {
            count++;
        }
        prev = curr;
    }

    cout << "Результат: " << count << endl;
    return 0;
}