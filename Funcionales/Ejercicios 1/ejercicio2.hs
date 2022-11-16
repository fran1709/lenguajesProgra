sub_conjunt:: Eq a => [a] -> [a] -> Bool
sub_conjunt [] _ =  True
sub_conjunt _  [] =  False
sub_conjunt(x:xs) (y:ys) =  
	x == y && sub_conjunt xs ys