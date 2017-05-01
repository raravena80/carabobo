# logp -- Log Parser

Problem is given a log file like this:
```
<timestamp> Non event
<timestamp> Begin Recording
<timestamp> event a
<timestamp> event b
<timestamp> event c
<timestamp> event d
<timestamp> End Recording
<timestamp> Non event
<timestamp> Non event
<timestamp> Non event
<timestamp> Non event
<timestamp> Begin Recording
<timestamp> event e
<timestamp> event f
<timestamp> End Recording
```
Capture all the 'event x' lines only


Solutions are in Go:

- logp.go - Solution in go. Run with: go run logp.go
- logpalt.go - Alternative solution in go using for (like while). Run with: go run logpalt.go
