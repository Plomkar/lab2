section .data
    fmt_in db "%d", 0
    fmt_out db "Результат: %d", 10, 0

section .bss
    N resd 1
    prev resd 1
    curr resd 1
    count resd 1
    i resd 1

section .text
    extern scanf, printf
    global main

main:
    push rbp
    mov rbp, rsp
    
    ; Ввод N
    mov rdi, fmt_in
    mov rsi, N
    call scanf

    ; Ввод первого числа
    mov rdi, fmt_in
    mov rsi, prev
    call scanf

    mov dword [count], 0
    mov dword [i], 1

.loop:
    mov eax, [i]
    cmp eax, [N]
    jge .end_loop

    ; Ввод следующего числа
    mov rdi, fmt_in
    mov rsi, curr
    call scanf

    ; Сравнение
    mov eax, [curr]
    cmp eax, [prev]
    jne .not_equal

    inc dword [count]

.not_equal:
    ; prev = curr
    mov eax, [curr]
    mov [prev], eax

    inc dword [i]
    jmp .loop

.end_loop:
    ; Вывод результата
    mov rdi, fmt_out
    mov esi, [count]
    call printf

    mov rsp, rbp
    pop rbp
    ret