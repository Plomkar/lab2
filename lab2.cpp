#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

struct Meeting {
    int start, end, index;
};

bool compareByEnd(const Meeting &a, const Meeting &b) {
    if (a.end == b.end) return a.start < b.start;
    return a.end < b.end;
}

int main() {
    int N;
    cout << "Введите количество заявок: ";
    cin >> N;
    
    vector<Meeting> meetings(N);
    for (int i = 0; i < N; i++) {
        cin >> meetings[i].start >> meetings[i].end;
        meetings[i].index = i + 1; // нумерация с 1
    }
    
    // Сортируем по времени окончания
    sort(meetings.begin(), meetings.end(), compareByEnd);
    
    vector<int> selectedIndices;
    int lastEnd = -1;
    
    for (const auto &m : meetings) {
        if (m.start >= lastEnd) {
            selectedIndices.push_back(m.index);
            lastEnd = m.end;
        }
    }
    
    // Сортируем выбранные индексы по возрастанию (порядок поступления)
    sort(selectedIndices.begin(), selectedIndices.end());
    
    cout << "Результат: " << selectedIndices.size() << endl;
    for (int idx : selectedIndices) {
        cout << idx << " ";
    }
    cout << endl;
    
    return 0;
}