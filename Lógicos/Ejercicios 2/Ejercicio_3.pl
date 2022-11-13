%Aplanando listas en una lista.

planish([],[]).
planish([Head|Tail], Rest):- 
	planish(Tail, Rest),
	append(Head, Rest, Result).