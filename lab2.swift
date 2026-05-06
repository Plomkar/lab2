import Foundation

struct Meeting {
    let start: Int
    let end: Int
    let index: Int
}

print("Введите количество заявок:", terminator: " ")
guard let N = Int(readLine() ?? "") else { fatalError() }

var meetings: [Meeting] = []

for i in 1...N {
    let input = readLine() ?? ""
    let parts = input.split(separator: " ").map { Int($0)! }
    meetings.append(Meeting(start: parts[0], end: parts[1], index: i))
}

meetings.sort { 
    if $0.end == $1.end {
        return $0.start < $1.start
    }
    return $0.end < $1.end
}

var selectedIndices: [Int] = []
var lastEnd = -1

for m in meetings {
    if m.start >= lastEnd {
        selectedIndices.append(m.index)
        lastEnd = m.end
    }
}

selectedIndices.sort()

print("Результат:", selectedIndices.count)
print(selectedIndices.map { String($0) }.joined(separator: " "))