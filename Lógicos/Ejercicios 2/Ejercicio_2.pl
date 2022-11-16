pertecene(E,[E|_]).
pertenece(E,[_|Tail]) :- 
	pertenece(E,Tail),!.

subConjunto([],_).
subConjunto([F|L],X):- 
	pertenece(H,X), subConjunto(T,X),!.