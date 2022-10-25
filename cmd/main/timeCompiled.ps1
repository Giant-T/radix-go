Measure-Command -Expression { go build -o a.exe }

Measure-Command -Expression { ./a.exe | Out-Default }