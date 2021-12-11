# aoc2021

https://adventofcode.com/2021

| Day                                                                   | Source             |
|-----------------------------------------------------------------------|--------------------|
| [Day 1: Sonar Sweep](https://adventofcode.com/2021/day/1)             | [day1.go](day1.go) |
| [Day 2: Dive](https://adventofcode.com/2021/day/2)                    | [day2.go](day2.go) |
| [Day 3: Binary Diagnostic](https://adventofcode.com/2021/day/3)       | [day3.go](day3.go) |
| [Day 4: Giant Squid](https://adventofcode.com/2021/day/4)             | [day4.go](day4.go) |
| [Day 6: Lanternfish](https://adventofcode.com/2021/day/6)             | [day6.go](day6.go) |
| [Day 7: The Treachery of Whales](https://adventofcode.com/2021/day/7) | [day7.go](day7.go) |

```shell
goos: linux
goarch: amd64
pkg: github.com/karlovskiy/aoc2021
cpu: Intel(R) Core(TM) i7-4770K CPU @ 3.50GHz
BenchmarkSonarSweep-8                         	   20122	     57504 ns/op
BenchmarkSonarSweepThreeMeasurement-8         	   18186	     65947 ns/op
BenchmarkDive-8                               	   10000	    123018 ns/op
BenchmarkDiveWithAim-8                        	    9742	    120727 ns/op
BenchmarkBinaryDiagnostic-8                   	   21632	     54027 ns/op
BenchmarkBinaryDiagnosticPowerConsumption-8   	    1746	    692155 ns/op
BenchmarkGiantSquid-8                         	    6088	    201350 ns/op
BenchmarkGiantSquidLast-8                     	    2637	    428137 ns/op
BenchmarkLanternfish-8                        	  190245	      5832 ns/op
BenchmarkLanternfish256-8                     	   93782	     11553 ns/op
BenchmarkTreacheryOfWhales-8                        1086	   1126557 ns/op
BenchmarkTreacheryOfWhalesExpensive-8                 54	  22171851 ns/op
```