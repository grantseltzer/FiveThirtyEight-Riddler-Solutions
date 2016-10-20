-- 0 means heads up
allHeadsUp = take 100 [0,0..]

coinFlip :: (Eq int, Num int) => int ->  int
coinFlip coin
    | coin == 0    = 1
    | coin == 1    = 0
    | otherwise    = -1

mapFlip :: (Integral int, Num int, Eq int) => int -> int -> [int] -> [int]
mapFlip _ _ [] = []
mapFlip x counter (c:cs)
    | counter `mod` x == 0 = (coinFlip c):(mapFlip x (counter+1) cs)
    | otherwise = c:(mapFlip x (counter+1) cs)

mapMapFlip counter list prevList
    | counter == 1 = mapMapFlip (counter+1) (mapFlip counter 1 allHeadsUp) list
    | counter < 99 = mapMapFlip (counter+1) (mapFlip counter 1 prevList) list
    | otherwise    = mapFlip counter 1 prevList

solution = mapMapFlip 1 allHeadsUp []
