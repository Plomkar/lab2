defmodule CountEqual do
  def main do
    _n = IO.gets("") |> String.trim() |> String.to_integer()

    numbers = IO.gets("") |> String.trim() |> String.split() |> Enum.map(&String.to_integer/1)

    {_, count} = Enum.reduce(numbers, {nil, 0}, fn curr, {prev, acc} ->
      if prev != nil and curr == prev do
        {curr, acc + 1}
      else
        {curr, acc}
      end
    end)

    IO.puts("Результат: #{count}")
  end
end

CountEqual.main()
