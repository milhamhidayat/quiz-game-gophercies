
# quiz-game-gophercies

Quiz Game from gophercies where user must answer question

  

# Local Installation

1. Clone this project

2. Install all package

	```sh
	go get ./...
	```

3. Build this project

	```sh
	go build .
	```

4. Run project

	```sh
	./quiz-game-gophercies -limit=time limit -csv=csv file -random=flag to randomize quiz 
	```
	| Flag | Description | Example |
	|--|--| -- |
	| limit | Quiz time limit in second, default (30s) | ```-limit=5``` |
	| csv | Csv file | ```-csv=problem.csv```
	| random | boolean random (true, false) | ```-random=true```

