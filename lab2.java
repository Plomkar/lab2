import java.util.*;

class Meeting {
    int start, end, index;
    Meeting(int start, int end, int index) {
        this.start = start;
        this.end = end;
        this.index = index;
    }
}

public class lab2 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        
        System.out.print("Введите количество заявок: ");
        int N = sc.nextInt();
        
        List<Meeting> meetings = new ArrayList<>();
        for (int i = 0; i < N; i++) {
            int start = sc.nextInt();
            int end = sc.nextInt();
            meetings.add(new Meeting(start, end, i + 1));
        }
        
        meetings.sort((a, b) -> {
            if (a.end == b.end) return a.start - b.start;
            return a.end - b.end;
        });
        
        List<Integer> selectedIndices = new ArrayList<>();
        int lastEnd = -1;
        
        for (Meeting m : meetings) {
            if (m.start >= lastEnd) {
                selectedIndices.add(m.index);
                lastEnd = m.end;
            }
        }
        
        Collections.sort(selectedIndices);
        
        System.out.println("Результат: " + selectedIndices.size());
        for (int idx : selectedIndices) {
            System.out.print(idx + " ");
        }
        System.out.println();
        
        sc.close();
    }
}