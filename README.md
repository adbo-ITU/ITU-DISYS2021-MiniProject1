# ITU-DISYS2021-MiniProject1
## Group members
Group name: *Point of Go return*

- Adrian Borup (adbo@itu.dk)
- Andreas Wachs (ahja@itu.dk)
- Anne M. Bartholdy-Falk (anmb@itu.dk)
- Joachim Borup (aljb@itu.dk)

## How to run
```
$ go run ./src
```

## Example output
```
[P #0]: Has eaten 1003 times. Has thought 1306 times. Is eating: false
[P #1]: Has eaten 972 times. Has thought 1336 times. Is eating: true  
[P #2]: Has eaten 916 times. Has thought 1392 times. Is eating: false 
[P #3]: Has eaten 1091 times. Has thought 1219 times. Is eating: true 
[P #4]: Has eaten 633 times. Has thought 1673 times. Is eating: false 
Total eats: 4615, total thinks: 6926.

[F #0]: Has been picked up 1636 times, Is picked up: false.
[F #1]: Has been picked up 1976 times, Is picked up: false.
[F #2]: Has been picked up 1889 times, Is picked up: false.
[F #3]: Has been picked up 2008 times, Is picked up: true.
[F #4]: Has been picked up 1725 times, Is picked up: true.
Total number of fork pickups: 9234.

pickups/eats ratio: 2.000867 (expected: 2).
eats/thinks ratio:  0.666330 (expected: 2/3 = 0.666..).
```
