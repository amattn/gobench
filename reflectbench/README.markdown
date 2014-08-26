

From this directory, run:

    go test -test.bench="."
    
or the equivalent but shorter:

    go test -bench .

The comments with results are not tested in a very rigid setting (controlling for cache, OS warmup, etc.).  General order of magnitude type comparison can be useful, but I wouldn't take to much stock in any variance less than 10x or so.