# GoTest2

TO RUN Q1.go:
$ go run Q1.go <int>
  
Expected outcome:
sum
time
parallel sum
parallel time
  
TO RUN Q2.go:
$ go run Q2.go <int>
  
Expected outcome:
sort time
stable time

When running my second program, Stable appeared to be faster. This does not line up with the documented Big O. Stable should take longer since it must traverse the slice to check for equal values before proceeding with the sort. 
