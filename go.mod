module example.com/go-learn

go 1.13

replace example.com/go-learn/fn => ./fn

replace example.com/go-learn/greet => ./greet

require (
	example.com/go-learn/fn v0.0.0-00010101000000-000000000000
	example.com/go-learn/greet v0.0.0-00010101000000-000000000000
)
