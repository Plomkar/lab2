import java.util.Scanner

fun main() {
    val scanner = Scanner(System.`in`)
    val N = scanner.nextInt()

    var prev = scanner.nextInt()
    var count = 0

    for (i in 1 until N) {
        val curr = scanner.nextInt()
        if (curr == prev) {
            count++
        }
        prev = curr
    }

    println("Результат: $count")
}